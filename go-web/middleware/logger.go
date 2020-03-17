package middleware

import (
	"github.com/simple-framework-golang/go-web/ges"
	"log"
	"time"
)

// Logger
func Logger() ges.HandlerFunc {
	return func(c *ges.Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf(" %s in %v", c.Req.RequestURI, time.Since(t))
	}
}
