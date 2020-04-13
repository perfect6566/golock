# golock
golock is a distribute lock developed by golang.

it provide Lock AND Unlock http rest Api for clients. 

clients should post it's pid to obtain the lock, once one client has been locked, others will be bocked until the locked one call unlock or time-out(10s) expired.

客户端post方式发送pid到/lock api获取锁，如果没有获取锁就会阻塞，直到获取为止，获取锁之后有超时时间(超时时间可以配置)，在超时时间之内如果没有主动释放锁，就会自动释放。

Usage:	

go get github.com/perfect6566/golock 

直接运行api程序： go run github.com/perfect6566/golock

#Lock
curl -XPOST  -d pid=1235 127.0.0.1:9988/lock
it will return: 1235locked

#Unlock
curl -XPOST  -d pid=1235 127.0.0.1:9988/unlock 
it will return:  1235unlocked
