package jsonrpc

import (
	"testing"
)

// req: Client --> Server
var test_jsonrpc_method = `{"jsonrpc": "2.0", "method": "subtract", "params": {"subtrahend": 23, "minuend": 42}, "id": 3}`
var test_jsonrpc_notify = `{"jsonrpc": "2.0", "method": "update", "params": [1,2,3,4,5]}`

// res: Server --> Client
var test_jsonrpc_result = `{"jsonrpc": "2.0", "result": 19, "id": 1}`
var test_jsonrpc_errors = `{"jsonrpc": "2.0", "error": {"code": -32601, "message": "Method not found"}, "id": "1"}`

var test_batch_jsonrpc = `
[
	{"jsonrpc": "2.0", "method": "sum", "params": [1,2,4], "id": "1"},
	{"jsonrpc": "2.0", "method": "notify_hello", "params": [7]},
	{"jsonrpc": "2.0", "method": "subtract", "params": [42,23], "id": "2"},
	{"foo": "boo"},
	{"jsonrpc": "2.0", "method": "foo.get", "params": {"name": "myself"}, "id": "5"},
	{"jsonrpc": "2.0", "method": "get_data", "id": "9"}
]
`

func TestParse(t *testing.T) {
	rpc := ParseObject([]byte(test_jsonrpc_method))
	if rpc.Type != TypeRequest {
		t.Error(rpc)
	}
	if rpc.Method != "subtract" {
		t.Error(rpc.Method)
		t.Error(rpc)
	}
}

func TestNotify(t *testing.T) {
	rpc := ParseObject([]byte(test_jsonrpc_notify))
	if rpc.Type != TypeNotify {
		t.Error(rpc)
	}
	if rpc.Method != "update" {
		t.Error(rpc.Method)
		t.Error(rpc)
	}
}

func TestNewError(t *testing.T) {
	rpc := ParseObject([]byte("233"))
	if rpc.Type != TypeInvalid {
		t.Error(rpc)
	}
}

func TestNewErrors(t *testing.T) {
	rpc := NewErrors(nil)
	rpc.Errors.MethodNotFound(nil)
	if rpc.Errors.Code != CodeMethodNotFound {
		t.Error(rpc)
	}
}

func TestBatch(t *testing.T) {
	rpcs := Batch([]byte(test_batch_jsonrpc))
	if len(rpcs) != 6 {
		for i, rpc := range rpcs {
			t.Error(i, rpc)
		}
	}

	mapResult := make(map[int]int, 6)
	mapResult[0] = TypeRequest
	mapResult[1] = TypeNotify
	mapResult[2] = TypeRequest
	mapResult[3] = TypeInvalid
	mapResult[4] = TypeRequest
	mapResult[5] = TypeRequest

	for i, rpc := range rpcs {
		if rpc.Type != mapResult[i] {
			t.Errorf("index == %d, rpc.Type == %d, Should == %d\n", i, rpc.Type, mapResult[i])
			t.Error(i, rpc)
		}
	}

	rpcs2 := Batch([]byte(test_jsonrpc_method))
	if len(rpcs2) != 1 {
		for i, rpc := range rpcs2 {
			t.Error(i, rpc)
		}
	}
}

func TestNewRequest(t *testing.T) {
	//rpc, err := NewRequest("123", "test", []string{"sss", "zzz"})
	rpc := NewRequest("123", "test", []string{"sss", "zzz"})
	if rpc.Type != TypeRequest {
		t.Errorf("%s\n", *rpc.Params)
		//t.Error(err)
	}

	if rpc.Method != "test" {
		t.Error(rpc.Method)
	}
}

func TestRun(t *testing.T) {
	rpc := ParseObject([]byte(test_jsonrpc_method))
	if rpc.Type != TypeRequest {
		t.Error(rpc)
	}
	if rpc.Method != "subtract" {
		t.Error(rpc.Method)
		t.Error(rpc)
	}

	//res, err := NewSuccess(rpc.ID, []string{"sss", "zzz"})
	res := NewSuccess(rpc.ID, []string{"sss", "zzz"})
	if res.Type != TypeSuccess {
		t.Errorf("%s\n", *rpc.Params)
		//t.Error(err)
	}

	if res.Method != "" {
		t.Error(rpc.Method)
	}

	data, err := res.ToJSON()
	if err != nil {
		t.Error(err)
	}
	if ParseObject(data).Type != TypeSuccess {
		t.Errorf("%s\n", data)
	}
}
