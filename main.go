package main

import (
    "github.com/AaronO/gogo-proxy"
      "net/http"
      "fmt"
      "errors"
      "github.com/hashicorp/consul/api"
      "strings"
      "math/rand"
      "log"
    )

func main() {
  //consulUrl := os.Getenv("CONSUL_HTTP_ADDR")
  client, _ := api.NewClient(api.DefaultConfig())
  kv := client.KV()
  health := client.Health()

  //fmt.Printf("%s\n", consulUrl)
  p, _ := proxy.New(proxy.ProxyOptions{
    Balancer: func(req *http.Request) (string, error) {
      a := strings.Split(req.Host, ":")
      pair, _, err := kv.Get("domain/"+a[0], nil)
      if err != nil {
          fmt.Println(err)
      } else {
        if pair != nil {
          serviceConfig := strings.Split(string(pair.Value),":")
          protocol := "http"
          if len(serviceConfig)>1 {
            protocol=serviceConfig[1] 
          }
          services, _ , err:=health.Service(serviceConfig[0], "", true, nil)
          if err != nil {
          } else {
            serviceCount := len(services)
            if serviceCount > 0 {
              idx := rand.Intn(serviceCount)
              port :=services[idx].Service.Port
              address:=services[idx].Node.Address
              //fmt.Printf("%v\n", port)
              //fmt.Printf("%v\n", address)
              log.Printf("%s %s %s --> %s %d", req.RemoteAddr, req.Method, req.URL, address, port)
              return fmt.Sprintf("%s://%s:%d", protocol, address, port), nil
            }
          }
        }
      }
      log.Printf("%s %s %s --> NO MATCH", req.RemoteAddr, req.Method, req.URL)
      return "", errors.New("Domain not found") 
    },
    /*
    ErrorHandler: func(rw http.ResponseWriter, req *http.Request, err error){
      fmt.Fprintf(rw, "Domain not found: %s", err)
    },
    */
  })
  http.ListenAndServe(":8080", p)
}
