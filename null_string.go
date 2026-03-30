package x_null

import (
	"database/sql/driver"
	"encoding/json"
)

// String 支持前端传递null，int，string类型和不传值
// 前端传1，"1"都可以，都转换为string类型: String{Val: "1", Exist: true}
// 前端null值: String{Val: nil, Exist: true}
// 前端""值: String{Val: "", Exist: true}
// 前端没传值: String{Val: nil, Exist: false}
type String struct {
	Val   *string
	Exist bool
}

// NewString 创建一个带值的 String
func NewString(val string) String {
	return String{Val: &val, Exist: true}
}

// DecodeString 解码字符串值
func DecodeString(value any) (any, error) {
	switch v := value.(type) {
	case nil:
		return String{Val: nil, Exist: true}, nil
	case String:
		return v, nil
	default:
		result := ToString(v)
		return String{Val: result, Exist: true}, nil
	}
}

// Scan gorm实现Scanner,支持string, nil类型
func (i *String) Scan(value any) error {
	switch v := value.(type) {
	case nil:
		i.Val = nil
		i.Exist = true
		return nil
	case string:
		i.Val, i.Exist = &v, true
		return nil
	default:
		result := ToString(v)
		i.Val, i.Exist = result, true
		return nil
	}
}

// Value gorm实现 Valuer
func (i String) Value() (driver.Value, error) {
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
func (i String) String() string {
	if i.Val != nil {
		return *i.Val
	}
	return ""
}

// MarshalJSON 实现json序列化接口
func (i String) MarshalJSON() ([]byte, error) {
	if i.Exist {
		return json.Marshal(i.Val)
	}
	return json.Marshal(nil)
}

// UnmarshalText 实现 encoding.TextUnmarshaler 接口
func (i *String) UnmarshalText(text []byte) error {
	return i.Scan(string(text))
}

// UnmarshalParam 实现gin框架的参数绑定接口
func (i *String) UnmarshalParam(param string) error {
	return i.Scan(param)
}

// UnmarshalJSON 实现json反序列化接口
func (i *String) UnmarshalJSON(data []byte) error {
	var x any
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	switch v := x.(type) {
	case nil:
		i.Val = nil
		i.Exist = true
		return nil
	default:
		result := ToString(v)
		i.Val = result
		i.Exist = true
		return nil
	}
}

// SetValue 设置值
func (i *String) SetValue(value string) {
	i.Val = &value
	i.Exist = true
}

// SetNull 设置为null
func (i *String) SetNull() {
	i.Val = nil
	i.Exist = true
}

// GetValue 获取值指针
func (i *String) GetValue() *string {
	return i.Val
}

// ValueOr 获取值，不存在则返回默认值
func (i *String) ValueOr(v string) string {
	if i.Val == nil {
		return v
	}
	return *i.Val
}

// ValueOrZero 获取值，不存在则返回零值
func (i *String) ValueOrZero() string {
	if i.Val == nil {
		return ""
	}
	return *i.Val
}

// IsZero go to json时omitempty标签是否忽略该字段
func (i String) IsZero() bool {
	return !i.Exist
}

// IsExists 是否存在
func (i String) IsExists() bool {
	return i.Exist
}

// IsExistsAndNotNull 存在且不为null
func (i String) IsExistsAndNotNull() bool {
	return i.Exist && i.Val != nil
}

// IsExistsAndNull 存在且为null
func (i String) IsExistsAndNull() bool {
	return i.Exist && i.Val == nil
}
