package entities

type Tipo_Sesnores struct{	
	ID                int `json:"id"`
	Nombre            string `json:"nombre"`
	Descripcion       string `json:"descripcion"`
	Unidad_Medida     string `json:"unidad_medida"`
	Tipo_Dato         string `json:"tipo_dato"`
}
