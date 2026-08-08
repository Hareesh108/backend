#!/usr/bin/env bash
set -e
cd "$(dirname "$0")"
if [[ "$1" == "postgres" ]]; then
  ./mvnw spring-boot:run -Dspring-boot.run.profiles=postgres
else
  ./mvnw spring-boot:run
fi
