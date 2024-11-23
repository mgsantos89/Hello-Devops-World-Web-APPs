from flask import Flask
import socket
from datetime import datetime

app = Flask(__name__)

@app.route('/')
def hello():
    # Obtendo a data e hora atual
    current_time = datetime.now().strftime('%Y-%m-%d %H:%M:%S')

    # Obtendo o nome do host
    hostname = socket.gethostname()

    # Retornando a resposta
    return f"Hello DevOps World!!!\nCurrent Date and Time: {current_time}\nHostname: {hostname}"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
