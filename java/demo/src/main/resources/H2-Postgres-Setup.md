# H2 vs PostgreSQL for the Spring Boot Demo

This demo application can run with two different database setups:

1. **H2 in-memory database** (default)
2. **PostgreSQL** (optional profile)

---

## H2 (default)

H2 is an embedded database that runs inside the application process.

Why use H2?
- No separate database server needed
- Fast startup for local development and demos
- Easy to use with the default `application.properties`
- Good for temporary data and quick testing

Important behavior:
- Data is stored in memory only
- Data is lost when the application stops
- This is a safe default for local runs

Configured in `src/main/resources/application.properties`:
```properties
spring.datasource.url=jdbc:h2:mem:demo;DB_CLOSE_DELAY=-1;DB_CLOSE_ON_EXIT=FALSE
spring.datasource.username=sa
spring.datasource.password=
spring.datasource.driver-class-name=org.h2.Driver
spring.jpa.database-platform=org.hibernate.dialect.H2Dialect
spring.jpa.hibernate.ddl-auto=update
spring.jpa.show-sql=true
```

---

## PostgreSQL (optional)

For persistent storage using a real database, use PostgreSQL.

A separate profile file is provided:
- `src/main/resources/application-postgres.properties`

Start the app with this profile:
```bash
cd java/demo
./run.sh postgres
```

This requires a working local PostgreSQL server and matching credentials.

Example PostgreSQL profile settings:
```properties
spring.datasource.url=jdbc:postgresql://127.0.0.1:5432/nextdb
spring.datasource.username=postgres
spring.datasource.password=root
spring.datasource.driver-class-name=org.postgresql.Driver
spring.jpa.database-platform=org.hibernate.dialect.PostgreSQLDialect
spring.jpa.hibernate.ddl-auto=update
spring.jpa.show-sql=true
```

### Create the database and user

If your local PostgreSQL does not have the configured user or database, create them:

```sql
CREATE USER postgres WITH PASSWORD 'root';
CREATE DATABASE nextdb;
GRANT ALL PRIVILEGES ON DATABASE nextdb TO postgres;
```

---

## Which one should you use?

- Use **H2** when you want quick local startup without installing or configuring Postgres.
- Use **PostgreSQL** when you need data persistence between restarts.

---

## Run commands

Default H2 mode:
```bash
cd java/demo
./run.sh
```

PostgreSQL mode:
```bash
cd java/demo
./run.sh postgres
```
