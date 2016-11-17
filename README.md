# REST-server
*Can be used for Telegram-bots*


### Compiling and running:
```
  go build rest-server.go
  ./rest-server
```
or
```
go run rest-server.go
```

### `GET`-request on a client side:
```
$ curl -i localhost:8080/api/get
  HTTP/1.1 200 OK
  Content-Type: application/json; charset=utf-8
  X-Powered-By: go-json-rest
  Date: Mon, 14 Nov 2016 15:13:29 GMT
  Content-Length: 379

  {
    "dateTime": "1477298259",
    "deltarain": "NULL",
    "geiger": "NULL",
    "illumination": "NULL",
    "inTemp": "17.2222291675031",
    "interval": "0",
    "long_term_geiger": "NULL",
    "long_term_rain": "NULL",
    "outHumidity": "79.9999812477357",
    "outTemp": "100",
    "pressure": "1053.04597883494",
    "usUnits": "16",
    "windDir": "360",
    "windSpeed": "0.00000502980732661427"
}
```
