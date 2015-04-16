package handler

import (
	"testing"
)

func TestServer(t *testing.T) {
	t.Log("Initializing server...")

	Init()

	t.Log("Correctly initialized!")

	Start()

	t.Log("Started server...")
}
