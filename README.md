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

### 计算机结构

汇编语⾔是直⾯计算机的编程语⾔，因此理解计算机结构是掌握汇编语⾔的前提。当前流⾏的计算机
基本采⽤的是冯·诺伊曼计算机体系结构（在某些特殊领域还有哈佛体系架构）。冯·诺依曼结构也称为
普林斯顿结构，采⽤的是⼀种将程序指令和数据存储在⼀起的存储结构。冯·诺伊曼计算机中的指令和
数据存储器其实指的是计算机中的内存，然后在配合CPU处理器就组成了⼀个最简单的计算机了。


X86其实是是80X86的简称（后⾯三个字⺟），包括Intel 8086、80286、80386以及80486等指令集
合，因此其架构被称为x86架构。x86-64是AMD公司于1999年设计的x86架构的64位拓展，向后兼容
于16位及32位的x86架构。X86-64⽬前正式名称为AMD64，也就是Go语⾔中GOARCH环境变量指定
的AMD64。

#### 汇编中的伪寄存器

什么是寄存器？

寄存器（Register），是中央处理器内的其中组成部分。寄存器是有限存贮容量的高速存贮部件，它
们可用来暂存指令、数据和地址。在中央处理器的控制部件中，包含的寄存器有指令寄存器（IR）和程
序计数器。在中央处理器的算术及逻辑部件中，包含的寄存器有累加器。

在计算机体系结构里，处理器中的寄存器是少量且速度快的计算机存储器，借由提供快速共同地访问数
值来加速计算机程序的运行：典型地说就是在已知时间点所作的之计算中间的数值。

寄存器是存储器层次结构中的最顶端，也是系统操作数据的最快速途径。寄存器通常都是以他们可以保
存的比特数量来估量，举例来说，一个8位寄存器或32位寄存器。寄存器现在都以寄存器数组的方式来
实现，但是他们也可能使用单独的触发器、高速的核心存储器、薄膜存储器
以及在数种机器上的其他方式来实现出来。

GO汇编为了简化汇编代码的编写，引入PC，FP，SP，SB四个伪寄存器。四个伪寄存器加其它的通
⽤寄存器就是Go汇编语⾔对CPU的重新抽象，该抽象的结构也适⽤于其它⾮X86类型的体系结构。

![Aaron Swartz](https://github.com/zhan-liz/read-case-senior/blob/master/asm/asm1.png?raw=true)

在AMD64环境中，伪PC寄存器其实是IP指令计数器寄存器的别名。伪FP寄存器对应的是函数的帧指针，
一般用来访问函数的参数和返回值。










