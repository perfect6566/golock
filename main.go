
package main

import (
	"log"
	"net/http"
	"sync"
	"time"

)

type smap struct {
	l sync.Mutex
	m string
	lastchange time.Time


}

var s  =smap{}
var chlock=make(chan string,1)

var timerlock *time.Timer

func locksmap(w http.ResponseWriter,r *http.Request)  {
	pid:=r.FormValue("pid")
	if s.m=="free"{
		s.l.Lock()
		s.m=pid
		s.lastchange =time.Now()
		w.Write([]byte(pid+"locked\n"))
		log.Println("s.m locked"+s.m)
	}else {

		for{
			select {
			case  <- chlock:

				s.l.Lock()
				s.m=pid
				s.lastchange=time.Now()

				log.Println(pid+ " locked")



				w.Write([]byte(pid+"locked\n"))//必须break，否则for循环不能中断，无法返回http response
				goto END
			}

		}
	END:
		log.Println("RETURN CODE")
	}






}
func unlocksmap(w http.ResponseWriter,r *http.Request)  {
	pid:=r.FormValue("pid")
	if s.m==pid{
		s.m="free"
		s.lastchange=time.Now()
		log.Println(len(chlock))
		if len(chlock)==0{   //很重要，只有chlock长度为0才能往缓冲为1的channel里面send内容，否则send会阻塞
		chlock<-"free"}
		s.l.Unlock()
		w.Write([]byte(pid+"unlock"))
	}else {
		w.Write([]byte(pid+"has not been locked yet,can not perform unlock!"))
	}

}
func main() {
	timerlock = time.NewTimer(time.Second * 0)
	s.m = "free"
	s.lastchange=time.Now()
	http.HandleFunc("/lock", locksmap)
	http.HandleFunc("/unlock", unlocksmap)
	go http.ListenAndServe(":9988", nil)

	for {

		select {
		case <-timerlock.C:
			if s.m != "free" {
				log.Println("time-out release lock for pid:",s.m)
				s.l.Unlock()
				s.m = "free"
				chlock <- "free"
				s.lastchange = time.Now()
				goto END //跳出for里面的select循环
			}
		}
	END:
		{

			if s.m != "free" {
				timerlock.Reset(s.lastchange.Add(10 * time.Second).Sub(time.Now())) //重置timer
			} else {
				timerlock.Reset(10 * time.Second) //重置timer
			}

		}
	}
}


