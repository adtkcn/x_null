package x_null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

// NullFloat64 支持前端传递null，int，float，string类型和不传值
// 前端传1，"1"都可以，都转换为float64类型: NullFloat64{Val: 1.0, Exist: true}
// 前端null值: NullFloat64{Val: nil, Exist: true}
// 前端没传值: NullFloat64{Val: nil, Exist: false}
type NullFloat64 struct {
	Val   *float64
	Exist bool
}

// NewNullFloat64 创建一个带值的 NullFloat64
func NewNullFloat64(val float64) NullFloat64 {
	return NullFloat64{Val: &val, Exist: true}
}

// DecodeFloat 解码浮点数值
func DecodeFloat(value any) (any, error) {
	switch v := value.(type) {
	case nil:
		return NullFloat64{Val: nil, Exist: false}, nil
	case NullFloat64:
		return v, nil
	default:
		result, err := ToFloat64(value)
		if err != nil {
			return NullFloat64{Val: nil, Exist: false}, err
		}
		return NullFloat64{Val: &result, Exist: true}, nil
	}
}

// Scan gorm实现Scanner
func (f *NullFloat64) Scan(value any) error {
	result, err := ToFloat64(value)
	if err != nil {
		return err
	}
	f.Val, f.Exist = &result, true
	return nil
}

// Value gorm实现 Valuer
func (f NullFloat64) Value() (driver.Value, error) {
	if !f.Exist {
		return nil, nil
	}
	v := f.Val
	if v == nil {
		return nil, nil
	}
	return *v, nil
}

// String 实现fmt.Stringer接口
func (f NullFloat64) String() string {
	if f.Exist {
		if f.Val == nil {
			return ""
		}
		return strconv.FormatFloat(*f.Val, 'f', -1, 64)
	}
	return ""
}

// MarshalJSON 实现json序列化接口
func (f NullFloat64) MarshalJSON() ([]byte, error) {
	if f.Exist {
		return json.Marshal(f.Val)
	}
	return json.Marshal(nil)
}

// UnmarshalText 实现 encoding.TextUnmarshaler 接口
func (i *NullFloat64) UnmarshalText(text []byte) error {
	return i.Scan(string(text))
}

// UnmarshalParam 实现gin框架的参数绑定接口
func (i *NullFloat64) UnmarshalParam(param string) error {
	return i.Scan(param)
}

// UnmarshalJSON 实现json反序列化接口
func (f *NullFloat64) UnmarshalJSON(data []byte) error {
	var x any
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	switch v := x.(type) {
	case nil:
		f.Exist = true
		return nil
	case int64:
		f64 := float64(v)
		f.Val = &f64
		f.Exist = true
		return nil
	case float64:
		f.Val = &v
		f.Exist = true
		return nil
	case string:
		if v == "" {
			f.Val = nil
			f.Exist = true
			return nil
		}
		num, err := strconv.ParseFloat(v, 64)
		if err == nil {
			f.Val = &num
			f.Exist = true
		} else {
			f.Exist = false
		}
		return err
	default:
		return fmt.Errorf("不能将类型 %T 转换为 float64, 值为 %v", v, v)
	}
}

// SetValue 设置值
func (i *NullFloat64) SetValue(value float64) *NullFloat64 {
	i.Val = &value
	i.Exist = true
	return i
}

// SetNull 设置为null
func (i *NullFloat64) SetNull() *NullFloat64 {
	i.Val = nil
	i.Exist = true
	return i
}

// GetValue 获取值指针
func (i *NullFloat64) GetValue() *float64 {
	return i.Val
}

// ValueOr 获取值，不存在则返回默认值
func (i *NullFloat64) ValueOr(v float64) float64 {
	if i.Val == nil {
		return v
	}
	return *i.Val
}

// ValueOrZero 获取值，不存在则返回零值
func (i *NullFloat64) ValueOrZero() float64 {
	if i.Val == nil {
		return 0
	}
	return *i.Val
}

// IsZero go to json时omitempty标签是否忽略该字段
func (i NullFloat64) IsZero() bool {
	return !i.Exist
}

// IsExists 是否存在
func (i NullFloat64) IsExists() bool {
	return i.Exist
}

// IsExistsAndNotNull 存在且不为null
func (i NullFloat64) IsExistsAndNotNull() bool {
	return i.Exist && i.Val != nil
}

// IsExistsAndNull 存在且为null
func (i NullFloat64) IsExistsAndNull() bool {
	return i.Exist && i.Val == nil
}
