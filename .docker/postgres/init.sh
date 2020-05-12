#!/bin/bash
set -e
for database in stocks stocks-test;
do
psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = '$database'" | grep -q 1 || \
psql -U postgres <<-EOSQL
CREATE DATABASE "$database" WITH owner=postgres;
EOSQL
done
