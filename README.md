# glt

## Description

Global Location Tracking

GLT est une API écrite en Go qui permet à un utilisateur de se connecter, de prendre un véhicule et d'envoyer sa position avec un timestamp. En appelant l'API à intervalle régulier, cela enregistre la position de l'utilisateur. Un endpoint permet de récupérer l'utilisateur et ses déplacements dans un appareil, et un autre endpoint permet de récupérer les appareils avec leurs positions et leur utilisateur.

## Endpoint

### Users

- POST /api/v1/users/login : Cet endpoint permet à l'utilisateur de se connecter. Il accepte un login et un mot de passe dans le corps de la requête et renvoie un token d'authentification.

| Paramètre | Type | Description |
| :--------------- |:---------------:| -----:|
| login | string | Le nom d'utilisateur |
| password | string | Le mot de passe de l'utilisateur |

return : string


- POST /api/v1/users/register : Cet endpoint permet à un nouvel utilisateur de s'enregistrer. Il accepte login et un mot de passe dans le corps de la requête. Il renvoie un message de confirmation si l'enregistrement est réussi.

| Paramètre | Type | Description |
| :--------------- |:---------------:| -----:|
| login | string | Le nom d'utilisateur |
| password | string | Le mot de passe de l'utilisateur |

return : string


### Véhicules

- POST /api/v1/vehicles : Cet endpoint permet à l'utilisateur creer un véhicule. Il accepte le nom et le type du vehicule dans le corps de la requête. Il renvoie un message de confirmation si l'enregistrement est réussi.

| Paramètre | Type | Description |
| :--------------- |:---------------:| -----:|
| name | string | Le nom du véhicule |
| type | string | Le type de véhicule |

return : string


- GET /api/v1/vehicles : Cet endpoint permet à l'utilisateur de recuperer la liste des véhicules. il retourne la liste des véhicules.

return : JSON
'''json
{
  [
    {
        id : integer,
        name : string,
        type: string
    }
  ]
}

'''

- POST /api/v1/vehicles/take : Cet endpoint permet à l'utilisateur de prendre un véhicule. Il accepte l'ID du véhicule. Il renvoie un message de confirmation si l'emprunt est réussi.


| Paramètre | Type | Description |
| :--------------- |:---------------:| -----:|
| id | integer | id du vehicule |

return : string


- POST /api/v1/vehicles/leave : Cet endpoint permet à un utilisateur de sortir d'un véhicule. Il accepte l'ID du véhicule dans le corps de la requête. Il renvoie un message de confirmation si l'action est réussie.

| Paramètre | Type | Description |
| :--------------- |:---------------:| -----:|
| id | integer | id du vehicule |

return : string

- GET /api/v1/vehicles : Cet endpoint permet de récupérer les véhicules avec leurs positions et leur utilisateur pour une période donnée. Il accepte le token d'authentification dans l'en-tête de la requête, et deux paramètres de requête : start et end qui représentent respectivement le timestamp de début et de fin pour la récupération des données.

### Mouvement

- POST /api/v1/positions : Cet endpoint permet à l'utilisateur d'envoyer sa position avec un timestamp. Il accepte la position de l'utilisateur et le timestamp dans le corps de la requête. Il renvoie un message de confirmation si l'action est réussie.

| Paramètre | Type | Description |
| :--------------- |:---------------:| -----:|
| id | integer | id de l'utilisateur |
| lat | double precision | latitude |
| lon | double precision | longitude |

return : string


- GET /api/v1/users/{userId}/movements : Cet endpoint permet de récupérer les déplacements d'un utilisateur pour une période donnée. Il accepte l'ID de l'utilisateur dans l'URL et deux paramètres de requête : start et end qui représentent respectivement le timestamp de début et de fin pour la récupération des données dans le corps de la requête.

| Paramètre | Type | Description |
| :--------------- |:---------------:| -----:|
| start | timestamps | debut de l'affichage des données |
| lon | timestamps | fin de l'affichage des données |



// todo
'''json
{
  [
    {
        id : integer,
        name : string,
        type: string
    }
  ]
}

'''
