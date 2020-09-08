package jsonrpc

import (
	"encoding/json"
	"testing"
)

func TestErrorsNil(t *testing.T) {
	var rpcErr *Errors
	if m := rpcErr.Error(); m != "" {
		t.Error(m)
	}
}

func TestParseError(t *testing.T) {
	rawData := "23333333"
	rpcErr := &Errors{}
	rpcErr.ParseError(rawData)
	if rpcErr.Code != CodeParseError {
		t.Error("Code is: ", rpcErr.Code)
	}
	if rpcErr.Message != MessageParseError {
		t.Error("Message is: ", rpcErr.Message)
	}

	if rpcErr.Error() != MessageParseError {
		t.Error("Message is: ", rpcErr.Message)
	}

	var data string
	raw, err := rpcErr.Data.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(raw, &data)
	if err != nil {
		t.Error(err)
	}
	if rawData != data {
		t.Errorf("data is: %s", data)
	}
}

func TestInvalidRequest(t *testing.T) {
	rpcErr := &Errors{}
	rpcErr.InvalidRequest(nil)
	if rpcErr.Code != CodeInvalidRequest {
		t.Error("Code is: ", rpcErr.Code)
	}
	if rpcErr.Message != MessageInvalidRequest {
		t.Error("Message is: ", rpcErr.Message)
	}
}

func TestMethodNotFound(t *testing.T) {
	rpcErr := &Errors{}
	rpcErr.MethodNotFound(nil)
	if rpcErr.Code != CodeMethodNotFound {
		t.Error("Code is: ", rpcErr.Code)
	}
	if rpcErr.Message != MessageMethodNotFound {
		t.Error("Message is: ", rpcErr.Message)
	}
}

func TestInvalidParams(t *testing.T) {
	rpcErr := &Errors{}
	rpcErr.InvalidParams(nil)
	if rpcErr.Code != CodeInvalidParams {
		t.Error("Code is: ", rpcErr.Code)
	}
	if rpcErr.Message != MessageInvalidParams {
		t.Error("Message is: ", rpcErr.Message)
	}
}

func TestInternalError(t *testing.T) {
	rpcErr := &Errors{}
	rpcErr.InternalError(nil)
	if rpcErr.Code != CodeInternalError {
		t.Error("Code is: ", rpcErr.Code)
	}
	if rpcErr.Message != MessageInternalError {
		t.Error("Message is: ", rpcErr.Message)
	}
}
