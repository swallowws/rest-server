# rest-server

### Compiling and running:
```
  go run rest-server.go
```
### `GET`-request on a client side:
```
  $ curl <IP>:8080/api/get
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
