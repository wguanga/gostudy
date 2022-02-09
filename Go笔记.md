# Go

## 基本语法

### 指针

>类型指针不能进行偏移和运算。
>
>&：取地址
>
>*：根据地址取值
>
>

#### 普通指针定义

```go
    var b *int
    a:=10
    b=&a
    //或者
    a := 10
	b := &a//b指向a
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)
```

输出结果：

```txt
a: 10
b: 0xc000014088
a: int
b: *int
```

c语言定义指针对比

```c
    int a = 10;
    int *b = &a;//b指向a
```

#### 指向数组的指针

>
>
>```go
>	//指向数组的指针定义
>	var array_addr []*int
>	array_1 := []int{10, 20}
>	array_addr = make([]*int, len(array_1))
>	for i := 0; i < len(array_1); i++ {
>		array_addr[i] = &array_1[i]
>	}
>
>	for i := 0; i < len(array_addr); i++ {
>		fmt.Printf("*array_addr[%v]: %v\n", i, *array_addr[i])
>	}
>```
>
>输出结果：
>
>```
>*array_addr[0]: 10
>*array_addr[1]: 20
>```
>
>

### 特别数据类型

#### 数组

数组定义

```go
//数组定义
	array_1 := [2]int{1, 2}
	fmt.Printf("array_1: %T\n", array_1)//输出array_1的数据类型
	fmt.Printf("array_1[1]: %T\n", array_1[1])//输出array_1数组第一个元素的类型
    fmt.Printf("array_1: %v\n", array_1)//输出array_1数组所有元素
	fmt.Printf("array_1[1]: %v\n", array_1[1])//输出array_1数组第一个元素
```

>
>
>根据索引进行初始化：
>
>```go
>	array_2 := []string{0: "sdf", 5: "trt"}
>	fmt.Printf("array_2: %v\n", array_2)
>```
>
>输出结果：
>
>```
>array_2: [sdf     trt]
>```
>
>

[注]:数组长度是固定的，无法使用append()函数扩容。

#### 切片

理解为动态数组,可以自动扩容

```go
//切片定义
	array_2 := []int{1, 2, 3}
	fmt.Printf("array_2: %T\n", array_2)//输出array_2的数据类型
```

初始化

>
>
>1.数组进行初始化的方法同样可用于切片
>
>2.除此之外，还可以用数组进行初始化,可以理解为，数组切片形成了切片
>
>>```go
>>	//用数组对切片进行初始化
>>	slice_1 := array_2[:] //前后都没有数表示取所有元素，前闭后开
>>	fmt.Printf("slice_1: %T\n", slice_1)
>>```
>>
>>

切片是引用类型，可以使用make函数来创建切片

>
>
>var slice1 []type = make([]type,len)
>
>可简写为：
>
>slice1 := make([type,len])

增加与删除

```go
	//切片增加元素
	slice_1 = append(slice_1, "100")
	fmt.Printf("slice_1: %v\n", slice_1)
	//切片删除索引为i的元素
	i := 6
	slice_1 = append(slice_1[:i])
```

#### 字典

> 1.定义
>
> ```go
> var map_1 map[string]string//前键后值
> ```
>
> 2.初始化
>
> ```go
> 	map_1 = make(map[string]string, 5)//使用make函数，长度为5，可以不设长度
> 	map_1["name"] = "guang"
> 	map_1["age"] = "22"
> ```
>
> 简化，直接创建
>
> ```go
> //方法一
> 	map_2 := map[string]string{"name": "guang", "age": "22"}
> ```
>
> ```go
> //方法二
> 	var map_3 = map[string]string{"name": "guang", "age": "22"}
> ```
>
> 



### 变量类型

