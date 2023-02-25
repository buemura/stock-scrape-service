FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build -o app
CMD ["./app"]
