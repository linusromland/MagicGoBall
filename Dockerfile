FROM golang:1.18

WORKDIR /app

COPY ./bot .

RUN go get

RUN go build -o magicgoball

CMD ["./magicgoball"]
