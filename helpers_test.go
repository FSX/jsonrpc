package jsonrpc

import (
	"errors"
	"testing"
)

func TestNewMethodNoFunc(t *testing.T) {
	nofunc := 1
	if _, err := NewMethod("test", nofunc); err == nil {
		t.Fatal("expected an error, but got none")
	}
}

func TestNewMethodNoArgs(t *testing.T) {
	fn := func() (interface{}, error) {
		return nil, nil
	}

	if _, err := NewMethod("test", fn); err == nil {
		t.Fatal("expected an error, but got none")
	}
}

func TestNewMethodNoReturnValues(t *testing.T) {
	fn := func(_ interface{}, _ bool) {}

	if _, err := NewMethod("test", fn); err == nil {
		t.Fatal("expected an error, but got none")
	}
}

func callMethod(t *testing.T, fn interface{}) (interface{}, error) {
	method, err := NewMethod("test", fn)
	if err != nil {
		return nil, err
	}

	return method.Call("bleep", false)
}

func TestCallMethod(t *testing.T) {
	expected := "bloop"
	fn := func(_ string, _ bool) (string, error) {
		return expected, nil
	}

	result, err := callMethod(t, fn)

	if result != expected {
		t.Fatalf("expected %s, but got %s", expected, result)
	}
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestCallMethodWrongErrorType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected a panic, but it did not happen")
		}
	}()

	type wrongError struct {
		Message string
	}

	fn := func(_ string, _ bool) (string, wrongError) {
		return "", wrongError{"shouldn't work"}
	}

	callMethod(t, fn)
}

func TestCallMethodError(t *testing.T) {
	fn := func(_ string, _ bool) (string, error) {
		return "", errors.New("error terror")
	}

	_, err := callMethod(t, fn)

	if err == nil {
		t.Fatal("expected an error, but got none")
	}
}
