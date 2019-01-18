#  ubuntu golang环境搭建
## 下载安装golang
首先可以直接尝试
```
sudo apt-get install golang-go
```
如果这个命令过时的话，可以改用下面的
这里用的是从命令行下载
```
$ sudo add-apt-repository ppa:gophers/archive
$ sudo apt-get update
$ sudo apt-get install golang-1.10-go
```

测试安装
```
$ go version
```
## 配置环境变量
Go代码必须放在工作空间内。它其实就是一个目录，其中包含三个子目录：

src 目录包含Go的源文件，它们被组织成包（每个目录都对应一个包），
pkg 目录包含包对象，
bin 目录包含可执行命令。
go 工具用于构建源码包，并将其生成的二进制文件安装到 pkg 和 bin 目录中。

src 子目录通常包会含多种版本控制的代码仓库（例如Git或Mercurial）， 以此来跟踪一个或多个源码包的开发。

1. 创建工作空间
```
$ mkdir $HOME/gowork
```
2. 配置环境变量，对ubuntu 在~/.profile 下添加
```
export GOPATH=$HOME/gowork
export PATH=$PATH:$GOPATH/bin
```
使文件生效
```
$ source $HOME/.profile
```

3. 检查环境变量是否设置成功
```
$ go env
...
GOPATH = ...
...
GOROOT = ...
...
```
看到GOPATH 的变量出现就说明成功了
##  编写第一个包，做一次测试
1. 创建包目录
```
$ mkdir $GOPATH/src/github.com/user/stringutil
```
user 替换为你自己的github用户名
创建reverse.go 文件
```
// stringutil 包含有用于处理字符串的工具函数。
package stringutil

// Reverse 将其实参字符串以符文为单位左右反转。
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```
用 go build 命令测试该包的编译
2. 测试所写的文件
go 拥有轻量级的测试框架 go test 命令 和testing 包构成。

你可以通过创建一个名字以 _test.go 结尾的，包含名为 TestXXX 且签名为 func (t *testing.T) 函数的文件来编写测试。 测试框架会运行每一个这样的函数；若该函数调用了像 t.Error 或 t.Fail 这样表示失败的函数，此测试即表示失败。

测试刚才所写的代码
```
package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
```

使用go test 命令测试
如果成功，则会显示ok
