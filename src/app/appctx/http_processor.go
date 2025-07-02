// Package appctx
package appctx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/gorilla/schema"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/validates"
)

// Data context untuk http UseCase
type Data struct {
	Request    *http.Request
	BytesValue []byte
}

// Cast data berdasarkan dari http.Request atau MessageProcessor
func (d *Data) Cast(target interface{}) error {
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("target %T cannot addressable, must pointer target", target)
	}

	return d.httpCast(target)
}

func (d *Data) httpCast(target interface{}) error {
	if d.Request == nil {
		return fmt.Errorf("unable to cast http data, null request")
	}

	// httpCast transform request payload data
	// GET -> params-query-string
	// POST -> json-body
	if err := d.grabMethod(target); err != nil {
		return err
	}
	// validate payload request or params
	validator := validates.New()
	return validator.Request(target)

	// return nil
}

// Transform query-string into json struct
func (d *Data) transform(target interface{}, src map[string][]string) error {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	decoder.SetAliasTag("url")
	if err := decoder.Decode(target, src); err != nil {
		return fmt.Errorf("unable to decode query string:%s", err.Error())
	}
	return nil
}

// Grab request method
// Take a destination source of struct
func (d *Data) grabMethod(target interface{}) error {
	switch d.Request.Method {
	case http.MethodPost, http.MethodPut:
		cType := d.Request.Header.Get("Content-Type")
		if !d.isJSON(cType) {
			return fmt.Errorf("unsupported http content-type=%s", cType)
		}
		return d.decodeJSON(d.Request.Body, target)

	case http.MethodGet:
		return d.transform(target, d.Request.URL.Query())
	default:
		return fmt.Errorf("unsupported method or content-type")
	}
}

func (d *Data) isJSON(cType string) bool {
	return cType == "application/json"
}

func (d *Data) decodeJSON(body io.ReadCloser, dst interface{}) error {
	if body == nil {
		return nil
	}
	err := json.NewDecoder(body).Decode(dst)
	if err != nil {
		return fmt.Errorf("unable decode request body, err:%s", err.Error())
	}

	return nil
}
