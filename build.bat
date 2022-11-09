docker pull mysql:latest
docker run -itd --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=P@ssw0rd mysql
docker build -t UI/docker .
docker run --rm -p 4041:4041 -d UI/docker