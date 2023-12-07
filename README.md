# NetDoop
## Generate Swagger
```
./update-swagger.sh
```

## Build Image
```
docker build -t netdoop/netdoop:latest -f ./Dockerfile .
```

## Deploy
```
mkdir -p /opt/netdoop
cp ./compose/compose.yml /opt/netdoop/compose.yml
cd /opt/netdoop
docker compose up -d
```

