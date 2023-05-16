
package halyard

import (
  "time"
)

/**

Lifecycle is along the following, depending on update beacons received

LIVE -> WARN -> CRIT -> DEAD
  \      /       /       /
    -> DONE   <-      <-

**/

type status string

const (
  statLIVE status = "LIVE"
  statWARN status = "WARN"
  statCRIT status = "CRIT"
  statDEAD status = "DEAD"
  statDONE status = "DONE"
)

type Burgee struct {
  uid         string
  createdTime  time.Time
  warnTTL    time.Duration
  critTTL    time.Duration
  deadTTL    time.Duration
}

func (b *Burgee) UID() string {
  return b.uid
}

func (b *Burgee) CreatedTime() time.Time {
  return b.createdTime
}

func (b *Burgee) WarnTTL() time.Duration {
  return b.warnTTL
}

func (b *Burgee) CritTTL() time.Duration {
  return b.critTTL
}

func NewBurgee(ttl string, tags []string) (Burgee, error) {
  b := Burgee{}
  b.uid = "asdf"
  b.createdTime = time.Now()

  if ttl == "" {
    b.warnTTL, _ = time.ParseDuration("60m")
  } else {
    var err error
    b.warnTTL, err = time.ParseDuration(ttl)
    if err != nil {
      return b, err
    }
  }

  b.critTTL = 3 * b.warnTTL

  return b,nil
}

