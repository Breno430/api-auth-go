# Branch: feature/password-reset-email

## 📋 Resumo das Funcionalidades Implementadas

Esta branch implementa a funcionalidade completa de recuperação de senha por email, seguindo as melhores práticas de segurança e arquitetura limpa.

## 🚀 Funcionalidades Adicionadas

### 1. **Entidade PasswordReset**
- **Arquivo:** `internal/domain/entities/password_reset.go`
- **Funcionalidades:**
  - Geração segura de PIN de 6 dígitos
  - Validação de expiração (15 minutos)
  - Validação de uso único
  - Validação de dados de entrada
  - Centralização de toda lógica de validação

### 2. **Repositório PasswordReset**
- **Arquivo:** `internal/domain/repositories/password_reset_repository.go`
- **Interface:** Define métodos para CRUD de tokens de reset
- **Arquivo:** `internal/infrastructure/repositories/password_reset_repository_impl.go`
- **Implementação:** GORM para PostgreSQL

### 3. **Serviços de Comunicação**
- **Arquivo:** `internal/infrastructure/services/email_service.go`
  - Envio de emails HTML para recuperação
  - Configurável via variáveis de ambiente
  - Suporte a SMTP (Gmail, Outlook, etc.)
- **Arquivo:** `internal/infrastructure/services/sms_service.go`
  - Simulação de envio de SMS
  - Preparado para integração com APIs reais

### 4. **Use Cases de Recuperação**
- **Arquivo:** `internal/domain/usecases/user_usecase.go`
- **Funcionalidades:**
  - `RequestPasswordReset`: Solicita recuperação por email
  - `ResetPassword`: Redefine senha com token
  - Validações de segurança antes de consultas no banco
  - Centralização de validações na entidade

### 5. **Handlers Consolidados**
- **Arquivo:** `internal/presentation/handlers/user_handler.go`
- **Métodos adicionados:**
  - `RequestPasswordReset`: Handler para solicitar recuperação
  - `ResetPassword`: Handler para redefinir senha
- **Consolidação:** Todos os handlers de usuário em um local

### 6. **Rotas de Recuperação**
- **Arquivo:** `internal/presentation/routes/routes.go`
- **Endpoints:**
  - `POST /api/v1/password-reset/request`: Solicitar recuperação
  - `POST /api/v1/password-reset/reset`: Redefinir senha

### 7. **Banco de Dados**
- **Arquivo:** `internal/infrastructure/database/database.go`
- **Migração:** Adicionada tabela `password_resets`
- **Campos:** ID, UserID, Token, Email, Used, ExpiresAt, timestamps

### 8. **Validações de Segurança**
- **Entidade User:** Validações centralizadas
- **Entidade PasswordReset:** Validações específicas
- **Use Cases:** Validações antes de consultas no banco
- **Segurança:** Prevenção de ataques de injeção

## 🔧 Arquivos Modificados

### **Entidades**
- `internal/domain/entities/user.go`
  - Removido campo `phone`
  - Adicionadas validações centralizadas
  - Função `ValidateResetPasswordInput`

### **Repositórios**
- `internal/domain/repositories/user_repository.go`
  - Adicionado método `Update`
- `internal/infrastructure/repositories/user_repository_impl.go`
  - Implementação do método `Update`

### **Use Cases**
- `internal/domain/usecases/user_usecase.go`
  - Adicionados use cases de recuperação
  - Validações de segurança
  - Integração com serviços de email

### **Infraestrutura**
- `internal/infrastructure/server/server.go`
  - Integração com repositório de password reset
  - Configuração de handlers consolidados

### **Documentação**
- `README.md`: Atualizado com novos endpoints
- `ENV_VARIABLES.md`: Adicionadas variáveis de email
- `.env.example`: Exemplo de configuração

## 🛡️ Melhorias de Segurança

### **Validações Implementadas**
1. **Token de Reset:**
   - Exatamente 6 dígitos
   - Apenas números (0-9)
   - Não vazio
   - Validação antes de consultar banco

2. **Email:**
   - Formato válido (regex)
   - Comprimento máximo (255 chars)
   - Não vazio
   - Sanitização de espaços

3. **Senha:**
   - Comprimento mínimo (6 chars)
   - Comprimento máximo (128 chars)
   - Não vazia
   - Sanitização de espaços

4. **Nome:**
   - Comprimento mínimo (2 chars)
   - Comprimento máximo (100 chars)
   - Não vazio
   - Sanitização de espaços

### **Prevenção de Ataques**
- Validação de entrada antes de consultas no banco
- Tokens expiram em 15 minutos
- Tokens de uso único
- Não revela se email existe ou não
- Prevenção de múltiplas solicitações

## 📊 Endpoints Disponíveis

### **Autenticação**
- `POST /api/v1/users/signup` - Cadastro de usuário
- `POST /api/v1/users/login` - Login
- `GET /api/v1/profile` - Perfil (autenticado)

### **Recuperação de Senha**
- `POST /api/v1/password-reset/request` - Solicitar recuperação
- `POST /api/v1/password-reset/reset` - Redefinir senha

### **Saúde**
- `GET /health` - Health check

## 🔧 Configuração Necessária

### **Variáveis de Ambiente**
```env
# Email Configuration
EMAIL_FROM=seu-email@gmail.com
EMAIL_PASSWORD=sua-senha-de-app
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
```

### **Para Gmail:**
1. Ativar autenticação de 2 fatores
2. Gerar senha de app em: https://myaccount.google.com/apppasswords
3. Usar a senha de app (16 caracteres)

## 🧪 Como Testar

### **1. Iniciar Aplicação**
```bash
make up
```

### **2. Testar Fluxo Completo**
```bash
# 1. Cadastrar usuário
curl -X POST http://localhost:8080/api/v1/users/signup \
  -H "Content-Type: application/json" \
  -d '{"name": "Teste", "email": "seu-email@gmail.com", "password": "123456"}'

# 2. Solicitar recuperação
curl -X POST http://localhost:8080/api/v1/password-reset/request \
  -H "Content-Type: application/json" \
  -d '{"email": "seu-email@gmail.com"}'

# 3. Verificar email recebido
# 4. Redefinir senha
curl -X POST http://localhost:8080/api/v1/password-reset/reset \
  -H "Content-Type: application/json" \
  -d '{"token": "123456", "password": "novaSenha123"}'
```

## 📈 Métricas de Implementação

- **Arquivos criados:** 5
- **Arquivos modificados:** 11
- **Linhas adicionadas:** 571
- **Linhas removidas:** 18
- **Funcionalidades:** 8 principais
- **Endpoints:** 2 novos
- **Validações:** 4 tipos
- **Serviços:** 2 (email + SMS)

## 🎯 Próximos Passos

1. **Testes:** Implementar testes unitários
2. **Integração SMS:** Conectar com provedor real
3. **Rate Limiting:** Implementar limitação de tentativas
4. **Logs:** Melhorar logging de segurança
5. **Monitoramento:** Adicionar métricas de uso

## ✅ Status da Branch

- ✅ **Funcionalidade completa** implementada
- ✅ **Validações de segurança** implementadas
- ✅ **Documentação** atualizada
- ✅ **Configuração** documentada
- ✅ **Testes manuais** funcionando
- ✅ **Código limpo** e bem estruturado

A branch está pronta para merge na `develop` após revisão de código. 