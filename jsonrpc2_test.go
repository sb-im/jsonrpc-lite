package jsonrpc2

import (
	"encoding/json"
	"testing"
)

// req: Client --> Server
var test_jsonrpc_method = `{"jsonrpc": "2.0", "method": "subtract", "params": {"subtrahend": 23, "minuend": 42}, "id": 3}`
var test_jsonrpc_notify = `{"jsonrpc": "2.0", "method": "update", "params": [1,2,3,4,5]}`

// res: Server --> Client
var test_jsonrpc_result = `{"jsonrpc": "2.0", "result": 19, "id": 1}`
var test_jsonrpc_errors = `{"jsonrpc": "2.0", "error": {"code": -32601, "message": "Method not found"}, "id": "1"}`

func Test_IsResponse_not(t *testing.T) {
	jsonrpcs := [2]Jsonrpc{}
	for index, msg := range []string{test_jsonrpc_method, test_jsonrpc_notify} {
		err := json.Unmarshal([]byte(msg), &jsonrpcs[index])
		if err != nil {
			t.Errorf("Not JSON")
		}

		if jsonrpcs[index].IsResponse() {
			t.Errorf("Not Is Response")
		}
	}
}

func Test_IsResponse(t *testing.T) {
	jsonrpcs := [2]Jsonrpc{}
	for index, msg := range []string{test_jsonrpc_result, test_jsonrpc_errors} {
		err := json.Unmarshal([]byte(msg), &jsonrpcs[index])
		if err != nil {
			t.Errorf("Not JSON")
		}

		if !jsonrpcs[index].IsResponse() {
			t.Errorf("It is Response")
		}
	}
}

func Test_IsNotify(t *testing.T) {
	notify := WireRequest{}
	json.Unmarshal([]byte(test_jsonrpc_notify), &notify)
	if !notify.IsNotify() {
		t.Errorf("IsNotify")
	}

	method := WireRequest{}
	json.Unmarshal([]byte(test_jsonrpc_method), &method)
	if method.IsNotify() {
		t.Errorf("Not Notify")
	}
}

func Test_IsSuccess(t *testing.T) {
	result := WireResponse{}
	json.Unmarshal([]byte(test_jsonrpc_result), &result)
	if !result.IsSuccess() {
		t.Errorf("Is Success")
	}

	errors := WireResponse{}
	json.Unmarshal([]byte(test_jsonrpc_errors), &errors)
	if errors.IsSuccess() {
		t.Errorf("Not Success")
	}
}
