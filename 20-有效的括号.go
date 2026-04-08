package cdyleetcode

// *************************************************
// 用for range，所以从string里取出的是rune。注意栈顶弹出之前首先判断是否为空，还有循环结束后要看栈里是否还有左括号
// *************************************************
func isValid(s string) bool {
    stack := make([]rune, 0, len(s)) // 不可以是byte

    for _, v := range s {
        if v == '(' || v == '[' || v == '{' {
            stack = append(stack, v)
        } else {
            if len(stack) == 0 {
                return false
            }
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if v == ')' && top !=  '(' {
                return false
            }
            if v == ']' && top !=  '[' {
                return false
            }
            if v == '}' && top !=  '{' {
                return false
            }
        }
    }

    if len(stack) > 0 {
        return false
    }

    return true
}

// 直接用栈

// *************************************************
// 用s[i]，优化内存，因为rune是4个字节，byte是一个字节
// *************************************************
func isValid(s string) bool {
    stack := make([]byte, 0, len(s)) // 不可以是byte

    // 下标访问 s[i]：返回的是第 i 个 byte。
    // for range 迭代：返回的是字符的 Unicode 码点（Code Point），也就是 rune。

    for i:=0; i<len(s); i++ {
        v := s[i]
        if v == '(' || v == '[' || v == '{' {
            stack = append(stack, v)
        } else {
            if len(stack) == 0 {
                return false
            }
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if v == ')' && top !=  '(' {
                return false
            }
            if v == ']' && top !=  '[' {
                return false
            }
            if v == '}' && top !=  '{' {
                return false
            }
        }
    }

    if len(stack) > 0 {
        return false
    }

    return true
}

// 直接用栈

// *************************************************
// 一开始判断一下是不是偶数长度，优化一下
// *************************************************
func isValid(s string) bool {
    // 优化：s的长度必须是偶数
    if len(s)%2 != 0 {
        return false
    }
    stack := make([]byte, 0, len(s)) // 不可以是byte

    // 下标访问 s[i]：返回的是第 i 个 byte。
    // for range 迭代：返回的是字符的 Unicode 码点（Code Point），也就是 rune。

    for i:=0; i<len(s); i++ {
        v := s[i]
        if v == '(' || v == '[' || v == '{' {
            stack = append(stack, v)
        } else {
            if len(stack) == 0 {
                return false
            }
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if v == ')' && top !=  '(' {
                return false
            }
            if v == ']' && top !=  '[' {
                return false
            }
            if v == '}' && top !=  '{' {
                return false
            }
        }
    }

    if len(stack) > 0 {
        return false
    }

    return true
}

// 直接用栈
