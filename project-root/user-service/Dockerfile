FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod init user-service
RUN go build -o main .

EXPOSE 8081

CMD ["./main"]
