FROM golang:alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY ./configs/.env ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /crux-docker ./cmd/app

FROM alpine:3.21

COPY --from=builder /crux-docker /crux-docker 
COPY --from=builder /app/configs/.env .

CMD ["/crux-docker"]

