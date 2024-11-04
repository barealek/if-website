FROM denoland/deno:alpine AS tw
WORKDIR /app
COPY . .
RUN deno i npm:tailwindcss
RUN deno run -A --node-modules-dir npm:tailwindcss -o css/compiled.css

FROM nginx:alpine

COPY --from=tw /app /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]
