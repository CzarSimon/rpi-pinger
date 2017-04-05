package main

type Config struct {
  certPath, keyPath, startIp, port string
}

func getConfig() Config {
  return Config{
    certPath: "./certs/server.rsa.crt",
    keyPath: "./certs/server.rsa.key",
    startIp: "0.0.0.0",
    port: ":3333",
  }
}
