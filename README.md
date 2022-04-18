# Hash Challenge

Baseado [nessa](https://github.com/hashlab/hiring/tree/master/challenges/pt-br/new-backend-challenge) proposta, desenvolvi essa solução para o proposto desafio.

## Testes unitários

Para executar os testes unitáros da aplicação, é necessário executar o seguinte comando:

```sh
make test
```

Isso irá mostrar os casos de teste, além de um resumo final com as porcentagens de testes de cada package.

## Variáveis de ambiente

```sh
> Obrigatórias
    START_DATE_BLACK_FRIDAY   #layout: DD/MM/YYYY HH24:MM:SS
    END_DATE_BLACK_FRIDAY   #layout: DD/MM/YYYY HH24:MM:SS

> Opcionais
    DISCOUNT_SERVICE_URL
    SERVER_PORT
```

## Subir toda a aplicação localmente

A aplicação será executada via docker-compose, para isso será preciso executar o seguinte comando:

```sh
make run-all-local
```

Esse comando irá subir tanto a aplicação como o serviço de desconto.

> **Nota:** Os serviços não serão executados em background, mas pode ser facilmente alterado no arquivo de [Makefile](./Makefile)  caso prefira

##

Para simular o serviço de desconto off, é possível parar seu serviço com:

```sh
make stop-discount-service
```

E se for necessário iniciá-lo novamente:

```sh
make start-discount-service
```

## Parar toda a aplicação

Há também um comando para parar toda a aplicação:

```sh
make down-all-local
```
