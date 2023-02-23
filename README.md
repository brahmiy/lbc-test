## Contexte

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by ""fizz"", all multiples of 5 by ""buzz"", and all multiples of 15 by ""fizzbuzz"". 
The output would look like this: ""1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,..."".

Your goal is to implement a web server that will expose a REST API endpoint that:
- Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

## Points d'attention
- Ce projet a été créé en implémentant l'architecture hexagonal.
- Le service des statistiques n'est pas très avancé, il permet tout de même de récupérer les requêtes les plus appelés pendant le démarrage du serveur (in memory service).

## Démarrage du projet en local

* `make start`: démarrage du serveur en local
* `make start-with-docker`: démarrage d'un container docker

### Exemple d'appels : 
- Pour FizzBuzz : 
```curl
    curl --location --request POST 'http://localhost:8088/fizzbuzz' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "limit": 20,
            "int1": 3,
            "int2": 5,
            "str1": "fizz",
            "str2": "buzz"
        }'
```
- Pour les stats :
```curl
    curl --location --request GET 'http://localhost:8088/stats'
```

### Exécution des tests
* `make test`