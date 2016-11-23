/*
问题的提出：
很多问题的解决涉及到高精度的数学计算。
有时候，高级语言提供的浮点数类型(float,double等）的精度不满足实际应用的需要，
所以很有必要用其它方法实现高精度的计算。
此问题要求编写一个程序，准确计算一个实数R(0.0 < R < 99.999)的n（0 < n <= 25)次方。

输入数据：输入由一系列R和n值组成，R占据第1到6列，n占据8和9列，第7列是空格。

输出数据：对于每一行输入数据，输出R的n次方的准确值。
不得输出前导零，也不得输出无意义的末尾的零；如果结果是整数，不得输出小数点。
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	strpow("6.4", 2)
	strpow("3.7", 8)
}

func strpow(str string, n int) {

	var ps int

	for _, digit := range str {

		if digit == '.' {
			ps = 0
		} else if ps >= 0 {
			ps++
		}
	}

	result := pow(strings.Replace(str, ".", "", -1), int64(n))
	result = reverse(result, ps*n)

	fmt.Printf("pow(%v,%v)=%v\n", str, n, result)
}

// 求 val 的 n 次幂
func pow(val string, n int64) string {

	op := reverse(val, 0)
	result := op

	for i := int64(0); i < n-1; i++ {
		/*
			当前结果的每一位与操作数相乘
					  操作数:   6 4
				×	当前结果: 2 5 6
				---------------------
				              3 8 4
			                3 2 0
						  1 2 8
				---------------------
						  1 6 3 8 4
		*/
		tmpres := ""

		for idx, digit := range result {

			tmpstr := ""
			carry := 0

			/*
				   6 4
				×   6
				------
				 3 8 4
			*/
			for _, char := range op {
				num := int(digit-'0')*int(char-'0') + carry // 两数字相乘,加上进位
				tmpstr += string(rune(num%10) + '0')        // 颠倒表示
				carry = num / 10                            // 进位
			}

			if carry > 0 {
				tmpstr += string(rune(carry) + '0')
			}

			/*
				移位累加
			*/
			if idx == 0 {
				tmpres = tmpstr
			} else {
				tmpres = progression(tmpres, tmpstr, idx)
			}
		}

		result = tmpres
	}

	return result
}

func progression(op1, op2 string, pos int) (result string) {

	var carry int

	max1 := len(op1)
	max2 := len(op2)
	idx1 := pos
	idx2 := 0

	result = op1

	// 任何一个操作数没有处理完所有数位,或者有进位
	for idx1 < max1 || idx2 < max2 || carry > 0 {

		// bit = 进位 + op1 + op2
		bit := carry

		if idx1 < max1 {
			bit += int(op1[idx1] - '0')
		}

		if idx2 < max2 {
			bit += int(op2[idx2] - '0')
		}

		// 更新数位
		if idx1 < max1 {
			result = result[:idx1] + string(rune(bit%10)+'0') + result[idx1+1:]
		} else {
			result += string(rune(bit%10) + '0')
		}

		// 进位
		carry = bit / 10

		idx1++
		idx2++
	}

	return
}

func reverse(str string, pos int) (ret string) {

	for idx, char := range str {

		if len(ret) == 0 && char == '0' && pos > 0 {
			continue
		} else {
			ret = string(char) + ret
		}

		if idx+1 == pos && len(ret) > 0 {
			ret = "." + ret
		}
	}
	return
}
