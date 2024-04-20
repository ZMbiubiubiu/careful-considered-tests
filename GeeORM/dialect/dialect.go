package dialect

import "reflect"

// orm框架需要兼容多种数据库
// 每种数据库都有不同的部分，为了兼容，抽象出这一部分，然后各数据库去实现
var dialectMap = map[string]Dialect{}

type Dialect interface {
	// DataTypeOf 用于将 Go 语言的类型转换为该数据库的数据类型
	DataTypeOf(tye reflect.Value) string
	// TableExistSQL 表是否存在的语句
	TableExistSQL(table string) (string, []interface{})
}

func RegisterDialect(name string, d Dialect) {
	dialectMap[name] = d
}

func GetDialect(name string) (Dialect, bool) {
	dial, ok := dialectMap[name]
	return dial, ok
}
