# Vormir

## tl;dr

Keep track of people that will be leaving the company. When identified, it tags the person in a Slack channel.

### Database Migrations

We use [**golang-migrate**](https://github.com/golang-migrate/migrate) to create and run migrations. It is pre-installed
in the docker container for easier use. Install it locally, if you're not using Docker for local development.

#### Initial Setup

```
# create the initial database
# using the psql helpers installed via `postgresql-client`
createdb vormir
```

#### Creating migrations

```
migrate create -seq -dir db/migrations -ext sql <migration-name>
```

#### Running database migrations

```
migrate -path db/migrations -database "$DATABASE_URL" up

# dump the schema
pg_dump -s $DATABASE_URL > db/structure.sql
```
