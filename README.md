# Distributed-Social-Network
The project consists of implementing a distributed social network for communication following the idea on which Twitter is based. Project of the Distributed Systems subject of the Computer Science career, Faculty of Mathematics and Computing (MATCOM), University of Havana.

## Requirements
Run
```bash
go mod download
```
Next, to run the project execute
```bash
go run ./cmd/api
```
to get the necesary dependecies for this project to run succesfully.

## Endpoints
| **Method** | **Route**                        | **Action**                                                   |
|------------|----------------------------------|-------------------------------------------------------------|
| **POST**   | `/users`                        | Create a new user.                                          |
| **GET**    | `/users/{id}`                   | Get details of a specific user.                            |
| **PUT**    | `/users/{id}`                   | Update the profile of a specific user.                     |
| **DELETE** | `/users/{id}`                   | Delete a specific user.                                     |
| **GET**    | `/users`                        | Retrieve a list of users (supports search and pagination).  |
| **POST**   | `/users/{id}/follow`            | Follow a specific user.                                     |
| **DELETE** | `/users/{id}/follow`            | Unfollow a specific user.                                   |
| **GET**    | `/users/{id}/followers`         | Retrieve followers of a specific user.                     |
| **GET**    | `/users/{id}/following`         | Retrieve users followed by a specific user.                |
| **POST**   | `/tweets`                       | Create a new tweet.                                         |
| **GET**    | `/tweets/{id}`                  | Get details of a specific tweet.                           |
| **GET**    | `/users/{id}/tweets`            | Retrieve tweets from a specific user.                      |
| **GET**    | `/timeline`                     | Retrieve the timeline of a specific user.                  |
| **DELETE** | `/tweets/{id}`                  | Delete a specific tweet.                                    |
| **POST**   | `/tweets/{id}/retweet`          | Retweet a specific tweet.                                   |
| **DELETE** | `/tweets/{id}/retweet`          | Undo a retweet of a specific tweet.                        |
| **POST**   | `/tweets/{id}/favorite`         | Mark a specific tweet as favorite.                         |
| **DELETE** | `/tweets/{id}/favorite`         | Remove a specific tweet from favorites.                    |
| **GET**    | `/users/{id}/notifications`     | Retrieve notifications for a specific user.                |
| **POST**   | `/messages`                     | Send a private message to another user.                    |
| **GET**    | `/users/{id}/messages/received` | Retrieve received private messages for a specific user.    |
| **GET**    | `/users/{id}/messages/sent`     | Retrieve sent private messages for a specific user.        |
| **POST**   | `/auth/login`                   | Log in a user and return an authentication token.          |
| **POST**   | `/auth/logout`                  | Log out the current user.                                  |
| **POST**   | `/auth/register`                | Register a new user and return an authentication token.    |
| **GET**    | `/users/{id}/stats`             | Get statistics for a specific user (e.g., tweets, likes).  |


docker run -it --network test_kademlia --network-alias client --dns 10.0.10.2 -v "$(pwd)":/app -w /app --name client test

docker run -d --network server-net --network-alias node1 --cap-add=NET_ADMIN -v "$(pwd)":/app -v db_node1:/app/data -w /app --name node1 test go run ./server/cmd/api

### Comandos para levantar los contenedores en redes distintas (en modo interactivo, pero esto no es imprescindible):

#### Redes (crearlas si es necesario):
docker network create --subnet=10.0.10.0/24 test_kademlia
docker network create --subnet=10.0.11.0/24 client

#### Servers:
docker run -it --network client-net --network-alias node1 --cap-add=NET_ADMIN -v "$(pwd)":/app -v db_node1:/app/data --name node1 test sh

**requisitos de conectividad**
docker exec -it node1 sh (para interactuar con el contenedor)
ip route del default via 10.0.10.1 dev eth0 (en caso de no ser este el ip, consultarlo mediante el comando **ip route** dentro del contenedor)
ip route add default via 10.0.10.254 dev eth0

#### DNS:
docker run -it --network test_kademlia --name dns1 --ip 10.0.10.5 dns

**requisitos de conectividad**
docker exec -it dns1 sh
ip route del default via 10.0.10.1 dev eth0
ip route add default via 10.0.10.254 dev eth0

#### Router:
docker run -d --name router --network test_kademlia --ip 10.0.10.254 --cap-add=NET_ADMIN router-image

docker network connect --ip 10.0.11.254 client router


#### Clients:
docker run -it --network client --network-alias client1 --cap-add=NET_ADMIN --dns 10.0.10.5 -v "$(pwd)":/app -w /app --name client1 test

**requisitos de conectividad**
docker exec -it client1 sh (para interactuar con el contenedor)
ip route del default via 10.0.11.1
ip route add default via 10.0.11.254


## Deficiencias a Resolver:
- Al caerse nodos de la red, se intenta contactar con ellos no se recupera del pfallo, esto afecta el guardar valores nuevos en la red, por lo que se sospecha que el error esta en la funcion node.StoreValue de la implementacion de kademlia.

- Al caerse nodos de la red, el login falla, a pesar de el nodo tener la replica correctamente y devolver el usuario, la autenticación no se realiza correctamente.

- Debe bajarse el timpo de replicación dado que se debe esperar como minimo el doble del tiempo de recuperación para una recuperación correcta, (si un usuario no logra replicarse por el tiempo en el que entro el nodo es posterior a la llamada de replicación del usuario, los post asociados a el van a fallar por no encontrar la entidad principal)

- Al replicar tabla de usuarios se rehashea la contraseña corrompiendola y haciendola inaccesible para el usuario