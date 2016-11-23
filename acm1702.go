/*
【问题描述】
Eva有20个重量分别为1,3,9,27,...,3^19的砝码。Eva觉得用这20个砝码可以称量任何重量在1到(3^20-1)/2之间的物体的重量。
假设Eva在天平的左边托盘中放上一个物体，你的任务是在两边托盘中放入适当的砝码，称量出物体的重量。

【输入数据】
第一行含有一个整数T(1 <= T <= 20)，表示测试案例的个数。
随后的 T 行每一行包含一个整数 W (1 <= W <= (3^20-1)/2),表示一个物品的重量。

【输出数据】
对于每个测试案例输出一行，表示应该放在左边和右边托盘中的砝码的重量。
用空格分隔左边托盘和右边托盘中的砝码。同一边托盘中的砝码按升序排列，并且用逗号分隔各个砝码。
如果某边盘中没有砝码，则输出empty.

【输入样例】
3
9
5
20

【输出样例】
empty 9
1,3 9
1,9 3,27
*/
package main

import "fmt"

func weight(val int) (left, right []int) {

	cur := 1

	// 借鉴十进制转二进制时的辗转相除法
	for val > 0 {
		leftadd := 0
		switch rest := val % 3; rest {
		case 0:
		case 1:
			right = append(right, cur)
		case 2: // 余数为2时商增加1,使得余数变成负1,负1表示放在左边
			left = append(left, cur)
			leftadd = 1
		}
		val = val/3 + leftadd
		cur *= 3
	}

	return
}

func demo(val int) {

	left, right := weight(val)

	str := fmt.Sprintf("[%v]", val)

	for _, v := range left {
		str += fmt.Sprintf(" + %v", v)
	}

	for i, v := range right {
		if i == 0 {
			str += fmt.Sprintf(" = %v", v)
		} else {
			str += fmt.Sprintf(" + %v", v)
		}
	}

	fmt.Println(str)
}

func main() {

	for i := 1; i < 30; i++ {
		demo(i)
	}
}
