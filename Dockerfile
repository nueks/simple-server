FROM golang AS builder

WORKDIR /app
COPY * ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o simple-server

FROM scratch
COPY --from=builder /app/simple-server /
ENV HTTP_PORT 80
EXPOSE 80
CMD ["/simple-server"]
