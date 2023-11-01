### Require:
1. Installed Go
1. Installed Docker

---
### Run without Docker :
1. First we need to create .env file with copy some code:
    ```
    PORT=<PORT>
    DBPATH="<username>:<password>@tcp(<ip>:<port>)/<databases-name>"
    ```


2. Second, We run command :
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

### Run with Docker Compose:

```bash
docker-compose up -d
```

---
### To run access and create table in database:
1. You need to copy this line into your terminal:
    ```bash
    docker exec -it db bash
    ```
2. Next, access into Mysql Server with:
    ```bash
    mysql -u root -p
    ```

    After that, type `password` to input the password
3. Then, to use a database named golang. Copy this :
    ```sql
    use golang;
    ```
4. Finally, go to folder `database/db/init.sql/CreateUser.session.sql` or [click here!](https://github.com/billzayy/crud-go/blob/main/database/db/init.sql/CreateUser.session.sql) and copy all SQL Statements into your terminal.

5. To check table exists, type this command :
    ```sql
    desc User;
    ```
    If your mysql server still don't have a User table, rerun docker compose and to it again !!!



