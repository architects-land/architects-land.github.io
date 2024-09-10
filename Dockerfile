FROM node:22 AS builder

WORKDIR /app

COPY . .

RUN npm install -g bun

RUN bun i && bun run build

FROM golang:1.23

WORKDIR /app

COPY --from=builder /app .

RUN go build -o app .

EXPOSE 80

CMD ./app
