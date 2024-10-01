FROM docker.io/golang:1.23-alpine AS server

WORKDIR /server
COPY server .
RUN go build -o server

FROM node:20-alpine AS ui

WORKDIR /ui
COPY package.json .
COPY package-lock.json .
RUN npm install
COPY . .
RUN npm run build

FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=server /server/server server
COPY --from=ui /ui/dist dist
EXPOSE 3000

CMD ["./server"]
