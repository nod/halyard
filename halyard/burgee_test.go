package halyard

import (
    "testing"
    "time"
)

func TestNewBurgee(t *testing.T) {
  b,err := NewBurgee("60s", nil)
  if err != nil {
    t.Fatalf("error whiel creating burgee %v", err)
  }
  if b.UID() == "" {
    t.Fatalf("UID not assigned during creation")
  }
  if b.CreatedTime().IsZero() {
    t.Fatalf("createdtime not assigned")
  }
  min1,_ := time.ParseDuration("1m")
  if b.WarnTTL() != min1 {
    t.Fatalf("warn TTL not set %v", b.WarnTTL())
  }
  min3,_ := time.ParseDuration("3m")
  if b.CritTTL() != min3 {
    t.Fatalf("crit TTL not set")
  }

}
