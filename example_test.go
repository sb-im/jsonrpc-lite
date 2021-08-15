package jsonrpc_test

import (
	"fmt"
	"github.com/sb-im/jsonrpc-lite"
)

func ExampleParse() {
	rpc, err := jsonrpc.Parse([]byte(`
	{"jsonrpc": "2.0", "method": "subtract", "params": {"subtrahend": 23, "minuend": 42}, "id": 3}
	`))
	if err != nil {
		panic(err)
	}

	if rpc.Type != jsonrpc.TypeRequest {
		fmt.Println(rpc)
	}
	fmt.Println(rpc)
}

func ExampleParseObject() {
	rpc := jsonrpc.ParseObject([]byte(`
	{"jsonrpc": "2.0", "method": "subtract", "params": {"subtrahend": 23, "minuend": 42}, "id": 3}
	`))
	if rpc.Type != jsonrpc.TypeRequest {
		fmt.Println(rpc)
	}
	fmt.Println(rpc)
}

func ExampleBatch() {
	rpcs := jsonrpc.Batch([]byte(`
	[
		{"jsonrpc": "2.0", "method": "sum", "params": [1,2,4], "id": "1"},
		{"jsonrpc": "2.0", "method": "notify_hello", "params": [7]},
		{"jsonrpc": "2.0", "method": "subtract", "params": [42,23], "id": "2"},
		{"foo": "boo"},
		{"jsonrpc": "2.0", "method": "foo.get", "params": {"name": "myself"}, "id": "5"},
		{"jsonrpc": "2.0", "method": "get_data", "id": "9"}
	]
	`))

	for i, rpc := range rpcs {
		fmt.Println(i, rpc)
	}
}

func ExampleNewRequest() {
	data, err := jsonrpc.NewRequest("123", "test", []string{"sss", "zzz"}).ToJSON()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}

func ExampleNewSuccess() {
	data, err := jsonrpc.NewSuccess(1234, []string{"sss", "zzz"}).ToJSON()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}

func ExampleNewErrors() {
	res := jsonrpc.NewErrors(1234)
	res.Errors.ParseError([]string{"data", "data2"})
	data, err := res.ToJSON()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}
