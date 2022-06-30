FROM golang:1.18-alpine AS build

WORKDIR /go/src/go-rest-api

COPY go.mod /go/src/go-rest-api
COPY go.sum /go/src/go-rest-api
COPY .env /go/src/go-rest-api

RUN go mod download

COPY . /go/src/go-rest-api/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/go-rest-api cmd/main.go

FROM alpine

WORKDIR /usr/bin/go-rest-api

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=build /go/src/go-rest-api/build/go-rest-api /usr/bin/go-rest-api/main
COPY --from=build /go/src/go-rest-api/.env /usr/bin/go-rest-api/

EXPOSE 3000 3000

ENTRYPOINT [ "/usr/bin/go-rest-api/main" ]
