FROM golang:1.16.0-alpine3.13

ENV GOPATH /go

WORKDIR /backend

COPY .  .

RUN go mod download
RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]
