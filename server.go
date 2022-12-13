package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type memory struct {
	timestamp string
	ip        string
	free      string
}

var memory_history []memory

func main() {
	fmt.Print("Start server.")
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("/memory/record", record_memory)
	e.GET("/memory/show", show_memory)

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}

// Handler
// Main page
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, this is a server with webhook!")
}

// register an employee
func record_memory(c echo.Context) error {
	timestamp := c.FormValue("timestamp")
	ip := c.FormValue("ip")
	free := c.FormValue("free")
	memory_history = append(memory_history, memory{timestamp: timestamp, ip: ip, free: free})
	return c.String(http.StatusOK, "Register complete.")
}

// Show all registered employee
func show_memory(c echo.Context) error {
	allMem := ""
	for i, _ := range memory_history {
		allMem = allMem + memory_history[i].timestamp + " " + memory_history[i].ip + " " + memory_history[i].free + "\n"
	}
	return c.String(http.StatusOK, allMem)
}
