FROM mysql:8.0
ENV MYSQL_ROOT_PASSWORD="root"
ENV MYSQL_USER="test"
ENV MYSQL_PASSWORD="test"
ENV MYSQL_DATABASE="TCX"

COPY my.cnf /etc/mysql/conf.d/my.cnf
COPY init.sql /docker-entrypoint-initdb.d/init1.sql
COPY superUser.sql /docker-entrypoint-initdb.d/init2.sql

RUN apt-get update && \
    apt-get install -y locales && \
    rm -rf /var/lib/apt/lists/* && \
    echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen && \
    locale-gen ja_JP.UTF-8
ENV LC_ALL ja_JP.UTF-8
EXPOSE 3306
CMD ["mysqld"]