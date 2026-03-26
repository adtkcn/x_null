## Null Types
```go
core.NullFloat64
core.NullInt64
core.NullString
core.NullTime
```

## 解决了什么问题？


1. 数据库在int,string等类型时，同时可能允许null值，但是go中int,string等类型不允许null值，数据类型不一致就会报错。
2. 前端传递null、空，基础类型无法区分。
3. 前端传递字符串数字如："123"，基础类型int64等无法自动转换接收。


## 自定义类型介绍
```go
type NullInt64 struct {
	Val   *int64
	Exist bool
}
```
1. 前端传递空，Val=nil、Exist=false。
2. 前端传递null，Val=nil、Exist=true。
3. 前端传递字符串数字如："123"，Val=123、Exist=true。

## 自定义类型方法
暂时不要使用没有列出的方法。
```go
func (i *NullInt64) Scan(value any) error

func (i NullInt64) Value() (driver.Value, error)

func (i NullInt64) MarshalJSON() ([]byte, error)

func (i *NullInt64) UnmarshalJSON(data []byte) error
func (i *NullInt64) UnmarshalText(text []byte) error
func (i *NullInt64) UnmarshalParam(param string) error


func (i NullInt64) String() string


// ValueOr 获取值，不存在则返回默认值
func (i *NullInt64) ValueOr(v int64) int64
// ValueOrZero 获取值，不存在则返回零值
func (i *NullInt64) ValueOrZero() int64
// IsZero go to json时omitempty标签是否忽略该字段
func (i NullInt64) IsZero() bool 
// IsExists 是否存在
func (i NullInt64) IsExists() bool
// IsExistsAndNotNull 存在且不为null
func (i NullInt64) IsExistsAndNotNull() bool
// IsExistsAndNull 存在且为null
func (i NullInt64) IsExistsAndNull() bool
```
