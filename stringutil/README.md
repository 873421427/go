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

### 扩展： ubuntu 环境变量文件
/etc/profile
在用户登录时，操作系统定制用户环境时使用的第一个文件，此文件为系统的每个用户设置环境信息，当用户第一次登录时，该文件被执行。

/etc /environment
在用户登录时，操作系统使用的第二个文件， 系统在读取用户个人的profile前，设置环境文件的环境变量。

~/.profile
在用户登录时，用到的第三个文件 是.profile文件，每个用户都可使用该文件输入专用于自己使用的shell信息，当用户登录时，该文件仅仅执行一次！默认情况下，会设置一些环境变量，执行用户的.bashrc文件。

/etc/bashrc
为每一个运行bash shell的用户执行此文件，当bash shell被打开时，该文件被读取。

~/.bashrc
该文件包含专用于用户的bash shell的bash信息，当登录时以及每次打开新的shell时，该该文件被读取。

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
