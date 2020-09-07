package jsonrpc_test

import (
	"fmt"
	"jsonrpc"
)

func ExampleParseObject() {
	rpc := jsonrpc.ParseObject([]byte(`{"jsonrpc": "2.0", "method": "subtract", "params": {"subtrahend": 23, "minuend": 42}, "id": 3}`))
	if rpc.Type != jsonrpc.TypeRequest {
		fmt.Println(rpc)
	}
}
