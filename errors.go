package jsonrpc

import (
	"encoding/json"
)

const (
	// CodeParseError is used when invalid JSON was received by the server.
	CodeParseError = -32700
	//CodeInvalidRequest is used when the JSON sent is not a valid Request object.
	CodeInvalidRequest = -32600
	// CodeMethodNotFound should be returned by the handler when the method does
	// not exist / is not available.
	CodeMethodNotFound = -32601
	// CodeInvalidParams should be returned by the handler when method
	// parameter(s) were invalid.
	CodeInvalidParams = -32602
	// CodeInternalError is not currently returned but defined for completeness.
	CodeInternalError = -32603
)

// Error represents a structured error in a Response.
type Errors struct {
	// Code is an error code indicating the type of failure.
	Code int64 `json:"code"`
	// Message is a short description of the error.
	Message string `json:"message"`
	// Data is optional structured data containing additional information about the error.
	Data *json.RawMessage `json:"data,omitempty"`
}

func (err *Errors) Error() string {
	if err == nil {
		return ""
	}
	return err.Message
}

func (e *Errors) ParseError(data interface{}) {
	rawData, _ := marshalToRaw(data)
	e = &Errors{
		Code: CodeParseError,
		Message: "Parse error",
		Data:    rawData,
	}
}

func (e *Errors) InvalidRequest(data interface{}) {
	rawData, _ := marshalToRaw(data)
	e = &Errors{
		Code: CodeInvalidRequest,
		Message: "Invalid Request",
		Data:    rawData,
	}
}

func (e *Errors) MethodNotFound(data interface{}) {
	rawData, _ := marshalToRaw(data)
	e = &Errors{
		Code: CodeMethodNotFound,
		Message: "Method not found",
		Data:    rawData,
	}
}

func (e *Errors) InvalidParams(data interface{}) {
	rawData, _ := marshalToRaw(data)
	e = &Errors{
		Code: CodeInvalidParams,
		Message: "Invalid params",
		Data:    rawData,
	}
}

func (e *Errors) InternalError(data interface{}) {
	rawData, _ := marshalToRaw(data)
	e = &Errors{
		Code: CodeInternalError,
		Message: "Internal error",
		Data:    rawData,
	}
}
