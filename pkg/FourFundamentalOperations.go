package pkg

import (
	"fmt"
)
//生成四则运算式和对应的答案

func Add(x int,y int,x1 int,y1 int)(string,string){
	k := CommonDivisor(x*y1+y*x1,y*y1)
	var s1 string
	var s2 string

	if x < y {
		if x1 < y1 {
			s1 = fmt.Sprintf("%d/%d + %d/%d = \n",x,y,x1,y1)
		}else{
			if x1 % y1 == 0{
				s1 = fmt.Sprintf("%d/%d + %d = \n",x,y,x1/y1)
			} else{
				s1 = fmt.Sprintf("%d/%d + %d'%d/%d = \n",x,y,x1/y1,x1-(x1/y1)*y1,y1)
			}
		}
	}else{
		if x % y == 0 {
			if x1 < y1 {
				s1 = fmt.Sprintf("%d + %d/%d = \n", x/y, x1, y1)
			} else {
				if x1%y1 == 0 {
					s1 = fmt.Sprintf("%d + %d = \n", x/y, x1/y1)
				} else {
					s1 = fmt.Sprintf("%d + %d'%d/%d = \n", x/y, x1/y1, x1-(x1/y1)*y1, y1)
				}
			}
		}else{
			if x1 < y1 {
				s1 = fmt.Sprintf("%d'%d/%d + %d/%d = \n", x/y, x-(x/y)*y, y, x1, y1)
			} else {
				if x1%y1 == 0 {
					s1 = fmt.Sprintf("%d'%d/%d + %d = \n", x/y, x-(x/y)*y, y, x1/y1)
				} else {
					s1 = fmt.Sprintf("%d'%d/%d + %d'%d/%d = \n", x/y, x-(x/y)*y, y, x1/y1, x1-(x1/y1)*y1, y1)
				}
			}
		}
	}

	x2 := (x*y1+y*x1)/k
	y2 := (y*y1)/k
	if x2 < y2 {
		s2 = fmt.Sprintf("%d/%d\n", x2, y2)
	}else {
		if y2 == 1 {
			s2 = fmt.Sprintf("%d\n", x2)
		}else{
			s2 = fmt.Sprintf("%d'%d/%d\n", x2/y2, x2-(x2/y2)*y2, y2)
		}
	}

	return s1,s2
}

func Sub(x int,y int,x1 int,y1 int)(string,string) {
	var s1 string
	var s2 string

	if x*y1 == x1*y {
		s1 = fmt.Sprintf("%d - %d = \n", x/y, x1/y1)
		s2 = fmt.Sprintf("%d\n", 0)
		return s1,s2
	}
	if x*y1 < x1*y{
		var x2,y2 int
		x2 = x
		x = x1
		x1 = x2
		y2 = y
		y = y1
		y1 = y2
	}

	k := CommonDivisor(x*y1-y*x1, y*y1)
	if x < y {
		if x1 < y1 {
			s1 = fmt.Sprintf("%d/%d - %d/%d = \n", x, y, x1, y1)
		} else {
			if x1%y1 == 0 {
				s1 = fmt.Sprintf("%d/%d - %d = \n", x, y, x1/y1)
			} else {
				s1 = fmt.Sprintf("%d/%d - %d'%d/%d = \n", x, y, x1/y1, x1-(x1/y1)*y1, y1)
			}
		}
	} else {
		if x%y == 0 {
			if x1 < y1 {
				s1 = fmt.Sprintf("%d - %d/%d = \n", x/y, x1, y1)
			} else {
				if x1%y1 == 0 {
					s1 = fmt.Sprintf("%d - %d = \n", x/y, x1/y1)
				} else {
					s1 = fmt.Sprintf("%d - %d'%d/%d = \n", x/y, x1/y1, x1-(x1/y1)*y1, y1)
				}
			}
		} else {
			if x1 < y1 {
				s1 = fmt.Sprintf("%d'%d/%d - %d/%d = \n", x/y, x-(x/y)*y, y, x1, y1)
			} else {
				if x1%y1 == 0 {
					s1 = fmt.Sprintf("%d'%d/%d - %d = \n", x/y, x-(x/y)*y, y, x1/y1)
				} else {
					s1 = fmt.Sprintf("%d'%d/%d - %d'%d/%d = \n", x/y, x-(x/y)*y, y, x1/y1, x1-(x1/y1)*y1, y1)
				}
			}
		}
	}

	x2 := (x*y1-y*x1)/k
	y2 := (y*y1)/k
	if x2 < y2 {
		s2 = fmt.Sprintf("%d/%d\n", x2, y2)
	} else {
		if y2 == 1 {
			s2 = fmt.Sprintf("%d\n", x2)
		}else{
			s2 = fmt.Sprintf("%d'%d/%d\n", x2/y2, x2-(x2/y2)*y2, y2)
		}
	}

	return s1,s2
}

