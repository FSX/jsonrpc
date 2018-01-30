package jsonrpc

import (
	"testing"
)

type messageTypeTest struct {
	Name     string
	Expected messageType
	Input    string
}

var messageTypeTests = []messageTypeTest{
	{"request", requestMessageType, `{
		"method": "echo",
		"params": {"A": 1, "B": [1, 2, 3], "C": "a string"},
	}`},
	{"result", responseMessageType, `{
		"result": {"A": 1, "B": [1, 2, 3], "C": "a string"},
	}`},
	{"error", responseMessageType, `{
		"error": {"code": -32601, "message": "method not found"},
	}`},
	{"result + error", invalidMessageType, `{
		"result": {"A": 1, "B": [1, 2, 3], "C": "a string"},
		"error": {"code": -32601, "message": "method not found"},
	}`},
	{"null result + error", invalidMessageType, `{
		"result": null,
		"error": {"code": -32601, "message": "method not found"},
	}`},
	{"null result + null error", invalidMessageType, `{
		"result": null,
		"error": null,
	}`},
	{"empty object result + null error", invalidMessageType, `{
		"result": {},
		"error": null,
	}`},
	{"missing result and error", invalidMessageType, `{}`},
	{"garbage", invalidMessageType, `%^&*&^%$%^&*(*&^%$%^&HBVCFRT&&&EEEEE`},
	{"batch request", batchRequestMessageType, `[
		{"method": "echo", "params": "echo echo"},
		{"method": "echo", "params": "echo echo"},
	]`},
	{"batch response", batchResponseMessageType, `[
		{"result": "echo echo"},
		{"result": "echo echo"},
	]`},
	{"mixed batch", invalidBatchMessageType, `[
		{"result": "echo echo"},
		{"method": "echo", "params": "echo echo"},
	]`},
	{"garbage batch", invalidBatchMessageType, `[1, 2, 3, 4, 5]`},
	{"batch request prefixed with whitespace", invalidMessageType, `      [
		{"method": "echo", "params": "echo echo"},
		{"method": "echo", "params": "echo echo"},
	]`},
}

func TestGetMessageType(t *testing.T) {
	for _, test := range messageTypeTests {
		if result := getMessageType([]byte(test.Input)); result != test.Expected {
			t.Errorf("%s: got %d, expected %d", test.Name, result, test.Expected)
		}
	}
}
