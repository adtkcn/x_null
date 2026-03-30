package x_null

import (
	"fmt"
	"strconv"
	"time"
)

// ToFloat64 将任意值转换为 float64
func ToFloat64(value any) (*float64, error) {
	switch v := value.(type) {
	case nil:
		return nil, nil
	case float64:
		return &v, nil
	case float32:
		f64, err := strconv.ParseFloat(fmt.Sprintf("%f", v), 64)
		return &f64, err
	case int:
		return new(float64(v)), nil
	case int8:
		return new(float64(v)), nil
	case int16:
		return new(float64(v)), nil
	case int32:
		return new(float64(v)), nil
	case int64:
		return new(float64(v)), nil
	case uint:
		return new(float64(v)), nil
	case uint8:
		return new(float64(v)), nil
	case uint16:
		return new(float64(v)), nil
	case uint32:
		return new(float64(v)), nil
	case uint64:
		return new(float64(v)), nil
	case string:
		if v == "" {
			return nil, nil
		}
		f64, err := strconv.ParseFloat(v, 64)
		return &f64, err
	case []uint8:
		f64, err := strconv.ParseFloat(string(v), 64)
		return &f64, err
	default:
		return nil, fmt.Errorf("cannot convert %T to float64", value)
	}
}

// ToInt64 将任意值转换为 int64
func ToInt64(value any) (*int64, error) {
	switch v := value.(type) {
	case nil:
		return nil, nil
	case int:
		return new(int64(v)), nil
	case int8:
		return new(int64(v)), nil
	case int16:
		return new(int64(v)), nil
	case int32:
		return new(int64(v)), nil
	case int64:
		return &v, nil
	case uint:
		return new(int64(v)), nil
	case uint8:
		return new(int64(v)), nil
	case uint16:
		return new(int64(v)), nil
	case uint32:
		return new(int64(v)), nil
	case uint64:
		return new(int64(v)), nil
	case float32:
		i64 := int64(v)
		// 判断转换前后是否相等，防止精度丢失
		if float32(i64) != v {
			return nil, fmt.Errorf("cannot convert float64 %v to int64", value)
		}
		return new(int64(v)), nil
	case float64:
		i64 := int64(v)
		// 判断转换前后是否相等，防止精度丢失
		if float64(i64) != v {
			return nil, fmt.Errorf("cannot convert float64 %v to int64", value)
		}
		return new(int64(v)), nil
	case string:
		if v == "" {
			return nil, nil
		}
		i64, err := strconv.ParseInt(v, 10, 64)
		return &i64, err
	case []uint8:
		i64, err := strconv.ParseInt(string(v), 10, 64)
		return &i64, err
	default:
		return nil, fmt.Errorf("cannot convert %T to int64", value)
	}
}

// ToString 将任意值转换为 string
func ToString(value any) *string {
	switch v := value.(type) {
	case nil:
		return nil
	case string:
		return new(v)
	case []uint8:
		return new(string(v))
	case int:
		return new(strconv.FormatInt(int64(v), 10))
	case int8:
		return new(strconv.FormatInt(int64(v), 10))
	case int16:
		return new(strconv.FormatInt(int64(v), 10))
	case int32:
		return new(strconv.FormatInt(int64(v), 10))
	case int64:
		return new(strconv.FormatInt(v, 10))
	case uint:
		return new(strconv.FormatUint(uint64(v), 10))
	case uint8:
		return new(strconv.FormatUint(uint64(v), 10))
	case uint16:
		return new(strconv.FormatUint(uint64(v), 10))
	case uint32:
		return new(strconv.FormatUint(uint64(v), 10))
	case uint64:
		return new(strconv.FormatUint(v, 10))
	case float32:
		return new(strconv.FormatFloat(float64(v), 'f', -1, 32))
	case float64:
		return new(strconv.FormatFloat(v, 'f', -1, 64))
	case bool:
		return new(strconv.FormatBool(v))
	default:
		return new(fmt.Sprintf("%v", v))
	}
}

var layouts = []string{
	"2006-01-02 15:04:05",       // 最常见的数据库/日志格式
	"2006/01/02 15:04:05",       // 斜杠分隔
	time.RFC3339,                // 2024-01-15T14:30:00Z07:00 (标准网络格式)
	time.RFC3339Nano,            // 带纳秒的 RFC3339
	"2006-01-02",                // 纯日期
	"2006/01/02",                // 纯日期斜杠
	"2006-01-02 15:04:05Z07:00", // 带时区的自定义格式
	"02-01-2006",                // 日-月-年
	"01/02/2006",                // 美式日期
	// time.Kitchen,                // 3:04PM (仅时间)
	// time.RFC1123,                // 邮件/HTTP头常见格式
}

// ToTime
func ToTime(value any) (*time.Time, error) {
	switch v := value.(type) {
	case nil:
		return nil, nil
	case time.Time:
		return &v, nil
	case int64:
		// 解析秒级时间戳
		return new(time.Unix(v, 0)), nil
	case string:
		if v == "" {
			return nil, nil
		}
		for _, layout := range layouts {
			if t, err := time.ParseInLocation(layout, v, time.Local); err == nil {
				return &t, nil
			}
		}
		return nil, fmt.Errorf("cannot convert %T %v to time.Time", value, v)
	default:
		return nil, fmt.Errorf("cannot convert %T to time.Time", value)
	}

}
