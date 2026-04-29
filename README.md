# CRUD de Usuários em Go

Este projeto é uma API REST simples para gerenciamento de usuários, desenvolvida em Go (Golang).  
O objetivo principal deste repositório é **estudo e prática de conceitos de desenvolvimento backend**, como criação de APIs, organização de código e boas práticas.

> ⚠️ Este é um projeto de estudo, não sendo recomendado para uso em produção.

---

## 🚀 Funcionalidades

- Criar usuário
- Listar todos os usuários
- Buscar usuário por ID
- Atualizar usuário
- Deletar usuário

---

## 🛠️ Tecnologias utilizadas

- Go (Golang)
- chi v5.2.5
- JSON para comunicação
- Uuid v1.6.0
- slog
- net/http
---

## 📁 Estrutura do projeto

```
/project-root
│── api/api.go
│── database/db.go
│── go.mod
│── go.sum
│── main.go
```

---

## ⚙️ Configuração do ambiente

### 1. Clonar o repositório

```bash
git clone https://github.com/Mateus-R-De-Lima/CRUDUSUARIOS.git
cd CRUDUSUARIOS
```

### 2. Instalar dependências

```bash
go mod tidy
```

---

### 3. Rodar o projeto

```bash
go run .
```

A aplicação estará disponível em:

```
http://localhost:8080
```

---

## 📌 Endpoints

### 🔹 Criar usuário
```
POST /users
```

### 🔹 Listar usuários
```
GET /users
```

### 🔹 Buscar por ID
```
GET /users/{id}
```

### 🔹 Atualizar usuário
```
PUT /users/{id}
```

### 🔹 Deletar usuário
```
DELETE /users/{id}
```

---

## 📚 Objetivo do projeto

Este projeto foi desenvolvido com foco em aprendizado, com os seguintes objetivos:

- Praticar conceitos de API REST
- Entender melhor a estrutura de projetos em Go
- Aplicar boas práticas de organização de código
- Evoluir conhecimentos em backend

---

## 🔧 Melhorias futuras

- Integração com banco de dados (PostgreSQL, MySQL, etc)
- Implementação de autenticação (JWT)
- Validações mais robustas
- Testes automatizados
- Dockerização

---

## 👨‍💻 Autor

Desenvolvido por Mateus Lima  
GitHub: https://github.com/Mateus-R-De-Lima

---

## 📄 Licença

Este projeto está sob a licença MIT.
