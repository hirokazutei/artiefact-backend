### Run Server

```
go run cmd/artiefact-server/main.go cmd/artiefact-server/devel.toml
```

## Initial Set-up

#### Install python requirements

```
python3 -m venv venv
source venv/bin/activate
pip -r install requirements.txt
```

#### Set Up Env

```
export ARTIEFACT_DATABASE_URL="localhost"
export ARTIEFACT_DATABASE_NAME="{}"
export ARTIEFACT_DATABASE_PASS="{}"
export ARTIEFACT_DATABASE_USER="{}"
export ARTIEFACT_PASSWORD_PEPPER="{}"
```

#### Create new database

```
psql -U {}

CREATE ROLE artiefact WITH Superuser, CreateDB, LOGIN
```

#### Database migration

```
alembic upgrade head
```
