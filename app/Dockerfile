FROM golang:1.16-alpine

ENV PORT 8080
ENV HOST 0.0.0.0

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./

RUN go build -o /server

EXPOSE 8080/tcp

CMD [ "/server" ]
