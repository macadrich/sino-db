FROM alpine:latest

RUN apk add --no-cache git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

WORKDIR /sino-db

COPY go.mod ./

COPY . .

RUN GOOS=linux go build -o sino-db .

EXPOSE 8083

CMD ["./sino-db"]