# Desafio 6: Usando Banco de Dados
## Oque fazer
Usando esse guia: [criando api para postgress database](https://www.vultr.com/docs/how-to-create-a-golang-web-api-with-fiber-postgresql-and-gorm/#Conclusion) como material de apoio, adicione as seguintes funcionalidades ao código fornecido:
- nova função ```QueryByName```: Crie uma função que busque um livro no banco de dados pelo seu nome, retornando o livro em questão caso entrado e retornando um erro caso contrário.
- alterar função ```CreateBook```: Antes de criar armazenar um novo livro no banco de dados valide se já existe outro com esse mesmo nome, caso isso ocorra a função deve retornar um erro.

## Oque está feito
A parte principal da API descrita no [material de apoio](https://www.vultr.com/docs/how-to-create-a-golang-web-api-with-fiber-postgresql-and-gorm/#Conclusion) já está pronta, além disso, foi adicionado um container docker para rodar o banco de dados e uma página html simples para ser possível visualizar os livros inseridos. Os arquivos, estão organizados da seguinte forma:
- ```database.go```: Definição das funções utilizadas para gerenciar o banco de dados. Aqui é onde suas mudanças devem ser aplicadas.
- ```librarySystem.go```: Definição das funções para setup do banco de dados e início do programa.
- ```books.go```: Definição dos tipos e estruturas usadas no banco de dados, além da função ```MigrateBooks``` definir esses formatos no banco de dados.
- ```views/```: Essa é a pasta onde ficam os arquivos ```html```, não é esse o padrão de frontend que usaremos no projeto então não se preocupe em entender como ele funciona.
- ```docker-compose.yml```: Esse aquivo permite que você crie um banco de dados de maneira simples usando docker, os outros arquivos já estão configurados com base no banco de dados criado aqui.

## Como rodar o código
Segue o passo a passo:
- Importar o pakage ```library``` no seu arquivo ```main.go``` e chamar a função ```RunSystem```.
- Instale todas as dependencia do código, isso pode ser feito usando o comando ```go get [url da dependecia]``` para cada depencia ou simplesmente ```go get```para instalar todas de uma vez.
- Baixar a imagem do postgress para o docker usando o comando ```docker pull postgres``` e depois iniciar o container com o comando ```docker compose up db```.
- Rodar seu programa ```main.go```.