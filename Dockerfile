# [ 1 ]
FROM docker.io/golang:1.23-alpine AS server

WORKDIR /server
COPY server .
RUN go build -ldflags "-s -w" -o server

# [ 2 ]
FROM node:20-alpine AS ui

WORKDIR /ui
COPY package.json .
COPY package-lock.json .
RUN npm install
COPY . .
RUN npm run build

# [ 3 ]
FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=server /server/server server
COPY --from=ui /ui/dist dist

EXPOSE 80

CMD ["./server"]
