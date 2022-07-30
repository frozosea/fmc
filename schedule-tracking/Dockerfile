FROM golang:latest
RUN mkdir /app
WORKDIR /app/
COPY schema/ /app/
COPY  cmd/ /app/
ADD . /app/
ENV TZ="Asia/Vladivostok"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN go build -o main
EXPOSE 8005
CMD ["/app/main"]