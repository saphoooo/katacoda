## Le code

Le code de votre application est dans votre *home*:

`ls ~`{{execute}}

C'est du Go.

`kubectl` et `minikube` sont déjà téléchargés, et `minikube` est déjà démarré (peut-être):

`minikube status`{{execute}}

Si `host`, `kubelet`, et `apiserver` ne sont pas encore `running`, il vous faudra encore un peu de partience.

## Dockerfile

Comment, à patir de notre code, allons-nous créer une image Docker ?

Tout d'abord, mettons par convention notre code dans un répertoire portant le nom de l'application que nous développons ; ici, c'est une application pour générer des grilles du loto :

`mkdir loto`{{execute}}

`mv main.go loto`{{execute}}

`cd loto`{{execute}}

Pour créer une image, la méthode la plus répandue à ce jour c'est d'utiliser un Dockerfile : un fichier qui décrit les étapes d'installation de notre application en somme.

Essayons de faire ça de manière empirique :

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

### Attention

> Vérifiez que le copié-collé dans `nano` a bien pris les retours à la ligne (ou n'en a pas ajouté là où nous n'en voulons pas).

Notre Dockerfile étant prêt, il est temps de passer au build :

`docker build -t loto .`{{execute}}

- Pour créer l'image, nous exécutons `docker build`, ce qui est somme toute logique.

- L'argument `-t` (ou --tag) permet de donner un nom à l'image, ainsi qu'un tag, sous la forme de `nom:tag`. Ici, comme nous ne précisons le précisons pas, le tag `latest` sera ajouté à notre image.

- Enfin le `.` désigne le chemin où se trouve notre Dockerfile.

Une fois l'image créée, il ne nous reste plus qu'à essayer de l'exécuter :

`docker run -d --name loto loto`{{execute}}

Avec l'argument `-d`, nous envoyons notre conteneur en background. Vérifions s'il tourne correctement :

`docker ps -f name=loto`{{execute}}

`docker logs loto`{{execute}}

Et ça tourne !