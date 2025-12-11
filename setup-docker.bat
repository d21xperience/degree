@echo off
echo Creating network if not exists...
docker network create app-network 2>nul

echo Starting Auth Service...
cd backend\auth_service
docker-compose up -d
cd ..

REM echo Starting SC Service...
REM cd backend\sc_service
REM docker-compose up -d
REM cd ..

REM echo Starting Sekolah Service...
REM cd backend\sekolah
REM docker-compose up -d
REM cd ..

REM echo Starting Frontend...
REM cd sakai-vue
REM docker-compose up -d
REM cd ..

echo Starting Nginx...
cd nginx
docker-compose up -d
cd ..

echo All services started!
echo Checking running containers...
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

pause