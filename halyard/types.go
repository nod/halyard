package halyard

import (
  // "github.com/goccy/go-json"
  "time"
)


type SFUser struct {
  uid string `json:"uid" clover:"_id"`
  uname string `json:"uname" clover:"name"`
  secret string `json:"-" clover:"secret"` // not exported to json
  token string `clover:"token"`
}


type SigFlag struct {
  Owner string `json:"-" clover:"uid"`
  Slug string `json:"slug" clover:"slug"`
  Label string `json:"label" clover:"label"`
  // does it really need to store status? or just calc if live at query time?
  Status int   `json:"status" clover:"status"`

  Ctime time.Time
  Utime time.Time
  // xtime   expires time
  // ttl     1x how long after updated to WARN
  //         3x after update is FAIL unless status is set to FIN

}


func NewSigFlag() *SigFlag {
  return &SigFlag{Ctime: time.Now().UTC()}
}


func (sf *SigFlag) Save(s *Storage) {
}
