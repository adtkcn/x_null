package x_null

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

// Int64 支持前端传递null，int，string类型和不传值
// 前端传1，"1"都可以，都转换为int64类型: Int64{Val: 1, Exist: true}
// 前端null值: Int64{Val: nil, Exist: true}
// 前端""值: Int64{Val: nil, Exist: true}
// 前端没传值: Int64{Val: nil, Exist: false}：结构体字段都是零值，并且Value接口返回nil，会忽略更新
type Int64 struct {
	Val   *int64
	Exist bool
}

// NewInt64 创建一个带值的 Int64
func NewInt64(val int64) Int64 {
	return Int64{Val: &val, Exist: true}
}

// DecodeInt 解码整数值
func DecodeInt64(value any) (any, error) {
	result, err := ToInt64(value)
	if err != nil {
		return Int64{Val: nil, Exist: false}, err
	}
	return Int64{Val: result, Exist: true}, nil

	// switch v := value.(type) {
	// case nil:
	// 	return Int64{Val: nil, Exist: true}, nil
	// case Int64:
	// 	return v, nil
	// default:
	// 	result, err := ToInt64(value)
	// 	if err != nil {
	// 		return Int64{Val: nil, Exist: false}, err
	// 	}
	// 	return Int64{Val: result, Exist: true}, nil
	// }
}

// Scan gorm实现Scanner
func (i *Int64) Scan(value any) error {
	result, err := ToInt64(value)
	if err != nil {
		return err
	}
	i.Val = result
	i.Exist = true
	return nil
	// case nil:
	// 	i.Exist = true
	// 	return nil
	// case int64:
	// 	i.Val, i.Exist = &v, true
	// 	return nil
	// case string:
	// 	num, err := strconv.ParseInt(v, 10, 64)
	// 	if err == nil {
	// 		i.Val = &num
	// 		i.Exist = true
	// 	} else {
	// 		i.Exist = false
	// 	}
	// 	return err
	// default:
	// 	return fmt.Errorf("不能将类型 %T 转换为 int64", v)
	// }
}

// Value gorm实现 Valuer
func (i Int64) Value() (driver.Value, error) {
	if !i.Exist {
		return nil, nil
	}
	v := i.Val
	if v == nil {
		return nil, nil
	}
	return *v, nil
}

// String 实现fmt.Stringer接口
func (i Int64) String() string {
	if i.Exist {
		if i.Val == nil {
			return ""
		}
		return strconv.FormatInt(*i.Val, 10)
	}
	return ""
}

// MarshalJSON 实现json序列化接口
func (i Int64) MarshalJSON() ([]byte, error) {
	if i.Exist {
		return json.Marshal(i.Val)
	}
	return json.Marshal(nil)
}

// UnmarshalText 实现 encoding.TextUnmarshaler 接口
func (i *Int64) UnmarshalText(text []byte) error {
	return i.Scan(string(text))
}

// UnmarshalParam 实现gin框架的参数绑定接口
func (i *Int64) UnmarshalParam(param string) error {
	return i.Scan(param)
}

// UnmarshalJSON 实现json反序列化接口,支持 int64, string，null类型，对于float64类型，判断转换前后是否相等，防止精度丢失
func (i *Int64) UnmarshalJSON(data []byte) error {
	var x any
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	result, err := ToInt64(x)
	if err != nil {
		return err
	}
	i.Val = result
	i.Exist = true
	return nil
	// switch v := x.(type) {
	// case nil:
	// 	i.Exist = true
	// 	return nil
	// case int64:
	// 	i.Val = &v
	// 	i.Exist = true
	// 	return nil
	// case float64:
	// 	i64 := int64(v)
	// 	// 判断转换前后是否相等，防止精度丢失
	// 	if float64(i64) != v {
	// 		i.Exist = false
	// 		return errors.New("int64转换失败，" + fmt.Sprintf("%f", v) + "精度丢失")
	// 	}
	// 	i.Val = &i64
	// 	i.Exist = true
	// 	return nil
	// case string:
	// 	if v == "" {
	// 		i.Val = nil
	// 		i.Exist = true
	// 		return nil
	// 	}
	// 	num, err := strconv.ParseInt(v, 10, 64)
	// 	if err == nil {
	// 		i.Val = &num
	// 		i.Exist = true
	// 	} else {
	// 		i.Exist = false
	// 	}
	// 	return err
	// default:
	// 	return fmt.Errorf("不能将类型 %T 转换为 int64, 值为 %v", v, v)
	// }
}

// SetValue 设置值
func (i *Int64) SetValue(value int64) {
	i.Val = &value
	i.Exist = true
}

// SetNull 设置为null
func (i *Int64) SetNull() {
	i.Val = nil
	i.Exist = true
}

// GetValue 获取值指针
func (i *Int64) GetValue() *int64 {
	return i.Val
}

// ValueOr 获取值，不存在则返回默认值
func (i *Int64) ValueOr(v int64) int64 {
	if i.Val == nil {
		return v
	}
	return *i.Val
}

// ValueOrZero 获取值，不存在则返回零值
func (i *Int64) ValueOrZero() int64 {
	if i.Val == nil {
		return 0
	}
	return *i.Val
}

// IsZero go to json时omitempty标签是否忽略该字段
func (i Int64) IsZero() bool {
	return !i.Exist
}

// IsExists 是否存在
func (i Int64) IsExists() bool {
	return i.Exist
}

// IsExistsAndNotNull 存在且不为null
func (i Int64) IsExistsAndNotNull() bool {
	return i.Exist && i.Val != nil
}

// IsExistsAndNull 存在且为null
func (i Int64) IsExistsAndNull() bool {
	return i.Exist && i.Val == nil
}
