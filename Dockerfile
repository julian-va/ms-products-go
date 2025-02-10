FROM  golang:1.23.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY src/ ./src/
RUN CGO_ENABLED=0 GOOS=linux go build -o /ms-products ./src

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY application.yml ./
COPY --from=builder ms-products .
EXPOSE 8080
CMD ["./ms-products"]