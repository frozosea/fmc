FROM golang:latest
RUN mkdir /app
WORKDIR /app/
COPY  cmd/ /app/
ADD . /app/
RUN go build -o main
EXPOSE 9005
CMD ["/app/main"]