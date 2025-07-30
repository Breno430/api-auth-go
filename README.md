# API Auth Go

Uma API de autenticação desenvolvida em Go seguindo a arquitetura Clean Architecture.

## 🚀 Tecnologias

- **Go** - Linguagem de programação
- **Gin** - Framework web
- **GORM** - ORM para Go
- **PostgreSQL** - Banco de dados
- **JWT** - Autenticação

## 📁 Estrutura do Projeto

```
api-auth-go/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── domain/
│   │   ├── entities/
│   │   ├── repositories/
│   │   └── usecases/
│   ├── infrastructure/
│   │   ├── config/
│   │   ├── database/
│   │   ├── repositories/
│   │   └── server/
│   └── presentation/
│       ├── handlers/
│       └── routes/
├── go.mod
├── go.sum
├── .env.example
└── README.md
```

## 🏗️ Arquitetura

O projeto segue os princípios da Clean Architecture:

- **Domain Layer**: Entidades, repositórios e casos de uso
- **Infrastructure Layer**: Implementações concretas (banco de dados, servidor)
- **Presentation Layer**: Handlers e rotas da API

## 🛠️ Configuração

### Pré-requisitos

- Go 1.21+
- PostgreSQL
- Git

### Instalação

1. Clone o repositório:
```bash
git clone <url-do-repositorio>
cd api-auth-go
```

2. Instale as dependências:
```bash
go mod download
```

3. Configure as variáveis de ambiente:
```bash
cp .env.example .env
# Edite o arquivo .env com suas configurações
```

4. Execute o projeto:
```bash
go run cmd/api/main.go
```

## 📝 Variáveis de Ambiente

Copie o arquivo `.env.example` para `.env` e configure:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=api_auth
DB_SSL_MODE=disable
JWT_SECRET=seu_jwt_secret
SERVER_PORT=8080
```

## 🚀 Endpoints

### Health Check
- `GET /health` - Verificar status da API

### Usuários
- `POST /users` - Criar usuário
- `GET /users/:id` - Buscar usuário por ID
- `PUT /users/:id` - Atualizar usuário
- `DELETE /users/:id` - Deletar usuário

## 🧪 Testes

Execute os testes:
```bash
go test ./...
```

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes. 