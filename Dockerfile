FROM golang

WORKDIR /go/src/kafkaesque
COPY . .

RUN apt-get update
RUN apt-get install -y vim

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD ["go-wrapper", "run"] # ["kafkaesque"]