## Namespces

L'exemple le plus simple et qui illustre le mieux la notion de namespace, c'est le namespace PID (vous vous souvenez de Calvin dans son openspace ?).

Votre shell `bash` est l'un des processus exécuté sur la machine sur laquelle vous vous trouvez. Il a certainement comme parent un processus `sshd`, qui lui-même a pour parent `/usr/bin/sshd -D` qui a finalement pour parent `/sbin/init` qui est le processus 1.

Vous pouvez le constater simplement en exécutant la commande :

```
ps -ef
```{{execute}}

Vous pouvez également par cette commande trouver l'ID de votre processus (PID) :

```
echo $$
```{{execute}}

Il est temps de s'isoler en créant un namespace pour votre shell :

```
unshare --fork --pid --mount-proc bash
```{{execute}}

A présent, vous êtes seul au monde :

```
ps -ef
```{{execute}}

Et votre shell se voit maintenant comme le premier processus du système, avec le PID 1 :

```
echo $$
```{{execute}}

## cgroups

Commençons par nous créer un cgroup :

```
cgcreate -a root -g memory:calvin
```{{execute}}

Pour le moment nous n'y avons aposé aucune limite, mais jetons un oeil à ce qui a été créé :

```
ls -l /sys/fs/cgroup/memory/calvin
```{{execute}}

Parmi tous ces fichiers, il y en a un qui s'appelle `memory.kmem.limit_in_bytes`. C'est celui que nous allons utiliser pour définir une limite de 10M (ce qui n'est pas beaucoup) :

```
echo 10000000 >/sys/fs/cgroup/memory/calvin/memory.kmem.limit_in_bytes
```{{execute}}

Et finalement, nous allons placer notre shell dans ce cgroup, le limitant par là même à l'utilisation de 10M de mémoire :

```
cgexec -g memory:calvin bash
```{{execute}}

Là où ça va devenir croustillant, c'est quand nous allons vouloir installer openjdk-8-jdk : 

```
apt update && apt -y install openjdk-8-jdk
```{{execute}}

Oups, nous venons de rencontrer une erreur nos signifiant que nous manquions de mémoire : `Cannot allocate memory` ; cgroups ou coïncidence ?