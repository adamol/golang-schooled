#!/bin/bash

echo "dropping old database tables..."
mysql -uroot -p go_school < drop_tables.sql
echo "database tables dropped!"
echo "creating new tables..."
mysql -uroot -p go_school < create_tables.sql
echo "tables created!"
