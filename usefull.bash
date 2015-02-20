GOPATH=$GOPATH:/home/juno/git/imagehoster go test -v



 find . -type f |grep '^./0' |awk '{print "mv "$0" dest"$0  }' 
 find . -type f |grep '^./0' |awk '{print "mkdir -p `dirname "$0"`; mv "$0" dest"$0  }'
 
 :1,96s/dest.\/0/.\//
1,96s/dirname .\/0/dirname .\//

find . -type d -empty -delete


scp 10.132.97.23:git/imagehoster/weber_production.db .
scp 10.132.97.23:git/imagehoster/upload.tar.gz .

tar zxfv  upload.tar.gz
rm upload.tar.gz

SQL
use mysql
drop database weber_production;
create database weber_production;
use weber_production;

update characters set moto='',description ='',region_id=0,adv_phone_id=0 where Topic='sex' and sex='female';

mysqldump --order-by-primary --compress --single-transaction --skip-triggers -uasterisk -p weber_production >weber_production.db

cat weber_production.db |mysql -uasterisk -p weber_production

rm weber_production.db




