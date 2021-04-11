#!/bin/bash
set -ex

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE uber;
    GRANT ALL PRIVILEGES ON DATABASE uber TO postgres;
EOSQL

set +ex
