package pkg

import (
	"fmt"
	"os"
)
//文件写入
func Output (str string,name string){
	f, err := os.Create(name) //创建文件
	check(err)
	defer f.Close()
	n, err := f.WriteString(str)
	fmt.Printf("文件%s 写入 %d 个字节\n", name,n)
	f.Sync()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
