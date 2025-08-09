#!/bin/bash
set -e

# Inicia o MySQL em segundo plano
mysqld --user=mysql --skip-networking --socket=/var/run/mysqld/mysqld.sock &

# Aguarda o MySQL estar pronto para aceitar conexões
while ! mysqladmin ping -hlocalhost -uroot -proot --silent; do
    sleep 1
done

# Cria o banco de dados 'orders' se ele não existir
mysql -u root -proot -e "CREATE DATABASE IF NOT EXISTS orders;"

# Encerra o MySQL em segundo plano
mysqladmin -u root -proot shutdown

# Inicia o MySQL no primeiro plano (para que o container continue rodando)
exec mysqld --user=mysql