FROM golang:1.16
COPY server.go server.go
RUN go build server.go
EXPOSE 8000

ENTRYPOINT ["./server"]