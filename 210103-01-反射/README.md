## 反射

#### 一、基本介绍

1. 反射可以再运行时动态获取变量的各种信息，比如变量的类型 `type` ，类别 `kind` 
2. 如果是结构体变量，可以获取到结构体本身的信息，字段和方法
3. 通过反射，可以修改变量的值，可以调用关联的方法
4. 使用反射，需要 `import ("reflect")` 



#### 二、`package reflect` 

> 典型用法是用静态类型 `interface{}` 保存一个值，
>
> 通过调用 `TypeOf` 获取其动态类型信息，该函数返回一个 `Type` 类型值。
>
> 调用 `ValueOf` 函数返回一个 `Value` 类型值，该值代表运行时的数据。



#### 三、反射的重要的函数和概念

1.  `reflect.TypeOf(变量名)` ，获取变量的类型，返回 `reflect.Type` 类型
2.  `reflect.ValueOf(变量名)` ，获取变量的值，返回 `reflect.Value` 类型， `reflect.Value` 是一个结构体类型
3. 变量、 `interface{}` 和 `reflect.Value` 是可以相互转换的



#### 四、反射的注意事项和细节说明

1.  `reflect.Value.Kind` 获取变量的类型，返回的是一个常量；

2.  `Type` 是类型， `Kind` 是类别， `Type` 和 `Kind` 可能相同，也可能不同：

   ```go
   var num int = 10 // num 的 Type 是 int ， Kind 也是 int
   var stu Student // stu 的 Type 是 包名.Student， Kind 是 struct
   ```

3. 通过反射可以让变量在 `interface{}` 和 `reflect.Value` 之间相互转换；

4. 使用反射获取变量的值，要求数据类型匹配，比如 `int` 类型：

   ```go
   var x int
   xNum := reflect.Value(x).Int()
   ```

   

