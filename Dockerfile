FROM golang

EXPOSE 8080

COPY . /go/src/github.com/bigbluebutton/bbb-api-meetings
WORKDIR /go/src/github.com/bigbluebutton/bbb-api-meetings

RUN go get github.com/beego/bee
RUN go get -u github.com/kardianos/govendor
RUN govendor add +external

CMD ["bee", "run"]
