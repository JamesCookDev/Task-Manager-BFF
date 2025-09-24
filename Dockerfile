# Estágio de Build
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copia apenas os arquivos de módulo primeiro para aproveitar o cache do Docker.
# Se go.mod e go.sum não mudarem, o Docker reutilizará a camada de dependências.
COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]