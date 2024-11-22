# Etapa 1: Builder - Instalação de dependências do Laravel
FROM php:8.1-cli AS builder

# Instalação de dependências do sistema
RUN apt-get update && apt-get install -y \
    unzip \
    git \
    libzip-dev \
    && docker-php-ext-install zip

# Instalar o Composer
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

# Configurar o diretório de trabalho
WORKDIR /app

# Copiar os arquivos do projeto
COPY . .

# Instalar as dependências do Laravel
RUN composer install --no-dev --optimize-autoloader

# Etapa 2: Runtime - Imagem de execução do Laravel
FROM php:8.1-apache

# Habilitar o módulo rewrite do Apache
RUN a2enmod rewrite

# Configurar o diretório de trabalho
WORKDIR /var/www/html

# Copiar os arquivos do projeto para o diretório público do Laravel
COPY --from=builder /app /var/www/html

# Copiar o arquivo .env
COPY .env /var/www/html/.env

# Ajustar permissões para o Apache
RUN chown -R www-data:www-data /var/www/html \
    && chmod -R 755 /var/www/html \
    && chmod -R 775 /var/www/html/storage /var/www/html/bootstrap/cache

# Gerar a chave de aplicação
RUN php artisan key:generate

# Configuração do VirtualHost do Apache para o Laravel
RUN echo '<VirtualHost *:80>\n\
    DocumentRoot /var/www/html/public\n\
    <Directory /var/www/html/public>\n\
        Options Indexes FollowSymLinks\n\
        AllowOverride All\n\
        Require all granted\n\
    </Directory>\n\
</VirtualHost>' > /etc/apache2/sites-available/000-default.conf

# Expor a porta 80
EXPOSE 80

# Comando padrão para iniciar o servidor
CMD ["apache2-foreground"]
