# pg-casbin
A simple implementation with


## Start the example

```
make complete
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
