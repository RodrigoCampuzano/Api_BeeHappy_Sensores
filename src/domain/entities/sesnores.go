package entities

type Sensores struct {
	ID              int `json:"id"`
	ID_Tipo_Sensor  int `json:"id_tipo_sensor"`
	ID_Raspberry    int `json:"id_raspberry"`
	Pin_conexion    string `json:"pin_conexion"`
	Direccion_i2c   string `json:"direccion_i2c"`
	Estado          string `json:"estado"`
	Fecha_Registro  string `json:"fecha_registro"`
}