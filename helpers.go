package jsonrpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
)

func Reply(w io.Writer, id Id, result interface{}) error {
	b, err := json.Marshal(result)
	if err != nil {
		return err
	}

	response := &Response{
		JSONRPC: ProtocolVersion,
		Id:      id,
		Result:  (*json.RawMessage)(&b),
	}

	return writeResponse(w, response)
}

func ReplyWithErrorObject(w io.Writer, id Id, obj *Error) error {
	response := &Response{
		JSONRPC: ProtocolVersion,
		Id:      id,
		Error:   obj,
	}

	return writeResponse(w, response)
}

func ReplyWithErrorData(w io.Writer, id Id, code int64, message string, data interface{}) error {
	rerror, err := ErrorWithData(code, message, data)
	if err != nil {
		return errors.New("cannot marshal error data: " + err.Error())
	}

	response := &Response{
		JSONRPC: ProtocolVersion,
		Id:      id,
		Error:   rerror,
	}

	return writeResponse(w, response)
}

func ReplyWithError(w io.Writer, id Id, code int64, message string) error {
	return ReplyWithErrorData(w, id, code, message, nil)
}

func writeResponse(w io.Writer, response *Response) error {
	b, err := json.Marshal(response)
	if err != nil {
		return err
	}

	if _, err := w.Write(b); err != nil {
		return err
	}

	return nil
}

type method struct {
	Name       string
	Function   reflect.Value
	ParamsType reflect.Type
}

func NewMethod(name string, function interface{}) (*method, error) {
	funcType := reflect.TypeOf(function)

	if funcType.Kind() != reflect.Func {
		return nil, fmt.Errorf("%s: function must be a function", name)
	} else if funcType.NumIn() != 2 {
		return nil, fmt.Errorf("%s: function must have two arguments", name)
	} else if funcType.NumOut() != 2 {
		return nil, fmt.Errorf("%s: function must have two return values", name)
	}

	// TODO: Check type of last In. Must be a boolean.
	// TODO: Check out of last Out. Must be an error.

	funcValue := reflect.ValueOf(function)
	paramsType := funcType.In(0)

	return &method{
		Name:       name,
		Function:   funcValue,
		ParamsType: paramsType,
	}, nil
}

func (m *method) Call(args ...interface{}) (interface{}, error) {
	vArgs := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		if value, ok := args[i].(reflect.Value); ok {
			vArgs[i] = value
		} else if args[i] != nil {
			vArgs[i] = reflect.ValueOf(args[i])
		} else {
			return nil, fmt.Errorf("parameter %d is nil", i)
		}
	}

	ret := m.Function.Call(vArgs)

	if err := ret[1].Interface(); err != nil {
		return nil, err.(error)
	}

	return ret[0].Interface(), nil
}

func (m *method) ParamsFromJSON(b *json.RawMessage) (interface{}, error) {
	params := reflect.New(m.ParamsType)

	if b != nil {
		if err := json.Unmarshal(*b, params.Interface()); err != nil {
			return nil, err
		}
	}

	return params.Elem(), nil
}

type Router struct {
	methods map[string]*method
}

func NewRouter(functions map[string]interface{}) (*Router, error) {
	methods := make(map[string]*method)

	for name, function := range functions {
		if m, err := NewMethod(name, function); err == nil {
			methods[m.Name] = m
		} else {
			return nil, err
		}
	}

	return &Router{methods}, nil
}

func (r *Router) Handle(w io.Writer, req *Request) {
	// TODO: Process possible errors returned by Reply* functions.
	// TODO: Log errors that are not usable for users.

	method, exists := r.methods[req.Method]
	if !exists {
		ReplyWithErrorObject(w, req.Id, &Error{
			Code:    CodeMethodNotFound,
			Message: "method not found",
		})
		return
	}

	params, err := method.ParamsFromJSON(req.Params)
	if err != nil {
		ReplyWithErrorObject(w, req.Id, &Error{
			Code:    CodeInvalidParams,
			Message: "cannot parse parameters",
		})
		return
	}

	result, err := method.Call(params, req.IsNotification())
	if err != nil {
		if rerr, ok := err.(*Error); ok {
			ReplyWithErrorObject(w, req.Id, rerr)
		} else {
			ReplyWithError(w, req.Id, CodeInternalError, err.Error())
		}
	} else {
		Reply(w, req.Id, result)
	}
}
