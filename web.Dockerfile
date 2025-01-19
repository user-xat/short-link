FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN go build -o web ./cmd/web

FROM alpine
LABEL maintainer="alex.s.kolesnikov@vk.com"
EXPOSE 8110
WORKDIR /app
COPY ./ui .
COPY --from=builder /build/web ./web
ENTRYPOINT [ "./web", "-html-templ-dir", "./html", "-static-dir", "./static"]