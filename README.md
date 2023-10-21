### Require:
1. Installed Go
1. Installed Docker

---
### Run without Docker :
First we need to create .env file with :
```
PORT=<PORT>
DBPATH="<username>:<password>@tcp(<ip>:<port>)/<databases-name>"
```


Second, We run command :
```bash 
air
```

If you can't run `air`, so please install *air* :
```bash
go install github.com/cosmtrek/air@latest
```

Run `air -v` to check air in your local.

---
or

### Run with Docker:

```bash
docker-compose up -d
```