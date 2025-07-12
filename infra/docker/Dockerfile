# Stage 1: Build frontend
FROM node:22-alpine AS frontend-builder
WORKDIR /app
COPY apps/frontend/package*.json ./
RUN npm install
COPY apps/frontend/ ./
RUN npm run build

# Stage 2: Build backend
FROM golang:1.24-alpine AS backend-builder
WORKDIR /app
COPY apps/static-server/go.mod apps/static-server/go.sum ./
RUN go mod download
COPY apps/static-server/ ./
RUN go build -o /server/main .

# Stage 3: Final image
FROM alpine:latest
WORKDIR /app

RUN apk --no-cache add ca-certificates

RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

COPY --from=backend-builder /server/main .
COPY --from=frontend-builder /app/dist ./dist/
RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

CMD ["/app/main", "-appDir", "/app/dist"]
