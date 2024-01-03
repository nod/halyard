package halyard

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

/**
  *
  * /flags      GET return all flags attached to session
  * /flag       POST add new flag
  * /flag/<fid> PUT edit existing flag
  * /flag/<fid> DELETE
  *
  */

func getFlags(c *gin.Context) {
  flags := []SigFlag{
    {Owner: "b23", Label: "cool one", Status: 0},
    {Owner: "b23", Label: "other one", Status: 1},
    {Owner: "b23", Label: "my one", Status: 1},
  }
  // XXX return an envelope
  c.IndentedJSON(http.StatusOK, flags)
}

func makeFlag(c *gin.Context) {
  log := GetLogger()
  newFlag := NewSigFlag()
  if err := c.BindJSON(&newFlag); err != nil {
    // XXX return an error payload
    return
  }
  // bogus XXX
  newFlag.Owner = "somedude"
  stor,_ := GetStorage()
  fid,err := stor.SaveFlag(newFlag)
  log.Info("saved doc", "id", fid)
  gotDoc := stor.GrabFlag(fid)
  log.Info("docfdb", "doc", gotDoc)

  // XXX return an envelope
  c.IndentedJSON(http.StatusOK, newflag)
}


func StartHTTPServer(uri string) {


  // defaults to localhost:6488
  logger := GetLogger()
  logger.Info("starting http server", "uri", uri)
  router := gin.Default()
  router.GET("/flags", getFlags)
  router.POST("/flag", makeFlag)
  router.Run(uri)
}



