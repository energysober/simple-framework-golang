package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type sqlite3 struct{}

var _ Dialect = (*sqlite3)(nil)

func init() {
	RegisterDialect("sqlite3", &sqlite3{})
}

func (s *sqlite3) DataTypeOf(tye reflect.Value) string {
	switch tye.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := tye.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invaild sql type %s (%s)", tye.Type().Name(), tye.Kind()))
}

func (s *sqlite3) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "SELECT name FROM sqlite_master WHERE type='table' and name = ?", args
}
