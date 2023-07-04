package halyard

import (
	"testing"
)

func TestMakeLogger(t *testing.T) {
  loggy := NewLogger("")
  if loggy == nil {
    t.Fatalf("failed to make logger")
  }
}
