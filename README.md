# tp-analyzer-go

## Groupe :
- DERENSY Dany
- GRANIER Antoine

## Structure :
cmd/ (Contient les commandes CLI Cobra)  
&nbsp;&nbsp;&nbsp;&nbsp;root.go (Commande racine)  
&nbsp;&nbsp;&nbsp;&nbsp;analyze.go (Sous-commande "analyze")  

internal/ (Code interne structuré par fonctionnalités)  
&nbsp;&nbsp;&nbsp;&nbsp;config/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;config.go (Lecture du fichier de configuration JSON)  

&nbsp;&nbsp;&nbsp;&nbsp;analyzer/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;analyzer.go (Analyse concurrente des logs)  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;errors.go (Erreurs personnalisées (fichier/parsing))  

&nbsp;&nbsp;&nbsp;&nbsp;reporter/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;reporter.go (Export des résultats en JSON)  

go.mod   
main.go 

## Option de la commande analyze :

  -c : Chemin du fichier de configuration JSON (obligatoire)
  
  -o : Chemin pour exporter le rapport JSON

## Exemple de commande :
```code
  go run main.go analyze -c ./config.json -o ./results.json  
```
