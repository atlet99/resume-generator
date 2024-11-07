FROM golang:1.23.3-alpine AS builder

RUN apk --no-cache add build-base git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o resume-generator

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/resume-generator .

RUN chmod +x resume-generator

EXPOSE 8080

CMD ["./resume-generator"]