# Sign-up gRPC service with verification

**Description:** Sign-up user and verify by email.

## Build and run

Update `.env` file:

```dosini
REDIS_ADDR=redis:6379
REDIS_PASS= <Password to access Redis>

PG_HOST=pg
PG_PORT=5432
PG_USER= <Username to access Redis>
PG_PASS= <Password to access Redis>
PG_DB= <Default database>
PG_SSL= <Example: disable>
PG_TIMEZONE= <Example: Asia/Ho_Chi_Minh>

EMAIL_SENDER= <Email address of sender that send the verification link>
EMAIL_PASS= <Email password (App password)>
EMAIL_SMTP=smtp.gmail.com
SMTP_PORT=587

GRPC_GW= <Host that run the gRPC gateway. Example: 0.0.0.0:8081>
GRPC= <Host that run the gRPC service. Example: 0.0.0.0:8080>
```

**Note:** 
* Docker container has its own DNS that allow to use domain name (usually `0.0.0.0` or `localhost`) as name of the Docker service (here is `redis` or `pg`)
* Email password or app password need to be configure. Search keywork: `app password gmail smtp`

Then, build and run Docker image with the command below:

```docker compose up --build -d```

## Test

Tests are located in `test/` folder. Run test with the command below:

```go test . -v```