## **Python Application**

### **Descrição**
Esta aplicação Python exibe "Hello DevOps World!!!" junto com a data, hora e o nome do host.

---

### **Como Rodar Localmente**

1. **Instale as dependências:**
   ```bash
   pip install flask
   ```
2. **Execute a aplicação:**
   ```bash
   python python_hello-devops-world.py
   ```
3. **Acesse no navegador:**
   ```text
   http://localhost:5000
   ```

---

### **Como Buildar a Imagem Docker**

1. **Crie a imagem Docker:**
   ```bash
   docker build -t <dockerhub-username>/python_hello-devops-world:v1 .
   ```
2. **Rode o contêiner:**
   ```bash
   docker run -d -p 5000:5000 <dockerhub-username>/python_hello-devops-world:v1
   ```

---

### **Como Deployar no Kubernetes**

1. **Aplique o manifesto:**
   ```bash
   kubectl apply -f python_hello-devops-world-deployment.yaml
   ```
2. **Verifique os pods:**
   ```bash
   kubectl get pods
   ```
3. **Acesse a aplicação via o IP do serviço:**
   ```bash
   kubectl get svc python-hello-devops-world-service
   ```

