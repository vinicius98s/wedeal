FROM golang:1.15-alpine

WORKDIR /go/src/github.com/vinicius98s/wedeal

COPY . .

RUN go get -u github.com/cosmtrek/air

CMD [ "/go/bin/air" ]