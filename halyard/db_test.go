package halyard

import (
  "io/ioutil"
  "log"
  "os"
  "testing"

	"github.com/ostafen/clover/v2"
	doc "github.com/ostafen/clover/v2/document"
)


type test_with_db func(*clover.DB)


func testWithDb(t *testing.T, testfn test_with_db) {
  dir, err := ioutil.TempDir("/tmp", "test_hy_")
  if err != nil {
    log.Fatal(err)
  }
  defer os.RemoveAll(dir)
  err = ConnectDB(dir)
  if err != nil {
    t.Fatalf("unable to create db at %s.  ERR %v", dir, err)
  }
  db,err := GetDB()
  if err != nil {
    t.Fatalf("unable to create db at %s.  ERR %v", dir, err)
  }
  defer CloseDB()

  testfn(db)
}


func TestGetDBbeforeConnect(t *testing.T) {
  db,err := GetDB()
  if db != nil || err == nil {
    t.Fatalf("db creation failed %v", err)
  }
}


func TestSaveSomethingInDb(t *testing.T) {
  testWithDb(t, func(db *clover.DB) {
    db.CreateCollection("test")
    doc := doc.NewDocument()
    doc.Set("hello", "clover!")

    EnsureCollection("testcol")
    // InsertOne returns the id of the inserted document
    docId, _ := db.InsertOne("testcol", doc)
    if docId == "" {
      t.Fatalf("docId is %s, should be nil", docId)
    }
    _, err := GrabDoc("testcol", docId)
    if err != nil {
      t.Fatalf("unable to grab from db: %v", err)
    }
  })
}


func TestSaveDoc(t *testing.T) {
  testWithDb(t, func(db *clover.DB) {
    doc := doc.NewDocument()
    doc.Set("hello", "clover!")
    EnsureCollection("junk")
    docId,err := SaveDoc("junk", doc)
    if docId == "" || err != nil {
      t.Fatalf("docId is %s, should be nil.  err: %s", docId, err)
    }
  })
}

