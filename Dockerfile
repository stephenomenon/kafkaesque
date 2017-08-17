FROM golang:1.8

MAINTAINER Steve Wilson <ten.ten.steve@hotmail.com>

ENV APP kafkaesque

WORKDIR /go/src/$APP
COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD ["go-wrapper", "run"] # ["kafkaesque"]