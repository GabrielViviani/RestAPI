API Explicada

Primeiro, foi baixado o Framework que seria utilizado no projeto, no caso, o Gin e Gorm.

Foi setado o servidor "hello world" dentro da função main() no arquivo main.go
Dentro dessa função, foi inicializado o roteador Gin dentro da variável r.
Foi utilizado o roteador padrão (Default).

Para definir a rota GET, é necessário que especifiquemos o ponto final e o handler (uma função que pegue algo e passe a frente).

Para passar a resposta em JSON, podemos usar o método JSON que requer um código de status HTTP e uma resposta em JSON como parâmetros.

Configurando a base de dados

Agora vamos construir nossos modelos para a base de dados.
No caso, "Model" vai ser uma struct que nos permite comunicar com uma tabela especifica na base de dados.
Sendo assim, o modelo do struct Book ficou da seguinte maneira:

type Book struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Title  string `json:"title"`
  Author string `json:"author"`
}

A notação ``