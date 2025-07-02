package entities

type Calibracion struct{
	ID                 int `json:"id"`
	ID_Colmena         int `json:"id_colmena"`
	ID_Sensor          int `json:"id_sensor"`
	Valor_Minimo       float64 `json:"valor_minimo"`
	Valor_Maximo       float64 `json:"valor_maximo"`
	Factor_Calibracion float64 `json:"factor_calibracion"`
	Offset_Calibracion float64 `json:"offset_calibracion"`
	Fecha_Calibracion  string `json:"fecha_calibracion"`
}