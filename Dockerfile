FROM  golang:1.23.5-alpine
WORKDIR /app
COPY go.mod go.sum ./
COPY application.yml ./
RUN go mod download
COPY src/ ./src/
RUN CGO_ENABLED=0 GOOS=linux go build -o /ms-products ./src
EXPOSE 8080
CMD ["/ms-products"]