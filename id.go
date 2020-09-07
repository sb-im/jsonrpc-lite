package jsonrpc

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Inspired by: https://github.com/golang/tools/blob/master/internal/jsonrpc2/wire.go#L94

// ID is a Request identifier.
// Only one of either the Name or Number members will be set, using the
// number form if the Name is the empty string.
type ID struct {
	Name   string
	Number int64
}

// String returns a string representation of the ID.
// The representation is non ambiguous, string forms are quoted, number forms
// are preceded by a #
func (id *ID) String() string {
	if id == nil {
		return ""
	}
	if id.Name != "" {
		return strconv.Quote(id.Name)
	}
	return "#" + strconv.FormatInt(id.Number, 10)
}

func (id *ID) MarshalJSON() ([]byte, error) {
	if id.Name != "" {
		return json.Marshal(id.Name)
	}
	return json.Marshal(id.Number)
}

func (id *ID) UnmarshalJSON(data []byte) error {
	*id = ID{}
	if err := json.Unmarshal(data, &id.Number); err == nil {
		return nil
	}
	return json.Unmarshal(data, &id.Name)
}

// int(8, 16, 32, 64), string, *ID, ID
// Other type to (string): fmt.Sprint
func NewID(raw interface{}) (id *ID) {
	switch raw.(type) {
	case *ID:
		id = raw.(*ID)
		return
	case ID:
		a := raw.(ID)
		id = &a
		return
	}

	id = &ID{}
	switch raw.(type) {
	case string:
		id.Name = raw.(string)
	case int:
		id.Number = int64(raw.(int))
	case int8:
		id.Number = int64(raw.(int8))
	case int16:
		id.Number = int64(raw.(int16))
	case int32:
		id.Number = int64(raw.(int32))
	case int64:
		id.Number = raw.(int64)
	default:
		id.Name = fmt.Sprint(raw)
	}
	return
}
