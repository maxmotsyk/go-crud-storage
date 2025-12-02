# -----------------------------
#   Stage 1 — Build Go binary
# -----------------------------
    FROM golang:1.23.5 AS builder
    WORKDIR /app
    
    # Копіюємо модулі окремо — прискорює збірку
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Копіюємо увесь код
    COPY . .
    
    # Збираємо сервер
    # Твій main знаходиться в cmd/api/main.go
    RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/api
    
    # -----------------------------
    #   Stage 2 — Final image
    # -----------------------------
    FROM alpine:3.19
    
    WORKDIR /app
    
    RUN apk add --no-cache ca-certificates
    
    # Копіюємо зібраний Go-бінарник
    COPY --from=builder /app/server .
    
    EXPOSE 8080
    
    CMD ["./server"]