FROM golang:1.20

WORKDIR /app

COPY . /app/

RUN go mod tidy

RUN bash build.sh

CMD ["/app/webApp" ]