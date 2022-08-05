FROM golang:1.18-alpine
WORKDIR /app

user root

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /delivery-system

CMD [ "/delivery-system" ]