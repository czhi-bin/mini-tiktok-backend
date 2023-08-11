FROM golang:1.20-alpine3.18

COPY biz/ /go/src/biz/
COPY pkg/ /go/src/pkg/
COPY go.mod go.sum *.go /go/src/

WORKDIR "/go/src/"
RUN go env -w GO111MODULE=on \
    && go env -w GOOS=linux \
    && go env -w GOARCH=amd64
RUN go mod tidy
RUN go build -o mini-tiktok-backend

RUN chmod +x /go/src/mini-tiktok-backend

EXPOSE 18000
ENTRYPOINT ["/go/src/mini-tiktok-backend"]