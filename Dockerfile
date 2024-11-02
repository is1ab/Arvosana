FROM node:20-alpine3.20 AS frontend-builder
WORKDIR /app

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

COPY ./web/package.json ./web/pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

COPY ./web ./
RUN pnpm build

FROM golang:1.23.0-alpine3.20 AS backend-builder
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

COPY --from=frontend-builder /app/build ./web/build
COPY --from=backend-builder /app/arvosana .

ENTRYPOINT [ "/app/arvosana" ]
