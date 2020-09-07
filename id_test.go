package jsonrpc

import (
	"fmt"
	"testing"
)

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
	numberID := 233
	id := &ID{
		Number: int64(numberID),
	}

	if fmt.Sprint(id) != fmt.Sprint(NewID(numberID)) {
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
