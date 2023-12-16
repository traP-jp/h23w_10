FROM golang:1.21.5-alpine AS builder
WORKDIR /server
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -tags timetzdata -o /server/server ./cmd/server/main.go

FROM gcr.io/distroless/base-debian12
WORKDIR /server
COPY --from=builder /server/server .
COPY --from=builder /server/db/ ./db/
EXPOSE 8080
ENTRYPOINT ["./server"]