package jsonrpc

import (
	"encoding/json"
)

const (
	// CodeParseError is used when invalid JSON was received by the server.
	CodeParseError    = -32700
	MessageParseError = "Parse error"
	//CodeInvalidRequest is used when the JSON sent is not a valid Request object.
	CodeInvalidRequest    = -32600
	MessageInvalidRequest = "Invalid Request"
	// CodeMethodNotFound should be returned by the handler when the method does
	// not exist / is not available.
	CodeMethodNotFound    = -32601
	MessageMethodNotFound = "Method not found"
	// CodeInvalidParams should be returned by the handler when method
	// parameter(s) were invalid.
	CodeInvalidParams    = -32602
	MessageInvalidParams = "Invalid params"
	// CodeInternalError is not currently returned but defined for completeness.
	CodeInternalError    = -32603
	MessageInternalError = "Internal error"
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

// Set Errors ParseError
func (e *Errors) ParseError(data interface{}) {
	rawData, _ := marshalToRaw(data)
	e.Code = CodeParseError
	e.Message = MessageParseError
	e.Data = rawData
}

// Set Errors InvalidRequest
func (e *Errors) InvalidRequest(data interface{}) {
	rawData, _ := marshalToRaw(data)
	e.Code = CodeInvalidRequest
	e.Message = MessageInvalidRequest
	e.Data = rawData
}

// Set Errors MethodNotFound
func (e *Errors) MethodNotFound(data interface{}) {
	rawData, _ := marshalToRaw(data)
	e.Code = CodeMethodNotFound
	e.Message = MessageMethodNotFound
	e.Data = rawData
}

// Set Errors InvalidParams
func (e *Errors) InvalidParams(data interface{}) {
	rawData, _ := marshalToRaw(data)
	e.Code = CodeInvalidParams
	e.Message = MessageInvalidParams
	e.Data = rawData
}

// Set Errors InternalError
func (e *Errors) InternalError(data interface{}) {
	rawData, _ := marshalToRaw(data)
	e.Code = CodeInternalError
	e.Message = MessageInternalError
	e.Data = rawData
}
