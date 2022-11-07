FROM golang
MAINTAINER User="i20072004@gmail.com"
RUN mkdir /init
WORKDIR /app/
COPY  cmd/ /app/
ADD . /app/
RUN go build -o main
EXPOSE 8080
CMD ["/app/main"]