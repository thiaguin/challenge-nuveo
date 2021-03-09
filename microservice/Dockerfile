FROM golang:1.16.0-alpine3.13

ENV GOPATH /go

WORKDIR /microservice

COPY . .

RUN go mod download
RUN go build -o main .

EXPOSE 8000

CMD [ "./main" ]
