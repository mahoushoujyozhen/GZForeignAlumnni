#!/bin/bash
exitOnErr(){
  exitCode=$?
  echo failed to deploy with exit code $exitCode on deploy.sh $1
  exit $exitCode
}

trap 'exitOnErr $LINENO' ERR

# ----------------- 

export backEndFileName=alumni
export deployDst=/var/deploy/alumni
export workspace=/var/jenkins_home/workspace/alumni-association
export websrc=/var/jenkins_home/workspace/alumni-association/web/dist

# correct go build syntax
# go build -o kzz.io -ldflags \
#  "-X main.buildVer=r0.b_nonCI(2018-02-21_09:46:15) '-extldflags=-v -static'"

echo "build from source"
export GOPATH=/var/deploy/kUser/goUser:$workspace
export GOROOT=/usr/local/go
# export GO111MODULE=off
# ./b.sh
export BE_VER=$(git rev-parse HEAD)_$(date '+%Y-%m-%dT%H:%M:%S')

# go env -w GOPROXY=https://goproxy.cn,direct
# go env -w GOSUMDB="sum.golang.google.cn"
cd serve
go build \
  -o $backEndFileName \
  -ldflags "-X main.buildVer=${BE_VER} '-extldflags=-v -static'" \
  main/*.go

echo "  build complete"


echo "stop alumnisrv container"
d stop alumnisrv > /dev/null 2> /dev/null || :
echo "  alumnisrv stoped"

echo "publishing artifact for remote"



sudo mkdir -p $deployDst/f


rm -f "$backEndFileName"_r
mv $backEndFileName "$backEndFileName"_r
chmod +x "$backEndFileName"_r

sudo rsync -u -E "$backEndFileName"_r $deployDst

sudo rsync -r -u -E ${websrc}/ $deployDst/f

echo "  artifact published"

echo "drop old alumni container"
d container rm alumnisrv > /dev/null 2> /dev/null || :

export appName="$backEndFileName"_r
echo "start new alumnisrv container"
d run --name=alumnisrv \
  --restart=always -d \
  -p 9090:9090 \
  -v /etc/ssl/certs:/etc/ssl/certs \
  -v deploy:/var/deploy:z \
  --network qnear \
  k:alpineTZ $deployDst/$appName
echo "deploy done"


