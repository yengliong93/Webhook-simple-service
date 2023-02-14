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
    e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {}))

	// Routes
	e.GET("/", hello)
	e.POST("/memory/record", record_memory)
	e.GET("/memory/show", show_memory)
	e.POST("/hvs/v2/reports", generate_report)
	e.POST("/node/node1/assertion/assertion1", tcaas_assertion)

	// Start server
	e.Logger.Fatal(e.Start(":8443"))
}

// Handler
// Main page
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, this is a server with webhook!")
}

// Dummy method for generate attestation report
func generate_report(c echo.Context) error {
    fmt.Print("\nReceive report generation request.\n")
	return c.String(http.StatusOK, "{\"result\": \"Report generation success.\"}")
	//return c.String(http.StatusBadRequest, "{\"result\": \"Report generation failed.\"}")
}

// Dummy method for generate attestation report
func tcaas_assertion(c echo.Context) error {
    fmt.Print("\nReceive TCaaS assertion request.\n")
    fmt.Print(c.Request().Body)
    fmt.Print("\n\n")
	return c.String(http.StatusOK, "Assertion is successfully stored.")
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
