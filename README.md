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


**GET /users/{userId}**
```
curl -v -X GET http://localhost:3000/users/user1
```

returns `GET /users/{userId}: user id is user1`


**GET /users?userId=user2**
```
curl -v -X GET http://localhost:3000/users?userId=user2
```

returns `GET /users: user id param is user2`
# bmeenan-go-sandbox
