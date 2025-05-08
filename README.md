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

## TODO Roadmap
In order to fully run Lightning Everywhere backend as this service we need to implement multiple features and components.  
These things are: 
- Implement authentication first via e-mail+password with some kind of bearer token. 
- MongoDB for document database for `eshops` and `merchants`.
- Authentication via SSO Google (implemented to accomodate Go backend).
- Some storage (like S3) for photo uploads.

These cloud services first will be implemented to run in `DigitalOcean` infrastructure. We might migrate then to `AWS` later on. Web application runs in DOKS (DigitalOcean Kubernetes), might be easily migrated to Amazon EKS (AWS Elastic Kubernetes).
