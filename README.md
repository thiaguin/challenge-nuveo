# Nuveo

This is the code developed on nuveo backend test, to see the specification click [(here)](/API-N2.pdf).

The stack used in this project was the following:

- Golang: The language used
- RabbitMQ: The message service used
- Postgres: The database used
- Docker: To manager the images and containers

## Prerequisites

-   Docker
-   Docker Compose
## How To Run

Set the enviroment variable NOVOS_CLIENTES at the docker-compose.yml to be the path to upload the user .json file

Clone the gosidekick project [(here)](https://github.com/gosidekick/migration), build the ./migration executable file on the project root, and run:

``` bash
    # To load the images
    $ docker-compose build

    # To start the service (This may take a while)
    $ docker-compose up
```

If everything it's ok you'll have this:

- Postgres running on port: 5432
- RabbitMQ running on port: 5672
- The main backend service running on port: 8080
- The microservice running on port: 8000

## How To Test

To run the tests:

``` bash
    $ go test ./...
```