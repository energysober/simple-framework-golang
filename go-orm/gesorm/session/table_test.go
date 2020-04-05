package session

import (
	"database/sql"
	"github.com/simple-framework-golang/go-orm/gesorm/dialect"
	"testing"
)

type User struct {
	Name string `gesorm:PRIMARY KEY`
	Age  int
}

var (
	TestDB      *sql.DB
	TestDial, _ = dialect.GetDialect("sqlite3")
)

// NewSession
func NewSession() *Session {
	return New(TestDB, TestDial)
}

func TestSession_CreateTable(t *testing.T) {
	s := NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
