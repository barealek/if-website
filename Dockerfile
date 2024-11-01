FROM nginx:alpine

COPY . .

CMD ["nginx", "-g", "daemon off;"]
