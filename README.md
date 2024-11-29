# Distributed-Social-Network
The project consists of implementing a distributed social network for communication following the idea on which Twitter is based. Project of the Distributed Systems subject of the Computer Science career, Faculty of Mathematics and Computing (MATCOM), University of Havana.
## Requirements
Run
```bash
go mod download
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


