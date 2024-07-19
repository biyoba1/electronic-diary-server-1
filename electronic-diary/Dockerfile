FROM golang:latest

RUN go version

WORKDIR /cmd

COPY . .

RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o electronic-diary ./cmd/main.go

CMD ["./electronic-diary"]