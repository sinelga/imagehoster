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

#backup
scp 104.131.122.192:git/imagehoster/weber_production.db .
drop database weber_production;
create database weber_production; 

cat weber_production.db |mysql -uasterisk -p weber_production

rm weber_production.db

#Cuuma
insert into adv_phones(phone,created_at,updated_at) values('0700411308',now(),now());
insert into adv_phones(phone,created_at,updated_at) values('0700411343',now(),now());
insert into adv_phones(phone,created_at,updated_at) values('0700411344',now(),now());
insert into adv_phones(phone,created_at,updated_at) values('0600411464',now(),now());
insert into adv_phones(phone,created_at,updated_at) values('0600411465',now(),now());
insert into adv_phones(phone,created_at,updated_at) values('0600411466',now(),now());
insert into adv_phones(phone,created_at,updated_at) values('0600411467',now(),now());
insert into adv_phones(phone,created_at,updated_at) values('0600411468',now(),now());
insert into adv_phones(phone,created_at,updated_at) values('0600411469',now(),now());
#insert into adv_phones(phone,created_at,updated_at) values('',now(),now());
#insert into adv_phones(phone,created_at,updated_at) values('',now(),now());

select * from characters;1442
select * from characters where adv_phone_id not in (85,86,87);958

select * from characters where adv_phone_id not in (85,86,87) limit 100;

update characters set adv_phone_id =89 where adv_phone_id not in (85,86,87) limit 100;

select * from characters where adv_phone_id=89;

update characters set adv_phone_id =90 where adv_phone_id not in (85,86,87,89) limit 100;


#OLD weber_production
select * from adv_phones where phone; 83

select * from adv_phones where phone not in ('0700411339','0700411342','0700411310','0700411308','0700411343','0700411344','0600411464','0600411465','0600411466','0600411467','0600411468','0600411469'); 71


update adv_phones set phone='0700411343' where phone not in ('0700411339','0700411342','0700411310','0700411308','0700411343','0700411344','0600411464','0600411465','0600411466','0600411467','0600411468','0600411469') limit 20;
update adv_phones set phone='0700411344' where phone not in ('0700411339','0700411342','0700411310','0700411308','0700411343','0700411344','0600411464','0600411465','0600411466','0600411467','0600411468','0600411469') limit 20;
update adv_phones set phone='0700411308' where phone not in ('0700411339','0700411342','0700411310','0700411308','0700411343','0700411344','0600411464','0600411465','0600411466','0600411467','0600411468','0600411469') limit 7;
update adv_phones set phone='0600411464' where phone not in ('0700411339','0700411342','0700411310','0700411308','0700411343','0700411344','0600411464','0600411465','0600411466','0600411467','0600411468','0600411469') limit 7;
update adv_phones set phone='0600411465' where phone not in ('0700411339','0700411342','0700411310','0700411308','0700411343','0700411344','0600411464','0600411465','0600411466','0600411467','0600411468','0600411469') limit 7;
update adv_phones set phone='0600411466' where phone not in ('0700411339','0700411342','0700411310','0700411308','0700411343','0700411344','0600411464','0600411465','0600411466','0600411467','0600411468','0600411469') limit 7;
update adv_phones set phone='0600411467' where phone not in ('0700411339','0700411342','0700411310','0700411308','0700411343','0700411344','0600411464','0600411465','0600411466','0600411467','0600411468','0600411469') limit 7;
update adv_phones set phone='0600411468' where phone not in ('0700411339','0700411342','0700411310','0700411308','0700411343','0700411344','0600411464','0600411465','0600411466','0600411467','0600411468','0600411469') limit 7;
update adv_phones set phone='0600411469' where phone not in ('0700411339','0700411342','0700411310','0700411308','0700411343','0700411344','0600411464','0600411465','0600411466','0600411467','0600411468','0600411469');



 update adv_phones set phone='0700411310' where phone='0700440194'
 update adv_phones set phone='0700411339' where phone='0700440192';
 update adv_phones set phone='0700411342' where phone='0700440193';



mogrify -chop 0x20+0+0 -gravity South teen.amateur.pics.homemade.004.jpg 

mogrify -chop 0x10%+0+0 -gravity South teen.amateur.pics.homemade.*.jpg 







