# Organização de projetos

## Modelo padrão

Uma proposta de Organização de projeto bem aceita pela comunidade Go.

[Veja aqui mais informações](convencao#organização-de-projeto)

## Modelo 1 Clean Architecture

Este exemplo de modelo foi obtido em <https://github.com/amitshekhariitbhu/go-backend-clean-architecture> e documentado [aqui](https://amitshekhar.me/blog/go-backend-clean-architecture). O código baixado do projeto pode ser acessado [aqui](assets/organizacao-projeto-modelo1-clean-architecture.zip)

Os benefícios da implementação de arquitetura limpa são os seguintes:

- __Framework independent__: Mais fácil substituir um pacote por outro pacote, se necessário, uma vez que tudo é desacoplado. Por exemplo, podemos alterar o pacote de banco de dados que usamos ou adicionar outro se precisarmos.
- __Altamente Testável__: Mais fácil de escrever testes. Escrevi o teste para as camadas de caso de uso, repositório e controlador.
- __Fácil adição de recursos e alterações__: A adição de um novo recurso torna-se fácil assim como modificar o código para quaisquer alterações necessárias.

Como este projeto segue o princípio da arquitetura limpa, você pode substituí-los muito facilmente por pacotes que melhor se ajustem às suas necessidades. No entanto, os principais pacotes que usei são os seguintes:

- __gin__: Gin é um framework web HTTP escrito em Go (Golang). Ele possui uma API semelhante ao Martini com desempenho muito melhor - até 40 vezes mais rápido.
- __mongo go driver__: O driver oficial do Golang para o MongoDB.
- __jwt__: Os Web Tokens JSON são um método RFC 7519 aberto e padrão do setor para representar declarações com segurança entre duas partes. Usado para Token de Acesso e Token de Atualização.
- __viper__: Para carregar a configuração do arquivo. Encontre, carregue e desmarque um arquivo de configuração nos formatos de propriedades JSON, TOML, YAML, HCL, INI, envfile ou Java..env
- __bcrypt__: O pacote bcrypt implementa o algoritmo de hash adaptativo bcrypt de Provos e Mazières.
- __testify__: Um kit de ferramentas com afirmações e simulações comuns que funciona bem com a biblioteca padrão.
- __mockery__: Um autogerador de código simulado para Golang usado em testes.

Agora, vamos discutir todas as camadas usadas no projeto:

1. Primeiro de tudo, o pedido chega ao __roteador__.
    - Entre o roteador, um middleware é adicionado para verificar a validade do token de acesso.
1. Em seguida, o roteador chamará seu __controlador__ correspondente.
    - Primeiro, ele validará os dados presentes dentro da solicitação. Se algo for inválido, ele retornará um "400 Bad Request" como a resposta ao erro.
    - Se tudo for válido dentro da solicitação, ele chamará a camada de caso de uso para executar uma operação.
    - Dependências:
        - Caso de uso, pois o controlador depende do caso de uso.
        - Repositório, pois o caso de uso depende do repositório.
        - Domínio
1. A camada de __caso de uso__ usa a camada de repositório para executar uma operação. Cabe completamente ao repositório como ele vai executar uma operação.
    - Dependências:
        - Repositório.
        - Domínio.
1. A camada de __repositório__ é livre para escolher qualquer banco de dados, na verdade, ela pode chamar qualquer outro serviço independente com base no requisito.
    - Dependências:
        - Domínio.
1. Na camada de __dominio__ temos:
    - Modelos de solicitação e resposta.
    - Entidades para o banco de dados.
    - Interfaces para casos de uso e repositórios.
    - Domínio, modelo e entidade são usados no controlador, no caso de uso e no repositório.

Abaixo a estrutura de diretórios do projeto:

```txt
.
├── Dockerfile
├── api
│   ├── controller
│   │   ├── login_controller.go
│   │   ├── profile_controller.go
│   │   ├── profile_controller_test.go
│   │   ├── refresh_token_controller.go
│   │   ├── signup_controller.go
│   │   └── task_controller.go
│   ├── middleware
│   │   └── jwt_auth_middleware.go
│   └── route
│       ├── login_route.go
│       ├── profile_route.go
│       ├── refresh_token_route.go
│       ├── route.go
│       ├── signup_route.go
│       └── task_route.go
├── bootstrap
│   ├── app.go
│   ├── database.go
│   └── env.go
├── cmd
│   └── main.go
├── docker-compose.yaml
├── domain
│   ├── error_response.go
│   ├── jwt_custom.go
│   ├── login.go
│   ├── profile.go
│   ├── refresh_token.go
│   ├── signup.go
│   ├── success_response.go
│   ├── task.go
│   └── user.go
├── go.mod
├── go.sum
├── internal
│   └── tokenutil
│       └── tokenutil.go
├── mongo
│   └── mongo.go
├── repository
│   ├── task_repository.go
│   ├── user_repository.go
│   └── user_repository_test.go
└── usecase
    ├── login_usecase.go
    ├── profile_usecase.go
    ├── refresh_token_usecase.go
    ├── signup_usecase.go
    ├── task_usecase.go
    └── task_usecase_test.go
```

## Modelo 2 Clean Architecture

Este exemplo de modelo foi obtido em <https://github.com/bxcodec/go-clean-arch> e documentado [aqui](https://medium.com/@imantumorang/trying-clean-architecture-on-golang-2-44d615bf8fdf). O código baixado do projeto pode ser acessado [aqui](assets/organizacao-projeto-modelo2-clean-architecture.zip)

Neste modelo a estrutura é separada por agregação, onde cada agregação possui as camadas _Model_, _Repository_, _Usecase_ e _Delivery_. Abaixo a estrutura de diretórios da agregação _article_:

```txt
domain
├── mocks
│   ├── ArticleRepository.go
│   ├── AuthorRepository.go
│   └── ArticleUsecase.go
├── article.go
├── author.go
└── errors.go 

article
├── delivery
│   └── http
│       ├── article_handler.go
│       └── article_test.go
├── repository //Encapsulated Implementation of Repository Interface
│   └── mysql
│       ├── mysql_article.go
│       └── mysqlarticle_test.go
└── usecase //Encapsulated Implementation of Usecase Interface
    ├── articleucase_test.go
    └── artilce_ucase.go
```

## Modelo 3 Clean Architecture

Este exemplo de modelo foi obtido em <https://github.com/evrone/go-clean-template> e documentado [aqui](https://evrone.com/go-clean-template?utm_source=github&utm_campaign=go-clean-template). O código baixado do projeto pode ser acessado [aqui](assets/organizacao-projeto-modelo3-clean-architecture.zip)

Neste projeto é usado os diretórios Go `internal` e `pkg`

## Modelo 4 Clean Architecture

Este exemplo de modelo foi obtido em <https://github.com/manakuro/golang-clean-architecture> e documentado [aqui](https://manakuro.medium.com/clean-architecture-with-go-bce409427d31). O código baixado do projeto pode ser acessado [aqui](assets/organizacao-projeto-modelo4-clean-architecture.zip)

As camadas usadas no projeto:

- __Entidades__: Entidades é um modelo de domínio que tem regras de negócios corporativas amplas e pode ser um conjunto de estruturas e funções de dados. (ex: um tipo struct de usuário, livro e autor).
- __Casos de uso__: Os casos de uso contêm regras de negócios de aplicativo usando um modelo de domínio e têm Porta de Entrada e Porta de Saída.
    - A Porta de Entrada é responsável por lidar com dados da camada externa e definida como abstrata.
    - A Porta de Saída é responsável por manipular dados de casos de uso para a camada externa e definida como abstrata.
- __Adaptador de interface__: O adaptador de interface lida com a comunicação com a camada interna e externa. Tem apenas preocupações com a lógica tecnológica, não com a lógica empresarial.
    - __Controladores__ são um conjunto de implementações específicas da Porta de Entrada em Casos de Uso (ex: converta dados de um formulário antes de salvá-lo no banco de dados).
    - __Presenter__ é um conjunto de implementações específicas da Porta de Saída em Casos de Uso (ex: converter dados do banco de dados antes de passá-los para exibir).
- __Frameworks e Drivers__: Frameworks e drivers contêm ferramentas como bancos de dados, frameworks ou API e basicamente não têm muito código (ex: API, banco de dados e estrutura da Web)

Abaixo temos a estrutura do projeto:

```txt
.
├── domain
│   └── model
│       └── user.go
├── infrastructure
│   ├── datastore
│   │   └── db.go
│   └── router
│       └── router.go
├── interface
│   ├── controller
│   │   ├── app_controller.go
│   │   ├── context.go
│   │   └── user_controller.go
│   ├── presenter
│   │   └── user_presenter.go
│   └── repository
│       └── user_repository.go
├── main.go
├── registry*
│   ├── registry.go
│   └── user_registry.go
├── usecase
│   ├── presenter
│   │   └── user_presenter.go
│   ├── repository
│   │   └── user_repository.go
│   └── interactor
│       └── user_interactor.go

* Em registry acontecerá a resolução de dependências usando injeção via construtor.
```

## Referências

- <https://amitshekhar.me/blog/go-backend-clean-architecture>
- <https://manakuro.medium.com/clean-architecture-with-go-bce409427d31>
- <https://medium.com/@imantumorang/trying-clean-architecture-on-golang-2-44d615bf8fdf>
- <https://evrone.com/go-clean-template?utm_source=github&utm_campaign=go-clean-template>