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
var isHealthy = true

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
	if dbUp && isHealthy {
		data, _ := json.Marshal(healthCheckResponse{
			Status: "UP",
		})
		writeJsonResponse(w, http.StatusOK, data)

	} else {
		data, _ := json.Marshal(healthCheckResponse{
			Status: "Database unaccessible",
		})
		writeJsonResponse(w, http.StatusServiceUnavailable, data)
	}

}

func SetHealthyState(w http.ResponseWriter, r *http.Request) {
	state, err := strconv.ParseBool(mux.Vars(r)["state"])
	if err != nil {
		log.Println("Invalid request to SetHealthyState, expected true or false")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	isHealthy = state
	w.WriteHeader(http.StatusOK)
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
