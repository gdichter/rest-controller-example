package stuff

import (
	"net/http"
	"io/ioutil"
	"github.com/gdichter/rest-controller-example/request"
	"encoding/json"
)

func HandleGetStuff() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("from Get: here is some stuff"))
	}
}

func HandlePostStuff() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		response := handleRequest(body)
		w.Write(response)
	}
}

func handleRequest(bytes []byte) []byte {
	incomingRequest := request.MyRequest{}
	unmarshallErr := json.Unmarshal(bytes, &incomingRequest)
	if unmarshallErr != nil {
		panic(unmarshallErr)
	}

	originalArg1 := incomingRequest.Arg1
	incomingRequest.Arg1 = originalArg1 + "-NEW"

	originalA1 := incomingRequest.Arr[0].A1
	incomingRequest.Arr[0].A1 = originalA1 + 1

	originalB1 := incomingRequest.Arr[0].B1
	incomingRequest.Arr[0].B1 = originalB1 + "-UPDATED"

	data, err := json.Marshal(&incomingRequest)
	if err != nil {
		panic(err)
	}
	return data


}

