package main

import (
  "log"
  "os"
  "io"
  //"fmt"

  //"github.com/bitly/go-simplejson"
  "github.com/widuu/gojson"
  //"github.com/ant0ine/go-json-rest/rest"
  "net/http"
  "github.com/gin-gonic/gin"
)


var configStr = 
`
{
    "_id" : ObjectId("5a01769ce0ad42690f50f2f5"),
    "code" : "S0001",
    "key" : "ftp_host",
    "value" : "10.28.13.224",
    "remark" : "DPMS的FTP服务器地址"
}
`

var (
    Info    *log.Logger     //重要的信息
)

func init() {
    file, err := os.OpenFile("c:/dp-dpms/slidewatch.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln("Failed to open log file:", err)
    }
    
    Info = log.New(io.MultiWriter(file, os.Stderr),  "Info: ", log.Ltime|log.Lshortfile)
}

func main() {
  Info.Println("SlideDog startup!")
  
  router := gin.Default()
  router.POST("/webhook/slideReady", func(c *gin.Context) {
            buf := make([]byte, 1024)  
            n, _ := c.Request.Body.Read(buf)  
            Info.Println("slideReady: " + string(buf[0:n]))
            c.String(http.StatusOK, "SlideReady")
  } ).POST("/webhook/snapshotAdded", func(c *gin.Context) {
            buf := make([]byte, 1024)  
            n, _ := c.Request.Body.Read(buf)  
            Info.Println("snapshotAdded: " + string(buf[0:n]))
            c.String(http.StatusOK, "snapshotAdded")
  } ).POST("/webhook/snapshotDeleted", func(c *gin.Context) {
            buf := make([]byte, 1024)  
            n, _ := c.Request.Body.Read(buf)  
            Info.Println("snapshotDeleted: " + string(buf[0:n]))
            c.String(http.StatusOK, "snapshotDeleted")
  } )
  
  
  
  
  /*
  js, err := simplejson.NewJson([]byte(configStr))
  
  if err != nil {
    panic(err.Error())
  }
  
  value := js.Get("value").MustString()
  log.Println("value=%s", value)
  */
  
  //json1 := `{"from":"en","to":"zh","trans_result":[{"src":"today","dst":"\u4eca\u5929"},{"src":"tomorrow","dst":"\u660e\u5929"}]}`
  //c8 := gojson.Json(json1).Get("trans_result").Getkey("src", 1).Tostring()
  
  json1 := `{"value": "123"}`
  c8 := gojson.Json(json1).Get("value").Tostring()
  log.Println(c8) 
  
  router.Run()
}
