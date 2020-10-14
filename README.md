# DB_DUFOURFAKHOURI_P01
Database interaction with basic SQL go module exercice 

## Requirements

If you use docker you will only need:
* Docker;
* Docker-Compose;

Refer to [Docker-Setup](#docker-setup) to install with docker.

To run this project, you will also need to install the following dependencies on your system:

- [go](https://golang.org/doc/install)

## How to launch the project
* To run the project
 `docker-compose up --build`

* Once the project is up, import the DB with:
 `docker exec -i db_dufourfakhouri_p01_db_1 sh -c 'exec mysql -ugobdd -p"gobdd" image_gobdd' < ./docker/data/database.sql`

