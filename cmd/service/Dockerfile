FROM golang:alpine AS builder
# RUN apk add bash
WORKDIR /build
COPY . .
# CMD [ "bash" ]
RUN go build -o service ./cmd/service/

FROM alpine
LABEL maintainer="alex.s.kolesnikov@vk.com"
EXPOSE 50555
WORKDIR /app
COPY --from=builder /build/service ./service
CMD [ "./service" ]