FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git make
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Bangkok

RUN adduser -D -g '' appuser

WORKDIR /

COPY --from=builder /bin/task-management-backend /

USER appuser

EXPOSE 3000

CMD ["/task-management-backend", "start"]
