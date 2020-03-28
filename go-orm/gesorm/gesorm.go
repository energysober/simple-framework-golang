package gesorm

import (
	"database/sql"
	"github.com/simple-framework-golang/go-orm/gesorm/log"
	"github.com/simple-framework-golang/go-orm/gesorm/session"
)

// Engine db engine
type Engine struct {
	db *sql.DB
}

// New Engine
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	e = &Engine{db: db}
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
	return session.New(engine.db)
}
