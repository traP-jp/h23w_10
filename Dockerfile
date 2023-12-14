FROM golang:1.21.5-alpine
WORKDIR /server
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -tags timetzdata -o /server/server ./cmd/server/main.go
CMD [ "/server/server" ]