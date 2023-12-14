FROM golang:1.21

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y make

RUN go mod download

RUN make build/nats

EXPOSE 4222

ENTRYPOINT ["./bin/nats"]