Development:

```
docker build -t bbb-api-meetings .
docker run -ti --rm -p 8080:8080 bbb-api-meetings
```

Test requests with:

```
curl -X POST -H "Content-Type: application/json" -d '{ "attributes": { "name": "-- POSTED NAME --", "roomId": "my-amazing-room" } }' -k "http://localhost:8080/meetings"

curl -X GET -H "Content-Type: application/json" -k "http://localhost:8080/meetings"
```
