package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type employee struct {
	name string
	age  string
	gender string
}

var employee_list []employee

func main() {
	fmt.Print("Start server.")
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("/register", register)
	e.GET("/show", show)

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}

// Handler
// Main page
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, this is a server with webhook!")
}

// register an employee
func register(c echo.Context) error {
	name := c.FormValue("name")
	age := c.FormValue("age")
	gender := c.FormValue("gender")
	employee_list = append(employee_list, employee{name: name, age: age, gender: gender})
	return c.String(http.StatusOK, "Register complete.")
}

// Show all registered employee
func show(c echo.Context) error {
	ey := ""
	for i, _ := range employee_list{
		ey = ey + employee_list[i].name + " " + employee_list[i].age + " " + employee_list[i].gender + "\n"
	}
	return c.String(http.StatusOK, ey)
}

