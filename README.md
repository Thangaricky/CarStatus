# Car Status Application v1.0

This project was generated with GoLang v1.15.2[https://blog.golang.org/go1.15]
Pre requisite for running this application 
 * Download and Install Go [https://golang.org/dl/]
 * Set the GOPATH to you project directory
 * Make sure to set the GOROOT to where GO is installed

## Development server

In the terminal go to /CarStatus/main
Run `go build core.go` to start the application.
In the browser navigate to `http://localhost:8231/car/overview` to list all the car details.

## Build

Run `docker build . -t go-carstatus` to build the project.
Run `docker images` to verify that image `go-carstatus` exists.
Use `docker run -p 8321:8321 -it go-carstatus` command to run the application

Access the application in the browser `http://localhost:8231/car/overview`
