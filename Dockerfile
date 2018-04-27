FROM golang:1.10

EXPOSE 8080
ENV PORT 8080

COPY . /go/src/github.com/bigbluebutton/bbb-api-meetings
WORKDIR /go/src/github.com/bigbluebutton/bbb-api-meetings

RUN go get -u github.com/beego/bee \
    && go get -u github.com/kardianos/govendor \
    && govendor add +external

CMD ["bee", "run"]
