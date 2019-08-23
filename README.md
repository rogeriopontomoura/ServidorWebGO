# Servidor Web em Go com CI no Cloud Build

O servidor escuta a porta 8000 e responde com uma mensagem ao ser acessado.

Ao fazer um pull para o github o Cloud Build faz o teste da aplicação.


## Executar no docker

1. Clonar repositório

> git clone git@github.com:rogeriopontomoura/ServidorWebGO.git

2. Fazer build da imagem

> docker build -t webserver .

3. Rodar o container

> docker run --publish 8000:8000 -t webserver

4. Abrir um novo terminal e executar o comando:

> curl localhost:8000

Deve retornar

<b>Code.education Rocks!</b>

# Endereço da Imagem no Docker Hub

[Acesse aqui](https://hub.docker.com/r/rogeriosims/web-server-go)