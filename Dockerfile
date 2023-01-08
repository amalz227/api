FROM golang:1.19

WORKDIR /app

COPY *.go ./ 

RUN go build main.go

CMD ["./main"]
