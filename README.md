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

Run docker compose with MongoDB
```
docker compose up --build
```
## Local setup with `docker compose`   
Rebuild and start detached.
```
sudo docker compose up --build -d
```  
Connect via `MongoDB Compass` to MongoDB with following url (setup must be up).
```
mongodb://localhost:27017
```  
Other stuff
```
sudo docker ps
```  
### Containerized Local Development
docker-compose.yaml consists of:
- mongo
- go-backend (this repo)
- minio (TODO)
### Forecasted Deployment  
We will use managed Mongo and managed storage with S3 interface.  
Then we will use appropriate CI/CD with different .env parametrization to deploy go-backend into Kubernetes. 

## Database migrations run
We manage MongoDB database migrations using [migrate](https://aur.archlinux.org/packages/migrate) on Arch Linux. Our migrations are defined in JSON format, which are executed in MongoDB via db.`runCommand()`.  

We run migration as showed in the attached snippet.  
```
[stepo@archlinux LightningEverywhere-backend]$ migrate -path db/migrations -database "mongodb://localhost:27017/mydb" up
1/u initial (122.823423ms)
2/u seed_merchants (212.237493ms)
```  
Migrate is de-facto migration standard and can be read about in the [official documentation](https://pkg.go.dev/github.com/golang-migrate/migrate/v4@v4.18.3#section-readme). Another resource is [golang-migrate/migrate](https://github.com/golang-migrate/migrate/blob/v4.18.3/database/mongodb/README.md) with specific MongoDB + Go duo description.
## TODO Roadmap
In order to fully run Lightning Everywhere backend as this service we need to implement multiple features and components.  
These things are: 
- Implement authentication first via e-mail+password with some kind of bearer token. 
- MongoDB for document database for `eshops` and `merchants`.
- Authentication via SSO Google (implemented to accomodate Go backend).
- Use S3 (compatible) storage for photo uploads on DigitalOcean [Spaces Object Storage](https://www.digitalocean.com/products/spaces).
- We will use [Minio](https://github.com/minio/minio) (as S3 compatible substitute on localhost) for Docker compose local setup.

These cloud services first will be implemented to run in `DigitalOcean` infrastructure. We might migrate then to `AWS` later on. Web application runs in DOKS (DigitalOcean Kubernetes), might be easily migrated to Amazon EKS (AWS Elastic Kubernetes).  

## Backend programming
- auth e-mail/password (send "click it's me" e-mail) for now, then SSO Google interoperably
- simple HTTP CRUD: e-shops, merchants, profile
- Minio S3 stuff for images handling
- make sure Mongo migrations (if empty -> run...if populated -> don't run), maybe store some version of database in the database
- session syncing via Redis
