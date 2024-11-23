# Usar uma imagem base do Node.js
FROM node:16

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar o package.json e o package-lock.json para o diretório de trabalho
COPY package*.json ./

# Instalar as dependências do Node.js
RUN npm install

# Copiar o restante do código-fonte para o diretório de trabalho
COPY . .

# Expor a porta que o aplicativo vai rodar
EXPOSE 8080

# Comando para rodar o aplicativo Node.js
CMD ["node", "index.js"]
