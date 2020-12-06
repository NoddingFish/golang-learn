package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
	"strconv"
	"strings"
)

func main() {
	//TODO　string
	str := "hello背"
	// len 按字节，字母和数字占一个字节，汉字占用3个字节
	fmt.Println("str len=", len(str))

	//! 字符串遍历，同时处理有中文的问题 str22:=[]rune(str2)
	// str2 := "hello 世界"
	// for i := 0; i < len(str2); i++ {
	// 	fmt.Println("str2 len=", str2[i])
	// }

	// str2 := "hello 世界"
	// for i := 0; i < len(str2); i++ {
	// 	fmt.Printf("str2 len=%c \n", str2[i])
	// }

	str2 := "hello 世界"
	str22 := []rune(str2)
	for i := 0; i < len(str22); i++ {
		fmt.Printf("str2 len=%c \n", str22[i])
	}

	//TODO 3、字符串转整数 strconv.Atoi
	num3, err := strconv.Atoi("123")
	// num3, err := strconv.Atoi("123a")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换结果：", num3)
	}

	//TODO 4、整数转字符串
	str4 := strconv.Itoa(12378)
	fmt.Printf("转换结果：%v, 类型是：%T", str4, str4)

	//TODO 5、字符串转 []byte
	var bytes = []byte("hello go")
	fmt.Printf("转换结果：%c \n %v \n", bytes, bytes)

	//TODO 6、[]byte 转字符串
	var str6 = string([]byte{97, 98, 99})
	fmt.Printf("转换结果：%v \n", str6)

	//TODO 7、10进制转 2,8,16进制
	str71 := strconv.FormatInt(123, 2)
	fmt.Println("10进制转2进制结果:", str71)
	str72 := strconv.FormatInt(123, 8)
	fmt.Println("10进制转8进制结果:", str72)
	str73 := strconv.FormatInt(123, 16)
	fmt.Println("10进制转16进制结果:", str73)

	//TODO 8、查找字符串
	bool8 := strings.Contains("hello world", "wo1")
	fmt.Println("查找字符串的结果:", bool8)

	//TODO 9、统计一个字符串有几个指定的子串
	num9 := strings.Count("changese", "e")
	fmt.Println("统计一个字符串有几个指定的子串的结果:", num9)

	//TODO 10、不区分大小写比较 , == 是区分大小写的
	fmt.Println("不区分大小写比较：", strings.EqualFold("abc", "ABC"));
	fmt.Println("区分大小写比较：", "abc" == "ABC");

	//TODO 11、返回子串在字符串第一次出现的 index 值，如果没有，返回 -1
	num11 := strings.Index("Test_abcabcabc", "abc")
	fmt.Println("返回子串在字符串第一次出现的 index 值：", num11);

	//TODO 12、返回子串在字符串最后一次出现的 index 值，如果没有，返回 -1
	num12 := strings.LastIndex("Test_abcabcabc", "abc")
	fmt.Println("返回子串在字符串最后一次出现的 index 值：", num12);

	//TODO 13、将指定的子串替换成另一个子串， n 数字代表替换几个，1-一个 -1 代表所有
	fmt.Println("将指定的子串替换成另一个子串：", strings.Replace("go go go go hello", "go", "go语言", 2));
	fmt.Println("将指定的子串替换成另一个子串：", strings.Replace("go go go go hello", "go", "go语言", -1));

	//TODO 14、将字符串按照指定的子串，拆分成数组
	str14 := "hello,world,ok"
	strArr14 := strings.Split(str14, ",")
	for i := 0; i < len(strArr14); i++ {
		fmt.Printf("数组：%v ：%v \n", i, strArr14[i])
	}
	fmt.Printf("将字符串按照指定的子串，拆分成数组： %v\n", strArr14)

	//TODO 15、将字符串的字母进行大小写转换
	fmt.Println("将字符串的字母进行大小写转换：", strings.ToLower("goLang Hello"))
	fmt.Println("将字符串的字母进行大小写转换：", strings.ToUpper("goLang Hello"))

	//TODO 16、将字符串左右两边的空格去掉
	fmt.Printf("将字符串左右两边的空格去掉：%q \n", strings.TrimSpace("    goLang Hello "))

	//TODO 17、将字符串左右两边的指定的字符串去掉
	fmt.Printf("将字符串左右两边的空格去掉：%q \n", strings.Trim(" $$ ,,1, goLang Hello ,, ", " ,$"))

	
	//TODO 18\19、将字符串左或右两边的指定的字符串去掉
	fmt.Printf("将字符串左或右两边的指定的字符串去掉：%q \n", strings.TrimLeft(" $$ ,,1, goLang Hello ,, ", " ,$"))
	fmt.Printf("将字符串左或右两边的指定的字符串去掉：%q \n", strings.TrimRight(" $$ ,,1, goLang Hello ,, ", " ,$"))

	//TODO 20\21、判断字符串是否已指定的字符串开头或结尾
	fmt.Printf("判断字符串是否已指定的字符串开头或结尾：%v \n", strings.HasPrefix("ftp://192.168.10.1", "ftp"))
	fmt.Printf("判断字符串是否已指定的字符串开头或结尾：%v \n", strings.HasPrefix("ftp://192.168.10.1", "fix"))
	fmt.Printf("判断字符串是否已指定的字符串开头或结尾：%v \n", strings.HasSuffix("ftp://192.168.10.1/main.go", ".go"))

}