package entities

type ColmenaRaspberryPi struct {
	ID				   int `json:"id"`
	ID_Colmena         int `json:"id_colmena"`
	ID_Raspberry       int `json:"id_raspberry"`
	Fecha_Asignacion   string `json:"fecha_asignacion"`
	Estado             string `json:"estado"`
}