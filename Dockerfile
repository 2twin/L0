FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN go build -o /l0 cmd/main.go

WORKDIR /app
COPY --from=builder /l0 /app/l0

ADD /templates /app/templates

CMD ["./l0"]