# nuveo

## Database Migration

Clone the gosidekick, build the migration and run:

```bash
# To create the database:
./migration exec -url "postgres://backend-username:pwd2021@localhost:5432/nuveo-db?sslmode=disable" -dir ./migrations -action up

```