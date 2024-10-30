# Sistema de Stress test
Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.
O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

## Entrada de Parâmetros via CLI:

- --url (-u): URL do serviço a ser testado.
- --requests (-r): Número total de requests.
- --concurrency (-c): Número de chamadas simultâneas.

## Execução do Teste:

- Realizar requests HTTP para a URL especificada.
- Distribuir os requests de acordo com o nível de concorrência definido.
- Garantir que o número total de requests seja cumprido.
- Geração de Relatório:

## Apresentar um relatório ao final dos testes contendo:
- Tempo total gasto na execução
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Executar Localmente 
```bash
go run main.go --url <url> --requests <requests> --concurrency <concurrency>
```
ou
```bash
go run main.go -u <url> -r <requests> -c <concurrency>
```

## Executar com Docker
### Criar a imagem com o Dockerfile
```bash
docker build -t gostress .
```
### Executar o teste
```bash
docker run stress_test_go -u https://example.com -r 1000 -c 100
```
### Exemplo de saída do teste
```bash
Initializing test, please wait...
Load test complete:
- Total time: 2.37949489s
- Total requests: 1000
- Successful requests: 1000
- Other status codes:
```
