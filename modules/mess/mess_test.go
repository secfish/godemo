package mess

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestGetMessage(t *testing.T) {
	want := "GetMessage Welcome!"
	msg := GetMessage()
	if want != msg {
		t.Fatalf(`GetMessage() = %q, want for %q, nil`, msg, want)
	}
}
