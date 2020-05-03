# JOKE MAKER

A dead simple api for generating jokes.


## HOW TO RUN SERVER

just simply run from terminal under project root
```
make run-rest-api
```
it will start web-server on port 5000


or you can run cli

```
make run
```


## HOW TO TEST

```
curl -v 'http://localhost:5000/v1/joke'
```

Successful JSON Response should look like:

```
HTTP/1.1 200 OK

{"code":200,"message":"Ivan Odden solved the Travelling Salesman problem in O(1) time. Here's the pseudo-code: Break salesman into N pieces. Kick each piece to a different city."}
```

on error:
```
HTTP/1.1 500 Internal Server Error

{"code":500,"error":"Results in random joke response is empty"}
```

## IN ADDITION
API health check
```
curl -v 'http://localhost:5000/ping'
```

Successful JSON Response should look like:
```
HTTP/1.1 200 OK

{"message":"pong"}
```