#!/bin/bash
# Skrip untuk load ke database

# Password: root

docker exec -i tubes3_13520126_postgres_1 psql -U root -d postgres < dump.sql
