# Webhook-simple-service (WSS)

## Purpose:
Create two sample services and demonstrate data sharing by registering webhooks from one service to another. Based on event webhook should be invoked.

## Server and Collector
In this repo, there are two services written in Golang. 
- Server: A simple service that runs the server. It exposes the webhook to **record** and **show** the free memory recorded by collector. The wss-server always installed on the control node.
- Collector: A simple service to collect the free memory on the system and sends it to the server every 10 seconds.

## Instructions:
1. Ensure Helm has been installed. 
2. Update the **serverAddr** of [wss-collector](helm_chart/wss-collector/values.yaml#L13) to the ip of control node. 
3. Run helm command to install wss-server and wss-collector. 
NOTE: Ensure wss-server is running before installing wss-collector. 
   ```
   sudo helm install wss-server helm_chart/wss-server/
   sudo helm install wss-collector helm_chart/wss-collector/
   ```
4. You can login to _http://localhost:5000/memory/show_ (control-node) to view the free memory of the worker node where the wss-collector installed. 
   You can also run ```curl http://localhost:5000/memory/show``` on control node to check the collected data.

## Build Image:
1. Ensure docker has been installed. 
2. To build server: 
```sudo docker build -t wss-server -f Dockerfile_server .```
2. To build collector:
 ```sudo docker build -t wss-collector -f Dockerfile_collector .```