```go
    //变量数据类型
	var a1 int8 =10//有符号整型
	var a2 int16=10
	var a3 int32=10//32位系统上int默认是int32
	var a4 int64=10//64位系统上int默认是int64
	var b1 uint8=10//无符号整型
	var b2 uint16=10
	var b3 uint32=10//同上
	var b4 uint64=10//同上
	var c1 float32=10//相当于c中的float
	var c2 float64=10//相当于c中的double
```

### 进制输出

```go
    //进制输出
	output := 10
	fmt.Printf("output: %d\n", output)//十进制输出
	fmt.Printf("output: %b\n", output)//二进制输出
	fmt.Printf("output: %o\n", output)//八进制输出
	fmt.Printf("output: %x\n", output)//十六进制输出
```

### 字符串

#### 字符串定义

~~~go
	//字符串定义
	var ss1 string = "abcdefg"
	var ss2 = "abcdefg"
	ss3 := "abcdefg"
	//字符串连接
	ss4 := ss1 + ss2                             //方法1
	ss5 := fmt.Sprint("111%s 222%s", ss1, ss2)   //方法2
	ss6 := strings.Join([]string{ss1, ss2}, ",") //方法3
	longss := `
	line 1
	line 2
	line 3
	`
~~~

#### 字符串切片

```go
	//字符串切片
	sl1 := "playground"
	ord := 3
	fmt.Printf("sl1[ord]: %c\n", sl1[ord])
	fmt.Printf("%v\n", sl1[0:4]) //前闭后开[,),此处等价于sl1[:4]
```

输出结果：

```
sl1[ord]: y
play
```

#### 常用字符串函数

##### len()

输出字符串长度

```go
	fmt.Printf("len(sl1): %v\n", len(sl1))
```

##### +或者fmt.Sprint

