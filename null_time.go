package x_null

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

const (
	// DateFormat 日期格式
	DateFormat = "2006-01-02"
	// TimeFormat 时间格式
	TimeFormat = "2006-01-02 15:04:05"
)

// Time 自定义时间格式
type Time struct {
	Val   *time.Time
	Exist bool
}

// NewTime 创建一个带值的 Time
func NewTime(val time.Time) Time {
	return Time{Val: &val, Exist: true}
}

// DecodeTime 解码字符串
func DecodeTime(value any) (any, error) {
	switch v := value.(type) {
	case nil:
		return Time{Val: nil, Exist: true}, nil
	case Time:
		return v, nil
	default:
		result, err := ToTime(v)
		return Time{Val: result, Exist: true}, err
	}
}

// Scan 读取数据gorm调用
func (t *Time) Scan(value any) error {
	result, err := ToTime(value)
	if err != nil {
		return err
	}
	t.Val = result
	t.Exist = true
	return nil

	// switch val := value.(type) {
	// case nil:
	// 	t.Val = nil
	// 	t.Exist = true
	// 	return nil
	// default:
	// 	result, err := ToTime(val)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	t.Val = &result
	// 	t.Exist = true
	// 	return nil
	// }

	// switch val := v.(type) {
	// case nil:
	// 	t.Val = nil
	// 	t.Exist = true
	// 	return nil
	// case string:
	// 	tt, err := time.ParseInLocation(TimeFormat, val, time.Local)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	t.Val = &tt
	// 	t.Exist = true
	// 	return nil
	// case time.Time:
	// 	tt := val.Format(TimeFormat)
	// 	if tt == "0001-01-01 00:00:00" {
	// 		t.Val = nil
	// 		t.Exist = true
	// 	} else {
	// 		t.Val = &val
	// 		t.Exist = true
	// 	}
	// 	return nil
	// default:
	// 	return fmt.Errorf("不能将类型 %T 转换为 time.Time, 值为 %v", v, v)
	// }
}

// Value 写入数据库gorm调用
func (t Time) Value() (driver.Value, error) {
	timeStr := t.String()
	if timeStr == "" {
		return nil, nil
	}
	return timeStr, nil
}

// String 实现fmt.Stringer接口
func (t Time) String() string {
	if !t.Exist {
		return ""
	}
	if t.Val == nil {
		return ""
	}
	tt := *t.Val
	return tt.Format(TimeFormat)
}

// MarshalJSON 将Time类型的时间转化为JSON字符串格式
func (t Time) MarshalJSON() ([]byte, error) {
	if t.Exist {
		if t.Val == nil {
			return json.Marshal(nil)
		}
		tt := *t.Val
		tStr := tt.Format(TimeFormat)
		return json.Marshal(tStr)
	}
	return json.Marshal(nil)
}

// UnmarshalText 实现 encoding.TextUnmarshaler 接口
func (i *Time) UnmarshalText(text []byte) error {
	return i.Scan(string(text))
}

// UnmarshalParam 实现gin框架的参数绑定接口
func (i *Time) UnmarshalParam(param string) error {
	return i.Scan(param)
}

// UnmarshalJSON 实现json反序列化接口
func (t *Time) UnmarshalJSON(bs []byte) error {
	var date string
	err := json.Unmarshal(bs, &date)
	if err != nil {
		return err
	}

	return t.Scan(date)
	// if date == "" {
	// 	*t = Time{
	// 		Val:   nil,
	// 		Exist: true,
	// 	}
	// 	return nil
	// }
	// tt, err := ToTime(date)
	// // tt, err := time.ParseInLocation(TimeFormat, date, time.Local)
	// if err != nil {
	// 	return err
	// }
	// *t = Time{
	// 	Val:   &tt,
	// 	Exist: true,
	// }
	// return nil
}

// GormDBDataType gorm数据类型
// func (Time) GormDBDataType(db *gorm.DB, field *schema.Field) string {
// 	return "DATETIME"
// }

// SetValue 设置值
func (i *Time) SetValue(value time.Time) {
	i.Val = &value
	i.Exist = true
}

// SetNull 设置为null
func (i *Time) SetNull() {
	i.Val = nil
	i.Exist = true
}

// GetValue 获取值指针
func (i *Time) GetValue() *time.Time {
	return i.Val
}

// ValueOr 获取值，不存在则返回默认值
func (i *Time) ValueOr(v time.Time) time.Time {
	if i.Val == nil {
		return v
	}
	return *i.Val
}

// ValueOrZero 获取值，不存在则返回零值
func (i *Time) ValueOrZero() time.Time {
	if i.Val == nil {
		return time.Time{}
	}
	return *i.Val
}

// IsZero go to json时omitempty标签是否忽略该字段
func (i Time) IsZero() bool {
	return !i.Exist
}

// IsExists 是否存在
func (i Time) IsExists() bool {
	return i.Exist
}

// IsExistsAndNotNull 存在且不为null
func (i Time) IsExistsAndNotNull() bool {
	return i.Exist && i.Val != nil
}

// IsExistsAndNull 存在且为null
func (i Time) IsExistsAndNull() bool {
	return i.Exist && i.Val == nil
}
