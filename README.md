# first.steps.go

#### Iniciando em Go

Esta é a página oficial [golang.org](golang.org), aqui encontraremos todas informações que poderíamos saber sobre Go.
Nosso start em Go inicia-se por este site, aqui temos os docs, packages, blog e nosso queridinho play.golang.org, specs da lingugem, download, tour em Go e muito mais.

Então primeiro passo é conhecer o SITE e como podemos usufruir ainda mais de tudo que está ai.
O site é o primeiro passo antes de avançarmos, ele é bem completo e muito grande para vermos tudo de uma única vez então o objetivo aqui é apresentar onde encontra o que precisa e aos poucos vai navegando e aprendendo.


#### Qual é o objetivo do projeto?

Na época do início do Go, apenas uma década atrás, o mundo da programação era diferente de hoje. O software de produção era geralmente escrito em C++ ou Java, o GitHub não existia, a maioria dos computadores ainda não eram multiprocessadores e, além do Visual Studio e do Eclipse, havia poucos IDEs ou outras ferramentas de alto nível disponíveis, muito menos gratuitamente na Internet.

Enquanto isso, ficamos frustrados com a complexidade indevida necessária para usar as linguagens com que trabalhamos para desenvolver software de servidor. Os computadores se tornaram enormemente mais rápidos desde que linguagens como C, C++ e Java foram desenvolvidas pela primeira vez, mas o próprio ato de programar não avançou tanto. Além disso, estava claro que os multiprocessadores estavam se tornando universais, mas a maioria das linguagens oferecia pouca ajuda para programá-los com eficiência e segurança.

Decidimos dar um passo para trás e pensar sobre os principais problemas que dominariam a engenharia de software nos próximos anos, à medida que a tecnologia se desenvolvesse, e como uma nova linguagem poderia ajudar a resolvê-los. Por exemplo, o surgimento de CPUs multicore argumentou que uma linguagem deveria fornecer suporte de primeira classe para algum tipo de simultaneidade ou paralelismo. E para tornar o gerenciamento de recursos tratável em um grande programa simultâneo, era necessária a coleta de lixo ou pelo menos algum tipo de gerenciamento de memória automático seguro.

Essas considerações levaram a uma série de discussões das quais Go surgiu, primeiro como um conjunto de ideias e desideratos, depois como uma linguagem. Um objetivo geral era que Go fizesse mais para ajudar o programador ativo, habilitando ferramentas, automatizando tarefas rotineiras, como formatação de código, e removendo obstáculos para trabalhar em grandes bases de código.

