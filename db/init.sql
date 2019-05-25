USE TCX;
CREATE TABLE users(
    joid INT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    pass TEXT NOT NULL,
    token TEXT NOT NULL
);

CREATE TABLE syubetsu_list(
    name TEXT NOT NULL,
    syubetsu TEXT NOT NULL,
    salary INT NOT NULL
);

-- this table is a test table. table create should be managed by api server
CREATE TABLE `5_rireki`(
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    joid INT NOT NULL,
    syubetsu INT NOT NULL,
    about TEXT NOT NULL,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL
);

-- this insert string is test data.
INSERT INTO `5_rireki` (joid,syubetsu,about,start_time,end_time) VALUES(63,12,"ビラ",cast('2019-05-16 12:34:00' as DATETIME),cast('2019-05-16 12:56:00' as DATETIME));
INSERT INTO `5_rireki` (joid,syubetsu,about,start_time,end_time) VALUES(63,12,"会議",cast('2019-05-17 13:34:00' as DATETIME),cast('2019-05-17 13:56:00' as DATETIME));
INSERT INTO `5_rireki` (joid,syubetsu,about,start_time,end_time) VALUES(94,12,"ポスタースタンド",cast('2019-05-16 12:34:00' as DATETIME),cast('2019-05-16 12:56:00' as DATETIME));
INSERT INTO `5_rireki` (joid,syubetsu,about,start_time,end_time) VALUES(94,12,"会議",cast('2019-05-17 12:34:00' as DATETIME),cast('2019-05-17 12:56:00' as DATETIME));

