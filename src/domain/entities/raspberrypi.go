package entities

type Raspberrypi struct {
	ID                     int `json:"id"`
	Mac                    string `json:"mac"`
	Nombre                 string `json:"nombre"`
	Modelo                 string `json:"modelo"`
	IP_Address             string `json:"ip_address"`
	Estado                 string `json:"estado"`
	Fecha_Registro         string `json:"fecha_registro"`
	Fecha_Ultima_Conexion  string `json:"fecha_ultima_conexion"`
}