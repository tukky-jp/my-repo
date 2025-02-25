FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.* ./
RUN go version
RUN go mod tidy
ARG pkg="github.com/tukky-jp/my-repo/cmd"
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server ${pkg}

FROM alpine:3 

WORKDIR /cloudrun
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server .

CMD ["/cloudrun/server"]
