#!/usr/bin/env bash

PurgeDB() {
    echo "Purging DB: $1"
    sudo -u postgres psql -c "DROP DATABASE IF EXISTS $1;" 2>/dev/null
}

# Ensure PostgreSQL service is running
sudo systemctl start postgresql

# Call the PurgeDB function for each database
PurgeDB auth
PurgeDB payments
PurgeDB reviews
PurgeDB books
PurgeDB lending
