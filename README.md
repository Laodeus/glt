# glt

## Description

CGLT est une API écrite en Go qui permet à un utilisateur de se connecter, de prendre un véhicule et d'envoyer sa position avec un timestamp. En appelant l'API à intervalle régulier, cela enregistre la position de l'utilisateur. Un endpoint permet de récupérer l'utilisateur et ses déplacements dans un appareil, et un autre endpoint permet de récupérer les appareils avec leurs positions et leur utilisateur.

## Endpoint

- POST /api/v1/users/login : Cet endpoint permet à l'utilisateur de se connecter. Il accepte un nom d'utilisateur et un mot de passe dans le corps de la requête et renvoie un token d'authentification.

- POST /api/v1/users/register : Cet endpoint permet à un nouvel utilisateur de s'enregistrer. Il accepte un nom d'utilisateur, un mot de passe et d'autres informations - pertinentes (comme l'email, le nom complet, etc.) dans le corps de la requête. Il renvoie un message de confirmation si l'enregistrement est réussi.

- POST /api/v1/vehicles/take : Cet endpoint permet à l'utilisateur de prendre un véhicule. Il accepte l'ID du véhicule et le token d'authentification dans le corps de la requête.

- POST /api/v1/vehicles/leave : Cet endpoint permet à un utilisateur de sortir d'un véhicule. Il accepte l'ID du véhicule et le token d'authentification dans le corps de la requête. Il renvoie un message de confirmation si l'action est réussie.

- POST /api/v1/positions : Cet endpoint permet à l'utilisateur d'envoyer sa position avec un timestamp. Il accepte la position de l'utilisateur, le timestamp et le token d'authentification dans le corps de la requête.

- GET /api/v1/users/{userId}/movements : Cet endpoint permet de récupérer les déplacements d'un utilisateur pour une période donnée. Il accepte l'ID de l'utilisateur dans l'URL, le token d'authentification dans l'en-tête de la requête, et deux paramètres de requête : start et end qui représentent respectivement le timestamp de début et de fin pour la récupération des données.

- GET /api/v1/vehicles : Cet endpoint permet de récupérer les véhicules avec leurs positions et leur utilisateur pour une période donnée. Il accepte le token d'authentification dans l'en-tête de la requête, et deux paramètres de requête : start et end qui représentent respectivement le timestamp de début et de fin pour la récupération des données.