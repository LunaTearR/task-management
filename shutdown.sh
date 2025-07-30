#!/bin/bash

echo "🛑 Stopping gateway ..."
(cd gateway && docker-compose down)

echo "🛑 Stopping backend ..."
(cd backend && docker-compose down)

echo "🛑 Stopping frontend ..."
(cd frontend && docker-compose down)

echo "💡 Docker containers stopped. Network remains available."