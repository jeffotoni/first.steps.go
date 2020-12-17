# first.steps.go

#### Iniciando em Go

Esta é a página oficial [golang.org](golang.org), aqui encontraremos todas informações que poderíamos saber sobre Go.
Nosso start em Go inicia-se por este site, aqui temos os docs, packages, blog e nosso queridinho play.golang.org, specs da lingugem, download, tour em Go e muito mais.

Então primeiro passo é conhecer o SITE e como podemos usufruir ainda mais de tudo que está ai.
O site é o primeiro passo antes de avançarmos, ele é bem completo e muito grande para vermos tudo de uma única vez então o objetivo aqui é apresentar onde encontra o que precisa e aos poucos vai navegando e aprendendo.


#### Doc Go

Neste link teremos exatamente o que precisaremos para iniciar  [doc golang](https://golang.org/doc)

_A linguagem de programação Go é um projeto de código aberto para tornar os programadores mais produtivos.
Go é expressivo, conciso, limpo e eficiente. Seus mecanismos de simultaneidade facilitam a escrita de programas que aproveitam ao máximo as máquinas multicore e em rede, enquanto seu novo tipo de sistema permite a construção flexível e modular de programas. Go compila rapidamente para código de máquina, mas tem a conveniência da coleta de lixo e o poder de reflexão em tempo de execução. É uma linguagem rápida, compilada estaticamente, que parece uma linguagem interpretada dinamicamente._

#### Dando Start em Go

- [Installing Go](https://golang.org/doc/install)

- [Tutorial Getting started](https://golang.org/doc/tutorial/getting-started.html)

- [Tutorial: Create a module](https://golang.org/doc/tutorial/create-module.html)


#### Aprendendo mais sobre Go

**Uma introdução interativa ao Go em três seções.**

   1) Apresenta a sintaxe básica e estruturas de dados

   2) Discute métodos e interfaces

   3) É apresentado os primitivos de simultaneidade do Go. 

   Cada seção termina com alguns exercícios para que você possa praticar o que aprendeu. Você pode fazer o tour online ou instalá-lo localmente com:

```bash
$ go get golang.org/x/tour
$ tour
```
Isto irá abrir no seu browser 127.0.0.1:3999 todo conteúdo do Tour... Lindo não é ? ❤️

- [A Tour of Go](https://tour.golang.org)



#### 3 Pilares

 - *Simplicidade*
 
 - *Legibilidade*
 
 - *Produtividade*


#### Goroutines e alguns patterns

Olha a fera ai, *Goroutine* é um dos pontos mais fortes da linguagem Go. Possibilitando a utilização do paradigma concorrente em sua a programação, ou seja podemos programar de forma concorrente na linguagem, neste quesito Go é poderoso pois trouxe um arsenal para possibilitar a programação concorrente de forma flexível e produtiva.

A concorrência integrada com base no documento CSP de Tony Hoare. Go foi projetado com a simultaneidade em mente e nos permite construir pipelines simultâneos complexos. Mas você já se perguntou omo são os vários padrões de simultaneidade?


### Patterns de concorrência

#### Fan-In

Um dos padrões populares no mundo concorrente é o chamado padrão fan-in. É o oposto do padrão de fan-out. Fan-in é uma função de leitura de várias entradas e multiplexação em um único canal.

#### Fan-Out

*Workers*

O padrão oposto ao *fan-in* é um padrão *fan-out* ou workers. Vários *goroutines* podem ler de um único canal, distribuindo uma quantidade de worker entre os núcleos da CPU, daí o nome dos workes. Em Go, esse padrão é fácil de implementar basta iniciar uma série de goroutines com canal como parâmetro e apenas enviar valores para esse canal a distribuição e a multiplexação serão feitas pelo tempo de execução Go, de forma automática❤️ 

#### Servers

Muito semelhante ao fan-out, mas com goroutines gerados por um curto período de tempo, apenas para realizar alguma tarefa. É normalmente usado para implementar servidores. É muito expressivo e permite implementar manipuladores de servidor da forma mais simples possível.

Não é muito interessante parece que não acontece nada em termos de simultaneidade. Claro, sob o capô há uma tonelada de complexidade, que é deliberadamente escondida. "Simplicidade é complicada".

Mas vamos voltar à simultaneidade e adicionar alguma interação ao nosso servidor. Digamos que cada manipulador deseja gravar de forma assíncrona no logger. O próprio Logger, é uma goroutine separada que faz o trabalho.