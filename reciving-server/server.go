package main

import (
  "net/http"
  "log"
)

type Env struct {
  Ip string
}

func main() {
  config := getConfig()
  env := &Env{
    Ip: config.startIp,
  }

  server := &http.Server{
    Addr: config.port,
    Handler: setupRoutes(env),
  }

  err := server.ListenAndServe()
  checkErr(err)
}

func checkErr(err error) {
  if err != nil {
    log.Println(err.Error())
  }
}
