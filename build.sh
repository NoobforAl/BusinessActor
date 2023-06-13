#!/bin/bash

CSV_FILE_NAME="business-financial-data-mar-2022-quarter-csv.csv"
CSV_FILE="./src/businessActorCsv/$CSV_FILE_NAME"

mv $CSV_FILE /tmp/$CSV_FILE_NAME

go build -o /tmp/webApp ./main.go

rm -rf /app/*

mv /tmp/$CSV_FILE_NAME  /app/$CSV_FILE_NAME 
mv /tmp/webApp /app/webApp