Um pouco da história aqui [10 anos escalando](https://commandcenter.blogspot.com/2017/09/go-ten-years-and-climbing.html)


#### História do mascote Gophers 

O Go gopher é um mascote icônico e uma das características mais distintivas do projeto Go. Neste post vamos falar sobre suas origens, evolução e comportamento.

[história do gopher](https://blog.golang.org/gopher)


#### Qual é a história do projeto?

Robert Griesemer, Rob Pike e Ken Thompson começaram a esboçar as metas para uma nova linguagem no quadro branco em 21 de setembro de 2007. Em poucos dias, as metas se estabeleceram em um plano para fazer algo e uma ideia justa do que seria. O design continuou a tempo parcial em paralelo com o trabalho não relacionado. Em janeiro de 2008, Ken começou a trabalhar em um compilador para explorar ideias; ele gerou código C como sua saída. Em meados do ano, a linguagem se tornou um projeto em tempo integral e se estabeleceu o suficiente para tentar um compilador de produção. Em maio de 2008, Ian Taylor começou de forma independente em um front-end GCC para Go usando as especificações preliminares. Russ Cox entrou no final de 2008 e ajudou a mover a linguagem e as bibliotecas do protótipo para a realidade.

Go tornou-se um projeto de código aberto público em 10 de novembro de 2009. Inúmeras pessoas da comunidade contribuíram com ideias, discussões e códigos.

Existem agora milhões de programadores Go - Gophers - em todo o mundo, e há mais a cada dia. O sucesso de Go excedeu em muito nossas expectativas.


#### Por que criaram uma nova linguagem?

Go nasceu da frustração com os idiomas e ambientes existentes para o trabalho que estávamos fazendo no Google. A programação havia se tornado muito difícil e a escolha das línguas era parcialmente culpada. Era preciso escolher entre compilação eficiente, execução eficiente ou facilidade de programação; todos os três não estavam disponíveis no mesmo idioma principal. Os programadores que poderiam escolher facilidade em vez de segurança e eficiência, mudando para linguagens tipadas dinamicamente, como Python e JavaScript, em vez de C ++ ou, em menor grau, Java.

Não estávamos sozinhos em nossas preocupações. Depois de muitos anos com um cenário bastante tranquilo para linguagens de programação, Go estava entre as primeiras de várias novas linguagens - Rust, Elixir, Swift e mais - que tornaram o desenvolvimento de linguagens de programação um campo ativo, quase dominante novamente.

Go tratou dessas questões tentando combinar a facilidade de programação de uma linguagem interpretada e digitada dinamicamente com a eficiência e segurança de uma linguagem compilada estaticamente. Também pretendia ser moderno, com suporte para computação em rede e multicore. Finalmente, trabalhar com Go pretende ser rápido : deve demorar no máximo alguns segundos para construir um grande executável em um único computador. Para atender a esses objetivos, é necessário abordar uma série de questões linguísticas: um sistema de tipos expressivo, mas leve; concorrência e coleta de lixo; especificação de dependência rígida; e assim por diante. Eles não podem ser bem tratados por bibliotecas ou ferramentas; uma nova linguagem foi necessária.

O artigo Go at Google discute o histórico e a motivação por trás do design da linguagem Go, além de fornecer mais detalhes sobre muitas das respostas apresentadas neste FAQ.


#### Doc Go

Neste link teremos exatamente o que precisaremos para iniciar  [doc golang](https://golang.org/doc)

_A linguagem de programação Go é um projeto de código aberto para tornar os programadores mais produtivos.
Go é expressivo, conciso, limpo e eficiente. Seus mecanismos de simultaneidade facilitam a escrita de programas que aproveitam ao máximo as máquinas multicore e em rede, enquanto seu novo tipo de sistema permite a construção flexível e modular de programas. Go compila rapidamente para código de máquina, mas tem a conveniência da coleta de lixo e o poder de reflexão em tempo de execução. É uma linguagem rápida, compilada estaticamente, que parece uma linguagem interpretada dinamicamente._

#### Dando Start em Go

- [Installing Go](https://golang.org/doc/install)

```bash
	$ sudo go get golang.org/dl/go1.16beta1
	$ sudo tar -C /usr/local -xzf /tmp/go1.15.5.linux-amd64.tar.gz
```

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


 - Programa para compilar uma imagem Go
 - Criação do hello World commo Api
 - Game
 - mostrar build , run 
 - build para gerar binarios estáticos
 - build para gerar binarios dinâmics
 - build usando gcc para C 
 - gerando imagemm docker
 - Executando imagem docker
 - Apresentando retrocompatibilidade
 - palestras -> https://docs.google.com/presentation/d/1rHGhzeptr7SA4P2F0wjh6_9VsvLgAAe_awVjSG3oUyo/edit#slide=id.p

 - Boteco Go Canal de Youtube fala sobre Go de forma humorada..
  https://www.youtube.com/channel/UCyfu7Q4ZjmW94iDnfTGvmBA


### Compilação CRUZADA

```bash
$ GOARCH=386 GOOS=windows go build -o myapp.exe main.go
```

```go
package main

import (
	"fmt"
)

func main() {
        fmt.Println("Sejam bem vindo!")
}
```
 

#### rEST api

```go
package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/ping",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("ok")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong"))
		})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
```