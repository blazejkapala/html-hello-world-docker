# Etap budowania (build stage)
FROM golang:1.21-alpine AS builder

# Zainstaluj git dla go mod download
RUN apk add --no-cache git

# Ustaw katalog roboczy
WORKDIR /app

# Skopiuj pliki go.mod i go.sum (jeśli istnieją)
COPY go.mod ./

# Pobierz zależności
RUN go mod download

# Skopiuj kod źródłowy
COPY . .

# Zbuduj aplikację
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Etap produkcyjny (production stage)
FROM alpine:latest

# Zainstaluj ca-certificates i wget dla health checks
RUN apk --no-cache add ca-certificates wget

# Utwórz użytkownika bez uprawnień root
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /root/

# Skopiuj skompilowaną aplikację z etapu budowania
COPY --from=builder /app/main .

# Skopiuj pliki statyczne
COPY --from=builder /app/index.html .

# Zmień właściciela plików na appuser
RUN chown -R appuser:appgroup /root/

# Przełącz na użytkownika bez uprawnień root
USER appuser

# Otwórz port 8080
EXPOSE 8080

# Dodaj health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Uruchom aplikację
CMD ["./main"] 