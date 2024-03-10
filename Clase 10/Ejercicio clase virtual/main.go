package main

/*
Para utilizar Gin se requiere la versión 1.13+ de Go, una vez
instalada, utilizamos el siguiente comando para instalar Gin:
go get -u github.com/gin-gonic/gin

Luego lo importamos  a nuestro código:
import "github.com/gin-gonic/gin"
*/

// go mod init actividad/virtual

// go mod tidy
import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Es un slice de parámetros dados por la URL.
//Estos “Param” ocupan el mismo orden que en la URL.
type Param struct {
	Key   string
	Value string
}

//Definición de “Params”
type Params []Param

//Definimos nuestra estructura de información
type Empleado struct {
	// Una etiqueta de struct se cierra con caracteres de acento grave `
	Id     int    `form:"id" json:"id"`
	Nombre string `form:"name" json:"name"`
	Activo bool   `form:"active" json:"active" binding:"required"`
}

var empleados = LoadEmployees()

func main() {
	server := gin.Default()
	server.GET("/", PaginaPrincipal)
	server.GET("/employees/", GetEmployees)
	server.GET("/employees/:id", SearchEmployee)
	server.GET("/employeesparams", GetEmployeesParams)
	server.GET("/employeesactive", FilterEmployees)

	server.Run(":8080")

	//Ejecutar con: go run main.go
	//Para para el servidor hacer: Ctrl + C
}

//Este handler se encargará de responder a /.
func PaginaPrincipal(ctxt *gin.Context) {
	ctxt.String(200, "¡Bienvenido a la Empresa Gophers!")
}

//Este handler devolvera todos los empleados
func GetEmployees(ctxt *gin.Context) {
	ctxt.JSON(200, empleados)
}

//Este handler verificará si la id que pasa el cliente existe en nuestra base de datos.
func SearchEmployee(ctxt *gin.Context) {
	var empleadoReturn Empleado
	var encontrado bool
	idParam := ctxt.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctxt.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	for _, em := range empleados {
		if em.Id == id {
			empleado := Empleado{
				Id:     em.Id,
				Nombre: em.Nombre,
				Activo: em.Activo,
			}
			empleadoReturn = empleado
			encontrado = true
		}
	}
	if encontrado {
		ctxt.JSON(200, gin.H{
			"empleado": empleadoReturn,
		})
	}
	if !encontrado {
		ctxt.JSON(400, gin.H{
			"mensaje": "El empleado no ha sido encontrado.",
		})
	}

}

//Este handler devolvera un empleado tomado por parametro
func GetEmployeesParams(ctxt *gin.Context) {
	var empleado Empleado

	if ctxt.BindQuery(&empleado) == nil {
		ctxt.JSON(200, gin.H{
			"EmpleadoNuevo": empleado,
		})
	}
}

//Esta función solo mostrará aquellos empleados activos o inactivos, dependiente del parámetro active.
func FilterEmployees(ctxt *gin.Context) {
	var filtrados []Empleado

	active := ctxt.Query("active")
	parseActive, err := strconv.ParseBool(active)
	if err != nil {
		ctxt.JSON(400, gin.H{"error": "Invalid active."})
		return
	}

	for _, e := range empleados {
		if e.Activo == parseActive {
			filtrados = append(filtrados, e)
		}
	}
	ctxt.JSON(200, filtrados)
}

//  Devuelve una lista de empleados
func LoadEmployees() []Empleado {
	var listaDeEmpleados []Empleado

	empl1 := Empleado{
		Id:     1,
		Nombre: "Juan",
		Activo: false,
	}
	empl2 := Empleado{
		Id:     2,
		Nombre: "Maria",
		Activo: true,
	}
	empl3 := Empleado{
		Id:     3,
		Nombre: "Esteban",
		Activo: true,
	}
	empl4 := Empleado{
		Id:     4,
		Nombre: "Estefania",
		Activo: false,
	}

	listaDeEmpleados = append(listaDeEmpleados, empl1, empl2, empl3, empl4)

	return listaDeEmpleados
}
