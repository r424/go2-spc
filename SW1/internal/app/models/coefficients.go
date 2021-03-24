package models

type Coefficients struct {
	A int `json:"A"`
	B int `json:"B"`
	C int `json:"C"`
}

type CoefficientsAndSolution struct {
	Coefficients
	Nroots int
}
