package shema

import (
	"github.com/simple-framework-golang/go-orm/gesorm/dialect"
	"testing"
)

// User test
type User struct {
	Name string `gesorm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.FieldNames) != 2 {
		t.Fatal("Failed to parse user struct")
	}
	if schema.GetFiled("Name").Tag != "PRIMARY KEY" {
		t.Fatal("Failed to parse primary key")
	}
}
