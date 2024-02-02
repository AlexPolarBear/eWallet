FROM golang:1.21.6

WORKDIR /app

COPY . .

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go get -d -v ./...

RUN cd cmd/app && go build -o /api

EXPOSE 8000

CMD ["/api"]
