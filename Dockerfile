FROM golang:1.14.6-alpine3.12

WORKDIR /go/src/square

COPY /src/. .

ENV CGO_ENABLED 0
RUN GOOS=linux go build main.go

EXPOSE 80

CMD [ "./main" ]