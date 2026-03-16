FROM golang:1.26-alpine AS build

WORKDIR /workspace

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o my-backend ./cmd/main.go

RUN chmod +x my-backend


FROM alpine:latest

WORKDIR /app

COPY --from=build /workspace/my-backend /app

ENTRYPOINT ["app/my-backend"]