
```bash
brew install postgresql@15
brew services start postgresql@15
brew install nats-streaming-server
brew services start nats-streaming-server
```
after checkout
```bash 
cd L0
go get .
pdql
```
```psql
\ir conf/create_db.sql
\q
```
```bash
psql -U marla -d db_wb_test -q -f conf/fill_db.sql
```

Локально: 
1. создать psql bd  + юзер + таблицы 
2. развернуть  nats-streaming

