FROM golang:1.14.4
WORKDIR /app
COPY main.go /app/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=0 /app/ .
ENTRYPOINT ["./app"]
CMD [""]