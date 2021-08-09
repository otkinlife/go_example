#!/bin/bash
mkdir -p ~/.countDb
cd ~/.countDb
[[ -f count.db  ]] || touch count.db && echo "db is ok"
cd -
sqlite3 -init ./sql/sqllite3_init.sql ~/.countDb/count.db <<- EOF
.exit
EOF