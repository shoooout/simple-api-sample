FROM golang:1.20

WORKDIR /app
COPY . /app

RUN go mod tidy
RUN go install github.com/cosmtrek/air@v1.27.3

CMD ["air"]
