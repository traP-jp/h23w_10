FROM golang:1.21.5-alpine
WORKDIR /server
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go build -o /server/server ./cmd/server/main.go
CMD [ "/server/server" ]