#!/bin/bash
[[ -f count.db  ]] || touch count.db && echo "db is ok"
sqlite3 -init ./sql/sqllite3_init.sql count.db <<- EOF
.exit
EOF