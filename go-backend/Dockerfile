FROM golang:1.23-alpine3.20 AS builder
WORKDIR /build
COPY . .
RUN go mod download && \
	go build -o ./binary

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /build/binary ./binary
CMD [ "/app/binary" ]
