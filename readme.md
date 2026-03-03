# Gin Gonic API

API REST em Go com Gin Gonic e MongoDB, baseada na playlist **"Meu Primeiro CRUD em GoLang"** do canal HunCoding:
https://www.youtube.com/playlist?list=PLm-xZWCprwYQ3gyCxJ8TR1L2ZnUOPvOpr

Proxima aula planejada no estudo:
- Criando repository de busca do usuario - AULA #18
- https://www.youtube.com/watch?v=Dzmz9mS3U7g&list=PLm-xZWCprwYQ3gyCxJ8TR1L2ZnUOPvOpr&index=19

## Objetivo do projeto

Construir um CRUD de usuarios com foco em:
- organizacao em camadas (controller -> service -> repository)
- validacao de entrada
- persistencia em MongoDB
- padronizacao de erros e logs

## Status atual

Implementado:
- bootstrap da aplicacao com Gin
- conexao com MongoDB via variaveis de ambiente
- endpoint `POST /user` completo (controller, service, repository)
- validacao de payload com `go-playground/validator`
- resposta padrao de erro (`RestErr`)
- logging estruturado com Zap

Em andamento / pendente:
- `GET /userById/:id`
- `GET /userByEmail/:email`
- `PUT /user`
- `DELETE /user`
- JWT/autenticacao (dependencia existe no `go.mod`, fluxo ainda nao implementado)

## Tecnologias e bibliotecas

- Go `1.25.0`
- [Gin Gonic](https://github.com/gin-gonic/gin)
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)
- [Godotenv](https://github.com/joho/godotenv)
- [go-playground/validator](https://github.com/go-playground/validator)
- [Zap Logger](https://github.com/uber-go/zap)
- [JWT v5](https://github.com/golang-jwt/jwt) (instalado, ainda nao usado)

## Arquitetura e padroes

O projeto segue uma separacao por responsabilidades (estilo camadas, inspirada em DDD simples / Clean-ish):

1. `controller`
- recebe request HTTP
- faz bind e validacao inicial
- chama a camada de servico
- converte dominio para response

2. `models/services`
- contem regras de negocio
- orquestra operacoes (ex.: criptografar senha antes de persistir)
- depende de interfaces de repositorio

3. `models/repository`
- acesso a dados no MongoDB
- converte dominio <-> entidade de banco

4. `models` (dominio)
- entidade de dominio (`UserDomainInterface`)
- getters/setters e comportamento (`EncryptPassword`)

5. `view`
- monta DTO de resposta para API

6. `config`
- conexao com banco
- logger
- handler de erros
- validacao

### Fluxo atual do Create User

1. `POST /user` chega no controller
2. Gin faz `ShouldBindJSON` no `UserRequest`
3. erros de bind/validacao sao convertidos para `RestErr`
4. service chama `EncryptPassword()`
5. repository converte para entidade Mongo e executa `InsertOne`
6. retorno e convertido para dominio e depois para `UserResponse`
7. API responde `200 OK` com `id`, `name`, `email`, `age`

## Estrutura de pastas

```text
.
|-- main.go
|-- init_dependencies.go
|-- .env.example
|-- test.http
`-- src
    |-- config
    |   |-- database/mongodb
    |   |-- errorHandler
    |   |-- logger
    |   `-- validation
    |-- controller
    |   |-- model/request
    |   |-- model/response
    |   `-- routes
    |-- models
    |   |-- repository
    |   |   `-- entity/converter
    |   `-- services
    `-- view
```

## Variaveis de ambiente

Arquivo base: `.env.example`

Obrigatorias:
- `MONGODB_URI`
- `MONGODB_DATABASE_NAME`
- `MONGODB_COLLECTION_NAME`
- `APP_PORT`

Opcionais (logger):
- `LOG_OUTPUT` (`stdout` por padrao)
- `LOG_LEVEL` (`info` por padrao; opcoes: `debug`, `info`, `error`)

Exemplo:

```env
MONGODB_URI=mongodb+srv://<user>:<password>@<cluster>/<params>
MONGODB_DATABASE_NAME=gin_gonic_api
MONGODB_COLLECTION_NAME=users
APP_PORT=8787
LOG_OUTPUT=stdout
LOG_LEVEL=info
```

## Como rodar o projeto

1. Instalar dependencias:
```bash
go mod tidy
```

2. Criar `.env` a partir do `.env.example` e preencher valores.

3. Rodar em desenvolvimento:
```bash
go run main.go
```

Opcional com hot reload (Air):
```bash
air
```

## Endpoints

### 1) Criar usuario
- Metodo: `POST`
- Rota: `/user`
- Status: implementado

Body:

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "123456!",
  "age": 25
}
```

Regras de validacao:
- `name`: obrigatorio, min 3, max 50
- `email`: obrigatorio, formato de email
- `password`: obrigatorio, min 6, max 100, deve conter caractere especial de `!@#$%^&*()_+-=~`
- `age`: obrigatorio, min 1, max 200

Sucesso (`200`):

```json
{
  "id": "<mongo_object_id>",
  "name": "John Doe",
  "email": "john@example.com",
  "age": 25
}
```

### 2) Buscar por ID
- Metodo: `GET`
- Rota: `/userById/:id`
- Status: parcial (retorna apenas o `id` recebido na URL)

### 3) Buscar por email
- Metodo: `GET`
- Rota: `/userByEmail/:email`
- Status: pendente

### 4) Atualizar usuario
- Metodo: `PUT`
- Rota: `/user`
- Status: pendente

### 5) Deletar usuario
- Metodo: `DELETE`
- Rota: `/user`
- Status: pendente

## Formato padrao de erro

Quando ha erro de validacao ou erro interno, a API usa a estrutura:

```json
{
  "message": "Campos invalidos",
  "error": "bad_request",
  "code": 400,
  "causes": [
    {
      "field": "Name",
      "message": "Name is a required field"
    }
  ]
}
```

Tipos principais no projeto:
- `bad_request`
- `internal_server_error`
- `not_found`
- `forbidden`

## Logs

O logger usa Zap em JSON, com campos como:
- `timestamp`
- `level`
- `message`
- tags por contexto (`controller`, `service`, `repository`, etc.)

## Arquivos uteis para estudo

- `main.go`: bootstrap da API
- `init_dependencies.go`: injecao manual de dependencias
- `src/controller/createUser.go`: fluxo HTTP de criacao
- `src/models/services/create_user.go`: regra de negocio
- `src/models/repository/create_user_repository.go`: persistencia MongoDB
- `src/config/validation/validate_user.go`: tratamento de erros de bind/validacao
- `test.http`: colecao simples para testes manuais

## Observacoes tecnicas importantes

- A criptografia de senha atual usa MD5 (`EncryptPassword`). Para producao, o recomendado e `bcrypt` ou `argon2`.
- Existem inconsistencias de mapeamento entre `name` e `email` no dominio/conversores que devem ser corrigidas antes de evoluir o CRUD.
- O arquivo `test.http` tem algumas rotas que nao refletem o estado atual das rotas implementadas.

## Proximos passos sugeridos

1. Implementar `FindUserByEmail` no repository/service/controller.
2. Implementar `UpdateUser` e `DeleteUser` ponta a ponta.
3. Corrigir mapeamento de campos no dominio (`NewUserDomain`) e conversores.
4. Adicionar testes unitarios por camada.
5. Implementar autenticacao JWT e rotas protegidas.
