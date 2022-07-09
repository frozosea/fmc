FROM golang:latest
MAINTAINER User="i20072004@gmail.com"
RUN mkdir /app
WORKDIR /app/
COPY schema/ /app/
COPY  cmd/ /app/
ADD . go.mod .env /app/
RUN go build -o main
EXPOSE ${GRPC_PORT}
CMD ["/app/main"]