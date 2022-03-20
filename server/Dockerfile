FROM golang:1.17

ARG PORT=8080
ENV PORT ${PORT}

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]

# To test
# docker build --build-arg PORT=${PORT_VALUE} -t test .
# docker run -p ${PORT_VALUE:${PORT_VALUE} -it test
