package jsonrpc

import (
	"encoding/json"
	"errors"
)

const (
	// Internal inference error, default
	TypeInvalid = iota
	// Request object
	TypeRequest
	// Response object result
	TypeSuccess
	// Notification
	TypeNotify
	// Response object error
	TypeErrors
)

// combined has all the fields of both Request and Response.
// We can decode this and then work out which it is.
type Jsonrpc struct {
	Type int `json:"-"`
	// VersionTag is always encoded as the string "2.0"
	VersionTag VersionTag `json:"jsonrpc"`
	// Method is a string containing the method name to invoke.
	Method string `json:"method,omitempty"`
	// Params is either a struct or an array with the parameters of the method.
	Params *json.RawMessage `json:"params,omitempty"`
	// Result is the response value, and is required on success.
	Result *json.RawMessage `json:"result,omitempty"`
	// Error is a structured error response if the call fails.
	Errors *Errors `json:"error,omitempty"`
	// The id of this request, used to tie the Response back to the request.
	// Will be either a string or a number. If not set, the Request is a notify,
	// and no response is possible.
	ID *ID `json:"id,omitempty"`
}

// Verification && json.Marshal object
func (t *Jsonrpc) ToJSON() ([]byte, error) {
	if t.Type == t.getType() {
		return json.Marshal(*t)
	}

	data, err := json.Marshal(*t)
	if err != nil {
		return data, errors.New("Invalid RPC object && " + err.Error())
	}
	return data, errors.New("Invalid RPC object")
}

// Parse message to jsonrpc
func Parse(raw []byte) (*Jsonrpc, error) {
	var jsonrpc Jsonrpc
	err := json.Unmarshal(raw, &jsonrpc)
	jsonrpc.Type = jsonrpc.getType()
	return &jsonrpc, err
}

// Ignore the parse error
func ParseObject(raw []byte) *Jsonrpc {
	jsonrpc, _ := Parse(raw)
	return jsonrpc
}

// Batch ParseObject
// if Parse error len == 0
func Batch(raw []byte) []*Jsonrpc {
	var jsonrpcs []*Jsonrpc
	json.Unmarshal(raw, &jsonrpcs)
	for _, v := range jsonrpcs {
		v.Type = v.getType()
	}

	// Try no batch
	if len(jsonrpcs) == 0 {
		jsonrpc, err := Parse(raw)
		if err == nil {
			jsonrpcs = append(jsonrpcs, jsonrpc)
		}
	}
	return jsonrpcs
}

// Inference jsonrpc type
func (t *Jsonrpc) getType() int {
	switch {
	case t.Method != "" && t.ID == nil && t.Result == nil && t.Errors == nil:
		return TypeNotify
	case t.Method != "" && t.ID != nil && t.Result == nil && t.Errors == nil:
		return TypeRequest
	case t.Method == "" && t.ID != nil && t.Result != nil && t.Errors == nil:
		return TypeSuccess
	case t.Method == "" && t.ID != nil && t.Result == nil && t.Errors != nil:
		return TypeErrors
	default:
		return TypeInvalid
	}
}

// New Request
func NewRequest(id interface{}, method string, params interface{}) *Jsonrpc {
	rawParams, _ := marshalToRaw(params)
	return &Jsonrpc{
		Type:   TypeRequest,
		ID:     NewID(id),
		Method: method,
		Params: rawParams,
	}
}

// New Notification, No ID
func NewNotify(method string, params interface{}) *Jsonrpc {
	rawParams, _ := marshalToRaw(params)
	return &Jsonrpc{
		Type:   TypeNotify,
		Method: method,
		Params: rawParams,
	}
}

// New Response result
func NewSuccess(id, result interface{}) *Jsonrpc {
	rawResult, _ := marshalToRaw(result)
	return &Jsonrpc{
		Type:   TypeSuccess,
		ID:     NewID(id),
		Result: rawResult,
	}
}

// New Response error
func NewError(id interface{}, code int64, message string, data interface{}) *Jsonrpc {
	rawData, _ := marshalToRaw(data)
	return &Jsonrpc{
		Type: TypeErrors,
		ID:   NewID(id),
		Errors: &Errors{
			Code:    code,
			Message: message,
			Data:    rawData,
		},
	}
}

// New a Null Error
// NewErrors(1234).Error.MethodNotFound("233").ToJSON()
func NewErrors(id interface{}) *Jsonrpc {
	return &Jsonrpc{
		Type:   TypeErrors,
		ID:     NewID(id),
		Errors: &Errors{},
	}
}

func marshalToRaw(obj interface{}) (*json.RawMessage, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	raw := json.RawMessage(data)
	return &raw, nil
}
