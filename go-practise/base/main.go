package main

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
)

var (
	c = []int{1, 2, 4, 5}
	p = []int{2, 4, 1, 5, 7, 9, 6, 3, 8, 0}
)

// 切片练习
func main() {
	fmt.Println("***********************************\n切片练习\n***********************************")
	fmt.Println("初始定义:", c)
	c = append(c, 6)
	fmt.Println("追加一个元素:", c)
	c = append(c, 0)
	fmt.Println("拓展一个空间:", c)
	copy(c[3:], c[2:len(c)-1]) //移位
	c[2] = 3                   //插入数据
	fmt.Println("移位 插入数据:", c)

	fmt.Println("排序前:", p)
	sort.Ints(p)
	fmt.Println("排序后:", p)

	path := []byte("AAAAA/BBBBB")
	fmt.Println("初始数据:", string(path))

	sepIndex := bytes.IndexByte(path, '/')
	tmp1 := path[:sepIndex]
	tmp2 := path[sepIndex+1:]
	fmt.Println("分割之后的数据:", string(tmp1), string(tmp2))
	tmp1 = append(tmp1, "suffix"...)
	fmt.Println("append之后的数据:", string(tmp1), string(tmp2)) //tmp1和tmp2共享内存,,因为cap足够,没有重新分配内存,所以导致新加数据拓展到tmp2上面去了

	demo := []byte("AAAAA/BBBBB")
	fmt.Println("初始数据:", string(demo)) //避免共享内存导致的数据覆盖,可设置切片的容量,最后一位的参数为Limited Capacity,后续的append就会重新分配内存
	tmp3 := demo[:sepIndex:sepIndex]
	tmp4 := demo[sepIndex+1:]
	fmt.Println("分割之后的数据:", string(tmp3), string(tmp4))
	tmp3 = append(tmp3, "suffix"...)
	fmt.Println("设置cap容量append之后的数据:", string(tmp3), string(tmp4))

	fmt.Println("***********************************\n深度比较\n***********************************")

	demo1 := map[string]string{"你好": "hello", "再见": "bye"} //深度比较:反射对比,可以用在数组，结构体，map……的内容比较中
	demo2 := map[string]string{"再见": "bye", "你好": "hello"}
	fmt.Println("map值:", demo1, demo2)
	fmt.Println("demo1 == demo2:", reflect.DeepEqual(demo1, demo2))

	fmt.Println("***********************************\n接口编程\n***********************************")

	var d = Demo{
		Name: "lilei",
		Age:  18,
	}
	PrintInfo(&d) //虽然输出内容相同,但是下面使用了Receiver
	d.Print()

	d1 := Country{"USA"}      //使用了一个叫Stringable 的接口，我们用这个接口把“业务类型” Country 和 City 和“控制逻辑” Print() 给解耦了。
	d2 := City{"Los Angeles"} //只要实现了Stringable 接口，都可以传给 PrintStr() 来使用。
	PrintStr(d1)              //面向对象编程方法的黄金法则——“Program to an interface not an implementation”
	PrintStr(d2)

	/*
		性能优化相关

		1. 如果需要把数字转字符串，使用 strconv.Itoa() 会比 fmt.Sprintf() 要快一倍左右
		2. 尽可能地避免把String转成[]Byte 。这个转换会导致性能下降。
		3. 如果在for-loop里对某个slice 使用 append()请先把 slice的容量很扩充到位，这样可以避免内存重新分享以及系统自动按2的N次方幂进行扩展但又用不到，从而浪费内存。
		4. 使用StringBuffer 或是StringBuild 来拼接字符串，会比使用 + 或 += 性能高三到四个数量级。
		5. 尽可能的使用并发的 go routine，然后使用 sync.WaitGroup 来同步分片操作
		6. 避免在热代码中进行内存分配，这样会导致gc很忙。尽可能的使用 sync.Pool 来重用对象。
		7. 使用 I/O缓冲，I/O是个非常非常慢的操作，使用 bufio.NewWrite() 和 bufio.NewReader() 可以带来更高的性能。
		8. 对于在for-loop里的固定的正则表达式，一定要使用 regexp.Compile() 编译正则表达式。性能会得升两个数量级。
		9. 如果你需要更高性能的协议，你要考虑使用 protobuf 或 msgp 而不是JSON，因为JSON的序列化和反序列化里使用了反射。
		10.你在使用map的时候，使用整型的key会比字符串的要快，因为整型比较比字符串比较要快。
	*/
}

type Demo struct {
	Name string
	Age  int
}

func PrintInfo(d *Demo) {
	fmt.Printf("Name=%s,Age=%d\n", d.Name, d.Age)
}

func (d *Demo) Print() {
	fmt.Printf("Name=%s,Age=%d\n", d.Name, d.Age)
}

type Country struct {
	Name string
}
type City struct {
	Name string
}
type Stringable interface {
	ToString() string
}

func (c Country) ToString() string {
	return "Country = " + c.Name
}
func (c City) ToString() string {
	return "City = " + c.Name
}
func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}
