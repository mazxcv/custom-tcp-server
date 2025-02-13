FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . ./

RUN go build -o /price-fetcher

EXPOSE 3000

CMD [ "/price-fetcher" ]