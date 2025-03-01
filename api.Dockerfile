FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN go build -o start ./cmd/api

FROM alpine
LABEL maintainer="alex.s.kolesnikov@vk.com"
EXPOSE 9090
WORKDIR /app
COPY --from=builder /build/start ./start
CMD [ "./start" ]