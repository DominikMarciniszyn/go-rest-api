FROM golang:1.18 AS build

WORKDIR /go/src/go-rest-api

COPY . .

RUN go mod download
RUN go build -o app /go/src/go-rest-api/cmd/main.go

# FROM alpine:latest

# WORKDIR /app

# COPY --from=build /go/src/go-rest-api/app .

EXPOSE 3000

CMD [ "./app" ]
