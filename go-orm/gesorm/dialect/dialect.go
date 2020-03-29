package dialect

import "reflect"

var defaultDialectMap = map[string]Dialect{}

// Dialect
type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

// RegisterDialect register dialect to default map
func RegisterDialect(name string, dialect Dialect) {
	defaultDialectMap[name] = dialect
}

// GetDialect return a dialect
func GetDialect(name string) (Dialect, bool) {
	dialect, ok := defaultDialectMap[name]
	return dialect, ok
}
