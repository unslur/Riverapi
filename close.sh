#!/bin/bash
#start $ClosePro
#2013/1/23 陈洋
#QQ:474595397
  
ClosePro=./RriverAPI

pid=`ps -ef|grep $ClosePro |grep -v grep|awk '{print $2}'`
if [ "x$pid" = "x" ] ; then
    echo "$ClosePro 没有启动！"
else
    kill -9 $pid
    sleep 1
    pid1=`ps -ef|grep $ClosePro|grep -v grep|awk '{print $2}'`
    if [ "x$pid1" = "x" ] ; then
            echo "成功杀死$ClosePro进程：" $pid
        else
            echo "$ClosePro进程杀死失败！"
            exit 1
    fi
fi

