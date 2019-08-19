package main

import (
	"fmt"
	"sort"
)

func main()  {
	//1.定义一个key打乱的字典
	games := map[int]string{3:"王者农药",1:"绝地求生",2:"连连看",4:"传奇霸业",5:"消消乐"}
	//2.定义一个切片
	s := make([]int,0,len(games))
	//3.遍历map获取key-->s1中
	for key := range games{
		s = append(s, key)
	}
	//4.给s进行排序
	sort.Ints(s)
	//4. 遍历s 来读取 games
	for _,k:=range s{ // 先下标，再数值
		fmt.Println(k, games[k])
	}
}
