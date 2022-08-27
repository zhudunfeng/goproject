package main

/*
1)编程创建一个Box结构体，在其中声明三个字段表示一个立方体的长、宽和高，长宽高要从终端获取
2)声明一个方法获取立方体的体积。
3)创建一个Box结构体变量，打印给定尺寸的立方体的体积
*/
type Box struct {
	len    float64
	width  float64
	height float64
}

func (box *Box) getVolumn() float64 {
	return box.len * box.width * box.height
}
