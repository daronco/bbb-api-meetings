## Development

Using the Makefile:

```
make image-dev
make run-dev
```

Or do it yourself:

```
docker build -t bigbluebutton/bbb-api-meetings .
docker run -ti --rm -p 8080:8080 bigbluebutton/bbb-api-meetings

# run in another port
docker run -ti --rm -p 9000:9000 -e PORT=9000 bigbluebutton/bbb-api-meetings
```

Test requests with:

```
curl -X POST -H "Content-Type: application/json" -d '{ "attributes": { "name": "-- POSTED NAME --", "roomId": "my-amazing-room" } }' -k "http://localhost:8080/meetings"

curl -X GET -H "Content-Type: application/json" -k "http://localhost:8080/meetings"
```
