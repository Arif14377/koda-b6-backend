# Stage 1 : Build App
FROM golang:1.24-alpine AS build

WORKDIR /workspace

COPY go.mod go.sum ./

RUN go mod download

COPY . .

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

ENTRYPOINT ["./my-backend"]