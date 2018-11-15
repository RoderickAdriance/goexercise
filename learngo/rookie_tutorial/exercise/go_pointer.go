package exercise

import "fmt"

func Po() {
	//一个指针变量指向了一个值的内存地址。
	var ip *int     /* 指向整型*/
	var fp *float64 /* 指向浮点型 */
	println(ip, fp)
}

/*指针使用流程：
定义指针变量。
为指针变量赋值。
访问指针变量中指向地址的值。*/
func PointerDemo() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */
	ip = &a        /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	fmt.Printf(" *ip 变量的值: %d\n", *ip)
}

func P2() {
	const MAX int = 3
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, ptr[i]) //输出的为地址符
	}
}

//指向指针的指针
//当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址：
func P3() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}

func P4() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值 : %d\n", a)
	fmt.Printf("交换前 b 的值 : %d\n", b)

	/* 调用函数用于交换值
	* &a 指向 a 变量的地址
	* &b 指向 b 变量的地址
	 */
	swap(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
