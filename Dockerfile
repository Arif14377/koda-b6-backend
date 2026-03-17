# Stage 1 : Build App
FROM golang:1.26-alpine AS build

WORKDIR /workspace

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

# RUN go build -o <nama-binary-file-bebas> <directory main.go>
# CGO_ENABLED=0 → static binary, tidak butuh C library
# -ldflags "-w -s" → hapus debug symbols, ukuran lebih kecil
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o my-backend ./cmd/main.go

RUN chmod +x my-backend

# Stage 2 : 
FROM alpine:latest

WORKDIR /app

# Ambil binary dari stage builder
COPY --from=build /workspace/my-backend .
# Ambil migration files (dipakai saat jalankan migrate)
COPY --from=build /workspace/migrations ./migrations

EXPOSE 8888
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 CMD wget -qO- http://localhost:8888/health || exit 1

ENTRYPOINT ["app/my-backend"]