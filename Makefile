PWD=$(shell pwd)

image-dev:
	docker build -t bigbluebutton/bbb-api-meetings:dev .

run-dev:
	docker run --rm -ti -v ${PWD}:/go/src/github.com/bigbluebutton/bbb-api-meetings \
		-p 8080:8080 --name bbb-api-meetings bigbluebutton/bbb-api-meetings:dev
