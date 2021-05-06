# Array
Pas besoin de pointer pour modifier la valeur d'une position dans une array
## Exemple 
func set(array []int, pos, value int) {
    array[pos] = value
}
## Contre-Exemple
Ne fonctionne pas avec un append
## Res
Si ajouter valeur -> passer pointer
Si modifier valeur ->