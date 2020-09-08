package jsonrpc

import (
	"fmt"
	"testing"
)

func TestIDString(t *testing.T) {
	var id *ID
	if m := id.String(); m != "" {
		t.Error(m)
	}

	id = &ID{
		Number: 233,
	}
	if m := id.String(); m != `#233` {
		t.Error(m)
	}

	id = &ID{
		Name: "233",
	}
	if m := id.String(); m != `"233"` {
		t.Error(m)
	}
}

func TestNewID(t *testing.T) {
	strID := "233"
	id := &ID{
		Name: strID,
	}

	if id.Name != NewID(id).Name {
		t.Errorf("Not ID: %s\n", id.Name)
	}
	if id.Name != NewID(*id).Name {
		t.Errorf("Not ID: %s\n", id.Name)
	}

	if id.Name != NewID(strID).Name {
		t.Errorf("Not ID: %s\n", id.Name)
	}

	if i := NewID(struct {
		TT string
	}{
		TT: "233",
	}).Name; id.Name == i {
		t.Errorf("Not ID: %s\n", i)
	}
}

func TestNewNumberID(t *testing.T) {
	numberID := 23
	id := &ID{
		Number: int64(numberID),
	}

	if fmt.Sprint(id) != fmt.Sprint(NewID(numberID)) {
		t.Error("Not ID: ", id)
	}
	if fmt.Sprint(id) != fmt.Sprint(NewID(int8(numberID))) {
		t.Error("Not ID: ", id)
	}
	if fmt.Sprint(id) != fmt.Sprint(NewID(int16(numberID))) {
		t.Error("Not ID: ", id)
	}
	if fmt.Sprint(id) != fmt.Sprint(NewID(int32(numberID))) {
		t.Error("Not ID: ", id)
	}
	if fmt.Sprint(id) != fmt.Sprint(NewID(int64(numberID))) {
		t.Error("Not ID: ", id)
	}
}
