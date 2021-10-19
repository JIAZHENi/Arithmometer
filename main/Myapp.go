package main

import (
	"GO_code/Arithmometer/pkg"
	"flag"
	"fmt"
)
// 名字，预设值，用法（-help，能够返回参数提示）
var n = flag.Int("n",10,"int类型")
var r = flag.Int("r",10,"int类型")
var e = flag.String("e","./exercisefile.txt","学生答案文件路径")
var a = flag.String("a","./answerfile.txt","正确答案文件路径")

func main() {
    //友好的控制台提示，可以看到自己的输入参数
	flag.Parse()                     //flag.Parse()把用户传递的命令行参数解析为对应变量的值，不然永远是预设值
	fmt.Println("-n:", *n)
	fmt.Println("-r:", *r)
	fmt.Println("-e:", *e)
	fmt.Println("-a:", *a)

	pkg.OperationGenerated(*n,*r)    //生成n个文件，范围0~r-1
	pkg.CompareDocuments(*e,*a)      //学生答案文件路径n和正确答案文件路径e

}