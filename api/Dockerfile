# Build stage
FROM golang:1.16-alpine AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app .

# Actual image
FROM alpine:3
LABEL maintainer="gabriel.gazola@poli.ufrj.br"
WORKDIR /app
COPY --from=build /app/app .
ENTRYPOINT [ "./app" ]