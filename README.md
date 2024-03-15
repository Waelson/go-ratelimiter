### Sobre o projeto
Simples implementação de um Rate Limit por meio de um middleware

### Pré Requisitos
- Go 1.22.1
- Redis
- Docker

### Parametrização
Todos as configurações podem ser feitas no arquivo '.env'
```
REDIS_ADDRESS=localhost:6379
IP_RATE_LIMIT=5
TOKEN_RATE_LIMIT=100
IP_BLOCK_DURATION=5
TOKEN_BLOCK_DURATION=1
```
Detalhes do arquivo de configuração
- REDIS_ADDRESS: Endereço do servidor Redis
- IP_RATE_LIMIT: Quantidade de requisições permitidas por segundos SEM o uso de token.
- TOKEN_RATE_LIMIT: Quantidade de requisições permitidas por segundos COM o uso de token.
- IP_BLOCK_DURATION=Tempo de bloqueio em segundos após ultrapassar a quantidade máxima de requisições por segundo SEM utilizar o token
- TOKEN_BLOCK_DURATION: Tempo de bloqueio em segundos após ultrapassar a quantidade máxima de requisições por segundo utilizando o token

### Como executar?

```
docker-compose up --build
```

### Fazendo Request
Usando o endereço IP no browser
```
http://localhost:8080
```
Usando o endereço IP com cURL
```
curl --location http://localhost:8080
```
Usando API_KEY
```
curl --location 'http://localhost:8080/' --header 'API_KEY: 111'
```