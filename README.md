# Go Events

Go Events é uma API para gerenciamento de eventos, desenvolvida em Go.

---

## **Pré-requisitos**

Certifique-se de que os seguintes softwares estejam instalados em sua máquina:

- **Docker**

---

## **Como Executar o Projeto**

### Clonar o Repositório

Clone o repositório do projeto para sua máquina local:

```bash
git clone https://github.com/marquescript/go-events.git
cd go-events
```

### Configurar o Arquivo `.env`

Crie e configure o arquivo `.env` com as variáveis de ambiente necessárias para a aplicação funcionar corretamente.

---

### Executar com Docker

Subir os serviços com Docker Compose:

```bash
docker-compose up --build
```

---

### Acessar a API

Após os serviços serem iniciados, a API estará disponível em:

[http://localhost:8000](http://localhost:8000)

---

### Acessar a Documentação Swagger

A documentação da API estará disponível em:

[http://localhost:8000/docs/index.html](http://localhost:8000/docs/index.html)
