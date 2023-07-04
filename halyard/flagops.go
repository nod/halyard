package halyard

import (
)

type OpBroker struct {

}


/*
NEW
SIG
FIN
GET
DEL
*/

type SFIdent struct {
  owner string
  ctx   string
  slug  string
}

func (sfi *SFIdent) FlagFromDB(db DB) *SignalFlag {
  // avoid a db hit here just create a placeholder ctx
  c := SFContext{ OwnerId: sfi.owner, CtxId: sfi.ctx }
  f,_ := FetchSignalFlag(db, c, sfi.slug)
  return f
}


// Operation: NEW
// Creates a new SignalFlag based on the params, mostly a wrapper around
// the function `NewSignalFlag`
// TODO - should do more checking to ensure not doubling up on an owner's flags
// for slugs
func OpNEW(sfi *SFIdent, slug string, tags []string) *SignalFlag {
  // XXX TODO - check for dup of slug, etc
  c := SFContext{ OwnerId: sfi.owner, CtxId: sfi.ctx }
  f,_ := NewSignalFlag(&c, slug, tags)
  return f
}


