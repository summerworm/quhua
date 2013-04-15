package main

import (
	"bufio"
	"io"
	"labix.org/v2/mgo"
	"os"
)

type Quhua struct {
	Id   string "_id"
	City string
	Pid  string
	Zip  string
	Call string
}

func main() {
	//读文件
	rf, _ := os.Open("quhua.txt")
	defer rf.Close()
	src := bufio.NewReader(rf)

	//数据库
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	//session.SetMode(mgo.Monotonic, true)
	c := session.DB("quhua").C("quhua")

	var Q Quhua

	for {
		lineByte, _, err := src.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		Q.Id = string(lineByte)

		lineByte, _, err = src.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		Q.City = string(lineByte)

		lineByte, _, err = src.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		Q.Pid = string(lineByte)

		lineByte, _, err = src.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		Q.Zip = string(lineByte)

		lineByte, _, err = src.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		Q.Call = string(lineByte)

		//写数据库
		err = c.Insert(&Q)
		if err != nil {
			panic(err)
		}
	}

}
