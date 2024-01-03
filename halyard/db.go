package halyard

import (
  "fmt"
  "log"

  "github.com/ostafen/clover/v2"
  cq "github.com/ostafen/clover/v2/query"
  cd "github.com/ostafen/clover/v2/document"
)


type Storage struct {
  Db *clover.DB
}

var storage *Storage 

const COLUSERS = "users"
const COLFLAGS = "flags"

// ConnectDB(fpath) should be called prior to calling GetDB() to instantiate and
// connect the database.  Calling with the the path to an existing database will
// open that database for use. Calling with a path to a non-existent database
// will create a new database.

func (s *Storage) ConnectDB(fpath string) (error) {
  // XXX TODO check for path, etc
  var err error
  s.Db,err = clover.Open(fpath)
  return err
}


func (s *Storage) Close() {
  s.Db.Close()
}


func GetStorage() (*Storage,error) {
  err := ensureStorage()
  if err == nil {
    return storage,nil
  }
  return nil,err
}


func ensureStorage() error {
  if storage == nil || storage.Db == nil {
    return fmt.Errorf("DBERR: attempted to GetDB on nil db instance")
  }
  return nil
}


func (s *Storage) ensureCollection(collection string) {
  // Check if collection already exists
  collectionExists, err := s.Db.HasCollection(collection)
  if err != nil {
    log.Panicf("Failed to check collection: %v", err)
  }
  if !collectionExists {
    // Create a collection named 'todos'
    s.Db.CreateCollection(collection)
  }
}


func (s *Storage) SaveDoc(col string, doc *cd.Document) (string, error) {
  docid, err := s.Db.InsertOne(col, doc)
  return docid, err
}


func (s *Storage) SaveFlag(flg *SigFlag) (string, error) {
  return s.SaveDoc(COLFLAGS,  cd.Document(flg))
}


func (s *Storage) SaveUser(usr *SFUser) (string, error) {
  return s.SaveDoc(COLUSERS, cd.Document(usr))
}


func (s *Storage) GrabDoc(col string, docid string) (*cd.Document, error) {
  q := cq.NewQuery(col).Where(cq.Field("_id").Eq(docid))
  return s.Db.FindFirst(q)
}


func (s *Storage) GrabUser(uid string) (*SFUser, error) {
  if err := ensureStorage(); err != nil { return nil, err }
  // var user *SFUser
  udoc,graberr := s.GrabDoc(COLUSERS, uid)
  if graberr != nil {
    return nil, graberr
  }
  var usr *SFUser
  udoc.Unmarshal(usr)
  return usr,nil
}


func (s *Storage) GrabFlag(fid string) (*SigFlag, error) {
  if err := ensureStorage(); err != nil { return nil, err }
  // var user *SFUser
  fdoc,graberr := s.GrabDoc(COLFLAGS, fid)
  if graberr != nil { return nil, graberr }
  var flg *SigFlag
  fdoc.Unmarshal(flg)
  return flg,nil
}

