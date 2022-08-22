FROM golang:1.12.0-alpine3.9
RUN mkdir /youtube-stats
ADD . /youtube-stats
WORKDIR /youtube-stats
RUN go mod download
RUN go build -o main .
CMD ["/youtube-stats/main"]