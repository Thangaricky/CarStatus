package database

type OrganizationDetails struct {
	Address     string `json:"address"`
	CompanyName string `json:"company_name"`
	VehicleList []VehicleDetails `json:"vehicle_list"`
}

type VehicleDetails struct {
	VehicleID          string `json:"VIN"`
	RegistrationNumber string `json:"reg_no"`
	Status			   string `json:"status"`
}