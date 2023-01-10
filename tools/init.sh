# !/bin/sh
command -v kratos >/dev/null 2>&1 || { echo >&2 "command kratos not found"; exit 1; }
CURDIR=$(pwd)
RUNDIR=$(dirname $CURDIR)
echo "rundir:${RUNDIR}"

if [ ! -n "$1" ];then
    echo "please input project name. eg: /.init.sh tact"
    exit 0
else
    PROJECT=$1
    PROTOBASE=$PROJECT
fi
echo "更新项目名称：${PROJECT}"
cd $RUNDIR
mv api/* api/${PROJECT}
if [ -n "$2" ];then
    PROTOBASE=$2
fi
echo "修改proto文件名：${PROTOBASE}"
rm api/${PROJECT}/v1/*.go
mv api/${PROJECT}/v1/greeter.proto api/${PROJECT}/v1/${PROTOBASE}.proto

for i in `grep "tact" -r --exclude-dir=local --exclude-dir=.git --exclude=*.exe | cut -d: -f1 | uniq`
do
    sed -i "s/tact/${PROJECT}/g" $i
done

for i in `grep "tact" -r --exclude-dir=local --exclude-dir=.git --exclude=*.exe | cut -d: -f1 | uniq`
do
    sed -i "s/tact/${PROJECT}/g" $i
done
kratos proto client api/

CPATH=local/dev/config.yaml
if [ ! -e $CPATH ];then
    echo "config file not exist" ; exit 1
fi
if [ -n "$3" ];then
    CPATH=$3
fi

set -x

cd cmd/${PROJECT}

trap "rm ${PROJECT};kill 0" EXIT

go build -o ${PROJECT}

./${PROJECT} -conf ${RUNDIR}/${CPATH}