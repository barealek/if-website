FROM denoland/deno:alpine AS tw
WORKDIR /app
COPY . .
RUN deno i npm:tailwindcss
RUN deno run -A --node-modules-dir npm:tailwindcss -o css/compiled.css

FROM nginx

COPY --from=tw /app /usr/share/nginx/html

ADD ./nginx/default.conf /etc/nginx/conf.d

CMD ["nginx", "-g", "daemon off;"]
