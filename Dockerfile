# ---------------------
# Etapa de build
# ---------------------
    FROM golang:1.22 AS builder

    WORKDIR /app
    
    # Copia arquivos necessários
    COPY go.mod ./
    COPY go.sum ./
    RUN go mod download
    
    COPY . .
    
    # Compila o binário estaticamente
    RUN CGO_ENABLED=0 GOOS=linux go build -o server ./main.go
    
    # ---------------------
    # Imagem final mínima
    # ---------------------
    FROM alpine:latest
    
    RUN apk --no-cache add ca-certificates
    
    WORKDIR /root/
    
    # Copia o binário da etapa de build
    COPY --from=builder /app/server .
    
    EXPOSE 8080
    
    CMD ["./server"]
    