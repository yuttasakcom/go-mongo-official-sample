# Go Mongo Official Sample

## Start Docker-Compose

```bash
docker-compose up -d
docker-compose exec mongo bash
mongo -u admin -p secret --authenticationDatabase admin
> use golang
db.createUser({user: 'user', pwd: 'password',roles: [{ role: 'readWrite', db: 'golang' }]})
```

```bash
go mod init github.com/yuttasakcom/go-mongo-official-sample
```
