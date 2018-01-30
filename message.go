//go:generate easyjson $GOFILE
package jsonrpc

import (
	"encoding/json"
	"errors"
	"fmt"

	jp "github.com/buger/jsonparser"
)

const ProtocolVersion = "2.0"

// Id represents a JSON-RPC 2.0 request ID.
// TODO: Support integers and null.
type Id string

// Request represents a JSON-RPC request or notification.
//easyjson:json
type Request struct {
	JSONRPC string           `json:"jsonrpc"`
	Method  string           `json:"method"`
	Params  *json.RawMessage `json:"params,omitempty"`
	Id      Id               `json:"id,omitempty"`
}

func (r *Request) IsNotification() bool {
	return len(r.Id) == 0
}

//easyjson:json
type BatchRequest []*Request

// Response represents a JSON-RPC response.
//easyjson:json
type Response struct {
	JSONRPC string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result,omitempty"`
	Error   *Error           `json:"error,omitempty"`
	Id      Id               `json:"id,omitempty"`
}

//easyjson:json
type BatchResponse []*Response

// Error represents a JSON-RPC response error.
//easyjson:json
type Error struct {
	Code    int64            `json:"code"`
	Message string           `json:"message"`
	Data    *json.RawMessage `json:"data,omitempty"`
}

func ErrorWithData(code int64, message string, data interface{}) (*Error, error) {
	rerror := &Error{Code: code, Message: message}

	if data != nil {
		if b, err := json.Marshal(data); err == nil {
			rerror.Data = (*json.RawMessage)(&b)
		} else {
			return nil, err
		}
	}

	return rerror, nil
}

func (e *Error) Error() string {
	return fmt.Sprintf("jsonrpc: code %v message: %s", e.Code, e.Message)
}

// JSONRPC 2.0 error codes.
const (
	CodeParseError       = -32700
	CodeInvalidRequest   = -32600
	CodeMethodNotFound   = -32601
	CodeInvalidParams    = -32602
	CodeInternalError    = -32603
	codeServerErrorStart = -32099
	codeServerErrorEnd   = -32000
)

type messageType int

const (
	undefinedMessage messageType = iota
	requestMessageType
	responseMessageType
	invalidMessageType
	batchRequestMessageType
	batchResponseMessageType
	invalidBatchMessageType
)

func getMessageType(data []byte) messageType {
	getType := func(data []byte) messageType {
		var method, result, rerror bool

		if v, t, _, err := jp.Get(data, "method"); err == nil && t != jp.NotExist {
			method = t == jp.Null || len(v) > 0
		}
		if v, t, _, err := jp.Get(data, "result"); err == nil && t != jp.NotExist {
			result = t == jp.Null || len(v) > 0
		}
		if v, t, _, err := jp.Get(data, "error"); err == nil && t != jp.NotExist {
			rerror = t == jp.Null || len(v) > 0
		}

		// The following conditions are valid:
		// 1. a method field and no result and error fields.
		// 2. a result field and no method and error fields.
		// 3. an error field and no method and result fields.
		if method && !result && !rerror {
			return requestMessageType
		} else if !method && ((result && !rerror) || (!result && rerror)) {
			return responseMessageType
		} else {
			return invalidMessageType
		}
	}

	if len(data) > 0 && data[0] == '[' {
		batchType := undefinedMessage

		jp.ArrayEach(data, func(v []byte, t jp.ValueType, _ int, err error) {
			if batchType == invalidBatchMessageType {
				return
			}

			messageType := getType(v)
			if messageType == invalidMessageType {
				batchType = invalidBatchMessageType
			} else if batchType == undefinedMessage {
				batchType = messageType
			} else if messageType != batchType {
				batchType = invalidBatchMessageType
			}
		})

		if batchType == requestMessageType {
			batchType = batchRequestMessageType
		} else if batchType == responseMessageType {
			batchType = batchResponseMessageType
		}

		return batchType
	} else {
		return getType(data)
	}
}

// Message holds one of the unmarshaled messages (request or response).
type Message struct {
	Request       *Request
	Response      *Response
	BatchRequest  BatchRequest
	BatchResponse BatchResponse
}

func (m *Message) UnmarshalJSON(data []byte) error {
	switch getMessageType(data) {
	case requestMessageType:
		return json.Unmarshal(data, &m.Request)
	case responseMessageType:
		return json.Unmarshal(data, &m.Response)
	case invalidMessageType:
		return errors.New("unable to determine message type")
	case batchRequestMessageType:
		return json.Unmarshal(data, &m.BatchRequest)
	case batchResponseMessageType:
		return json.Unmarshal(data, &m.BatchResponse)
	case invalidBatchMessageType:
		return errors.New("unable to determine batch message type")
	}

	return nil
}
