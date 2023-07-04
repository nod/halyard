package halyard

import (
	"bytes"
	"testing"

	badger "github.com/dgraph-io/badger/v4"
)

func NewMemDB() *BadgerDB {
	options := badger.DefaultOptions("").WithInMemory(true)
	dbm, _ := badger.Open(options)
	db := BadgerDB{}
	db.bdb = dbm
	return &db
}

func TestDBGetSet(t *testing.T) {
	db := NewMemDB()
	k := []byte("feh")
	v := []byte("mallow")
	db.Set(k, v)
	v2, _ := db.Get(k)
	if !bytes.Equal(v, v2) {
		t.Fatalf("failed set get")
	}
}

func TestSigSave(t *testing.T) {
	db := NewMemDB()
	o, _ := NewSFOwner("blah")
	c, _ := NewSFContext(o, "bleh", nil)
	s, _ := NewSignalFlag(&c, "something-cool", []string{"asdf", "adf"})
	s.Save(db)

	s2, _ := FetchSignalFlag(db, c, "something-cool")
	if s2 == nil || s2.Slug != "something-cool" {
		t.Fatalf("no sig recovered from db")
	}
	oflags, _ := o.AllSignalFlags(db)
	if len(oflags) != 1 {
		t.Fatalf("no flags found for owner")
	}
	sf2, _ := NewSignalFlag(&c, "another-cool", []string{"asdf", "adf"})
	if sf2 == nil {
		t.Fatalf("no flag made")
	}
	sf2.Save(db)
	oflags2, _ := o.AllSignalFlags(db)
	if len(oflags2) != 2 {
		t.Fatalf("incorrect num of flags found. exp 2, got %d", len(oflags2))
	}

}

func TestDBRange(t *testing.T) {
	db := NewMemDB()
	db.Set([]byte("feh1"), []byte("fellow"))
	db.Set([]byte("meh1"), []byte("mellow1"))
	db.Set([]byte("meh2"), []byte("mellow2"))
	db.Set([]byte("yeh1"), []byte("yellow"))

	// if !bytes.Equal(v, v2) {
	// t.Fatalf("failed set get")
	// }
	vals, _ := db.AllForPrefix([]byte("meh"))
	if 2 != len(vals) {
		t.Fatalf("wrong num returned")
	}
}
