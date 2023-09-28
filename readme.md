
```bash
brew install postgresql@15
brew services start postgresql@15
brew install nats-streaming-server
brew services restart nats-streaming-server
```
after checkout
```bash 
cd L0
go get .
psql
```
```psql
\ir conf/create_db.sql
\q
```
```bash
psql -U marla -d db_wb_test -q -f conf/fill_db.sql
```

---

тестовые данные в стрим

```bash
go build -o conf/pub  conf/nats_pub.go
bash conf/to_nats.sh
```



