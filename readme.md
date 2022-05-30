# GOLIROU 🐹 🤝 🦍
## Projeto Final - Computação Paralela

### Componentes do Grupo

#### [João Pedro Souza Araújo](https://github.com/aslirou)
#### [Ricardo Carvalho Ribeiro](https://github.com/RicardoCR68)

### Objetivos:

Executar um projeto de otimização para um problema cujo paralelismo é conhecido, nesse caso, uma sequência de somas infinitas, e comparar a performance entre execuções com diferentes quantidades de núcleos.

### Ferramentas

- Linguagem: Go 1.18.2
  - <https://go.dev/>
- Ambiente de Desenvolvimento:
  - [Fedora 36](https://getfedora.org/pt_BR/)
  - [WSL2 - Ubuntu 20.04](https://ubuntu.com/wsl)

### Instruções

1. **Verificar se a linguagem está instalada:**

Rode o comando a seguir:

```
$ go version
```
O output esperado deve ser semelhante à esse: 

```
go version go1.18.1 linux/amd64
```

2. **Rodar o programa com o comando 'go run [nome do arquivo].go'**

Para esse projeto:

```
$ go run parallel.go
```

### Resultados

#### Resultado Serial para 100 000 000 000 (10^11):
- Tempo:2 min 4.512 seg
- Resultado: **25.90565168**653643 (com 08 casas decimais de precisão)

#### Resultado Paralelo para 100 000 000 000 (10^11):
- Tempo: 21.213 seg
- Resultado: **25.905651687841**704 (com 12 casas decimais de precisão)

#### Resultado Paralelo para 1 000 000 000 000 (10^12):
- Tempo: 3 min 35.421 seg
- Resultado: **28.20823678083**1157 (com 11 casas decimais de precisão)

#### Resultado Paralelo para 10 000 000 000 000 (10^13):
- Tempo: 35 min 53.346 seg
- Resultado: **30.510821873824**76 (com 12 casas decimais de precisão)

#### Resultado Paralelo para 100 000 000 000 000 (10^14):
- Tempo: 358 min 50.996 seg
- Resultado: **32.81340696681**9706 (com 11 casas decimais de precisão)