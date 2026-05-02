# Create or open db
sqlite3 data/main.db

# Run migration
sqlite3 data/main.db < migrations/file.sql

# List tables
.tables

# Check schema
.schema table_name

# Storage Classes
NULL
INTEGER
REAL (floating point number)
TEXT
BLOB

