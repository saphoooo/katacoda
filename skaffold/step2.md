## Un peu de refactoring

Si nous y regardons de plus prêt, notre Dockerfile est loin d'être parfait : bien que notre code y tourne comme un charme, l'image que nous avons produit ne fait pas moins de 805MB... Si le but d'un conteneur est d'être léger, ici nous avons clairement manqué le pas.

Autre point : Docker est très pratique pour prototiper, mais au risque de ne pas toujours avoir le même résultat :

- Lorsque nous ne précisons pas le tag d'une image, comme je l'ai mentionné, Docker utilise le tag `latest`. L'inconvénient du tag `latest` est qu'il risque de ne pas correspondre au même contenu d'un jour sur l'autre.

> Prenons un exemple : imaginez que vous recevez chaque jour le journal. Vous pouvez dire le journal du jour, lorsque vous le recevez, ou le journal du 15 février 2020.

> Maintenant imaginez vous parliez à un ami d'un article que vous avez lu dans ce journal, en lui disant : "il y a un super article à la une du journal du jour".

> Le lendemain votre ami décide de se rendre dans un kiosque, et demande le journal du jour pour y lire l'article. Malheureusement, il ne le trouve pas.

> Si au lieu de dire le journal du jour, vous l'aviez appelé le journal du 15 février 2020, votre ami était assurer de retrouver l'article.

> C'est la même chose pour le tag `latest` : il est pratique a utiliser, mais ne garanti pas d'obtenir toujours le même résultat.

Corrigeons cela.

## Multistage build

Afin d'avoir une image plus petite, nous avons la possibilité de créer notre binaire dans un conteneur (qui a tous les outils nécessaire), et d'envoyer ce resultat dans un autre conteneur qui ne contiendra que notre binaire.

Afin de garantir à notre image d'être toujours la même, nous allons prendre une version précise de golang. Puisque notre binaire est statically-linked, nous allons utiliser `scratch` pour la seconde image ; c'est un image particulière qui ne contient rien.

Modifions notre Dockerfile comme suit :

```
FROM golang:1.13.7-stretch as builder

COPY main.go .

RUN go get -u github.com/pkg/errors
RUN CGO_ENABLED=0 go build -o /app main.go

FROM scratch
CMD ["./app"]
COPY --from=builder /app .
```{{copy}}

Puis répétons d'étape du build :

`docker build -t loto .`{{execute}}

Un petit coup d'oeil à la taille de l'image nous apprend qu'elle ne fait plus que 2MB, ce qui représente la taille de notre binaire compiler.

Supprimons notre ancien conteneur :

`docker rm -f loto`{{execute}}

Nous sommes obligé d'utiliser l'argument `-f` pour forcer la suppression du conteneur, car il est encore en cours d'exécution.

`docker run -d --name loto loto`{{execute}}

`docker ps -f name=loto`{{execute}}

`docker logs loto`{{execute}}

Cette fois notre image est prête.