#!/bin/bash
#start $ClosePro
#2013/1/23 陈洋
#QQ:474595397
  
ClosePro=./RriverAPI
suffix=_run
case "${1:-''}" in
'-u')
echo "更新程序！"	
    cp -f ./$ClosePro ./$ClosePro$suffix
esac

pid=`ps -ef|grep $ClosePro$suffix |grep -v grep|awk '{print $2}'`
if [ "x$pid" = "x" ] ; then
    echo "$ClosePro$suffix 没有启动！"
else
    kill -9 $pid
    sleep 1
    pid1=`ps -ef|grep $ClosePro$suffix|grep -v grep|awk '{print $2}'`
    if [ "x$pid1" = "x" ] ; then
            echo "成功杀死$ClosePro$suffix进程：" $pid
        else
            echo "$ClosePro$suffix进程杀死失败！"
            exit 1
    fi
fi
$ClosePro$suffix  1>> log.out 2>&1 &
pid2=`ps -ef|grep $ClosePro$suffix|grep -v grep|awk '{print $2}'`
if [ "x$pid2" = "x" ] ; then
        echo "$ClosePro$suffix服务启动失败！"
    else
        echo "$ClosePro$suffix服务成功启动:" $pid2
fi

