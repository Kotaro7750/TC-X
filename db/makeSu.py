import hashlib
import sys
import random

# get arguments from commandline
args = sys.argv

# check if number of commandline arg is 1
if len(args) != 2:
    print("Usage: python3 ./makeSu.py your_password")
    sys.exit(1)

name = "管理者"
joid = 0

password = str(args[1])
hashedPass = hashlib.sha256(password.encode('utf-8')).hexdigest()

random.seed()
rand = str(random.random())

token = hashlib.sha256((name+rand).encode('utf-8')).hexdigest()

# INSERT INTO TCX.users (joid,name,pass,token) VALUES(joid,name,hashedPass,token)
query = 'INSERT INTO TCX.users (joid,name,pass,token) VALUES({},"{}","{}","{}");\n'.format(joid,name,hashedPass,token)

f = open("./superUser.sql",mode='w')
f.write(query)
f.close()

