package halyard

import (
	"testing"
	"time"
)

func TestNewSignalFlag(t *testing.T) {
	owner, _ := NewSFOwner("yoyo")
	ctx, _ := NewSFContext(owner, "context", []string{"a"})
	s, err := NewSignalFlag(&ctx, "cool thing", []string{"blah"})
	if err != nil {
		t.Fatalf("error whiel creating sigflag %v", err)
	}
	if s.OwnerId == "" {
		t.Fatalf("UID not assigned during creation")
	}
	if s.TimeCreated.IsZero() {
		t.Fatalf("createdtime not assigned")
	}
	min1, _ := time.ParseDuration("1h")
	if s.WarnTTL != min1 {
		t.Fatalf("warn TTL not set %v", s.WarnTTL)
	}
	min3, _ := time.ParseDuration("3h")
	if s.CritTTL != min3 {
		t.Fatalf("crit TTL not set")
	}
}
