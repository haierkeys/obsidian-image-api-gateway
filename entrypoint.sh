#!/bin/sh
cd /${P_NAME}/
mkdir -p storage/logs
touch storage/logs/c.log
mv storage/logs/c.log storage/logs/c.log_$(date '+%Y%m%d%H%M%S%'| cut -b 1-17)
/${P_NAME}/${P_BIN} run 2>&1 | tee storage/logs/c.log