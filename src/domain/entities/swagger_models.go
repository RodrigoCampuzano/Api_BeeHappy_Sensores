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

// =============================
// Estructuras para RaspberryPi
// =============================

// CreateRaspberryPiRequest Modelo para crear un dispositivo Raspberry Pi
type CreateRaspberryPiRequest struct {
	MAC_Address string `json:"mac" example:"00:11:22:33:44:55" binding:"required" description:"Dirección MAC del dispositivo"`
	IP_Address  string `json:"ip_address" example:"192.168.1.100" binding:"required" description:"Dirección IP del dispositivo"`
	Nombre      string `json:"nombre" example:"Raspberry-001" binding:"required" description:"Nombre identificativo del dispositivo"`
	Modelo      string `json:"modelo" example:"Raspberry Pi 4B" binding:"required" description:"Modelo del dispositivo"`
	Estado      string `json:"estado" example:"activo" binding:"required,oneof=activo inactivo" description:"Estado del dispositivo"`
}

// UpdateRaspberryPiRequest Modelo para actualizar un dispositivo Raspberry Pi
type UpdateRaspberryPiRequest struct {
	MAC_Address string `json:"macj" example:"00:11:22:33:44:55" description:"Dirección MAC del dispositivo"`
	IP_Address  string `json:"ip_address" example:"192.168.1.100" description:"Dirección IP del dispositivo"`
	Nombre      string `json:"nombre" example:"Raspberry-001" description:"Nombre identificativo del dispositivo"`
	Estado      string `json:"estado" example:"activo" description:"Estado del dispositivo" binding:"required,oneof=activo inactivo"`
	Modelo      string `json:"modelo" example:"Raspberry Pi 4B" binding:"required" description:"Modelo del dispositivo"`
}

// UpdateEstadoRaspberryPiRequest Modelo para actualizar el estado de un Raspberry Pi
type UpdateEstadoRaspberryPiRequest struct {
	Estado string `json:"estado" example:"inactivo" description:"Nuevo estado del dispositivo" binding:"required,oneof=activo inactivo"`
}

// RaspberryPiResponse Modelo de respuesta para Raspberry Pi
type RaspberryPiResponse struct {
	ID                    int     `json:"id" example:"1" description:"ID del dispositivo"`
	MAC_Address           string  `json:"mac" example:"00:11:22:33:44:55" description:"Dirección MAC del dispositivo"`
	IP_Address            string  `json:"ip_address" example:"192.168.1.100" description:"Dirección IP del dispositivo"`
	Nombre                string  `json:"nombre" example:"Raspberry-001" description:"Nombre identificativo del dispositivo"`
	Modelo                string  `json:"modelo" example:"Raspberry Pi 4B" description:"Modelo del dispositivo"`
	Estado                string  `json:"estado" example:"activo" description:"Estado actual del dispositivo"`
	Fecha_Registro        string  `json:"fecha_registro" example:"2024-03-20 10:00:00" description:"Fecha de registro del dispositivo"`
	Fecha_Ultima_Conexion *string `json:"fecha_ultima_conexion,omitempty" example:"2024-03-20 10:00:00" description:"Última fecha de conexión"`
}

// RaspberryPiListResponse Modelo de respuesta para lista de Raspberry Pi
type RaspberryPiListResponse struct {
	RaspberryPis []RaspberryPiResponse `json:"raspberry_pis" description:"Lista de dispositivos Raspberry Pi"`
}



// =============================
// Estrcuturas para Sensor
// =============================

// CreateSensorRequest Modelo para crear un sensor
type CreateSensorRequest struct {
	ID_Tipo_Sensor  int `json:"id_tipo_sensor" example:"1" description:"ID del tipo de sensor"`
	ID_Raspberry    int `json:"id_raspberry" example:"1" description:"ID del dispositivo Raspberry Pi"`
	Pin_conexion    string `json:"pin_conexion" example:"1" description:"Pin de conexión del sensor"`
	Direccion_i2c   string `json:"direccion_i2c" example:"1" description:"Dirección I2C del sensor"`
	Estado          string `json:"estado" example:"activo" description:"Estado del sensor" binding:"required,oneof=activo inactivo"`
	Fecha_Registro  string `json:"fecha_registro" example:"2024-03-20 10:00:00" description:"Fecha de registro del sensor"`
}

