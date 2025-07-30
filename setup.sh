#!/bin/bash

# ชื่อ network กลางที่ใช้ร่วมกัน
NETWORK_NAME="task_management_network"

# สร้าง Docker network ถ้ายังไม่มี
echo "Checking Docker network: $NETWORK_NAME ..."
if ! docker network ls | grep -q "$NETWORK_NAME"; then
  echo "Creating shared Docker network: $NETWORK_NAME"
  docker network create "$NETWORK_NAME"
else
  echo "Docker network '$NETWORK_NAME' already exists"
fi

# รัน Backend Services
echo "Starting backend services ..."
(cd backend && docker-compose up -d)

# รัน frontend Services
echo "Starting frontend service ..."
(cd frontend && docker-compose up -d)

# รัน Gateway Services
echo "Starting gateway service ..."
(cd gateway && docker-compose up -d)



echo "All services are up and running!"
echo "Gateway is available at: http://localhost:8084"
echo "Frontend is available at: http://localhost:5173"