package halyard

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gosimple/slug"
	"github.com/lithammer/shortuuid/v4"
	"golang.org/x/exp/slog"
)

func NewUUID() string {
	return shortuuid.New()
}

func Slugify(instr string) string {
	return slug.Make(instr)
}

/**

state ->
  if not live:
    return DONE
  let e = now - ts_u
  - e > ttl_hist -> return REMV
  - e > ttl_dead -> return DEAD
  - e > ttl_crit -> return CRIT
  - e > ttl_warn -> return WARN
  - dflt -> return GOOD

keys are set with ttl to ensure garbage collection, on each update the ttl is
updated with the historical value

**/

type state string

const (
	stLIVE state = "LIVE"
	stWARN state = "WARN"
	stCRIT state = "CRIT"
	stDEAD state = "DEAD"
	stDONE state = "DONE"
)

type SFOwner struct {
	OwnerId string `json:oid`
	Name    string `json:nm`
}

func (sfo *SFOwner) Key() string {
	return fmt.Sprintf("own|o:%s", sfo.OwnerId)
}

func (sfo *SFOwner) AllSignalFlags(db DB) ([]SignalFlag, error) {
	keypfx := fmt.Sprintf("sig|o:%s", sfo.OwnerId)
	return SignalFlagsForPrefix(db, keypfx)
}

func NewSFOwner(name string) (SFOwner, error) {
	o := SFOwner{}
	o.OwnerId = NewUUID()
	o.Name = name
	return o, nil
}

type SFContext struct {
	OwnerId string   `json:oid`
	CtxId   string   `json:cid`
	Label   string   `json:lbl`
	Tags    []string `json:tags`
}

func (sfc *SFContext) Key() string {
	return fmt.Sprintf("ctx|o:%s|c:%s", sfc.OwnerId, sfc.CtxId)
}

func (sfc *SFContext) AllSignalFlags(db DB) ([]SignalFlag, error) {
	keypfx := fmt.Sprintf("sig|o:%s|c:%s", sfc.OwnerId, sfc.CtxId)
	return SignalFlagsForPrefix(db, keypfx)
}

func NewSFContext(owner SFOwner, label string, tags []string) (SFContext, error) {
	c := SFContext{}
	c.OwnerId = owner.OwnerId
	c.CtxId = NewUUID()
	c.Label = label
	return c, nil
}

type SignalFlag struct {
	Slug        string        `json:slug`
	OwnerId     string        `json:oid`
	CtxId       string        `json:ctx`
	Tags        []string      `json:tags`
	TimeCreated time.Time     `json:tsc`
	TimeUpdated time.Time     `json:tsu`
	WarnTTL     time.Duration `json:tw`
	CritTTL     time.Duration `json:tc`
	DeadTTL     time.Duration `json:td`
	Live        bool          `json:lv`
}

func (sf *SignalFlag) Key() string {
	return MakeSignalFlagKey(sf.OwnerId, sf.CtxId, sf.Slug)
}

func NewSignalFlag(ctx SFContext, slug string, tags []string) (*SignalFlag, error) {
	s := SignalFlag{}
	s.OwnerId = ctx.OwnerId
	s.CtxId = ctx.CtxId
	s.Slug = Slugify(slug)
	s.TimeCreated = time.Now()
	s.TimeUpdated = s.TimeCreated
	s.WarnTTL, _ = time.ParseDuration("60m")
	s.CritTTL = 3 * s.WarnTTL
	s.DeadTTL = 10 * s.WarnTTL
	s.Live = true
	return &s, nil
}

func (sf *SignalFlag) ToJSON() []byte {
	js, _ := json.Marshal(&sf)
	return js
}

func (sf *SignalFlag) Save(db DB) error {
	js := sf.ToJSON()
	db.Set([]byte(sf.Key()), js)
	slog.Info("saving signalflag", "key", sf.Key())
	return nil
}

func MakeSignalFlagKey(ownerId string, ctxId string, slug string) string {
	return fmt.Sprintf("sig|o:%s|c:%s|s:%s", ownerId, ctxId, slug)
}

func SignalFlagsForPrefix(db DB, keypfx string) ([]SignalFlag, error) {
	rawsf, _ := db.AllForPrefix([]byte(keypfx))
	var ret []SignalFlag
	var sig SignalFlag
	for i := 0; i < len(rawsf); i++ {
		if err := json.Unmarshal(rawsf[i], &sig); err != nil {
			return nil, err
		}
		ret = append(ret, sig)
	}
	return ret, nil
}

func FetchSignalFlag(db DB, ctx SFContext, slug string) (*SignalFlag, error) {
	k := MakeSignalFlagKey(ctx.OwnerId, ctx.CtxId, slug)
	r, _ := db.Get([]byte(k))
	var sig SignalFlag
	if err := json.Unmarshal(r, &sig); err != nil {
		return nil, err
	}
	return &sig, nil
}

