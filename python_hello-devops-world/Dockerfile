# Usando uma imagem base do Python
FROM python:3.9-slim

# Definindo o diretório de trabalho dentro do container
WORKDIR /app

# Copiando os arquivos do projeto para dentro do contêiner
COPY . /app

# Instalando as dependências
RUN pip install --no-cache-dir -r requirements.txt

# Expondo a porta 8080
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["python", "python_hello-devops-world.py"]
