FROM golang:1.23.0-alpine3.20 AS builder
WORKDIR /app

RUN apk add --no-cache \
    gcc \
    musl-dev

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o arvosana ./cmd

FROM alpine:3.20 AS runner
WORKDIR /app

RUN apk add -U tzdata
ENV TZ=Asia/Taipei
COPY --from=builder /app/arvosana .

ENTRYPOINT [ "/app/arvosana" ]
