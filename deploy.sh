#!/bin/bash
docker build -t cryptofolio-server . && docker-compose -f docker-compose-prod.yml up -d
