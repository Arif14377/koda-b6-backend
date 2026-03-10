# Backend Project Coffee Shop
Aplikasi backend untuk project ecommerce Coffee Shop

# Implementasi
1. [x] Middleware CORS
5. [x] Connection Database
2. [x] CRUD ke Database
3. [x] Hashing Password menggunakan Argon2
4. [x] Setting Environment Variable
6. [ ] Tambahkan fitur file upload
7. [ ] Tambahkan JWT ketika login

## Catatan Trainer
1. Kembalian delete, add, edit berupa data sebelum dihapus
2. Kardinalitas
3. Tes fase 3
4. RowToStruct

## Golang Migrate
- Dokumentasi
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

- Instalasi via CLI
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Jika langkah selanjutnya tidak bisa karena command migrate not found, perlu cek dulu apakah binary sudah terinstall? jika sudah terinstall (sudah ada folder migrate), berarti pathnya belum diset.

cara set path:
- biasanya go menyimpan di : ~/go/bin
- Jika pakai bash, maka bisa edit .bashrc:
```bash
    nano ~/.bashrc

    #tambahkan:
    export PATH=$PATH:$(go env GOPATH)/bin
    #atau
    export PATH=$PATH:$HOME/go/bin
```

- Cara pakai
```bash
migrate create -ext sql -dir migrations -seq init_db
```
    - setelah terinstall akan ada 2 file di folder yang dibuat, yaitu file up dan down.
    - file up untuk create DDL, file down untuk create ddl (drop).

- Menjalankan migration
```bash
# $ migrate -source file://path/to/migrations -database postgres://localhost:5432/nama_database up 2

migrate -source file://./migrations -database postgres://postgres:1@localhost:5432/weekly-db?sslmode=disable up

# ada flag (pahami penggunaannya)
# jika migrasi tidak pakai angka maka akan menjalankan semua migrasi sisanya yang belum dijalankan.
```

## Seed
- Tempat untuk initiate data awal untuk database (seperti data dummy).


##
```go
postgres://user:1@server:port/db_name?sslmode=disable
```