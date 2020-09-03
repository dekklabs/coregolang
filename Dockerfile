FROM golang:1.14

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build
EXPOSE 5000
CMD ["./apirest"]