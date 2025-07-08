package entities

import "time"

// =============================
// Estructuras para Calibracion
// =============================

// Estructura para la solicitud
type CreateCalibracionRequest struct {
	ID_Colmena         int       `json:"id_colmena" example:"1"`
	ID_Sensor          int       `json:"id_sensor" example:"1"`
	Valor_Minimo       float64   `json:"valor_minimo" example:"20.5"`
	Valor_Maximo       float64   `json:"valor_maximo" example:"30.5"`
	Factor_Calibracion float64   `json:"factor_calibracion" example:"1.5"`
	Offset_Calibracion float64   `json:"offset_calibracion" example:"0.5"`
	Fecha_Calibracion  time.Time `json:"fecha_calibracion" example:"2024-03-20T10:00:00Z"`
}

// Resultado de la calibracion
type CalibracionResponse struct {
	ID                 int       `json:"id" example:"1" description:"ID de la calibracion"`
	ID_Colmena         int       `json:"id_colmena" example:"1" description:"ID de la colmena"`
	ID_Sensor          int       `json:"id_sensor" example:"1" description:"ID del sensor"`
	Valor_Minimo       float64   `json:"valor_minimo" example:"20.5" description:"Valor minimo de la calibracion"`
	Valor_Maximo       float64   `json:"valor_maximo" example:"30.5" description:"Valor maximo de la calibracion"`
	Factor_Calibracion float64   `json:"factor_calibracion" example:"1.5" description:"Factor de calibracion"`
	Offset_Calibracion float64   `json:"offset_calibracion" example:"0.5" description:"Offset de la calibracion"`
	Fecha_Calibracion  time.Time `json:"fecha_calibracion" example:"2024-03-20T10:00:00Z" description:"Fecha de la calibracion"`
}

type UpdateCalibracionRequest struct {
	ID_Colmena         int       `json:"id_colmena" example:"1"`
	ID_Sensor          int       `json:"id_sensor" example:"1"`
	Valor_Minimo       float64   `json:"valor_minimo" example:"20.5"`
	Valor_Maximo       float64   `json:"valor_maximo" example:"30.5"`
	Factor_Calibracion float64   `json:"factor_calibracion" example:"1.5"`
	Offset_Calibracion float64   `json:"offset_calibracion" example:"0.5"`
	Fecha_Calibracion  time.Time `json:"fecha_calibracion" example:"2024-03-20T10:00:00Z"`
}

// ErrorResponse modelo para respuestas de error
type ErrorResponse struct {
	// Mensaje de error
	Error string `json:"error" example:"Credenciales inválidas. Por favor, verifique su usuario y contraseña" description:"Mensaje de error"`
}

// =============================
// Estructuras para Colmena
// =============================

// CreateColmenaRequest Modelo para crear una colmena
type CreateColmenaRequest struct {
	Identificador  string `json:"identificador" example:"COL-001" description:"Identificador único de la colmena"`
	Nombre         string `json:"nombre" example:"Colmena Principal" description:"Nombre de la colmena"`
	Area_Ubicacion string `json:"area_ubicacion" example:"Sector Norte" description:"Ubicación física de la colmena"`
	Tipo_Colmena   string `json:"tipo_colmena" example:"Langstroth" description:"Tipo o modelo de la colmena"`
	Estado         string `json:"estado" example:"activo" description:"Estado de la colmena (activo/inactivo)"`
}

// UpdateColmenaRequest Modelo para actualizar una colmena
type UpdateColmenaRequest struct {
	Identificador  string `json:"identificador" example:"COL-001" description:"Identificador único de la colmena"`
	Nombre         string `json:"nombre" example:"Colmena Principal" description:"Nombre de la colmena"`
	Area_Ubicacion string `json:"area_ubicacion" example:"Sector Norte" description:"Ubicación física de la colmena"`
	Tipo_Colmena   string `json:"tipo_colmena" example:"Langstroth" description:"Tipo o modelo de la colmena"`
	Estado         string `json:"estado" example:"activo" description:"Estado de la colmena"`
}

// UpdateEstadoColmenaRequest Modelo para actualizar el estado de una colmena
type UpdateEstadoColmenaRequest struct {
	Estado string `json:"estado" example:"inactivo" description:"Nuevo estado de la colmena" binding:"required,oneof=activo inactivo"`
}

// ColmenaResponse Modelo de respuesta para colmena
type ColmenaResponse struct {
	ID                  int    `json:"id" example:"1" description:"ID de la colmena"`
	Identificador       string `json:"identificador" example:"COL-001" description:"Identificador único de la colmena"`
	Nombre              string `json:"nombre" example:"Colmena Principal" description:"Nombre de la colmena"`
	Area_Ubicacion      string `json:"area_ubicacion" example:"Sector Norte" description:"Ubicación física de la colmena"`
	Tipo_Colmena        string `json:"tipo_colmena" example:"Langstroth" description:"Tipo o modelo de la colmena"`
	Estado              string `json:"estado" example:"activo" description:"Estado actual de la colmena"`
	Fecha_Registro      string `json:"fecha_registro" example:"2024-03-20 10:00:00" description:"Fecha de registro de la colmena"`
	Fecha_Actualizacion string `json:"fecha_actualizacion" example:"2024-03-20 10:00:00" description:"Última fecha de actualización"`
}

// ColmenasListResponse Modelo de respuesta para lista de colmenas
type ColmenasListResponse struct {
	Colmenas []ColmenaResponse `json:"colmenas" description:"Lista de colmenas"`
}

// =============================
// Estructuras para Colmena-Raspberry
// =============================

// CreateColmenaRaspberryRequest Modelo para crear una relación colmena-raspberry
type CreateColmenaRaspberryRequest struct {
	ID_Colmena   int    `json:"id_colmena" example:"1" description:"ID de la colmena"`
	ID_Raspberry int    `json:"id_raspberry" example:"1" description:"ID del dispositivo Raspberry Pi"`
	Estado       string `json:"estado" example:"activo" description:"Estado de la relación" binding:"required,oneof=activo inactivo"`
}

// UpdateColmenaRaspberryRequest Modelo para actualizar una relación colmena-raspberry
type UpdateColmenaRaspberryRequest struct {
	ID_Colmena   int    `json:"id_colmena" example:"1" description:"ID de la colmena"`
	ID_Raspberry int    `json:"id_raspberry" example:"1" description:"ID del dispositivo Raspberry Pi"`
	Estado       string `json:"estado" example:"activo" description:"Estado de la relación" binding:"required,oneof=activo inactivo"`
}

// ColmenaRaspberryResponse Modelo de respuesta para colmena-raspberry
type ColmenaRaspberryResponse struct {
	ID           int       `json:"id" example:"1" description:"ID de la relación"`
	ID_Colmena   int       `json:"id_colmena" example:"1" description:"ID de la colmena"`
	ID_Raspberry int       `json:"id_raspberry" example:"1" description:"ID del dispositivo Raspberry Pi"`
	Estado       string    `json:"estado" example:"activo" description:"Estado de la relación"`
	CreatedAt    time.Time `json:"created_at" example:"2024-03-20T10:00:00Z" description:"Fecha de creación"`
	UpdatedAt    time.Time `json:"updated_at" example:"2024-03-20T10:00:00Z" description:"Fecha de última actualización"`
}

// ColmenaRaspberryListResponse Modelo de respuesta para lista de relaciones colmena-raspberry
type ColmenaRaspberryListResponse struct {
	ColmenasRaspberry []ColmenaRaspberryResponse `json:"colmenas_raspberry" description:"Lista de relaciones colmena-raspberry"`
}
