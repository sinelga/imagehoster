GOPATH=$GOPATH:/home/juno/git/imagehoster go test -v

 find . -type f |grep '^./0' |awk '{print "mv "$0" dest"$0  }'
 
 find . -type f |grep '^./0' |awk '{print "mkdir -p `dirname "$0"`; mv "$0" dest"$0  }'
 
 :1,96s/dest.\/0/.\//
1,96s/dirname .\/0/dirname .\//

find . -type d -empty -delete