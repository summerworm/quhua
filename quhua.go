package main

import (
	"bufio"
	"io"
	"labix.org/v2/mgo"
	"os"
)

type Xingz struct {
	Code string
	City string
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

	var linea string
	var lineb string

	for {
		lineByte, _, err := src.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		linea = string(lineByte)

		lineByte, _, err = src.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lineb = string(lineByte)

		//写数据库
		_ = c.Insert(&Xingz{string(linea), string(lineb)})
	}

}
