[![codecov](https://codecov.io/gh/DanielFCubides/cerebro-meli/branch/main/graph/badge.svg?token=5DLN0WAQCL)](https://codecov.io/gh/DanielFCubides/cerebro-meli)
# Cerebro

This projects helps magneto to find mutants to defeat the X-MEN. 

## Getting Started

These instructions will give you a copy of the project up and running on
your local machine for development and testing purposes. See deployment
for notes on deploying the project on a live system.

### Prerequisites

Requirements for the software and other tools to build, test.
- [go](https://golang.org/)
- [mysql](https://www.docker.com/)
- [docker](https://www.mysql.com/)


### Installing

This is step by step series of examples that tell you how to get a development environment running. you will need the software mention in the Prerequisites.

First get clone this repository:

```sh
git clone https://github.com/DanielFCubides/cerebro-meli.git
```

Then run the mysql database
```shell
docker-compose up -d mysqldb
```
 Then run the application (the application will run in the port 80 by default if you need to set up other port change the docker-compose.yaml file):
```shell
docker-compose up -d app 
```
> Take in account that `app` needs that `mysqldb` is accepting connections, if the startup of the mysqldb is fast enough you could use:
:
```shell
docker-compose up -d
```
> This issue should be mitigated with the use of `depends-on` but it was deprecated for [docker-compose 3](https://docs.docker.com/compose/compose-file/compose-file-v3/#depends_on)


You will have the following endpoints:
```http request
GET /stats HTTP/1.1
Host: localhost:80
```

```http request
POST /mutant/ HTTP/1.1
Host: localhost:80
Content-Type: application/json

{
    "dna": [
        "AAAAGA",
        "CCGTGC",
        "TTCTGT",
        "AGAAAG",
        "CCTCTA",
        "TCGGGG"
    ]
}
```


## Running the tests

To run the test in the root folder of the project run
```shell
go test ./...
```

### code structure

We work with the following code structure inspired in clean architecture, using TDD and following SOLID principles making the code scalable and maintainable.
The project is divided in 5 folders:
- infrastructure: There are files that handle tools and details as the datasource, or the dependency injection tool.  
- domain: There are the core business objects of the system, for this project we only have the stats.
- usecase: These files are the core business logic of the system, where the mutant selector functions is, and the test that validate the expected behavior.
- repositories: These files handle the details of saving and retrieving objects from the database, we add an abstraction layer so we can have multiple implementation  of databases id needed (mysql, postgres, other web services, etc).
- adapters: there files handle how the systems get and send information to the outside world, we only have a rest implementation for now, and the logic of how the request is processed and how the response is build do not affect the core business logic.


### TODOs

1. Review bug with big entries.
2. Add some tool for logs and monitoring.
3. Add swagger documentation.
4. Kubernetes file deployment.
5. Add some diagrams to the documentations.
6. Add problem statement to repository.


## Authors

- **Daniel Fernando Cubides** - [DanielFCubides](https://github.com/DanielFCubides)
