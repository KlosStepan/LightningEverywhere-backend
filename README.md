# LightningEverywhere-backend
Build Docker image
```
sudo docker build -t stepanklos/lightningeverywhere-backend .
```
Run Docker image (with or without `-d`)
```
sudo docker run -d \
  --name lightning-backend \
  -p 8080:8080 \
  stepanklos/lightningeverywhere-backend
```
Delete Docker image
```
sudo docker rm lightning-backend
```

Run docker-compose with MongoDB
```
docker-compose up --build
```