连接字符串，详见[字符串定义](#字符串定义)

##### strings.Split()

分割字符串

```go
	fmt.Printf("strings.Split(sl1, \" \"): %v\n", strings.Split(sl1, "g")) //该函数返回一个数组
```

输出结果：

```
strings.Split(sl1, " "): [play round]
```

##### strings.contains()

判断原字符串是否包含提供的字符串

````go
	fmt.Printf("strings.Contains(sl1, \"lay\"): %v\n", strings.Contains(sl1, "lay")) //该函数返回bool型
````

输出结果：

```
strings.Contains(sl1, "lay"): true
```

##### strings.ToLower()与strings.ToUpper()

字符串转换大小写

```go
	fmt.Printf("strings.ToLower(sl1): %v\n", strings.ToLower(sl1))//变小写
	fmt.Printf("strings.ToUpper(sl1): %v\n", strings.ToUpper(sl1))//变大写
```

##### strings.HasPrefix()与strings.HasSuffix()

判断原字符串开头和结尾是否为所提供的字符串

```go
	fmt.Printf("strings.HasPrefix(sl1, \"play\"): %v\n", strings.HasPrefix(sl1, "play"))     //判开头
	fmt.Printf("strings.HasSuffix(sl1, \"Ground\"): %v\n", strings.HasSuffix(sl1, "Ground")) //判结尾
```

返回bool型

##### strings.Index()

返回提供的字符串的首字母在原字符串中的位置，返回第一个匹配值，没有则返回-1

strings.LastIndex()

返回最后一个strings.Index()的匹配值

```go
	fmt.Printf("strings.Index(sl1, \"oun\"): %v\n", strings.Index(sl1, "oun"))//oun的第一个字母在原字符串位置
	fmt.Printf("strings.Index(sl1, \"oo\"): %v\n", strings.Index(sl1, "oo"))
```

输出结果:

```
strings.Index(sl1, "oun"): 6
strings.Index(sl1, "oo"): -1
```

### 格式化输入输出

格式化输入

```go
	var name string
	var age int 
	fmt.Scan(&name,&age)
```

格式化输出

```go
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
```

输出占位符

```go
%v  var                                    %#v    输出该变量的所在路径格式和内容
%c  字符
%T  Type  输出类型
%t  bool

%b  二进制输出
%d  十进制
%o  二进制
%x  十六进制
%p  指针
```

### 条件判断

#### if else

1.条件不用加()

2.{}**必须**存在，即使只有一条语句，且左大括号必须与if或else在**同一行**

3.go中的if后的条件中，可以初始化变量，但是仅在该条件块中生效,效果如下：

```go
	if flagg := false; !flagg {
		fmt.Printf("flagg: %v\n", flagg)
	}
```

#### switch

````go
	switch b := 100; b {
	case 10, 20, 30:
		fmt.Println("1")
	case 50, 100:
		fmt.Println("2")

	}
````



```
	switch {
	case a > 80:
		fmt.Println("优秀")
	case a > 70:
		fmt.Println("良好")
	case a > 60:
		fmt.Println("及格")
	case a > 0:
		fmt.Println("不及格")
	default:
		fmt.Println("default")
	}
```

### 循环

#### for

>要求与go语言中的if类似，用法与c语言类似，此处不再赘述。
>
>[注]:go中的for的条件初始化和运算都可以省略，类似c中的while，也可以直接全部省略，无限循环





#### for range()

>
>
>通过range返回的值
>
>1.数组、切片、字符串返回**索引**和**值**
>
>2.map返回**键**和**值**
>
>3.通道(channel)只返回通道内的值
>
>
>
>```go
>	//for range
>	array_1 := []int{10, 20, 30}
>	fmt.Printf("array_1: %v\n", array_1)
>	for i, v := range array_1 {
>		fmt.Printf("index: %v  ", i)
>		fmt.Printf("value: %v\n", v)
>	} //i为索引，v为值
>	for _, v := range array_1 {
>		fmt.Printf("v: %v\n", v)
>	} //只要值不要索引
>```
>
>



### 函数

#### 普通函数

>定义：
>
>```go
>//无参有返回
>func func_1() {
>	fmt.Println("无参无返回")
>}
>//有参有返回
>func func_2(a int, b int) (c int, d int) {
>	return a + b, a - b
>}
>//不定长参数列表
>func func_3(args ...int) {
>	for _, v := range args {
>		fmt.Printf("v: %v\n", v)
>	}
>}
>```
>
>
>
>同一文件调用：
>
>```go
>	func_1()
>	c, d := func_2(4, 6)
>	_, f := func_2(4, 6) //可以用下划线丢弃返回值
>    func_3(1,2,3)
>    func_3(4,5,6,7,8)
>```
>
>函数传参：普通变量传参是按值传参。map、slice、interface、channel这些变量传参时，因其本身就是指针类型，所以传参时相当于传输的地址，在函数内改变时在函数外也会改变。

#### 定义一个函数类型和使用函数变量

>```go
>//定义commfun为函数类型
>type commfun func(int, int) int //也可以定义在函数内
>
>func sum(a int, b int) int {
>	return a + b
>}
>
>func dec(a int, b int) int {
>	return a - b
>}
>```
>
>使用：
>
>```go
>	var f1, f2 commfun
>	f1 = sum
>	fmt.Printf("f1(1, 2): %v\n", f1(1, 2))
>	f2 = dec
>	fmt.Printf("f2(1, 2): %v\n", f2(1, 2))
>```
>
>



#### 函数作为参数和返回值

>
>
>```go
>//工具函数
>func pr(name string) {
>	fmt.Printf("hello, %v !\n", name)
>}
>
>//函数作为参数和返回值
>func stringjoin(ss string, f func(string))func(string) {
>	f(ss)
>	return pr
>}
>```
>
>

#### 匿名函数

>
>
>go语言不允许函数嵌套，即在函数内部定义新的普通函数，但允许在函数内部定义匿名函数来实现简单的功能调用
>
>```go
>func main() {
>	//定义匿名函数
>	max := func(a int, b int) int {
>		if a >= b {
>			return a
>		} else {
>			return b
>		}
>	}
>	abmax := max(1, 2)
>	fmt.Printf("abmax: %v\n", abmax)
>}
>```
>
>

#### 闭包

>闭包可以理解为一个函数定义在另一个函数体内，是函数内外连接的桥梁。
>
>```go
>//闭包
>func closure_addd() func(int) int {
>	var x int
>	return func(y int) int {
>		x += y
>		return x
>	}
>}
>
>//使用闭包
>func main() {
>    cf := closure_addd()
>	fmt.Printf("cf(10): %v\n", cf(10))
>	fmt.Printf("cf(10): %v\n", cf(10))
>}
>```
>
>输出结果：
>
>```
>cf(10): 10
>cf(10): 20
>```
>
>可以看出，在原函数内，闭包外定义的变量会一直存在
>
>

#### 方法

>
>
>go中没有对象，所以也没有类中的方法，但存在可以定义于struct之上的方法。
>
>1.方法可以调用结构体中的变量。
>
>2.方法的接受对象也可以是数组，切片，指针，map，channel，func等。
>
>3.结构体中的方法相当于面向对象语言中类的方法，但是go语言中结构体的方法定义可以不在同一个文件下，但一定要在同一个包下。
>
>4.go语言中没有方法重载，但对于不同的接收者可以写名字相同的方法，类似方法重写。
>
>5.如果接受者是一个指针类型，则会自动解除引用。
>
>6.方法和type是分开的，通过接收者建立连接，接收者相当于对象。
>
>
>
>定义方法：
>
>type mystruct struct{}
>
>func (recv mystruct) my_method(para) return_type{}
>
>```go
>//定义结构体
>type struct_1 struct {
>	value int
>	next  *struct_1
>}
>
>
>
>//定义方法
>func (str_1 struct_1) my_mythod(a int) {
>	fmt.Printf("a: %v\n", a)
>	fmt.Printf("str_1.value: %v\n", str_1.value)
>}
>
>func main() {
>	var a struct_1
>	a = struct_1{
>		value: 11,
>	}
>	a.my_mythod(10)
>
>}
>```
>
>输出结果：
>
>```go
>a: 10
>str_1.value: 11
>```
>
>------
>
>指针做接收者
>
>会改变原来变量中的内容，且会自动解引用
>
>```go
>//接收者为值变量
>func (str_1 struct_1) my_mythod(a int) {
>	fmt.Printf("a: %v\n", a)
>	str_1.value += a
>}
>
>
>//接收者为指针变量
>func (p_str1 *struct_1) myp_mythod(a int) {
>	p_str1.value += a
>}
>```
>
>```go
>func main() {
>	var a struct_1
>	a = struct_1{
>		value: 11,
>	}
>	fmt.Printf("a.value: %v\n", a.value)
>
>	a.my_mythod(10) //接收者为值变量
>	fmt.Printf("a.value: %v\n", a.value)
>
>	a.myp_mythod(10) //接收者为指针变量，自动解引用
>	fmt.Printf("a.value: %v\n", a.value)
>
>}
>```
>
>输出结果：
>
>```
>a.value: 11
>a: 10
>a.value: 11
>a.value: 21
>```
>
>



### 特殊语句

#### defer

>defer语句会将后面跟随的语句进行延迟处理，在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行
>
>------------
>
>可用于锁资源释放，数据库连接释放等

````go
	fmt.Println("111")
	defer fmt.Println("222")
	defer fmt.Println("333")
	defer fmt.Println("444")
````

输出结果：

```
111
444
333
222
```

#### init

>1.init函数先于main函数执行，实现包级别的一些初始化操作，不能被其他函数调用。
>
>2.没有参数和返回值。
>
>3.每个包可以有多个init函数，包的每个源文件里也可以有多个init函数。
>
>4.同一个包内的init函数的执行顺序没有定义，所以定义时不要依赖于顺序。不同包内的init函数按包导入的顺序决定顺序。
>
> 
>
>[ 注 ]:golang的初始化顺序为：全局变量->init函数->main函数

### 类型定义与类型别名

>
>
>前者定义了一个新的类型，后者并没有定义新的类型，只是起了个别名。
>
>后者只会在代码中存在，编译完成后不存在。
>
>后者与原来的类型是完全相同的，所以可以使用原来的类型拥有的方法，而前者不能使用其方法。
>
>```go
>	//类型定义
>	type myint int
>	var a myint
>	a = 10
>	//类型别名
>	type myint2 = int
>	var b myint2
>	b = 10
>```
>
>

### 结构体

>
>
>go中的结构体类似于c中的结构体与java中的类的简化的结合。
>
>定义：
>
>```go
>type struct_1 struct {
>	value int
>	next  *struct_1
>}
>
>```
>
>
>
>调用：
>
>```go
>func main() {
>	var a struct_1    //定义a结点
>	var p_a *struct_1 //a结点的指针
>	p_a = &a          //连接a与p_a
>	a.value = 10
>	var b struct_1    //定义b结点
>	var p_b *struct_1 //b结点的指针
>	b.value = 20
>	p_b = &b     //连接b与p_b
>	a.next = p_b //b为a的下一个
>	c := a.next  //c=p_b
>	d := *c      //d=b
>	fmt.Printf("a: %v\n", a)
>	fmt.Printf("*p_a: %v\n", *p_a)
>	fmt.Printf("b: %v\n", b)
>	fmt.Printf("*p_b: %v\n", *p_b)
>	fmt.Printf("d: %v\n", d)
>}
>```
>
>输出结果：
>
>```go
>a: {10 0xc00003c230}
>*p_a: {10 0xc00003c230}
>b: {20 <nil>}
>*p_b: {20 <nil>}
>d: {20 <nil>}
>```
>
>指针定义方法二：
>
>```go
>	p_c := new(struct_1)
>	fmt.Printf("p_c: %p\n", p_c)
>```
>
>用new定义的变量为指针类型。
>
>
>
>[注]:访问结构体内部三种不同的方法，结果是一样的，如下：
>
>```go
>	fmt.Printf("a.value: %v\n", a.value)//用结构体本身
>	fmt.Printf("p_a.value: %v\n", p_a.value)//用结构体指针
>fmt.Printf("(*p_a).value: %v\n", (*p_a).value) //用结构体指针指向的内存
>```
>
>结果：
>
>```go
>a.value: 10
>p_a.value: 10
>```
>
>
>
>

> 结构体可以嵌套，可以通过结构体嵌套来实现继承。
>
> ```go
> type animal struct {
> 	name string 
> 	age int
> }
> 
> func (a animal)eat(){
> 	fmt.Println("eat...")
> }
> 
> type bird struct {
> 	animal
> }
> 
> type eager struct {
> 	animal
> }
> ```
>
> animal相当于父类，bird和eager相当于子类。



### 接口

#### 接口定义与要求

>
>
>#### 
>
>类似Java中的接口，但是是由结构体来定义而不是类。
>
>> 当一个结构体实现了一个接口，这个结构体就属于该接口了，也就是说，可以定义该接口来当做该结构体使用。
>
>[注]:接口中定义的全部都是方法，在接口中只写方法名。
>
>```go
>
>type usb interface {
>	read()
>	write()
>}
>
>type computer struct {
>	name string
>}
>
>type phone struct {
>	name string
>}
>
>func (com computer) read() {
>	fmt.Printf("com.name: %v\n", com.name)
>	fmt.Printf("computer read..\n")
>}
>
>func (com computer) write() {
>	fmt.Printf("com.name: %v\n", com.name)
>	fmt.Println("computer write..")
>}
>
>func (pho phone) read() {
>	fmt.Printf("pho.name: %v\n", pho.name)
>	fmt.Println("phone read")
>}
>
>func (pho phone) write() {
>	fmt.Printf("pho.name: %v\n", pho.name)
>	fmt.Println("phone write")
>}
>```
>
>实现接口则必须实现其内的所有方法。
>
>

接口内的方法同样可以以指针作为接收者，使用方法与[方法](#方法)中的使用发放一致，此处不再赘述。

#### 类型和接口的关系

>
>
>1.一个类型可以实现多个接口。
>
>2.一个接口可以被多个类型实现。

#### 接口嵌套

>
>
>```go
>type flyer interface {
>	fly()
>}
>
>type animer interface {
>	eat()
>	sleep()
>}
>
>type flyanimer interface {
>	flyer
>	animer
>}
>
>```
>
>

## 并发编程

### channel

>定义：
>
>```go
>	unbuffer := make(chan int)   //无缓冲池，相当于同步方式
>	buffer := make(chan int, 10) //有缓冲池
>```
>
>使用：
>
>```go
>unbuffer <- 10//写通道
>a:=<-unbuffer//读通道
>```
>
>关闭：
>
>```go
>defer close(ch_1)
>```
>
>

>#### channel遍历
>
> 方法一：for{
>
>r,ok:=<-chan
>
>if ok{
>
>
>
>}
>
>}
>
>方法二：for range





### 创建协程和协程并发

>
>
>使用go关键字创建协程。
>
>使用waitGroup控制协程并发。
>
>```go
>var wg sync.WaitGroup
>var ch_1 = make(chan int)//创建通道
>
>func setchan() {
>	defer wg.Done() //标记协程数减一
>	rand.Seed(time.Now().Unix())
>	a := rand.Intn(20)
>	ch_1 <- a
>	fmt.Printf("Set a: %v\n", a)
>}
>
>func getchan() {
>	defer wg.Done() //标记协程数减一
>	a := <-ch_1
>	fmt.Printf("Get a: %v\n", a)
>}
>
>func main() {
>
>	defer close(ch_1)
>
>	wg.Add(1) //标记协程数加1
>	go setchan()
>
>	wg.Add(1) //标记协程数加1
>	go getchan()
>
>	wg.Wait() //如果协程数不等于0，则让权等待
>
>}
>```
>
>

### runtime

>#### runtime.Gosched()
>
>有cpu使用权时，让出cpu使用权,让权等待

>#### runtime.Goexit()
>
>直接结束当前协程，类似于循环中的break。

> #### runtime.GOMAXPROCS()
>
> runtime.NumCPU()得到当前cpu核心数。
>
> runtime.GOMAXPROCS(int)要使用的核心数。

### 互斥锁

>
>
>```go
>var lock sync.Mutex//定义一个锁
>	lock.Lock()//上锁
>	lock.Unlock()//去锁
>```
>
>当一个协程使用lock.lock()时，如果本来没有锁住，则上锁然后继续执行后续语句，如果本来就是锁住的状态，则该协程堵塞。

### select

>
>
>类似于switch语句，区别在于select的case后跟的语句必须为通道。select会随机选择其中一个未被堵塞的通道执行，如果全部堵塞则执行default语句，如果没有default语句则返回堵塞错误。

## 使用mysql

### 使用sql

#### 建立连接

>
>
>```go
>    dsn := "root:852456@tcp(localhost:3306)/test1"
>//可选参数   charset=utf8mb4&parseTime=True&loc=Local
>	db, err := sqlx.Open("mysql", dsn)
>	if err != nil {
>		fmt.Printf("err: %v\n", err)
>	}
>	db.SetMaxOpenConns(10) //设置数据库连接池的最大连接数 10
>	db.SetMaxIdleConns(5)  //设置最大空闲连接数
>```
>
>

#### 增删改

>
>
>使用db.Exec()
>
>

#### 查

>查询一行
>
>```go
>	sql := "select * from t_stu where id=1"
>	stu_1 := new(stu)
>	erraa:=db.QueryRow(sql).Scan(&stu_1.Id, &stu_1.Name, &stu_1.Score)
>```
>
>查询多行
>
>```go
>	sql := "select * from t_stu"
>	row,errr:=db.Query(sql)
>	defer row.Close()
>	if errr!=nil {
>		fmt.Printf("errr: %v\n", errr)
>	}else{
>		for row.Next(){
>			row.Scan((&stu_1.Id, &stu_1.Name, &stu_1.Score)
>		}
>	}
>```
>
>
>
>

-----

### 使用sqlx

#### 建立连接

> 使用sqlx.Open()
>
> 

#### 增删改

>
>
>使用db.exec()

#### 查

>
>
>```go
>	sql := "select * from t_stu"
>
>	var stuss []stu
>
>	err1 := db.Select(&stuss, sql)
>	if err1 != nil {
>		fmt.Printf("err1: %v\n", err1)
>	}
>	fmt.Printf("stuu: %v\n", stuss)
>```
>
>也可以使用db.QueryRowx()和db.Queryx()，方法同[sql](####查)，此处不再赘述。

### 使用gorm

#### 获取gorm

>
>
>```cmd
>go get -u github.com/jinzhu/gorm
>```
>
>

#### 建表

>
>
>```go
>	type User struct{
>		gorm.Model
>	   Name string
>	   Age int
>	   Addr string
>}
>//
>db.CreateTable(&User{})
>//auto建表
>user:=User{"guang",20,"sss"}
>db.AutoMigrate(&user)
>//或者
>db.AutoMigrate(&User{})
>//可以一次建多个  db.AutoMigrate(&User{},&Stu{},&Tea{})
>```
>
>使用AutoMigrate建表的时候，结构体中的字段名一定要 
>
> **大写！！大写啊！！！**
>
>
>
>定义表名
>
>```go
>//通过实现接口定义表名
>func (u *User) TableName() string{
>	return "User"
>}
>
>//建表时定义表名
>db.Table("User").CreateTable(&User{})
>```
>
>

#### 增

> ```go
> var user User = User{Name: "guang", Age: 20, Addr: "sefe"}
> 	db.Create(&user)
> ```
>
> 

#### 删

>删表
>
>```go
>	db.DropTable("user")
>```
>
>删数据
>
>```go
>// 先查询数据 , 默认字段为 id
>var user User
>db.First(&user,1)
>// 删除数据
>db.Delete(&user)
>          
>//或者
>user := &User{ID:1}
>db.delete(user)
>```
>
>删除列
>
>```go
>// 删除User结构体对应表中的description字段
>db.Migrator().DropColumn(&User{}, "Name")
>```
>
>

#### 改

>
>
>```go
>// 先查询数据 , 默认字段为 id，如果按其他字段查询 db.First(&user,"name=?","小黑")
>db.First(&user,1)
>// 更新数据，链式调用Update即可
>db.Model(&user).Update("name":"小煤球")
>```
>
>

#### 查

>
>
>```go
>var user User
>	// 先查询数据 , 默认字段为 id
>	db.First(&user, 1)
>	fmt.Printf("user: %v\n", user)
>	// 如果按其他字段查询 db.First(&user,"name=?","小黑")
>	db.First(&user, "name=?", "bbb")
>	fmt.Printf("user: %v\n", user)
>	// 查询多条数据,最终是将数据存放在数组中,查询id>1的数据放在user地址上
>	var users []User
>	db.Where("id>?", 1).Find(&users)
>	fmt.Printf("users: %v\n", users)
>```
>
>

#### 使用原生sql

>
>
>```go
>db.Raw("select id,name,age from users where name=？","aaa").scan(&user)
>```
>
>

## 问题记录

### 类型断言

>
>
>```go
>	var in interface{}
>	in = "1111"
>	flag, ok1 := in.(string)
>	_, ok2 := in.(int)
>	fmt.Printf("flag: %v\n", flag)
>	fmt.Printf("ok1: %v   ok1: %v\n", ok1, ok2)
>```
>
>输出：
>
>![image-20220131173330465](C:\Users\DELL\AppData\Roaming\Typora\typora-user-images\image-20220131173330465.png)
>
>使用语法 变量.（类型）来进行断言，正确则ok返回true，错误则返回false

### 反射

>
>
>类似java中的反射，使用reflect包

### 错误处理

>
>
>实例
>
>```go 
>condi_1 := 0
>	fff := func(a int, b int) (float32, error) {
>		if b == 0 {
>			return 0, errors.New("除数为0")
>		} else {
>			return float32(a) / float32(b), nil
>		}
>	}
>
>	if _, err := fff(10, condi_1); err != nil {
>		panic(err)
>	}
>```
>
>也可以直接把错误上抛，直接输出err，不使用panic函数。

### 关于函数传参

注意：golang中的任何函数参数的传递，都是以**值传递**的形式，即实参的拷贝的形式传递！！！！

但是，在实际使用的过程中，发现了一些疑惑。

当使用数组作为函数参数传递，在函数中改变数组，函数外原数组不会改变，这是合理的。

但当使用切片作为函数参数传递时，在函数内改变切片，函数外的原切片也会改变，这一点很奇怪。

>
>
>对于go中的数组，实际上与c中的数组基本一致，不再分析。

```go
	a := []int{1}
	fmt.Printf("a.p: %p\n", &a)
	fmt.Printf("a[0].p: %p\n", &(a[0]))
```

```go
a.p: 0xc000004090
a[0].p: 0xc0000140b0
```

>
>
>可以看到，与数组不同，切片名对应的指针并不是切片第一个元素的地址，这是因为切片本质上是一个结构体，结构体中存放了切片元素的地址和长度。
>
>这意味着，虽然函数传递的参数是一个拷贝值，但是切片元素的地址作为该对象的值是直接被拷贝的，是不变的，所以当在函数中改变切片元素时，实际上是从切片对象中找到对应元素的地址进行改变，进而导致原切片也被改变。

### 关于切片

#### 测试切片定义方法

```go
	arry_1 := new([]int)
	arry_2 := make([]int, 0)
	arry_3 := []int{}
	fmt.Printf("arry_1: %T\n", arry_1)
	fmt.Printf("arry_2: %T\n", arry_2)
	fmt.Printf("arry_3: %T\n", arry_3)
```

输出：

```go
arry_1: *[]int
arry_2: []int
arry_3: []int
```

#### 测试切片与数组地址

>
>
>```go
>	a := [5]int{1, 2, 3, 4, 5}
>	b := a[:4]
>	fmt.Printf("a: %v\n", a)
>	fmt.Printf("b: %v\n", b)
>	b = append(b, 44)
>	fmt.Printf("b: %v\n", b)
>	fmt.Printf("a: %v\n", a)
>	fmt.Printf("a[3].p: %p   a[3].p: %p\n", &a[3], &b[3])
>	fmt.Printf("a[4]: %v    b[4]: %v\n", a[4], b[4])
>	fmt.Printf("a[4].p: %p   b[4].p: %p\n", &a[4], &b[4])
>```
>
>输出结果：
>
>```go
>a: [1 2 3 4 5]
>b: [1 2 3 4]
>b: [1 2 3 4 44]
>a: [1 2 3 4 44]
>a[3].p: 0xc00000a498   a[3].p: 0xc00000a498
>a[4]: 44    b[4]: 44
>a[4].p: 0xc00000a4a0   b[4].p: 0xc00000a4a0
>```
>
>结论：用数组创建的切片，是完全使用数组的地址和空间，切片改变，数组也会改变，数组改变，切片也会改变



### go的数据结构与C++STL的数据结构

>
>
>slice：：vector
>
>map：：unordered_map
>
>list：：list