package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	"strings"
)

type Xingz struct {
	Code string
	City string
}

func main() {
	//读文件
	rf, err := ioutil.ReadFile("quhuamin.txt")
	if err != nil {
		fmt.Println("file read error.")
	}
	src := string(rf)

	//数据库
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//session.SetMode(mgo.Monotonic, true)
	c := session.DB("quhua").C("quhua")

	//fmt.Print(src)
	srcbuf := bufio.NewReader(strings.NewReader(src))
	for err == nil {
		isP := true
		var linea []byte
		var lineb []byte
		for isP {
			linea, isP, err = srcbuf.ReadLine()
			lineb, isP, err = srcbuf.ReadLine()
			//fmt.Println(string(linea) + string(lineb))

			_ = c.Insert(&Xingz{string(linea), string(lineb)})
			//以下代码会使quhua.txt处理完之后进入无穷循环
			//从而在数据库中产生无穷空行
			//err = c.Insert(&Xingz{string(linea), string(lineb)})
			//if err != nil {
			//	panic(err)
			//}

		}
	}

}