func Mul(x int,y int,x1 int,y1 int)(string,string){
	k := CommonDivisor(x*x1,y*y1)
	var s1 string
	var s2 string

	if x < y {
		if x1 < y1 {
			s1 = fmt.Sprintf("%d/%d × %d/%d = \n",x,y,x1,y1)
		}else{
			if x1 % y1 == 0{
				s1 = fmt.Sprintf("%d/%d × %d = \n",x,y,x1/y1)
			} else{
				s1 = fmt.Sprintf("%d/%d × %d'%d/%d = \n",x,y,x1/y1,x1-(x1/y1)*y1,y1)
			}
		}
	}else{
		if x % y == 0 {
			if x1 < y1 {
				s1 = fmt.Sprintf("%d × %d/%d = \n", x/y, x1, y1)
			} else {
				if x1%y1 == 0 {
					s1 = fmt.Sprintf("%d × %d = \n", x/y, x1/y1)
				} else {
					s1 = fmt.Sprintf("%d × %d'%d/%d = \n", x/y, x1/y1, x1-(x1/y1)*y1, y1)
				}
			}
		}else{
			if x1 < y1 {
				s1 = fmt.Sprintf("%d'%d/%d × %d/%d = \n", x/y, x-(x/y)*y, y, x1, y1)
			} else {
				if x1%y1 == 0 {
					s1 = fmt.Sprintf("%d'%d/%d × %d = \n", x/y, x-(x/y)*y, y, x1/y1)
				} else {
					s1 = fmt.Sprintf("%d'%d/%d × %d'%d/%d = \n", x/y, x-(x/y)*y, y, x1/y1, x1-(x1/y1)*y1, y1)
				}
			}
		}
	}

	x2 := (x*x1)/k
	y2 := (y*y1)/k
	if x2 < y2 {
		s2 = fmt.Sprintf("%d/%d\n", x2, y2)
	}else {
		if y2 == 1 {
			s2 = fmt.Sprintf("%d\n", x2)
		}else{
			s2 = fmt.Sprintf("%d'%d/%d\n", x2/y2, x2-(x2/y2)*y2, y2)
		}
	}

	return s1,s2
}

func Div(x int,y int,x1 int,y1 int)(string,string){
	var s1 string
	var s2 string
	k := CommonDivisor(x*y1,y*x1)

	if x < y {
		if x1 < y1 {
			s1 = fmt.Sprintf("%d/%d ÷ %d/%d = \n",x,y,x1,y1)
		}else{
			if x1 % y1 == 0{
				s1 = fmt.Sprintf("%d/%d ÷ %d = \n",x,y,x1/y1)
			} else{
				s1 = fmt.Sprintf("%d/%d ÷ %d'%d/%d = \n",x,y,x1/y1,x1-(x1/y1)*y1,y1)
			}
		}
	}else{
		if x % y == 0 {
			if x1 < y1 {
				s1 = fmt.Sprintf("%d ÷ %d/%d = \n", x/y, x1, y1)
			} else {
				if x1%y1 == 0 {
					s1 = fmt.Sprintf("%d ÷ %d = \n", x/y, x1/y1)
				} else {
					s1 = fmt.Sprintf("%d ÷ %d'%d/%d = \n", x/y, x1/y1, x1-(x1/y1)*y1, y1)
				}
			}
		}else{
			if x1 < y1 {
				s1 = fmt.Sprintf("%d'%d/%d ÷ %d/%d = \n", x/y, x-(x/y)*y, y, x1, y1)
			} else {
				if x1%y1 == 0 {
					s1 = fmt.Sprintf("%d'%d/%d ÷ %d = \n", x/y, x-(x/y)*y, y, x1/y1)
				} else {
					s1 = fmt.Sprintf("%d'%d/%d ÷ %d'%d/%d = \n", x/y, x-(x/y)*y, y, x1/y1, x1-(x1/y1)*y1, y1)
				}
			}
		}
	}

	x2 := (x*y1)/k
	y2 := (y*x1)/k
	if x2 < y2 {
		s2 = fmt.Sprintf("%d/%d\n", x2, y2)
	}else {
		if y2 == 1 {
			s2 = fmt.Sprintf("%d\n", x2)
		}else{
			s2 = fmt.Sprintf("%d'%d/%d\n", x2/y2, x2-(x2/y2)*y2, y2)
		}
	}

	return s1,s2
}
//返回x和y的最大公约数
func CommonDivisor(x int,y int)int{
	for{
		if x == y{
			return x
		}
		if x > y{
			x = x - y
		}else{
			y = y - x
		}

		if x != y{
			continue
		}else{
			return x
		}
	}
}






