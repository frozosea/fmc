FROM golang:latest
RUN mkdir /init
WORKDIR /app/
COPY schema/ /app/
COPY  cmd/ /app/
ADD . /app/
RUN go build -o main
EXPOSE 51839
EXPOSE 8090
CMD ["/app/main"]