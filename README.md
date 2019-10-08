# 《Golang高级编程》读书笔记
这是一个关于读书的练习记录，这本书的名字是《Golang高级编程》


作者的地址：https://github.com/chai2010/advanced-go-programming-book


## 汇编部分的学习


- 调试汇编的命令
```
go tool compile -S pkg-data.go
```

- 将调试的命令放到s文件中

```
go tool compile -S pkg-data.go >> pkg-data.s
```

Go汇编中提供了DATA命令用于初始化包变量，下面是DATA的语法：

```
DATA symbol+offset(SB)/width, val
```

其中symbol为变量在汇编中对应的标示符，offset是符号的开始地址的偏移量，width是初始化内存的宽度大小，value是要初始化的值。其中当前包里中GO语言定义的symbol，在汇编代码中对应.symbol，其中“.”中点符号为一个特殊的unicode符号。

下面例子是个简单的例子：

```
DATA ·Id+0(SB)/1,$0x37
DATA ·Id+1(SB)/1,$0x25
```
给ID变量初始化了进制的0x2537,对应十进制的9527（需要以美元符号$开头表示）