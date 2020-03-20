# got milk? - Shops

## Prerequisites:
- go
- GNU make
#### Dev prerequisites
- [refresh](https://github.com/markbates/refresh)
- [wire](https://github.com/google/wire/) (optional - only needed to update DI containers)

## Setting up:
```
go get github.com/unicsmcr/hs_auth
cd $GOROOT/src/github.com/unicsmcr/hs_auth
cp app.env.example app.env
cp mongo.env.example mongo.env
cp mongo_express.env.example mongo_express.env
```

Prod:
```
make up
```

Dev:
```
make up-dev
```

### Shutting down:
```
make down
```