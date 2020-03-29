package session

import (
	"github.com/simple-framework-golang/go-orm/gesorm"
	"testing"
)

type User struct {
	Name string `gesorm:PRIMARY KEY`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	e, err := gesorm.NewEngine("sqlite3", "ges.db")
	if err != nil {
		t.Fatal("Failed connect ", err)
	}
	s := e.NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
