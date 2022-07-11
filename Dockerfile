FROM golang:1.18 AS builder
RUN apt-get update
ENV GO111MODULE=on \
    CGO=linux \
    GOARCH=amd64
WORKDIR /go/src/app
COPY go.mod .
RUN go mod download
COPY . .
RUN go install

FROM scratch
COPY --from=builder /go/bin/app .
ENTRYPOINT [ "./app" ]

# FROM golang:1.18 AS build

# WORKDIR /src/
# COPY main.go go.* /src/
# RUN CGO_ENABLED=0 go build -o /bin/demo

# FROM scratch
# COPY --from=build /bin/demo /bin/demo
# ENTRYPOINT ["/bin/demo"]