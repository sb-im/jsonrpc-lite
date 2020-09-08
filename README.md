# jsonrpc-lite

[![PkgGoDev](https://pkg.go.dev/badge/github.com/SB-IM/jsonrpc-lite)](https://pkg.go.dev/github.com/SB-IM/jsonrpc-lite)
[![Build Status](https://travis-ci.org/SB-IM/jsonrpc-lite.svg?branch=master)](https://travis-ci.org/SB-IM/jsonrpc-lite)
[![codecov](https://codecov.io/gh/SB-IM/jsonrpc-lite/branch/master/graph/badge.svg)](https://codecov.io/gh/SB-IM/jsonrpc-lite)
[![Documentation](https://godoc.org/github.com/SB-IM/jsonrpc-lite?status.svg)](http://godoc.org/github.com/SB-IM/jsonrpc-lite)
[![Go Report Card](https://goreportcard.com/badge/github.com/SB-IM/jsonrpc-lite)](https://goreportcard.com/report/github.com/SB-IM/jsonrpc-lite)
[![GitHub release](https://img.shields.io/github/tag/SB-IM/jsonrpc-lite.svg?label=release)](https://github.com/SB-IM/jsonrpc-lite/releases)
[![license](https://img.shields.io/github/license/SB-IM/jsonrpc-lite.svg?maxAge=2592000)](https://github.com/SB-IM/jsonrpc-lite/blob/master/LICENSE)

Parse and Serialize JSON-RPC2 messages in golang

[https://www.jsonrpc.org/specification](https://www.jsonrpc.org/specification)

Interface Inspired by [https://github.com/teambition/jsonrpc-lite](https://github.com/teambition/jsonrpc-lite)

ID Fork from [golang.org/x/tools/internal/jsonrpc2@feee8acb394c170fe9eb4eb9f538b8877d9a3cd4](https://github.com/golang/tools/commit/feee8acb394c170fe9eb4eb9f538b8877d9a3cd4)

## Features

- No depend
- Only Parse and Serialize JSON-RPC2 messages
- No `net.Conn`, `http`, `websocket` control, You can use any protocol
- Support batch

## Examples

```go
package main

import (
	"fmt"
	"github.com/SB-IM/jsonrpc-lite"
)

func main() {
	recv := []byte(`
	{"jsonrpc": "2.0", "method": "subtract", "params": [42, 23], "id": 1}
	`)

	fmt.Printf("Server RECV: %s\n", recv)
	rpc := jsonrpc.ParseObject(recv)
	if rpc.Type != jsonrpc.TypeRequest {
		panic("Not jsonrpc.TypeRequest")
	}

	send, err := jsonrpc.NewSuccess(rpc.ID, rpc).ToJSON()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server SEND: %s\n", send)
}
```


### Recv

```go
// alone
rpc := jsonrpc.ParseObject([]byte(`
{"jsonrpc": "2.0", "method": "subtract", "params": {"subtrahend": 23, "minuend": 42}, "id": 3}`
))

if rpc.Type != jsonrpc.TypeRequest {
	fmt.Println(rpc)
}


// Batch
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
```

### Request

```go
data, err := jsonrpc.NewRequest("123", "test", []string{"sss", "zzz"}).ToJSON()
if err != nil {
	panic(err)
}
fmt.Printf("%s\n", data)
```

#### Notify

```go
data, err := jsonrpc.NewNotify("test", []string{"sss", "zzz"}).ToJSON()
if err != nil {
	panic(err)
}
fmt.Printf("%s\n", data)
```

#### Success

```go
data, err := jsonrpc.NewSuccess(1234, []string{"sss", "zzz"}).ToJSON()
if err != nil {
	panic(err)
}
fmt.Printf("%s\n", data)
```

#### Error

Custom error

```go
// NewError
data, err := jsonrpc.NewError(233, 1, "This is Error", []string{"sss", "zzz"}).ToJSON()
if err != nil {
	panic(err)
}
fmt.Printf("%s\n", data)
```

#### Errors

The error codes from and including -32768 to -32000 are reserved for pre-defined error

```go
// NewErrors
res := jsonrpc.NewErrors(1234)
res.Errors.ParseError([]string{"data", "data2"})
data, err := res.ToJSON()
if err != nil {
	panic(err)
}
fmt.Printf("%s\n", data)
```

## TODO

- [ ] Automatic verification `ToJSON()` is [JSON-RPC 2.0 Specification](https://www.jsonrpc.org/specification)

