package easy_error

import (
	"testing"
	"errors"
)

var testError = errors.New("test error")

func testRecoverer() {
	if r := recover(); r != nil {
		if r == testError {
			return
		}
		t.Error("Got an unexpected error", r)
	}
}

func errorBytes() ([]byte, error) {
	return nil, testError
}

func successBytes() ([]byte, error) {
	return []byte{}, nil
}

func TestBytes(t *testing.T) {
	WrapBytes(successBytes())
	defer testRecoverer()
	t.Error("Should not reach this point", WrapBytes(errorBytes()))
}

func errorFile() (*os.File, error) {
	return nil, testError
}

func successFile() (*os.File, error) {
	return nil, nil
}

func TestFile(t *testing.T) {		
	WrapFile(successFile())
	defer testRecoverer()
	t.Error("Should not reach this point", WrapFile(errorFile()))
}

func errorInterface() (interface{}, error) {
	return nil, testError
}

func successInterface() (interface{}, error) {
	return "good stuff", nil
}

func TestInterface(t *testing.T) {
	t.Logf("Logging to test compatibility with type assertion %s", Wrap(successInterface()).(string))
	defer testRecoverer()
	t.Error("Should not reach this point", Wrap(errorInterface()))
}
