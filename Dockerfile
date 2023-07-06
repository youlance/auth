## Build
FROM golang:1.19-alpine AS build

RUN apk add --no-cache git ca-certificates

ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -installsuffix cgo -ldflags '-s -w' -o main .

## Deploy
FROM scratch as final

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /app .

ENTRYPOINT ["./main"]