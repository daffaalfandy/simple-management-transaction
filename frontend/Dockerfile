FROM php:8.2-fpm-alpine

WORKDIR /var/www/html

RUN apk add --no-cache bash shadow curl git unzip \
    && docker-php-ext-install pdo pdo_mysql

COPY --from=composer:latest /usr/bin/composer /usr/local/bin/composer

COPY . .

RUN composer install --no-dev --optimize-autoloader

CMD ["php-fpm"]
