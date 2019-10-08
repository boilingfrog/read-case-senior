# read-case-senior
这是一个关于读书的练习记录，这本书的名字是《Golang高级编程》


作者的地址：https://github.com/chai2010/advanced-go-programming-book


# Golang汇编中学习记录

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
