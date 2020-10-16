package handlers

import (
	"net/http"
	"log"
	"CarStatus/database"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)


func PingCarAndSaveStatus(t time.Time) {
	log.Println("Received a GetCar Status request\n",t)
	vinList := [] string{"YS2R4X20005399401", "VLUR4X20009093588", "VLUR4X20009048066"}

	for _, vin := range vinList {
		if ! pingCarToGetStatus(vin) {
			database.SaveDisconnectStatusOfCar(vin)
		}
	}
}

func SaveCarStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Received a POST request\n"))
}

func GetAllCarDetails(w http.ResponseWriter, r *http.Request) {
	log.Printf("Call for retrieving all the car details")
	if r.Method != http.MethodGet {
		log.Printf("Not a GET request")
		return
	}
	organizationDetails := database.GetAllCarDetails()

	jsonOrganizationDetails, err := json.Marshal(organizationDetails)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(jsonOrganizationDetails))
}

func pingCarToGetStatus(vin string) bool {
	return RandBool()
}

/*
RandBool
    This function returns a random boolean value based on the current time
*/
func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 1 {
		return true
	} else {
		return false
	}
}
