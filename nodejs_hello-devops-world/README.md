## **README - Node.js Application**

### **Descrição**
Esta aplicação Node.js exibe "Hello DevOps World!!!" junto com a data, hora e o nome do host.

---

### **Como Rodar Localmente**

1. **Instale as dependências:**
   ```bash
   npm install
   ```
2. **Execute a aplicação:**
   ```bash
   node index.js
   ```
3. **Acesse no navegador:**
   ```text
   http://localhost:8080
   ```

---

### **Como Buildar a Imagem Docker**

1. **Crie a imagem Docker:**
   ```bash
   docker build -t <dockerhub-username>/nodejs_hello-devops-world:v1 .
   ```
2. **Rode o contêiner:**
   ```bash
   docker run -d -p 8080:8080 <dockerhub-username>/nodejs_hello-devops-world:v1
   ```

---

### **Como Deployar no Kubernetes**

1. **Aplique o manifesto:**
   ```bash
   kubectl apply -f nodejs_hello-devops-world-deployment.yaml
   ```
2. **Verifique os pods:**
   ```bash
   kubectl get pods
   ```
3. **Acesse a aplicação via o IP do serviço:**
   ```bash
   kubectl get svc nodejs-hello-devops-world-service
   ```

