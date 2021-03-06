FROM golang:1.15.4-alpine3.12
RUN mkdir /app

ADD . /app
WORKDIR /app

RUN go mod vendor
RUN go mod download

RUN go build -o main .

CMD ["/app/main"]

EXPOSE 8000