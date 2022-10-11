# Media Sharing Platform - Backend

## Endpoints
#### Ping
| Method | Endpoint | Description |
| ------ | ------ | ------ |
| ```GET``` | /ping | responds with { "message": "pong" } |

#### Files
| Method | Endpoint | Description |
| ------ | ------ | ------ |
| ```GET``` | /file/{filename} | retrieve image / video file |

#### Users
| Method | Endpoint | Description |
| ------ | ------ | ------ |
| ```POST``` | /user/register | user registration |
| ```POST``` | /user/login | user login |
| ```DELETE``` | /user | delete logged-in user |
| ```GET``` | /user/posts/{page_num}?results_per_page | get logged-in user's posts paginated |
| ```GET``` | /user/favorites/{page_num}?results_per_page | get logged-in user's favorites paginated |

#### Posts
| Method | Endpoint | Description |
| ------ | ------ | ------ |
| ```POST``` | /post | create a post |
| ```GET``` | /post/{id} | retrieve a post |
| ```DELETE``` | /post/{id} | delete a post |
| ```POST``` | /post/{id}/favorite | favorite a post |
| ```DELETE``` | /post/{id}/favorite | unfavorite a post |

#### Building the container

```bash
docker build -t media-sharing-platform-api:1.0 .
```
