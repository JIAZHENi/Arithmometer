package pkg

import (
	"math/rand"
	"strconv"
	"time"
)
//生成n个表达式，范围0~r-1
func OperationGenerated(n int,r int){
	var Exe string
	var Ans string
	var s1 string
	var s2 string
	rand.Seed(time.Now().Unix())

	for i := 1; i <= n; i++{
		var x,y,x1,y1 int
		flag := rand.Intn(100)
        switch flag%3 {
        case 0:
			y = rand.Intn(r) + 1
			y1 = rand.Intn(r) + 1
			x = rand.Intn(r * y) + 1
			x1 = rand.Intn(r * y1) + 1

		case 1:
			y = 1
			y1 = rand.Intn(r) + 1
			x = rand.Intn(r * y) +1
			x1 = rand.Intn(r * y1) +1

		case 2:
			y = 1
			y1 = 1
			x = rand.Intn(r * y) + 1
			x1 = rand.Intn(r * y1) + 1
		}
		switch flag%4 {
		case 0:
			s1,s2 = Add(x,y,x1,y1)

		case 1:
			s1,s2 = Sub(x,y,x1,y1)

		case 2:
			s1,s2 = Mul(x,y,x1,y1)

		case 3:
			s1,s2 = Div(x,y,x1,y1)
		}
		Exe = Exe+"四则运算题目"+strconv.Itoa(i)+": "+s1
		Ans = Ans+"答案"+strconv.Itoa(i)+": "+s2
	}

	Output(Exe,"./Exercises.txt")
	Output(Ans,"./Answers.txt")
}











