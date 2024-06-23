FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod init product-service
RUN go build -o main .

EXPOSE 8082

CMD ["./main"]
