# Projet-Master
Code pour le projet de master année académique 2020-2021 dont le sujet porte sur l'étude des techniques de *string matching* utilisant les opérations bit à bit sur les mots machines

# Compilation
Le projet peut être compilé sur *linux* et *windows*. Pour compiler le projet sur *linux*, la commande *go build .* peut être utilisée.

## Compilation depuis *linux* pour *windows*
Il est possible de compiler le projet sur *linux* afin de le lancer sur *windows* avec la commande *env GOOS=windows GOARCH=amd64 go build .*

# Exécution du fichier main
Il existe deux manières de lancer le fichier main. Premièrement en exécutant le fichier binaire. Pour linux, il suffit de lancer la commande *./stringmatching* et pour windows, la commande *stringmatching.exe*. Il est aussi possible de lancer le fichier main grâce à la commande *go run main.go*.

# Exécution des tests unitaires
Pour lancer les tests unitaires, il suffit d'utiliser la commande *go test ./tests -v*.