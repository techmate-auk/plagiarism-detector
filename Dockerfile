FROM golang:1.23-alpine

WORKDIR /

COPY . .

RUN go mod download

RUN go build -o app .
EXPOSE 8080
CMD [ "./app" ]