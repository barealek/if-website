FROM denoland/deno:alpine AS tw
WORKDIR /app
COPY . .
RUN deno i npm:tailwindcss
RUN deno run -A --node-modules-dir npm:tailwindcss@3.4.2 -o css/compiled.css

FROM php:8.2-apache
RUN a2enmod rewrite
RUN docker-php-ext-install mysqli pdo pdo_mysql

WORKDIR /var/www/html

COPY --from=tw /app .
