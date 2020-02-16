## Le code

Le code de l'application est dans *home*:

`ls ~`{{execute}}

C'est une simple application en Go qui génére des grilles de loto.

`kubectl` et `minikube` sont déjà téléchargés, et `minikube` est déjà démarré (peut-être):

`minikube status`{{execute}}

Si `host`, `kubelet`, et `apiserver` ne sont pas encore `running`, il vous faudra faire preuve d'encore un peu de partience ; mais ne vous préoccuper pas de ça pour l'instant.

## Dockerfile

Comment, à patir de notre code, allons-nous créer une image OCI ?

Tout d'abord, par convention mettons notre code dans un répertoire portant le nom de l'application que nous développons ; ici, c'est une application pour générer des grilles du loto, aussi nous allons l'appeler loto :

`mkdir loto`{{execute}}

`cp main.go loto`{{execute}}

`cd loto`{{execute}}

Pour créer une image OCI, il existe à ce jour une grande variété de méthodes : `docker build`, `bazel`, `kaniko`, `buildha`, `buildpack` ; pour la plupart des ces méthodes, il faut créer un Dockerfile : un fichier qui décrit les étapes d'installation de notre application en quelque sorte.

Essayons de créer notre image de manière empirique :

`nano Dockerfile`{{execute}}

```
FROM golang

ADD main.go .

RUN go get -u github.com/pkg/errors

RUN GCO_ENABLE=0 go build -o /loto

CMD ["/loto"]
```{{copy}}

Quelques précisions sur la syntaxe du Dockerfile :

- Un Dockerfile commence toujours par l'argument **FROM**, il permet de définir l'image de base que nous souhaitons utiliser, dans notre cas `golang`.

- **ADD** permet de copier notre fichier source dans l'image.

- **RUN** nous sert à exécuter une commande.

- Finalement **CMD** est la commande que nous utilisons lorsque nous démarrons le conteneur.

### Astuce

> Les arguments de **CMD** peuvent être écris de deux manières :

> `CMD /loto` dans ce cas, la commande sera exécutée par le shell : `/bin/sh -c /loto`

> `CMD ["/loto"]`, le binaire sera exécuté directement.

### Attention

> Vérifiez que le copié-collé dans `nano` a bien pris les retours à la ligne (ou n'en a pas ajouté là où nous n'en voulons pas).

> Nano est un éditeur de texte simple a utiliser pour ceux qui ne sont pas familiarisé avec `vim` ; je vous recommande tout de même d'utiliser ce dernier si vous avez quelques familiarité avec lui, en n'oubliant pas d'utiliser `:set paste` avant de coller votre code.

Notre Dockerfile étant prêt, il est temps de passer au `build` :

`docker build -t loto .`{{execute}}

- Pour créer l'image, nous appelons la commande `docker build`.

- L'argument `-t` (ou `--tag`) permet de donner un nom à l'image, ainsi qu'un tag, sous la forme `nom:tag`. Ici, comme nous ne le précisons pas, le tag `latest` est automatiquement ajouté à notre image.

- Enfin le `.` désigne le chemin où se trouve notre Dockerfile.

Une fois l'image créée, il ne nous reste plus qu'à essayer de l'exécuter :

`docker run -d --name loto loto`{{execute}}

Avec l'argument `-d` ou `--detach`, nous envoyons notre conteneur en tâche de fond. Vérifions s'il tourne correctement :

`docker ps -f name=loto`{{execute}}

Avec l'argument `-f` ou `--filter` nous avons l'occasion de filtrer les conteneurs. Ici nous utilisons le nom du conteneur.

Et finalement capturons les logs ne notre conteneur :

`docker logs loto`{{execute}}

Et ça tourne !