FROM golang:1.26-alpine

WORKDIR /workspace

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o my-backend ./cmd/main.go

RUN chmod +x my-backend

ENTRYPOINT ["workspace/my-backend"]