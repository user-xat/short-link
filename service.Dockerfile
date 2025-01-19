FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN go build -o service ./cmd/service

FROM alpine
LABEL maintainer="alex.s.kolesnikov@vk.com"
EXPOSE 54321
WORKDIR /app
COPY --from=builder /build/service ./service
CMD [ "./service" ]