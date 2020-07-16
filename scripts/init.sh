#!/bin/sh
#replace 'asdf' with your own postgres password
export PGPASSWORD=asdf
redis-server &
#All captial characters in dbname will be converted to lowercase by postgress automatically
~postgres/bin/psql -U postgres -h localhost -p 5432 -c "create database usrdb;"
