#build frontend
FROM node:22 AS frontend-build
WORKDIR /app
COPY frontend/react_app/ .
RUN npm install && npm run build

#build backend
FROM golang:1.24.5 AS backend-build
WORKDIR /app
COPY backend/ .
ENV PORT=8080
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

#final image
FROM alpine:latest
WORKDIR /app

COPY --from=backend-build /app/server .
COPY --from=frontend-build /app/dist ./frontend

RUN chmod +x server
#serve static FE + API
EXPOSE 8080
CMD ["./server"]
