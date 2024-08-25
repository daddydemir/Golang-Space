package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	startTime := time.Now().UnixMilli()
	url := "http://sunucu1:8090/ping?time=" + fmt.Sprintf("%v", startTime)
	//fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	body, _ := io.ReadAll(resp.Body)
	var data response
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
	}
	endTime := time.Now().UnixMilli()
	/*	log.Println("start: ", startTime)
		log.Println("midd:  ", data.Endtime)
		log.Println("end:   ", endTime)

		log.Println("sunucuya gidis suresi  : ", data.Endtime-startTime, "ms")
		log.Println("sunucudan gelis suresi : ", endTime-data.Endtime, "ms")
		log.Println("toplam gecen sure      : ", endTime-startTime, "ms")
	*/
	log.Println((endTime-startTime)/2, "ms")
}

type response struct {
	Endtime int64
}
