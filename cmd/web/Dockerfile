FROM golang:alpine AS builder
WORKDIR /build
COPY . .
RUN go build -o web ./cmd/web

FROM alpine
LABEL maintainer="alex.s.kolesnikov@vk.com"
EXPOSE 8110
WORKDIR /app
COPY --from=builder /build/web ./web
COPY ./ui/static ./static
COPY ./ui/html ./html-templ
ENTRYPOINT [ "./web", "-html-templ", "./html-templ", "-static-dir", "./static"]