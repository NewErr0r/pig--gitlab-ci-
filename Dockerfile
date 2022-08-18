FROM golang:1.16

RUN mkdir /app
WORKDIR /app

COPY ./go.mod /app
COPY ./go.sum /app

RUN go mod download

COPY ./ /app

RUN go build 

EXPOSE 8000

CMD /app/pig
