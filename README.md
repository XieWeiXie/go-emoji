# Go-Emoji

> golang 版的 emoji 表情使用包

1. unicode 
2. emoji
3. hex --> decimal
4. unicode --> rune

```go
package main
import "fmt"
import "strconv"
func main(){
    fmt.Println(strconv.ParseInt("1F601", 16, 64))
}
```

### API

