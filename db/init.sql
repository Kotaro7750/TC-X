USE TCX;
CREATE TABLE users(
    joid INT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    pass TEXT NOT NULL,
    token TEXT NOT NULL
);

CREATE TABLE syubetsu(
    name TEXT NOT NULL,
    syubetsu INT NOT NULL UNIQUE,
    salary INT NOT NULL
);

-- this table is a test table. table create should be managed by api server
CREATE TABLE `2019_5_rireki`(
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    joid INT NOT NULL,
    syubetsu INT NOT NULL,
    about TEXT NOT NULL,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL
);

-- this insert string is test data.


INSERT INTO syubetsu (name,syubetsu,salary) VALUES("窓口",3,1000);
INSERT INTO syubetsu (name,syubetsu,salary) VALUES("広報局",12,1000);