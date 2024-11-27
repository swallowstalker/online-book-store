# Online book store - take home test

Hi there! This application requires 
- [Docker](https://www.docker.com/), 
- [docker-compose](https://docs.docker.com/compose/)

Please install it first to make the build easier

## Build up dependencies

First, copy env.sample file into .env, this should prepare the environment variables.

This application requires Postgres 13.7 but don't worry, it has already been integrated into docker-compose. Just run this command to see its magic.
```bash
make build-run-dependency
```
to stop the dependency temporarily, run this command
```bash
make stop-dependency
```
to continue, run this command
```bash
make continue-run-dependency
```

## Database migrations

This database migration will use [golang migrate](https://hub.docker.com/r/migrate/migrate) docker image.

After the postgres is up and running, we need to run this command to apply database migrations
```bash
make migrate
```

in case you want to check the migration version
```bash
make migrate-version
```

in case you want to rollback the migration by 1 step
```bash
make rollback
```

## Compile and run applications

After all required dependencies are set-up, run this command to compile and run the application

```bash
make compile
make run
```

The application server is up and running on port 8080 by default. If you wish to change it, please change the env `APP_PORT` on .env file

## Postman to test the application endpoints

To ease up testing, I've been using [Postman](https://www.postman.com/downloads/) with exported collection located in [gotu.postman_collection.json](doc%2Fgotu.postman_collection.json). You could import that on Postman and test the endpoints there 

## Testing the application code

Please run this command if you wish to test the application code
```bash
make test
```