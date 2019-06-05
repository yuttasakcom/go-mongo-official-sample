# Go Mongo Official Sample

## Docker-Compose

```bash
docker-compose up -d
docker-compose exec mongo bash
mongo -u admin -p secret --authenticationDatabase admin
> use golang
> db.createUser({user: 'user', pwd: 'password',roles: [{ role: 'readWrite', db: 'golang' }]})
```
