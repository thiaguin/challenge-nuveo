FROM ubuntu

COPY . .

CMD ./migration exec -url "postgres://backend-username:pwd2021@backend-postgres:5432/nuveo-db?sslmode=disable" -dir ./migrations -action up