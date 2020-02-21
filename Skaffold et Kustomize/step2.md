A ce stade, tout est parfait pour continuer à développer, mais pas pour passer en production ! En effet, il vous a été demandé de mettre d'autres limites que celles que vous utilisez en dev (la charge sur l'application en production sera plus élevée), mais égalament de passer à 3 réplicas pour en assurer la haute disponibilité.

Le fichier de déploiement de production doit donc être sensiblement différent de celui que vous utilisez pour la dev... Créer deux fichiers ? Personne, et surtout pas un codeur, n'aime dupliquer ; les dérives sont connues, et vous voulez éviter ça à tout prix.

C'est là que Kustomize entre en jeu : il permet de patcher à la volée un fichier yaml avant de le déployer, tout ça en laissant intact votre fichier originel.

Mais rien de tel qu'une démonstration !

## Préparation

Tout d'abord, nous avons besoin de Kustomize :

```
curl -s "https://raw.githubusercontent.com/\
kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash && chmod +x kustomize && sudo mv kustomize /usr/local/bin
```{{execute}}

Ensuite, du fichier de patch, qui ne va reprendre que les parties que nous voulons changer de notre déploiement :

```
cat << EOF >> patch.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  replicas: 3
  template:
    spec:
      containers:
      - name: myapp
        resources:
          requests:
            memory: "128Mi"
            cpu: "250m"
          limits:
            memory: "256Mi"
            cpu: "500m"
        
EOF
```{{execute}}

Nous avons ensuite besoin d'un fichier pour initialiser Kustomize, et lui dire quel fichier patcher :

```
cat << EOF >> kustomization.yaml
resources:
  - deployment.yaml
patches:
  - patch.yaml
EOF
```{{execute}}

Parfait ! Pour finir, nous allons créer un fichier spécifique pour exécuter Skaffold :

```
cat << EOF >> production.yaml
apiVersion: skaffold/v2alpha4
kind: Config
metadata:
  name: myapp
build:
  tagPolicy:
    sha256: {}
  artifacts:
  - image: myapp
    buildpack:
      builder: "cloudfoundry/cnb:tiny"
deploy:
  kustomize: {}
EOF
```{{execute}}

## Mise en oeuvre

Tout est en place, c'est le moment de vérité. Tout d'abord vérifions l'application qui est actuellement déployée correspond bien à notre application de dev :

```
kubectl get deployment myapp -ojsonpath={.spec.replicas}
```{{execute}}

Le nombre de réplicas doit être à 1.

```
kubectl get deployment myapp -ojson | jq .spec.template.spec.containers[0].resources
```{{execute}}

L'output doit ressembler à ceci :

```
{
  "limits": {
    "cpu": "200m",
    "memory": "128Mi"
  },
  "requests": {
    "cpu": "100m",
    "memory": "64Mi"
  }
}
```

Après ces vérifications, passons au déploiement de notre version de production :

```
skaffold run -f production.yaml
```{{execute}}

Constatez-le par vous-même, le fichier de déploiement n'a subit aucune altération :

```
cat deployment.yaml
```{{execute}}

Pourtant, si nous repassons nos commandes maintenant :

```
kubectl get deployment myapp -ojsonpath={.spec.replicas}
```{{execute}}

Le nombre de réplicas doit maintenant être à 3.

```
kubectl get deployment myapp -ojson | jq .spec.template.spec.containers[0].resources
```{{execute}}

L'output doit maintenant ressembler à ceci :

```
{
  "limits": {
    "cpu": "500m",
    "memory": "256Mi"
  },
  "requests": {
    "cpu": "250m",
    "memory": "128Mi"
  }
}
```

Grâce à Kustomize, vous avez personnalisé un déploiement de production, simplement en patchant votre fichier de dev.