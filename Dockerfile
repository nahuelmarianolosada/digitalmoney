FROM golang:1.16-alpine as builder
RUN mkdir /build
WORKDIR /build
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /digitalmoney .

FROM alpine
RUN apk add libc6-compat
COPY --from=builder /build/digitalmoney /digitalmoney
COPY --from=builder /build/.env.example /.env
ENTRYPOINT ["./digitalmoney"]