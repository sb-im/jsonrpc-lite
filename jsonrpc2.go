package jsonrpc2

import (
	"encoding/json"
)

func (r *WireRequest) IsNotify() bool {
	return r.ID == nil
}

func (r *WireResponse) IsSuccess() bool {
	return r.Error == nil
}

// combined has all the fields of both Request and Response.
// We can decode this and then work out which it is.
type Jsonrpc struct {
	VersionTag VersionTag       `json:"jsonrpc"`
	ID         *ID              `json:"id,omitempty"`
	Method     string           `json:"method,omitempty"`
	Params     *json.RawMessage `json:"params,omitempty"`
	Result     *json.RawMessage `json:"result,omitempty"`
	Error      *Error           `json:"error,omitempty"`
}

func (this *Jsonrpc) IsResponse() bool {
	return this.Method == ""
}

func (this *Jsonrpc) IsNotify() bool {
	return this.ID == nil
}

func (this *Jsonrpc) IsSuccess() bool {
	return this.Error == nil
}

func MarshalToRaw(obj interface{}) (*json.RawMessage, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	raw := json.RawMessage(data)
	return &raw, nil
}
