package errors

import "testing"

const errorMessage = "test error"

func TestNew(t *testing.T) {
	var err error = New(errorMessage)
	if err.Error() != errorMessage {
		t.Fail()
		return
	}
}
