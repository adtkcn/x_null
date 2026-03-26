package x_null

import (
	"database/sql/driver"
	"encoding/json"
)

// NullString 支持前端传递null，int，string类型和不传值
// 前端传1，"1"都可以，都转换为string类型: NullString{Val: "1", Exist: true}
// 前端null值: NullString{Val: nil, Exist: true}
// 前端没传值: NullString{Val: nil, Exist: false}
type NullString struct {
	Val   *string
	Exist bool
}

// NewNullString 创建一个带值的 NullString
func NewNullString(val string) NullString {
	return NullString{Val: &val, Exist: true}
}

// DecodeString 解码字符串值
func DecodeString(value any) (any, error) {
	switch v := value.(type) {
	case nil:
		var s string
		return NullString{Val: &s, Exist: true}, nil
	case NullString:
		return v, nil
	default:
		result := ToString(v)
		return NullString{Val: &result, Exist: true}, nil
	}
}

// Scan gorm实现Scanner,支持string, nil类型
func (i *NullString) Scan(value any) error {
	switch v := value.(type) {
	case nil:
		var s string
		i.Val = &s
		i.Exist = true
		return nil
	case string:
		i.Val, i.Exist = &v, true
		return nil
	default:
		result := ToString(v)
		i.Val, i.Exist = &result, true
		return nil
	}
}

// Value gorm实现 Valuer
func (i NullString) Value() (driver.Value, error) {
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
func (i NullString) String() string {
	if i.Val != nil {
		return *i.Val
	}
	return ""
}

// MarshalJSON 实现json序列化接口
func (i NullString) MarshalJSON() ([]byte, error) {
	if i.Exist {
		return json.Marshal(i.Val)
	}
	return json.Marshal(nil)
}

// UnmarshalText 实现 encoding.TextUnmarshaler 接口
func (i *NullString) UnmarshalText(text []byte) error {
	return i.Scan(string(text))
}

// UnmarshalParam 实现gin框架的参数绑定接口
func (i *NullString) UnmarshalParam(param string) error {
	return i.Scan(param)
}

// UnmarshalJSON 实现json反序列化接口
func (i *NullString) UnmarshalJSON(data []byte) error {
	var x any
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	switch v := x.(type) {
	case nil:
		var s string
		i.Val = &s
		i.Exist = true
		return nil
	default:
		result := ToString(v)
		i.Val = &result
		i.Exist = true
		return nil
	}
}

// SetValue 设置值
func (i *NullString) SetValue(value string) {
	i.Val = &value
	i.Exist = true
}

// SetNull 设置为null
func (i *NullString) SetNull() {
	i.Val = nil
	i.Exist = true
}

// GetValue 获取值指针
func (i *NullString) GetValue() *string {
	return i.Val
}

// ValueOr 获取值，不存在则返回默认值
func (i *NullString) ValueOr(v string) string {
	if i.Val == nil {
		return v
	}
	return *i.Val
}

// ValueOrZero 获取值，不存在则返回零值
func (i *NullString) ValueOrZero() string {
	if i.Val == nil {
		return ""
	}
	return *i.Val
}

// IsZero go to json时omitempty标签是否忽略该字段
func (i NullString) IsZero() bool {
	return !i.Exist
}

// IsExists 是否存在
func (i NullString) IsExists() bool {
	return i.Exist
}

// IsExistsAndNotNull 存在且不为null
func (i NullString) IsExistsAndNotNull() bool {
	return i.Exist && i.Val != nil
}

// IsExistsAndNull 存在且为null
func (i NullString) IsExistsAndNull() bool {
	return i.Exist && i.Val == nil
}
