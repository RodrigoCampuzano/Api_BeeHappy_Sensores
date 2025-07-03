package entities

type Colmenas struct{
	ID				   int `json:"id"`	
	Identificador      string `json:"identificador"`
	Nombre             string `json:"nombre"`
	Area_Ubicacion     string `json:"area_ubicacion"`
	Tipo_Colmena       string `json:"tipo_colmena"`
	Estado             string `json:"estado"`
	Fecha_Registro     string `json:"fecha_registro"`
	Fecha_Actualizacion string `json:"fecha_actualizacion"`
}