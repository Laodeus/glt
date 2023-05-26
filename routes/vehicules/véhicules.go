package routesVehicules

type Vehicules struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type DbVehicules struct {
	Vehicules
	Id int `json:"id"`
}