// UpdateSensorRequest Modelo para actualizar un sensor
type UpdateSensorRequest struct {
	ID_Tipo_Sensor  int `json:"id_tipo_sensor" example:"1" description:"ID del tipo de sensor"`
	ID_Raspberry    int `json:"id_raspberry" example:"1" description:"ID del dispositivo Raspberry Pi"`
	Pin_conexion    string `json:"pin_conexion" example:"1" description:"Pin de conexión del sensor"`
	Direccion_i2c   string `json:"direccion_i2c" example:"1" description:"Dirección I2C del sensor"`
	Estado          string `json:"estado" example:"activo" description:"Estado del sensor" binding:"required,oneof=activo inactivo"`
}

// SensorResponse Modelo de respuesta para sensor
type SensorResponse struct {
	ID              int `json:"id" example:"1" description:"ID del sensor"`
	ID_Tipo_Sensor  int `json:"id_tipo_sensor" example:"1" description:"ID del tipo de sensor"`
	ID_Raspberry    int `json:"id_raspberry" example:"1" description:"ID del dispositivo Raspberry Pi"`
	Pin_conexion    string `json:"pin_conexion" example:"1" description:"Pin de conexión del sensor"`
	Direccion_i2c   string `json:"direccion_i2c" example:"1" description:"Dirección I2C del sensor"`
	Estado          string `json:"estado" example:"activo" description:"Estado del sensor"`
	Fecha_Registro  string `json:"fecha_registro" example:"2024-03-20 10:00:00" description:"Fecha de registro del sensor"`
}

type UpdateEstadoSensorRequest struct{
	Estado string `json:"estado" example:"inactivo" description:"Nuevo estado del dispositivo" binding:"required,oneof=activo inactivo"`
}


// =============================
// Estrcuturas para Tipo Sensor
// =============================
type CreateTipoSensorRequest struct{
	Nombre            string `json:"nombre" example:"DS18B20" description:"Nombre del tipo de sensor"`
	Descripcion       string `json:"descripcion" example:"Sensor de temperatura DS18B20" description:"Descripción del tipo de sensor"`
	Unidad_Medida     string `json:"unidad_medida" example:"°C" description:"Unidad de medida del tipo de sensor"`
	Tipo_Dato         string `json:"tipo_dato" example:"decimal" description:"Tipo de dato del tipo de sensor"`
}

type UpdateTipoSensorRequest struct{
	Nombre            string `json:"nombre" example:"DS18B20" description:"Nombre del tipo de sensor"`
	Descripcion       string `json:"descripcion" example:"Sensor de temperatura DS18B20" description:"Descripción del tipo de sensor"`
	Unidad_Medida     string `json:"unidad_medida" example:"°C" description:"Unidad de medida del tipo de sensor"`
	Tipo_Dato         string `json:"tipo_dato" example:"decimal" description:"Tipo de dato del tipo de sensor"`
}


type TipoSensorResponse struct{
	ID              int `json:"id" example:"1" description:"ID del tipo de sensor"`
	Nombre            string `json:"nombre" example:"DS18B20" description:"Nombre del tipo de sensor"`
	Descripcion       string `json:"descripcion" example:"Sensor de temperatura DS18B20" description:"Descripción del tipo de sensor"`
	Unidad_Medida     string `json:"unidad_medida" example:"°C" description:"Unidad de medida del tipo de sensor"`
	Tipo_Dato         string `json:"tipo_dato" example:"decimal" description:"Tipo de dato del tipo de sensor"`
}


	