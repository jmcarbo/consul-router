Needs progrium/consul and progrium/registrator

```
docker run -d -e CONSUL_HTTP_ADDR=192.168.59.103:8500 -p 7799:8080 jmcarbo/consul-router
```

```
curl -XPUT http://192.168.59.103:8500/v1/kv/domain/www.192.168.59.103.xip.io -d "isawesome"
```
