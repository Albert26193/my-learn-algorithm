package main

import "fmt"

/*
1、数字为0
2、数字大于等于10
3、数字中间有多个0
4、个位数为0
*/

func toCN(num int) string {
	//1、数字为0
	if num == 0 {
		return "零"
	}
	var ans string
	//数字
	szdw := []string{"十", "百", "千", "万", "十万", "百万", "千万", "亿"}
	//数字单位
	sz := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	res := make([]string, 0)

	//数字单位角标
	idx := -1
	for num > 0 {
		//当前位数的值
		x := num % 10
		//2、数字大于等于10
		// 插入数字单位，只有当数字单位角标在范围内，且当前数字不为0 时才有效
		if idx >= 0 && idx < len(szdw) && x != 0 {
			res = append([]string{szdw[idx]}, res...)
		}
		//3、数字中间有多个0
		// 当前数字为0，且后一位也为0 时，为避免重复删除一个零文字
		if x == 0 && len(res) != 0 && res[0] == "零" {
			res = res[1:]
		}
		// 插入数字文字
		res = append([]string{sz[x]}, res...)
		num /= 10
		idx++
	}
	//4、个位数为0
	if len(res) > 1 && res[len(res)-1] == "零" {
		res = res[:len(res)-1]
	}
	//合并字符串
	for i := 0; i < len(res); i++ {
		ans = ans + res[i]
	}
	return ans
}

func main() {
	input := make([]int, 0)
	input = append(input, 8888)
	input = append(input, 100024)
	input = append(input, 800008888)
	fmt.Println(toCN(input[0]))
	fmt.Println(toCN(input[1]))
	fmt.Println(toCN(input[2]))
}
