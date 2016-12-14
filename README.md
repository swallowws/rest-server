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
Server: nginx/1.4.6 (Ubuntu)
Date: Wed, 14 Dec 2016 07:33:27 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 358
Connection: keep-alive
X-Powered-By: go-json-rest

{
  "dateTime": "1481700727",
  "deltarain": "0",
  "downfall": "0",
  "geiger": "1",
  "illumination": "187",
  "inTemp": "3",
  "interval": "0",
  "long_term_geiger": "135699",
  "long_term_rain": "97",
  "maxWind": "4.00",
  "outHumidity": "40.4",
  "outTemp": "-1.80",
  "pressure": "1012.3",
  "usUnits": "16",
  "windDir": "270",
  "windSpeed": "2.7"
}
```
