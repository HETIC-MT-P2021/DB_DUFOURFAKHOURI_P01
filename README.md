# DB_DUFOURFAKHOURI_P01
Database interaction with basic SQL go module exercice 

## How to launch the project
RUN docker-compose up --build

Once the project is up, import the DB with:

RUN docker exec -i db_dufourfakhouri_p01_db_1 sh -c 'exec mysql -ugobdd -p"gobdd" image_gobdd' < ./docker/data/database.sql

