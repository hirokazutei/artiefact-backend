# Database-Schema

## Declaring Database Tables

Done with `sqlalchemy` and `alembic`

#### Migrating Database

Generate migration script with **alembic**.

```
alembic revision --autogenerate -m "&{notes on migration}"
```

Apply migration to database.

```
alembic upgrade head
```
