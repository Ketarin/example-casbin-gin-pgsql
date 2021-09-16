# Example of use of Casbin with Gin and Pgsql
A simple implementation of [casbin/casbin](https://github.com/casbin/casbin) with a database 

## Database
[The database schema](docker/db/01-schema.sql) which contains the users, routes and roles tables.

[The database data](docker/db/01-schema.sql) : the fixtures for this example

The roles have access to the routes and we give these roles to a user. 

The roles can have multiple roles to inherit their rights.

The users can have multiple roles.

## Process
Casbin library execute the method **LoadPolicy** in the [Pgsql Casbin Adapter](app/mycasbin/adapter.go)

## Try yourself

### Requirements
- docker
- docker-compose

### Start the API and the database

```sh
docker-compose up
```

## Try some calls

GET 
```sh
#OK
curl http://localhost:8081/people -H "AUTH_IDENTITY:Sadmin"
curl http://localhost:8081/people -H "AUTH_IDENTITY:Liliread"
#NOK
curl http://localhost:8081/people -H "AUTH_IDENTITY:Moderator"
```

POST 
```sh
#OK
curl -X POST http://localhost:8081/people -H "AUTH_IDENTITY:Sadmin"
#NOK
curl -X POST http://localhost:8081/people -H "AUTH_IDENTITY:Liliread"
curl -X POST http://localhost:8081/people -H "AUTH_IDENTITY:Moderator"
```

PUT 
```sh
#OK
curl -X PUT http://localhost:8081/people/1 -H "AUTH_IDENTITY:Sadmin"
#NOK
curl -X PUT http://localhost:8081/people/1 -H "AUTH_IDENTITY:Liliread"
curl -X PUT http://localhost:8081/people/1 -H "AUTH_IDENTITY:Moderator"
```

DELETE 
```sh
#OK
curl -X DELETE http://localhost:8081/people/1 -H "AUTH_IDENTITY:Sadmin"
curl -X DELETE http://localhost:8081/people/1 -H "AUTH_IDENTITY:Moderator"
#NOK
curl -X DELETE http://localhost:8081/people/1 -H "AUTH_IDENTITY:Liliread"
```
