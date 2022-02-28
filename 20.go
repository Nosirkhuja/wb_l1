package main

import "fmt"

func ReverseWord(str string) string {
	var arr []string
	var tmp string = ""
	for i := 0; i < len(str); i++ {
		if string(str[i]) != " " {
			tmp = tmp + string(str[i])
		} else {
			arr = append(arr, tmp)
			tmp = ""
		}
	}
	if tmp != " " {
		arr = append(arr, tmp)
	}
	tmp = ""
	k := len(arr)
	for i := 0; i < k; i++ {
		tmp = tmp + arr[k-i-1] + " "
	}
	return tmp
}

func main() {

	s1 := "The hour of the samurai is near"
	fmt.Println(s1)
	fmt.Println(ReverseWord(s1))

}
