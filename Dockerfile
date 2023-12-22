# Adapted from https://stackoverflow.com/questions/68084178/react-client-and-golang-server-in-same-dockerfile

# Build the React application
FROM node:alpine as node_builder
WORKDIR /app
COPY frontend ./
RUN npm install
RUN npm run build

FROM golang:alpine AS go-build

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/cmd ./cmd
RUN go build -C /app/cmd/api -o /api

COPY --from=node_builder app/public ./public
CMD ["/api"]
