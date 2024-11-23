
## **Go Application**

### **Descrição**
Esta aplicação Go exibe "Hello DevOps World!!!" junto com a data, hora e o nome do host.

---

### **Como Rodar Localmente**

1. **Compile o projeto:**
   ```bash
   go build -o go_hello-devops-world main.go
   ```
2. **Execute o binário:**
   ```bash
   ./go_hello-devops-world
   ```
3. **Acesse no navegador:**
   ```text
   http://localhost:8080
   ```

---

### **Como Buildar a Imagem Docker**

1. **Crie a imagem Docker:**
   ```bash
   docker build -t <dockerhub-username>/go_hello-devops-world:v1 .
   ```
2. **Rode o contêiner:**
   ```bash
   docker run -d -p 8080:8080 <dockerhub-username>/go_hello-devops-world:v1
   ```

---

### **Como Deployar no Kubernetes**

1. **Aplique o manifesto:**
   ```bash
   kubectl apply -f go_hello-devops-world-deployment.yaml
   ```
2. **Verifique os pods:**
   ```bash
   kubectl get pods
   ```
3. **Acesse a aplicação via o IP do serviço:**
   ```bash
   kubectl get svc go-hello-devops-world-service
   ```


