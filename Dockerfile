FROM oven/bun:1-slim AS builder

WORKDIR /app

COPY . .

RUN bun i && bun run build

FROM golang:1.23-alpine

WORKDIR /app

COPY --from=builder /app .

RUN go mod tidy && go build -o app .

EXPOSE 80

CMD ./app
