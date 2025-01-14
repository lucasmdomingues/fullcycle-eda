FROM golang:1.23

WORKDIR /app/

RUN apt-get update && apt-get install -y librdkafka-dev
RUN go mod tidy

CMD ["tail", "-f", "/dev/null"]
