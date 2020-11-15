# telemedicine

```
docker build --no-cache --build-arg SSH_PRIVATE_KEY="$(cat ~/.ssh/id_rsa_github)" -t telemedecine:latest -f Dockerfile .
```

```
docker ps -a
TELEMEDICINE_DATA="./postgres-data" docker-compose -p go-dwh down
docker rmi <image>
sudo rm -rf postgres-data

mkdir postgres-data
TELEMEDICINE_DATA="./postgres-data" docker-compose build
TELEMEDICINE_DATA="./postgres-data" docker-compose up -d &
```