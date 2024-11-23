Aqui estão os modelos de README para cada aplicação. Cada README segue o mesmo formato para consistência.
## **Laravel Application**

### **Descrição**
Esta aplicação Laravel exibe "Hello DevOps World!!!" junto com a data, hora e o nome do host.

---

### **Como Rodar Localmente**

1. **Instale as dependências do Laravel:**
   ```bash
   composer install
   ```
2. **Configure o ambiente:**
   - Copie o arquivo `.env.example` para `.env`:
     ```bash
     cp .env.example .env
     ```
   - Gere a chave da aplicação:
     ```bash
     php artisan key:generate
     ```
3. **Rode o servidor local:**
   ```bash
   php artisan serve --host=0.0.0.0 --port=8000
   ```

4. **Acesse no navegador:**
   ```text
   http://localhost:8000
   ```

---

### **Como Buildar a Imagem Docker**

1. **Crie a imagem Docker:**
   ```bash
   docker build -t <dockerhub-username>/laravel_hello-devops-world:v1 .
   ```
2. **Rode o contêiner:**
   ```bash
   docker run -d -p 8000:8000 <dockerhub-username>/laravel_hello-devops-world:v1
   ```

---

### **Como Deployar no Kubernetes**

1. **Crie os Secrets e aplique o manifesto:**
   ```bash
   kubectl apply -f laravel_hello-devops-world-deployment.yaml
   ```
2. **Verifique os pods:**
   ```bash
   kubectl get pods
   ```
3. **Acesse a aplicação via o IP do serviço:**
   ```bash
   kubectl get svc laravel-hello-devops-world-service
   ```

