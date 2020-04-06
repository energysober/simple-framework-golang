package session

import (
	"github.com/simple-framework-golang/go-orm/gesorm/log"
	"reflect"
)

// Hooks constants
const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforInsert  = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

// CallMethod calls the registered hooks
func (s *Session) CallMethod(method string, value interface{}) {
	methodValue := reflect.ValueOf(s.RefTable().Model).MethodByName(method)
	if value != nil {
		methodValue = reflect.ValueOf(value).MethodByName(method)
	}
	param := []reflect.Value{reflect.ValueOf(s)}
	if methodValue.IsValid() {
		if v := methodValue.Call(param); len(v) > 0 {
			if err, ok := v[0].Interface().(error); ok {
				log.Error(err)
			}
		}
	}
	return
}
