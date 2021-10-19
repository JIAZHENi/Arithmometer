package pkg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)
//文件比较
func CompareDocuments(name1 string, name2 string) {
	file , err :=os.Open(name1)
	file2 , err :=os.Open(name2)
	if err != nil{
		fmt.Println("open file err=",err)
		return
	}

	defer file.Close()
	defer file2.Close()

	reader := bufio.NewReader(file)
	reader2 := bufio.NewReader(file2)
	i := 0
	j := 0
	k := 0
	var s1 string
	var s2 string

	for{
		i++
		str,err :=reader.ReadString('\n')
		str2,err :=reader2.ReadString('\n')
		if str==str2{
			j++
			s1 += fmt.Sprintf("%d, ",i)
		}else{
			k++
			s2 += fmt.Sprintf("%d, ",i)
		}
		if err == io.EOF {
			break
		}
	}

	j1:=strconv.Itoa(j)
	k1:=strconv.Itoa(k)
	var s4 = "Correct: "+j1+" ("
	var s5 = "Wrong: "+k1+" ("

	s1 = s4+s1+")"+"\n"
	s2 = s5+s2+")"
	s3 := s1+s2

	s3=strings.Replace(s3, ", )", ")", -1)

	Output(s3,"./Grade.txt")

}
