FROM golang:1.21 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .
COPY go.mod .

ENV CGO_ENABLED=0
RUN go mod download
RUN go get github.com/Warh40k/isrpo-rest
RUN go get github.com/Warh40k/isrpo-rest/go
RUN go build -a -installsuffix cgo -o openapi .

FROM scratch AS runtime
COPY --from=build /go/src/openapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./openapi"]
