#!/usr/bin/env bash

CreateDB() {
   echo "Creating DB $1"
   sudo -u postgres psql -c "CREATE ROLE $1admin WITH LOGIN PASSWORD '6y2pyLv31';"
   sudo -u postgres psql -c "CREATE DATABASE $1;"
   sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE $1 TO $1admin;"
}

sudo systemctl start postgresql

CreateDB auth
CreateDB payments
CreateDB reviews
CreateDB books
