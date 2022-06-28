# go-postgres

Basic Go application with Postgres.

1. Run PostgreSQL database with docker compose:
```
$ docker-compose up -d
```

2. Get PostgreSQL container IP address:
```
$ docker ps
CONTAINER ID   IMAGE      COMMAND                  CREATED             STATUS         PORTS                                       NAMES
29192497eff2   adminer    "entrypoint.sh docke…"   About an hour ago   Up 7 minutes   0.0.0.0:8080->8080/tcp, :::8080->8080/tcp   go-postgres_adminer_1
796d7e557340   postgres   "docker-entrypoint.s…"   About an hour ago   Up 7 minutes   5432/tcp                                    go-postgres_db_1

$ docker inspect 796d7e557340
```

3. Run Go application with `DATABASE_URL`:
```
DATABASE_URL="postgres://postgres:postgres@172.18.0.2:5432/postgres" go run main.go
```