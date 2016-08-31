package db

import "testing"

const (
	dataSourceName = "test_user:test_pass@tcp(127.0.0.1:13306)/stack_test"
)

func TestOpen(t *testing.T) {
	_, err := Open("mysql", dataSourceName)
	if err != nil {
		t.Error("open connection failed")
	}
}

func TestOpenWithInvalidDriver(t *testing.T) {
	_, err := Open("yoursql", dataSourceName)
	if err == nil {
		t.Error("open connection not failed")
	}
}

func TestOpenWithInvalidPass(t *testing.T) {
	input := "test_user:BAD_PASS@tcp(127.0.0.1:13306)/stack_test"
	_, err := Open("mysql", input)
	if err == nil {
		t.Error("open connection not failed")
	}
}
