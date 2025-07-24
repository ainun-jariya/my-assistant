#build frontend
FROM node:22 AS frontend-build
WORKDIR /app
COPY frontend/react_app/ .
RUN npm install && npm run build

#build backend
FROM golang:1.22 AS backend-build
WORKDIR /app
COPY backend/ .
RUN go build -o server main.go

#final image
FROM alpine:latest
WORKDIR /app

COPY --from=backend-build /app/server .
COPY --from=frontend-build /app/dist ./frontend

#serve static FE + API
EXPOSE 8001
CMD ["./server"]
