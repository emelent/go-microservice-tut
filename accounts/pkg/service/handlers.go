package service

import (
	"encoding/json"
	"log"
	"net"
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

	account.ServedBy = getIP()
	data, _ := json.Marshal(account)
	writeJsonResponse(w, http.StatusOK, data)
}

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "error"
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	panic("Unable to determine local IP address (non loopback). Exiting")
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
	Status string `json:"status"`
}
