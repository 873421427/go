# go 实现selpg


## go flag 包学习

主要说明flag包有什么用？
flag包就是可以帮助我们在终端输入
```
selpg -s 1 -l 2
```
命令时，可以读取 s的值为1 而l的值为2
相当于就是一个map 键值对

[标准库—命令行参数解析FLAG](http://blog.studygolang.com/2013/02/%E6%A0%87%E5%87%86%E5%BA%93-%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%8F%82%E6%95%B0%E8%A7%A3%E6%9E%90flag/)

[Go学习笔记：flag库的使用](https://studygolang.com/articles/5608)
## selpg实现逻辑学习

[开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)

程序顺序执行逻辑如下:
1. 读入参数
2. 检查参数的合法性
3. 读入文件或读取用户输入
4. 输出
## selpg 代码实现

这里主要介绍一下用到的包
1. 读取文件
```
fin := os.Stdin
fin, err := os.Open(sa.in_filename)
```
得到打开文件的指针

2.

```
cmd := exec.Command("grep", "-nf", "keyword")
```
帮助我们执行外部命令 ，返回cmd结构
```
cmd.Stdout = fout
```
将命令执行的标准输出设置为系统的标准输出

3.
```
  line := bufio.NewScanner(fin)
```
对标准输入得到扫描器
```  
  fout.Write([]byte(line.Text() + "\n"))
```
在标准输出写入每行的读入
```
  inpipe.Write([]byte(line.Text() + "\n"))
```
往管道中写入每行的数据

##测试
| < > 都是linux 的原生操作符，因此不用额外实现


## 遇到的问题
1. 我直接在centos上运行
```
go run selpg.go
```
时可以正常运行， 而用xshell连接时运行命令则会提示找不到flag包，我认为是我引用包的名字可能有问题
