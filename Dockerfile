#build frontend
FROM node:22 AS frontend-build
WORKDIR /app
COPY frontend/react_app/ .
RUN npm install && npm run build

#build backend
FROM golang:1.24.5 AS backend-build
WORKDIR /app
COPY backend/ .
COPY backend/env/production.env ./.env
RUN go build -o server main.go

#final image
FROM alpine:latest
WORKDIR /app

COPY --from=backend-build /app/server .
COPY --from=frontend-build /app/dist ./frontend

#serve static FE + API
EXPOSE 8080
CMD ["./server"]
