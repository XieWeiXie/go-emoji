# Go-Emoji

> golang 版的 emoji 表情使用包


### Usage

```
go get github.com/wuxiaoxiaoshen/go-emoji
```




### :open_mouth: 原理介绍

:email: 收集码点

资源网站：
- [github emoji](https://api.github.com/emojis)
- [Full emoji unicode](https://emojipedia.org/emoji/)
- [emoji table](https://apps.timwhitlock.info/emoji/tables/unicode)
- [emoji](https://www.webfx.com/tools/emoji-cheat-sheet/)


:love_letter: 处理码点

主要原理如下：

- :bookmark_tabs: 对应的码点为 U+1F4D1
- 则使用如下操作：

```go
package main

import "fmt"
import "strconv"
import "html"

func main(){
	x , _ := strconv.ParseInt("1F4D1", 16, 64)
    str := html.UnescapeString("&#" + strconv.Itoa(int(x)) + ";")
    fmt.Println(str)
}
// 📑
```

### API

- NewEmoji
> 初始化操作
```go 
emo := emoji.NewEmoji()
```

- ShortCodeList
> 获取所有的 shortCode

```
fmt.Println(emo.ShortCodeList())
```

- CodePoints
> 获取所有的 codePoints

```
fmt.Println(emo.CodePoints())

```

- Print
- Println
- Printf
- Fprint
- Fprintln
- Fprintf
- Sprint
- Sprintf
- Errorf

> 输出
```

	for _, i := range emo.ShortCodeList() {
		emo.Println(i)
	}
	for _, i := range emo.ShortCodeList() {
		fmt.Println(emo.Exists(i))
	}
	for _, i := range emo.ShortCodeList() {
		emo.Sprintf("%#v", i)
	}
	for _, i := range emo.ShortCodeList() {
		emo.Sprint(i)
	}

```

- Random
> 随机得出一个emoji
- Length
> 支持的 emoji 的个数
- Exists
> 根据 short code 判定是否支持 emoji

📃 License

MIT ©xiewei