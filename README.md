# nuveo

## Database Migration

```bash
# To create the database, on migration folder:
./migration exec -url "postgres://backend-username:pwd2021@localhost:5432/nuveo-db?sslmode=disable" -dir ./fixtures -action up

```