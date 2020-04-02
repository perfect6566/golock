# golock
golock is a distribute lock developed by golang.

it provide Lock AND Unlock http rest Api for clients. 

clients should post it's pid to obtain the lock, once one client has been locked, others will be bocked until the locked one call unlock or time-out(10s) expired.

Usage:	

go get github.com/perfect6566/golock

#Lock
curl -XPOST  -d pid=1235 127.0.0.1:9988/lock
it will return: 1235locked

#Unlock
curl -XPOST  -d pid=1235 127.0.0.1:9988/unlock 
it will return:  1235unlocked
