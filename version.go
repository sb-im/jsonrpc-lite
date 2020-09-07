package jsonrpc

import (
	"encoding/json"
	"fmt"
)

// VersionTag is a special 0 sized struct that encodes as the jsonrpc version
// tag.
// It will fail during decode if it is not the correct version tag in the
// stream.
type VersionTag struct{}

func (VersionTag) MarshalJSON() ([]byte, error) {
	return json.Marshal("2.0")
}

func (VersionTag) UnmarshalJSON(data []byte) error {
	version := ""
	if err := json.Unmarshal(data, &version); err != nil {
		return err
	}
	if version != "2.0" {
		return fmt.Errorf("Invalid RPC version %v", version)
	}
	return nil
}
