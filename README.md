# Car Status Application v1.0

This project is developed as a Coding test for the below problem and this is the backend service that would be connected to the frontend in  https://github.com/Thangaricky/connectVehicleStatus
# Scenario:
Imagine you are an Alten consultant and you got assigned to a project at one of our top partners.
They have a number of connected vehicles that belongs to a number of customers.
They have a need to be able to view the status of the connection among these vehicles on a monitoring display.

The vehicles send the status of the connection one time per minute.
The status can be compared with a ping (network trace); no request from the vehicle means no connection.

# Task:
Your task will be to create a data store that keeps these vehicles with their status and the customers who own them, as well as a GUI (preferably web-based) that displays the status.
How the GUI is designed is up to you, as well as how you chose to implement the features and use suitable technologies.

Obviously, for this task, there are no real vehicles available that can respond to your "ping" request.
This can either be solved by using static values or ??by creating a separate machinery that returns random fake status.

## To run this project
This project was generated with GoLang v1.15.2[https://blog.golang.org/go1.15]
Pre requisite for running this application 
 * Download and Install Go [https://golang.org/dl/]
 * Set the GOPATH to you project directory
 * Make sure to set the GOROOT to where GO is installed

## Development server

In your terminal go to "/CarStatus/main".
Run `go build core.go` to start the application.
In the browser navigate to `http://localhost:8231/car/overview` to list all the car details.

## Build

Run `docker build . -t go-carstatus` to build the project.
Run `docker images` to verify that image `go-carstatus` exists.
Use `docker run -p 8321:8321 -it go-carstatus` command to run the application

Access the application in the browser `http://localhost:8231/car/overview`
