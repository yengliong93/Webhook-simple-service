# Webhook-simple-service

## Purpose:
Create two sample services and demonstrate data sharing by registering webhooks from one service to another. Based on event webhook should be invoked.

## Server and Client
In this repo, there are two services written in Golang. 
- Server: A simple service that run the server. It exposes the webhook to **register** and **show** the employee.
- Client: The user can use the client app to register employee's detail with **register** option and query all the employee data with **show** option.

## Instructions:
1. Ensure Golang installed on the system. To install Golang, see https://go.dev/doc/install.
2. Run ```go run server.go``` to start the server. By default, it starts a Go web framework that run a server on localhost:5000.
3. Run ```go run client.go``` to start the client. 
   a. Select option 2 to register the employee. Provide the name, age and gender. The client will send these information to the server. 
   b. Select option 1 to show all the registered employee. 
   c. Select option 3 to quit.