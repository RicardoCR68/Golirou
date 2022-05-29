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

#### Resultado Serial para 100 000 000 000:
- Tempo:2 min 4.512 seg
- Resultado: **25.90565168**653643(com 10 casas de precisão)

#### Resultado Paralelo para 100 000 000 000:
- Tempo: 21.213 seg
- Resultado: **25.905651687841**704(com 14 casas de precisão)


#### Resultado Paralelo para 1 000 000 000 000:
- Tempo: 3 min 31.125 seg
- Resultado: (com 13 casas de precisão)

#### Resultado Paralelo para 100 000 000 000 000
- Tempo: 361 min 18.262 seg
- Resultado: 32.81336897130239