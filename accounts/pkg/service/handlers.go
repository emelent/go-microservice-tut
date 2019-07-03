package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/emelent/go-microservice-tut/accounts/pkg/dbclient"
	"github.com/gorilla/mux"
)

var DBClient dbclient.IBoltClient

func GetAccount(w http.ResponseWriter, r *http.Request) {
	accountID := mux.Vars(r)["accountId"]

	account, err := DBClient.QueryAccount(accountID)

	if err != nil {
		log.Println("Call to GetAccount failed:", err.Error())
		w.WriteHeader(http.StatusNotFound)

		return
	}

	data, _ := json.Marshal(account)
	writeJsonResponse(w, http.StatusOK, data)
	//w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	//w.WriteHeader(http.StatusOK)
	//w.Write(data)

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	dbUp := DBClient.Check()
	var data []byte
	if dbUp {
		data, _ = json.Marshal(healthCheckResponse{
			Status: "UP",
		})

	} else {
		data, _ = json.Marshal(healthCheckResponse{
			Status: "Database unaccessible",
		})
	}

	writeJsonResponse(w, http.StatusOK, data)
}

func writeJsonResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}

type healthCheckResponse struct {
	Status string `json: "status"`
}
