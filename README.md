
# XSIS-GO

Project ini di buat menggunakan golang versi 22rc2
- database menggunakan postgress
- silahkan export file atau buat database terlebih dahulu kemudian jalankan perintah ini
```bash
CREATE TABLE IF NOT EXISTS movie (
  id SERIAL not key,
  title varchar(100) NOT NULL,
  description varchar(200) NOT NULL,
  rating float NOT NULL,
  image varchar(100) default '',
  created_at timestamp,
  updated_at timestamp
)
```

- untuk request dan respons menggunakan generator dari proto
- menggunakan depedensi injection sebagai alur database
- Pattern yang di pakai project ini menggunakan MVVM
## Installation

Untuk menginstall go pergi ke tautan yang telah di sematkan

```bash
https://go.dev/doc/install
```
kemudian di project ini menggunakan Make perintah untuk menjalankan project

```bash
make server
```

## Endpoint
Endpoint yang tersedia di mini project ini diantaranya
- Insert data 
- Update data 
- Delete data 
- Get List Data 
- Get Detail Data
## Authors

- [@muhamadari10](https://github.com/muhamadari10)

