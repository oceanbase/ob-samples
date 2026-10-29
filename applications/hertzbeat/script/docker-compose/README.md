##  Docker-Compose Deploy

- Use Mysql + VictoriaMetrics as Hertzbeat dependent storage -> [HertzBeat+Mysql+VictoriaMetrics Solution](hertzbeat-mysql-victoria-metrics)
- Use Mysql + IoTDB as Hertzbeat dependent storage -> [HertzBeat+Mysql+IoTDB Solution](hertzbeat-mysql-iotdb)
- Use Mysql + Tdengine as Hertzbeat dependent storage -> [HertzBeat+Mysql+Tdengine Solution](hertzbeat-mysql-tdengine)
- Use Postgresql + IoTDB as Hertzbeat dependent storage -> [HertzBeat+PostgreSQL+IoTDB Solution](hertzbeat-postgresql-iotdb)

## Configure the JWT signing key

All four deployments require `SURENESS_JWT_SECRET`. In the selected deployment directory, generate a unique key into a private environment file:

```bash
umask 077
printf 'SURENESS_JWT_SECRET=%s\n' "$(openssl rand -hex 64)" > jwt.env
docker compose --env-file jwt.env up -d
```

Keep `jwt.env` outside version control and reuse it when recreating or upgrading this deployment. Compose rejects an unset or empty variable and passes it to the HertzBeat container; the mounted application configuration reads that environment variable without a fallback key. These examples use an external image, so they do not rely on this repository's automatic key generation.

For an existing deployment that used a public default, generate a replacement and recreate HertzBeat:

```bash
docker compose --env-file jwt.env up -d --force-recreate hertzbeat
```

Changing the signing key invalidates existing JWTs; users must sign in again. Do not reuse either public default or share one key between independent deployments.

## Configuration regression tests

With Python 3 and PyYAML installed, run from this directory:

```bash
python3 tests/test_jwt_configuration.py -v
```

These tests parse the packaged application configuration and all four Compose files, including the required environment variable and configuration mount. When Docker Compose is available, they also run `config --format json` to verify that unset/empty keys fail and a custom key reaches the container environment unchanged. No containers are started, and the rendering tests use a temporary empty Docker configuration directory.

To use a standalone Compose executable, set `COMPOSE=/path/to/docker-compose`. Set `COMPOSE="docker compose"` in CI to require rendering tests instead of allowing them to be skipped when Docker is absent.
