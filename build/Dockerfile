# syntax=docker/dockerfile:1
FROM golang:1.18-rc as build

##
## Build
##
WORKDIR /app
COPY go.mod ./
# COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /docker-go-awesome

##
## Deploy
##
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /docker-go-awesome /docker-go-awesome
EXPOSE 38000
ENTRYPOINT ["/docker-go-awesome"]