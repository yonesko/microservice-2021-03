FROM golang:1.10 AS build
WORKDIR /go/src/github.com/yonesko/microservice-2021-03/kuber_dz_2/
COPY go-server ./go-server
COPY vendor ./vendor
COPY go.mod ./go.mod

RUN ls -ltr

ENV CGO_ENABLED=0

RUN go build -a -installsuffix cgo -o swagger github.com/yonesko/microservice-2021-03/kuber_dz_2/go-server

FROM scratch AS runtime
COPY --from=build /go/src/github.com/yonesko/microservice-2021-03/kuber_dz_2/swagger ./
EXPOSE 8080/tcp
ENTRYPOINT ["./swagger"]
