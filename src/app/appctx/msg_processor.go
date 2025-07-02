// Package appctx
package appctx

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/validates"
)

type MessageDecoder struct {
	Body      []byte
	Key       []byte
	Message   string
	Error     string
	Topic     string
	Partition int32
	TimeStamp time.Time
	Offset    int64
	Context   context.Context
}

// DecodeJSON decode kafka message byte to struct
func (d *MessageDecoder) Cast(out interface{}) error {
	if reflect.TypeOf(out).Kind() != reflect.Ptr {
		return fmt.Errorf("%s", "output destination cannot addressable")
	}
	err := json.Unmarshal(d.Body, out)
	if err != nil {
		return err
	}
	validator := validates.New()
	return validator.Request(out)
}
