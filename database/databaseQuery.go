package database

import (
	"log"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func GetAllCarDetails() []OrganizationDetails{
	log.Printf("In database to get all the company and car details")

	// Open our jsonFile
	jsonFile, err := os.Open("./data.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Successfully Opened data.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// we initialize our companies array
	var comDetails []OrganizationDetails

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'listOrganization' which we defined above
	json.Unmarshal(byteValue, &comDetails)
	return comDetails
}


func SaveDisconnectStatusOfCar(vin string) {
	log.Printf("In database to save the car status")
}
