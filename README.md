# rest-controller-example

## Instructions for use
* go get github.com/gdichter/rest-controller-example
* in IDE File/Open and navigate to $GOPATH/src/github.com/gdichter/rest-controller-example

## Useful curls

**POST /stuff**
```
curl -v -X POST http://localhost:3000/stuff -d '{"this":"that","arr":[{"a1":1,"b1":"bee one"}]}'
```

returns `{"this":"that-NEW","arr":[{"a1":2,"b1":"bee one-UPDATED"}]}`


**GET /stuff**
```
curl -v -X GET http://localhost:3000/stuff'
```
returns from `Get: here is some stuff`


**GET /ping**
```
curl -v -X GET http://localhost:3000/ping -H "the-header:a new fun header value"
```

returns `pong.  you sent a new fun header value`
