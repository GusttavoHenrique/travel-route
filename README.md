# The travel-route application
Implementation of the minimum cost problem between an origin point and a destination

## Available operations
The following sections present the available REST and Shel operations

### Route register:
Creates a new route

```
POST http://localhost:8080/route
```

Request body example:
```JSON
{
	"origin": "Any string value. eg.: 'NAT'. This field is required",
	"destination": "Any string value. eg.: 'SAO'. This field is required",
	"price": "Any number value. eg.: 10. This field is required"
}
```

Response body:
```
No response body
```

Response code:
```
204 Created
```

### Find routes:
Retrieves all registered routes

```
GET http://localhost:8080/routes
```

Available query parameters:
```TEXT
    origin: the origin of route. Any string value. eg.: 'SAO'. This field is not mandatory.
    destination: the destination of route. Any string value. eg.: 'NAT'. This field is not mandatory.
    price: the price of route. Any number value. eg.: 10. This field is not mandatory.
```

Response body:
```JSON
[
    {
    	"origin": "Type string. eg.: 'SAO'. This field is mandatory.",
    	"destination": "Type string. eg.: 'NAT'. This field is mandatory.",
    	"price": "Type number. eg.: 10. This field is mandatory."
    }
]    
```

Response code:
```
200 Ok
```

### Find best route
Retrieves the best route from an informed origin to an informed destination

```
GET http://localhost:8080/best-route
```

Response body:
```JSON
{
    "origin": "Type string. eg.: 'SAO'. This field is mandatory.",
    "destination": "Type string. eg.: 'NAT'. This field is mandatory.",
    "price": "Type number. eg.: 10. This field is mandatory."
}
```

Response code:
```
200 Ok
```

### Find best route CLI
Retrieves the best calculated route

Shel input example:
```SHELL
$ please enter the route: GRU-CGD
```

Shel output example:
```SHELL
$ best route: GRU - BRC - SCL - ORL - CDG > $40
```

## Project structure
```MARKDOWN
.
├── http
│   └── `server.go` - implementation of http server
│   └── `server_test.go` - tests implementation of `server.go`
├── point
│   └── `point.go` - defines the point struct
│   └── `point_test.go` - tests implementation of `point.go`
├── resource
│   └── `input-routes.csv` - file with input routes to tests
├── route
│   └── `route.go` - defines the route struct
│   └── `route_test.go` - tests implementation of `route.go`
│   └── `controller.go` - defines the route controller handlers
│   └── `controller_test.go` - tests implementation of `controller.go`
│   └── `service.go` - defines the route procedures
│   └── `service_test.go` - tests implementation of `service.go`
│   └── `repository.go` - defines the route database operations
│   └── `repository_test.go` - tests implementation of `repository.go`
│   └── `db.go` - implementation of a singleton to simulate the database 
│   └── `db_test.go` - tests implementation of `db.go`
├── terminal
│   └── `handler.go` - defines the CLI handlers
│   └── `handler_test.go` - tests implementation of `handler.go`
│   └── `service.go` - defines the terminal procedures
│   └── `service_test.go` - tests implementation of `service.go`
├── `.gitignore` - the gitignore file
├── `application.go` - the main application file
├── `application_test.go` - tests implementation of application.go
├── `mysolution.sh` - a bash script to initialize the application
├── `Dockerfile` - a Docker file definition to application container
├── `go.mod` - with the application dependencies
├── `README.md` - contains this documentation
```

## Dependencies
* Linux
* Docker
* Golang

## Starting project
* Clone the project to host
* Open the CLI and navigate to the project's root directory
* Execute start command
```SHELL
$ ./mysolution.sh /<your_file_directory>/<your_input_file>.csv
```

Input file example
```CSV
GRU,BRC,10
BRC,SCL,5
GRU,CDG,75
GRU,SCL,20
GRU,ORL,56
ORL,CDG,5
SCL,ORL,20
```

After the command is being execute, a Docker container starts, and the application instance is ready. The input file is copied 
to the tmp directory (`/tmp/travel-route`) and the entire operation is performed using this copy file. When the application starts, 
the file copy is opened in gedit to display the new routes entered.

## Design of application
The initial flow of the application opens two new execution flows for the REST (`controller.go`) and CLI (`handler.go`) requests. 
These flows operate under the main entity, called the `Route`. The route consists of an origin point, a destination point, the price 
of the route and a list of routes that make up the best route.

The [dijkstra's](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm) algorithm was used to find the best route because:
 * all routes between two points have positive values as weights (prices);
 * low temporal complexity

The Golang programming language was used to obtain better performances searching the best route when there are many registered routes.

A Docker image has been configured to facilitate the application's execution making it unnecessary to install the dependencies on the host.

## Authors
* **Gusttavo Silva** - *Maintainer* - [gusttavohnssilva@gmail.com](mailto:gusttavohnssilva@gmail.com)