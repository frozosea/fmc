FROM golang:latest
MAINTAINER User="i20072004@gmail.com"
RUN mkdir /app
WORKDIR /app/
COPY  cmd/server/ /app/
ADD . go.mod /app/
RUN go build -o main
EXPOSE 8080
CMD ["/app/main"]