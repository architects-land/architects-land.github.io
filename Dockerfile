FROM node:22 AS builder

WORKDIR /app

COPY . .

RUN npm install -g bun

RUN bun i && bun run build

FROM golang:1.23-alpine

WORKDIR /app

COPY --from=builder /app .

RUN go mod tidy && go build -o app .

EXPOSE 80

CMD ./app
