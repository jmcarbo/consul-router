docker run -d -p 8400:8400 -p 8500:8500 -p 8600:53/udp -h node1 progrium/consul -server -bootstrap -advertise 192.168.59.103
docker run -d  -v /var/run/docker.sock:/tmp/docker.sock -h registrator progrium/registrator  consul://192.168.59.103:8500
curl -XPUT http://192.168.59.103:8500/v1/kv/domain/blabla.localtest.me -d "isawesome"
docker run -d -p 7778:80 rufus/isawesome
