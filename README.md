# miniWallet

ada dua project backend dan front end
di backend menggunakan golang framwork gin dan front end menggunakan vue
untuk cloning di localnya gunakan cara ini
go get github.com/akhamatvarokah/miniWallet atau bisa langsung dengan clone projectnya dan di taruh di goroot local anda


untuk instalasi backend
1. masuk ke directory backend dan tambahkan .env pada project kemudian create database local dan pada .env tambahkan code ini
DB_DRIVER=mysql
DB_USER=root
DB_PASSWORD=secret
DB_PORT=3306
DB_HOST=localhost
DB_NAME=mini_wallet
API_PORT=3000


2. download semua depedency yang dibutuhkan di golang ada beberapa depedency
-go get github.com/jinzhu/gorm
-go get github.com/badoux/checkmail
-go get golang.org/x/crypto/bcrypt
-go get github.com/gin-gonic/gin
-github.com/dgrijalva/jwt-go

3. run server backend dengan go run server.go
4. untuk schema table database otomatis kemigrate jika table belum ada sebelumnya



untuk front end
1. masuk ke directory frontend
2. npm install / yarn untuk mendownload semua package yang dibutuhkan untuk vue
3. yarn serve / npm serve untuk menjalankan server front end