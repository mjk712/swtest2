FROM golang:1.20-alpine


RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o swtest2 ./cmd/app/main.go

CMD [ "./swtest2" ]