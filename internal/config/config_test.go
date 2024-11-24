package config

import "testing"

func TestNewConfig(t *testing.T) {
	expected := "/some/socket.sock"

	t.Setenv("CSI_ENDPOINT", expected)

	config, err := NewConfig()

	if err != nil {
		t.Fatal("Expected NewConfig not to error")
	}

	actual := config.CSIEndpoint

	if actual != expected {
		t.Fatalf("Expected %v got %v", expected, actual)
	}
}

func TestNewConfigFailure(t *testing.T) {
	_, err := NewConfig()

	if err == nil {
		t.Fatal("Expected NewConfig to error when CSI_ENDPOINT is missing")
	}
}
