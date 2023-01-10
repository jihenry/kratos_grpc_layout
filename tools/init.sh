# !/bin/sh
command -v kratos >/dev/null 2>&1 || { echo >&2 "command kratos not found"; exit 1; }
CURDIR=$(pwd)
RUNDIR=$(dirname $CURDIR)
echo "rundir:${RUNDIR}"

function PrintHelp()
{
    echo "eg: ./init.sh [arg ...]"
    echo -e "\tOption:"
    echo -e "\t -n \t project name"
    echo -e "\t -p \t proto name"
    echo -e "\t -c \t path of config file"
    exit 1
}

while getopts 'p:c:n:h' OPT; do
    case $OPT in
        h) PrintHelp;;
        n) PROJECT="$OPTARG";PROTOBASE=$PROJECT;;
        p) PROTOBASE="$OPTARG";;
        c) CPATH="$OPTARG";;
        ?) PrintHelp;;
    esac
done

if [ ! -n "$PROJECT" ];then
    echo "projectname cannot be empty"
    PrintHelp
fi 
    
echo "更新项目名称：${PROJECT}"
cd $RUNDIR
mv api/* api/${PROJECT}
mv cmd/* cmd/${PROJECT}

echo "修改proto文件名：${PROTOBASE}"
rm api/${PROJECT}/v1/*.go
mv api/${PROJECT}/v1/greeter.proto api/${PROJECT}/v1/${PROTOBASE}.proto

for i in `grep "helloworld" -r --exclude-dir=local --exclude-dir=.git --exclude=*.exe --exclude=*.sh | cut -d: -f1 | uniq`
do
    sed -i "s/helloworld/${PROJECT}/g" $i
done

for i in `grep "layout" -r --exclude-dir=local --exclude-dir=.git --exclude=*.exe --exclude=*.sh | cut -d: -f1 | uniq`
do
    sed -i "s/layout/${PROJECT}/g" $i
done
kratos proto client api/

if [ ! -n "$CPATH" ];then
    CPATH=local/dev/config.yaml
fi 
if [ ! -e $CPATH ];then
    echo "config file not exist" ; exit 1
fi

set -x

cd cmd/${PROJECT}

trap "rm ${PROJECT};kill 0" EXIT

go build -o ${PROJECT}

./${PROJECT} -conf ${RUNDIR}/${CPATH}
