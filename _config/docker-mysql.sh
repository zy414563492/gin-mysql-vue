docker run --rm -d -p 3306:3306 -h 127.0.0.1 -v /Users/apple/Documents/Code/Go/gin-mysql-vue/_data/mysql:/var/lib/mysql -e MYSQL_DATABASE=mecha-glancer -e MYSQL_USER=glancer -e MYSQL_PASSWORD=glancer -e MYSQL_ROOT_PASSWORD=glancer --name mysql-mecha-glancer mysql:5.7 --character-set-server=utf8mb4 --collation-server=utf8mb4_general_ci
