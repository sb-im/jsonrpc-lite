# jsonrpc-lite

Parse and Serialize JSON-RPC2 messages in golang

[https://www.jsonrpc.org/specification](https://www.jsonrpc.org/specification)

Interface Inspired by [https://github.com/teambition/jsonrpc-lite](https://github.com/teambition/jsonrpc-lite)

ID Fork from http://golang.org/x/tools/internal/jsonrpc2

[![Build Status](https://travis-ci.org/SB-IM/jsonrpc2.svg?branch=master)](https://travis-ci.org/SB-IM/jsonrpc2)

JSON-RPC 2.0

Base: https://github.com/golang/tools/commit/feee8acb394c170fe9eb4eb9f538b8877d9a3cd4

## Features

- No depend
- Automatic verification [https://www.jsonrpc.org/specification](https://www.jsonrpc.org/specification)
- No `net.Conn` control, You can use any protocol
- Support batch
- Only Parse and Serialize JSON-RPC2 messages

## Use

### Recv

```go
// alone
rpc := Parse([]byte(`{"jsonrpc": "2.0", "method": "subtract", "params": {"subtrahend": 23, "minuend": 42}, "id": 3}`))
if rpc.Type != TypeRequest {
	fmt.Println(rpc)
}


// Batch
rpcs := Batch([]byte(`
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
```

### Send

```go
// Request
rpc, err := NewRequest("123", "test", []string{"sss", "zzz"})
if err != nil {
	fmt.Println(err)
}
data, err := rpc.ToJSON()
if err != nil {
	fmt.Println(err)
}
fmt.Printf("%s", data)
```

#### API

```go
notify, err := NewNotify("test", []string{"sss", "zzz"})
NewSuccess(233, []string{"sss", "zzz"})
NewError(233, CodeParseError, "This is Error", []string{"sss", "zzz"})
```

