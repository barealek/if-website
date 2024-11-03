FROM node:alpine
RUN npx tailwindcss - css/compiled.css

FROM nginx:alpine

COPY . /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]
