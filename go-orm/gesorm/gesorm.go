package gesorm

import (
	"database/sql"
	"github.com/simple-framework-golang/go-orm/gesorm/dialect"
	"github.com/simple-framework-golang/go-orm/gesorm/log"
	"github.com/simple-framework-golang/go-orm/gesorm/session"
)

// Engine db engine
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

// New Engine
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}
	e = &Engine{db: db, dialect: dial}
	log.Info("Connect database success")
	return
}

// Close database
func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Errorf("Failed to close database: %s", err)
	}
	log.Info("Close database success")
}

// NewSession
func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}

type TxFunc func(*session.Session) (interface{}, error)

// Transaction
func (e *Engine) Transaction(f TxFunc) (result interface{}, err error) {
	s := e.NewSession()
	if err := s.Begin(); err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = s.Rollback()
			panic(p)
		} else if err != nil {
			_ = s.Rollback()
		} else {
			err = s.Commit()
		}
	}()
	return f(s)
}
