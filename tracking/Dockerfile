FROM golang:1.18
RUN mkdir app
WORKDIR /app/
COPY schema/ /app/
COPY  cmd/ /app/
ADD . /app/
RUN go build -o main
EXPOSE 51372
CMD ["/app/main"]