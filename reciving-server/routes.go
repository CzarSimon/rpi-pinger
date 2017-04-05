package main

import (
  "encoding/json"
  "io"
  "log"
  "fmt"
  "net/http"
)

func setupRoutes(env *Env) *http.ServeMux {
  mux := http.NewServeMux()
  mux.HandleFunc("/api/ip", env.IpHandlers)
  return mux
}

func (env *Env) IpHandlers(res http.ResponseWriter, req *http.Request) {
  switch(req.Method) {
    case http.MethodGet:
      env.getIp(res, req)
    case http.MethodPost:
      env.updateStockInfo(res, req)
    default:
      env.getIp(res, req)
  }
}

type Ip struct {
  Address string
}

func (env *Env) updateStockInfo(res http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)
  var ip Ip
  if err := decoder.Decode(&ip); err != io.EOF {
    checkErr(err)
  }
  env.Ip = ip.Address
  msg := fmt.Sprintf("New ip: %s", env.Ip)
  log.Println(msg)
  fmt.Fprintln(res, msg)
}


func (env *Env) getIp(res http.ResponseWriter, req *http.Request) {
  fmt.Fprintln(res, "Ip address: " + env.Ip)
}
