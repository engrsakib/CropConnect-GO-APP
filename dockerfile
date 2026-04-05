# # Stage 1: Build
# FROM golang:1.25.5-alpine AS builder
# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .
# RUN go build -o server ./cmd/server

# # Stage 2: Final Image
# FROM alpine:latest
# WORKDIR /app

# COPY --from=builder /app/server .

# EXPOSE 8080
# CMD ["./server"]

# Stage 1: Build
FROM golang:1.22-alpine AS builder 

WORKDIR /app

# ক্যাশ অপ্টিমাইজেশন
COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN go clean -modcache
RUN go build -o server ./cmd/server

# Stage 2: Final Image
FROM alpine:latest
RUN apk --no-cache add ca-certificates 

WORKDIR /app

COPY --from=builder /app/server .

# COPY --from=builder /app/.env .
# COPY --from=builder /app/docs ./docs

EXPOSE 8080

# রেন্ডারে পোর্ট এনভায়রনমেন্ট ভেরিয়েবল সাপোর্ট নিশ্চিত করা
CMD ["./server"]