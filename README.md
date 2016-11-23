[![Build Status](https://travis-ci.org/swallowws/rest-server.svg?branch=master)](https://travis-ci.org/swallowws/rest-server)

# REST-server
*It can be used for Telegram-bots or other clients*


#### Compiling and running:
Set `$GOPATH` and `$GOBIN` before building 
```
  cd rest-server
  go get
  go build rest-server.go
  ./rest-server -c config.toml
```

#### `GET`-request on a client side:
```
$ curl -i weather.thirdpin.ru/api/get
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
