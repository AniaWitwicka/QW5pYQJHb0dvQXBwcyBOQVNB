FROM golang:latest

RUN mkdir /code

COPY . /url-collector
WORKDIR /url-collector/cmd

RUN go build .
#TODO: add flags to personalize
CMD [ "./cmd", "-PORT=6060"]
