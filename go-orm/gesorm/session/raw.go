package session

import (
	"database/sql"
	"github.com/simple-framework-golang/go-orm/gesorm/log"
	"strings"
)

// Session
type Session struct {
	db      *sql.DB
	sql     strings.Builder
	sqlVars []interface{}
}

// New a session
func New(db *sql.DB) *Session {
	return &Session{db: db}
}

// Clear session
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

// DB return db
func (s *Session) DB() *sql.DB {
	return s.db
}

// Raw write sql
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// Exec raw sql with exec
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// QueryRows gets a list of records from db
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
