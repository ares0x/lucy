# tools install
## MySQL
```shell
# pull mysql image
docker pull mysql

# run
docker run -p 3306:3306  \ 
  --name mysql  \
  -v /data/docker/mysql/conf:/etc/mysql/conf.d \ 
  -v /data/docker/mysql/logs:/var/lib/mysql  \
  -e MYSQL_ROOT_PASSWORD=123456  \
  -d mysql
```

## Kong
[kong docs](https://konghq.com/install#kong-community)
```shell
# create kong network
docker network create kong-net

# run postgres container
docker run -d --name kong-database \
  --network=kong-net \
  -p 5432:5432 \
  -e "POSTGRES_USER=kong" \
  -e "POSTGRES_DB=kong" \
  -e "POSTGRES_PASSWORD=kongpass" \
  postgres:13
  
# prepare kong database
docker run --rm --network=kong-net \
 -e "KONG_DATABASE=postgres" \
 -e "KONG_PG_HOST=kong-database" \
 -e "KONG_PG_PASSWORD=kongpass" \
  kong:3.3.0 kong migrations bootstrap 
  
# run
docker run -d --name kong-gateway \
 --network=kong-net \
 -e "KONG_DATABASE=postgres" \
 -e "KONG_PG_HOST=kong-database" \
 -e "KONG_PG_USER=kong" \
 -e "KONG_PG_PASSWORD=kongpass" \
 -e "KONG_PROXY_ACCESS_LOG=/dev/stdout" \
 -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" \
 -e "KONG_PROXY_ERROR_LOG=/dev/stderr" \
 -e "KONG_ADMIN_ERROR_LOG=/dev/stderr" \
 -e "KONG_ADMIN_LISTEN=0.0.0.0:8001, 0.0.0.0:8444 ssl" \
 -p 8000:8000 \
 -p 8443:8443 \
 -p 127.0.0.1:8001:8001 \
 -p 127.0.0.1:8444:8444 \
 kong:3.3.0 
 
# verify
curl -i -X GET --url http://localhost:8001/services 

# admin
http://localhost:8002
```