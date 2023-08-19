package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Id     int
	Nombre string
	Activo bool
}

var employees []Employee

func main() {

	//Crea un router con gin
	r := gin.Default()

	// Cargar empleados al iniciar la API
	loadEmployees()

	// Ruta de bienvenida
	r.GET("/", func(c *gin.Context) {
		c.String(200, "¡Bienvenido a la empresa Gophers!")
	})

	// Ruta para obtener todos los empleados en formato JSON
	r.GET("/employees", func(c *gin.Context) {
		c.JSON(200, employees)
	})

	// Ruta para obtener un empleado por su ID
	r.GET("/employees/:id", func(c *gin.Context) {
		// Extraer el valor del parámetro 'id' de la URL
		idStr := c.Param("id")
		// Convertir el valor del parámetro de cadena a un entero
		id, err := strconv.Atoi(idStr)
		if err != nil {
			// Si hubo un error al convertir a entero, enviar una respuesta con código 400
			c.String(400, "ID inválido")
			return
		}

		emp, found := findEmployeeById(id)
		if found {
			c.JSON(200, emp)
		} else {
			c.String(404, "Empleado no encontrado")
		}
	})

	// Ruta para crear un empleado a través de los parámetros
	r.GET("/employeesparams", func(c *gin.Context) {
		// Obtener los valores de los parámetros desde la URL
		idStr := c.Query("id")
		nombre := c.Query("nombre")
		activoStr := c.DefaultQuery("activo", "true")

		// Convertir el valor del parámetro 'id' de cadena a entero
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.String(400, "ID inválido")
			return
		}

		// Convertir el valor del parámetro 'activo' de cadena a booleano
		activo, err := strconv.ParseBool(activoStr)
		if err != nil {
			c.String(400, "Valor 'activo' inválido")
			return
		}

		// Crear una nueva instancia de Employee con los valores proporcionados
		newEmployee := Employee{
			Id:     id,
			Nombre: nombre,
			Activo: activo,
		}

		// Agregar el nuevo empleado al slice de empleados
		employees = append(employees, newEmployee)

		// Enviar una respuesta JSON con los detalles del nuevo empleado y código de respuesta 200
		c.JSON(200, newEmployee)
	})

	// Ruta para obtener empleados activos o inactivos
	r.GET("/employeesactive", func(c *gin.Context) {
		// Obtener el valor del parámetro 'activo' de la URL, con valor por defecto "true"
		activoStr := c.DefaultQuery("activo", "true")

		// Convertir el valor del parámetro 'activo' de cadena a booleano
		activo, err := strconv.ParseBool(activoStr)
		if err != nil {
			c.String(400, "Valor 'activo' inválido")
			return
		}

		// Crear un slice para almacenar los empleados filtrados
		filteredEmployees := []Employee{}

		// Recorrer todos los empleados en busca de los que coincidan con el valor de 'activo'
		for _, emp := range employees {
			if emp.Activo == activo {
				filteredEmployees = append(filteredEmployees, emp)
			}
		}

		c.JSON(200, filteredEmployees)
	})

	r.Run(":8080")
}

func loadEmployees() {
	// Cargar empleados de ejemplo
	employees = []Employee{
		{1, "Juan", true},
		{2, "María", true},
		{3, "Luis", false},
	}
}

func findEmployeeById(id int) (Employee, bool) {
	for _, emp := range employees {
		if emp.Id == id {
			return emp, true
		}
	}
	return Employee{}, false
}
