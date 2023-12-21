# Backend for teamchallange project

Backend can be available at:

- https://dev-backend-b4vo.onrender.com/

## Local run in Docker

First of all, install "Doker" and "Docker compose".
These URLs have instructions for installing "Doker" on any operating system like Linux, Windows, and Mac.

- https://docs.docker.com/engine/install/
- https://docs.docker.com/compose/install/

### docker-compose

Download the git repository with the code and switch to the branch named "dev".

Create a new network inside docker and named it as tch_network:

```bash
docker network create tch_network
```

Check whether the docker network has been created using the command:

```bash
docker network ls
```

Create a permanent volume inside docker and named it as tch_postgres_pg_data:

```bash
docker volume create tch_postgres_pg_data
```

Check whether the docker volume has been created using the command:

```bash
docker volume ls
```

Change the directory to dev_local, this directory should contain the docker-compose.yml file, then you can run local containers with the command:

```bash
docker-compose up -d
```

Check whether the docker containers have been created using the command:

```bash
docker-compose ps
```

If the containers are running, you need to enter inside the container using the command:

```bash
docker exec -it tch_backend bash
```

Before building project requires the presence of a .env file (example can be found in .env.example)

```bash
DATABASE_URL=postgresql://postgres:postgres@tch_postgres:5432/store
SERVER_HOSTNAME=":8000"
DOCS_HOSTNAME=":8000"
GOOGLE_CALLBACK_DOMAIN='your-own'
GOOGLE_OAUTH_CLIENT_ID='your-own'
GOOGLE_OAUTH_CLIENT_SECRET='your-own'
GOOGLE_OAUTH_REDIRECT_PAGE='your-own'
GOOGLE_EMAIL_ADDRESS='your-own'
GOOGLE_EMAIL_SECRET='your-own'
PASSWORD_RESET_REDIRECT_PAGE='your-own'
```

Run the application inside the container.

```bash
go mod download
go mod vendor
go run main.go
```

If everything is OK, you can use the application with any browser by this URL: http://localhost:8000

## Documentation

Swagger requires env variable **SERVER_HOSTNAME** defined in .env file

Documentation is available at:

- http://localhost:8000/api/docs/index.html
- https://dev-backend-b4vo.onrender.com/api/docs/index.html  

To updated swagger after changing the code run from the repo root:

1. Install swaggo:

```bash
go install github.com/swaggo/swag/cmd/swag@latest`
```

2. Generate docs:

```bash
swag init --parseDependency --parseInternal --parseDepth 1 -md ./documentation -o ./docs
   !! path to docs used in main.go: `_ "study_marketplace/app/docs"` !!
```

## DB code generation

For code generation from SQL to golang we use **sqlc**. Documentation https://docs.sqlc.dev/en/v1.23.0/

Install sqlc

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.23.0
```

Create SQL queries in:

```bash
./queries/{entity_name}.sql
```

In root folder with present `sqlc.yaml` file run:

```bash
sqlc generate
```

Generated code will be in `./queries/`

## Migrations

Database migrations written in Go. https://github.com/golang-migrate/migrate

Use it as CLI:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2
```

For migration file creation use:

```bash
migrate create -ext sql -dir migrations -seq {action_name}
```

Run migrations:

```bash
migrate -path migrations -database "postgresql://postgres:postgres@tch_postgres:5432/store?sslmode=disable" -verbose up
```

Migration using files from:

```bash
/app/db/migrations/000001_init_schema.{up/down}.sql -- up used for upstream migration
                                                    -- down for downstream migration
```

In case of failure and 'dirty database' connect to db with psql and run:

```bash
update schema_migrations set dirty=false;
```

## Database examination

From host machine, not from docker container, run:

```bash
docker exec -it tch_postgres bash
```

Inside of the docker container named tch_postgres, run:

```bash
psql -U postgres store
```

```bash
\dt
```

If table 'users' not present run migrations, else:

```bash
SELECT * FROM users;
```

## Request examples

User creation request example:

```bash
curl -v  POST http://localhost:8000/api/auth/register -H "Content-type: application/json" -d '{"email": "hello@world", "name":"world", "password":"hello"}'
```

Response body:

```json
{
  "Data": {
    "user": {
      "id": "143a5c7a-8e4b-4081-9da8-8005ee7b2e6f",
      "name": "world",
      "email": "hello@world",
      "photo": "default.jpeg",
      "verified": false,
      "password": "$2a$10$5bMojbONIM9KR3IsqvGQEO/eMOuKpRaJ8/qrSkphnaSKNkJeXbkw2",
      "role": "user",
      "created_at": "2023-09-25T07:42:26.808Z",
      "updated_at": "2023-09-25T07:42:26.804Z"
    }
  },
  "Status": "success"
}
```

User login request example:

```bash
curl -v  POST http://localhost:8000/api/auth/login -H "Content-type: application/json" -d '{"email": "hello@world", "password":"hello"}'
```

Response body:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNWZkNTU1MTUtZjYxMS00MTgyLTgwOGUtZjgwY2E1MjNkM2MzIiwidXNlcm5hbWUiOiJ3b3JsZCIsImV4cCI6MTY5NTcxNjk0MCwiaWF0IjoxNjk1NjMwNTQwfQ.JEJkT1vQs_WWFZ_fAPe2i1ScZavD0LgQOGzVJH-coXo"
}
```

Token usage for protected endpoint:

```bash
curl -X GET "http://localhost:8000/protected/userinfo" -H "Authorization: Bearer <token generated by login endpoint>"
```

Response body:

```json
{
  "user": {
    "id": "5fd55515-f611-4182-808e-f80ca523d3c3",
    "name": "world",
    "email": "hello@world",
    "photo": "default.jpeg",
    "verified": false,
    "password": "$2a$10$CqAghRSTwM3vNtGby4aDoOIf4ezGuJA4oUzNEV4oqgic3ORN9.RM2",
    "role": "user",
    "created_at": "2023-09-25T08:28:42.302Z",
    "updated_at": "2023-09-25T08:28:42.3Z"
  }
}
```

## API diagrams

![API Diagram](./api.drawio.svg)
