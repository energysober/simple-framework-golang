package session

import (
	"fmt"
	"github.com/simple-framework-golang/go-orm/gesorm/log"
	"github.com/simple-framework-golang/go-orm/gesorm/shema"
	"reflect"
	"strings"
)

// Model
func (s *Session) Model(value interface{}) *Session {
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = shema.Parse(value, s.dialect)
	}
	return s
}

// RefTable
func (s *Session) RefTable() *shema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set")
	}
	return s.refTable
}

// CreateTable
func (s *Session) CreateTable() error {
	table := s.RefTable()
	var columns []string
	for _, filed := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", filed.Name, filed.Type, filed.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s)", table.Name, desc)).Exec()
	return err
}

// DropTable
func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXIST %s", s.RefTable().Name)).Exec()
	return err
}

// HasTable
func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.RefTable().Name)
	row := s.Raw(sql, values...).QueryRow()
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == s.RefTable().Name
}
