# How to migration

## Install tool
`go get -v github.com/rubenv/sql-migrate/...`

## Usage
```
$ sql-migrate --help
usage: sql-migrate [--version] [--help] <command> [<args>]

Available commands are:
    down      Undo a database migration
    new       Create a new migration
    redo      Reapply the last migration
    status    Show migration status
    up        Migrates the database to the most recent version available
```

## Migration
### การสร้าง
`$ sql-migrate new create_users_table` หรือ 

`$ sql-migrate new add_email_to_users_table` หรือ

`$ sql-migrate new remove_key_from_users_table`

จะได้ไฟล์ 20180517132536-create_users_table โดยข้างหน้าจะเป็น วัน-เวลา ที่ไฟล์นั้นถูกสร้าง เรียงตามลำดับ

### การดูสถานะ
`$ sql-migrate status`

ระบบจะแสดงสถานะการรันไฟล์ sql

### การอัพเกรด
`$ sql-migrate up`

ระบบจะรันไฟล์ sql ทั้งหมดตามลำดับ

### การดาว์นเกรด
`$ sql-migrate down`

ระบบจะย้อนการรันไฟล์ sql ทั้งหมดตามลำดับ

### การ redo
`$ sql-migrate redo`

ระบบจะย้อนการรันไฟล์ sql ทีละไฟล์