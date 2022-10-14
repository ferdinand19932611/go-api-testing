FROM golang:1.19

WORKDIR /usr/src/app
COPY . .

RUN CGO_ENABLED=0 go get .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

EXPOSE 8080
CMD ["./main"]
