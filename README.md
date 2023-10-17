# Zapfast.go
API de consumo para apps de transporte.

## Sobre

Zapfast.go é uma API backend projetada para fornecer funcionalidades essenciais para aplicativos de transporte. A API permite que os usuários registrem e gerenciem corridas, clientes e motoristas.

## Estrutura

- `/app`: Contém o código fonte da aplicação.
- `/app/pkg`: Contém os pacotes e módulos necessários.
- `/app/cmd/server`: Contém o código para iniciar o servidor.
- `/app/pkg/delivery/http`: Contém os handlers HTTP.
- `/app/pkg/repository`: Contém a lógica de acesso ao banco de dados.
- `/app/pkg/usecase`: Contém a lógica de negócios.

## Como Rodar

### Pré-requisitos

- Go 1.16 ou superior.
- PostgreSQL 12 ou superior.

### Configuração

1. Clone o repositório para o seu ambiente local:
    ```bash
    git clone https://github.com/jcr04/Zapfast.go.git
    cd Zapfast.go
    ```

2. Configure o banco de dados PostgreSQL e atualize o arquivo de configuração com as credenciais do banco de dados.

3. Instale as dependências:
    ```bash
    go mod download
    ```

### Executando

1. Navegue até o diretório `app/cmd/server`:
    ```bash
    cd app/cmd/server
    ```

2. Execute o servidor:
    ```bash
    go run main.go
    ```

Agora, o servidor deve estar rodando em `http://localhost:8080`. Você pode usar o Postman ou qualquer outro cliente HTTP para interagir com a API.

## Endpoints

Os endpoints incluem:
* - ![Screenshot_2](https://github.com/jcr04/Zapfast.go/assets/70778525/81bdf190-3344-48bb-b2ba-3d924f91b654)


- `POST /rides`: Solicitar uma nova corrida.
- `PATCH /rides/{rideID}/accept`: Aceitar uma corrida.
- `PATCH /rides/{rideID}/complete`: Completar uma corrida.
- `POST /customers`: Registrar um novo cliente.
- `GET /customers/{customerID}`: Obter informações de um cliente.
- `PUT /customers/{customerID}`: Atualizar informações de um cliente.
- `DELETE /customers/{customerID}`: Deletar um cliente.
- `POST /drivers`: Registrar um novo motorista.
- `GET /drivers/{driverID}`: Obter informações de um motorista.
- `PUT /drivers/{driverID}`: Atualizar informações de um motorista.
- `DELETE /drivers/{driverID}`: Deletar um motorista.

Para mais informações sobre os endpoints e como utilizá-los, veja a seção [Como usar o Postman para testar a API](#como-usar-o-postman-para-testar-a-api).

## Como usar o Postman para testar a API

1. Abra o Postman.
2. Importe a coleção de endpoints fornecida no repositório.
3. Selecione o endpoint que deseja testar.
4. Configure os parâmetros e o corpo da requisição conforme necessário.
5. Clique em "Send" para enviar a requisição e ver a resposta.
