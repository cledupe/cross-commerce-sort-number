# cross-commerce-sort-number

## Implementação de um ETL na lingagem Go para a api

### <http://challenge.dienekes.com.br/api/numbers>

arquivo | caminho | função
--- | --- | ---
rest_api.go| /rest_api.go | rodando go run rest_api.go subir o servidor para consulta dos numeros
extract.go | /framework/extract/extract.go | rodando go run extract.go rodara o script de aquição e ordenação dos numeros
utils.go |/domain/utils/utils.go| arquivo possuis algoritmos de ordenação na etapa load foi utilizado o SortMultiThread algoritmo inspirado no famoso Merge Sort porém com uma etapa onde a primeira iteraçao do algoritimo é distribuida e no final uma etapa onde agrupa os valores dos processo paralelos

## Implementação foi realizada usando a arquitetura hexagonal

### Arquitetura
Pasta | Descricao
--- | --- 
application| possui os use cases e as interfaces para conexão
domain | possui a estrutura e as regras de negocios como a inserção e a ordenação
framework | possui o acesso ao mundo externo como conexão com o serviço rest, estrutura de storage dos dados e a exposiçao dos dados

