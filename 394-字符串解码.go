package cdyleetcode

// *************************************************
// ❌❌ 错误，最后构建的字符串应该是从stack里取的，而不是临时的结果
// *************************************************
func decodeString(s string) string {
    stack := make([]byte, 0, len(s))
    resultArray := make([]byte, 0)
    for i:=0; i<len(s); i++ { // byte类型，为什么可以和rune比较
        if s[i] != ']' { // byte和rune怎么可以比较
            // 压栈
            stack = append(stack, s[i])
        } else {
            // 处理]
            resultArray = []byte{}
            for {
                // 弹栈
                top := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                if top == '[' {
                    // 逆序处理
                    head, tail := 0, len(resultArray) - 1
                    for head < tail {
                        resultArray[head], resultArray[tail] = resultArray[tail], resultArray[head]
                        head++
                        tail--
                    }

                    // 再弹一个，也就是数字
                    top := stack[len(stack)-1]
                    stack = stack[:len(stack)-1]
                    count, _ := strconv.Atoi(string(top))

                    // 将新字符串压栈
                    for i:=0; i<count; i++ {
                        stack = append(stack, resultArray...)
                    }

                    break
                } else {
                    resultArray = append(resultArray, top)
                }
            }
        }
    }

    // 把resultArray转化为string
    result := ""
    for i:=0; i<len(resultArray); i++ {
        result = result + string(resultArray[i])
    }
    return result
}

// 对括号一个一个解码

// 遍历string，如果是数字、左括号、字符，就压栈
// 如果是右括号，就一直弹出，直到弹到左括号。中间弹出的字母，逆序，保存。然后下一个弹出的一定是数字。构造给定长度后，压栈
// 继续进行压栈

// 3[a2[c]]
// 压栈：3 [ a 2 [ c
// 弹栈: 3 [ a 2 逆序字符串为c
// 弹栈1个：3 [ a c c
// 继续扫描，又遇到右括号，弹栈：3 逆序字符串为acc
// 弹栈1个：accaccacc

// 如何处理数字：string。判断是否为数字：v, err := strconv.Atoi(s)

// *************************************************
// ❌❌ 修改的地方是，最后构建最终结果，用原本的stack。错误原因在于，只对单个数字有效，数字“100”在当前方法下无法识别
// *************************************************
func decodeString(s string) string {
    stack := make([]byte, 0, len(s))
    resultArray := make([]byte, 0)
    for i:=0; i<len(s); i++ { // byte类型，为什么可以和rune比较
        if s[i] != ']' { // byte和rune怎么可以比较
            // 压栈
            stack = append(stack, s[i])
        } else {
            // 处理]
            resultArray = []byte{}
            for {
                // 弹栈
                top := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                if top == '[' {
                    // 逆序处理
                    head, tail := 0, len(resultArray) - 1
                    for head < tail {
                        resultArray[head], resultArray[tail] = resultArray[tail], resultArray[head]
                        head++
                        tail--
                    }

                    // 再弹一个，也就是数字
                    top := stack[len(stack)-1]
                    stack = stack[:len(stack)-1]
                    count, _ := strconv.Atoi(string(top))

                    // 将新字符串压栈
                    for i:=0; i<count; i++ {
                        stack = append(stack, resultArray...)
                    }

                    break
                } else {
                    resultArray = append(resultArray, top)
                }
            }
        }
    }

    // 把resultArray转化为string
    result := ""
    for i:=0; i<len(stack); i++ {
        result = result + string(stack[i])
    }
    return result
}

// 对括号一个一个解码

// 遍历string，如果是数字、左括号、字符，就压栈
// 如果是右括号，就一直弹出，直到弹到左括号。中间弹出的字母，逆序，保存。然后下一个弹出的一定是数字。构造给定长度后，压栈
// 继续进行压栈

// 3[a2[c]]
// 压栈：3 [ a 2 [ c
// 弹栈: 3 [ a 2 逆序字符串为c
// 弹栈1个：3 [ a c c
// 继续扫描，又遇到右括号，弹栈：3 逆序字符串为acc
// 弹栈1个：accaccacc

// 如何处理数字：string。判断是否为数字：v, err := strconv.Atoi(s)

// *************************************************
// 修整了“数字前面不一定是右括号”的错误。这个方法只用一个栈，处理数字和逆序比较麻烦。
// *************************************************
func decodeString(s string) string {
    stack := make([]byte, 0, len(s))
    resultArray := make([]byte, 0)
    for i:=0; i<len(s); i++ { // byte类型，为什么可以和rune比较
    // 在 Go 中，'a' 这种单引号字面量是 untyped rune。当它与 byte (uint8) 比较时，如果其值在 0-255 范围内，编译器会自动处理，所以 s[i] != ']' 是合法且高效的。
        if s[i] != ']' { // byte和rune怎么可以比较
            // 压栈
            stack = append(stack, s[i])
        } else {
            // 处理]
            resultArray = []byte{}
            for {
                if len(stack) == 0 {
                    break
                }
                // 弹栈
                top := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                if top == '[' {
                    // 逆序处理
                    head, tail := 0, len(resultArray) - 1
                    for head < tail {
                        resultArray[head], resultArray[tail] = resultArray[tail], resultArray[head]
                        head++
                        tail--
                    }

                    // 把数字弹出来
                    temp := ""
                    for {
                        if len(stack) == 0 {
                            break
                        }
                        // 弹栈
                        top := stack[len(stack)-1]
                        if (top - '0' < 0) || (top - '0' > 9) {
                            break
                        }
                        // if top == ']' {
                        //     break
                        // }
                        stack = stack[:len(stack)-1]
                        temp = string(top) + temp
                    }
                    count, _ := strconv.Atoi(temp)

                    // 将新字符串压栈
                    for i:=0; i<count; i++ {
                        stack = append(stack, resultArray...)
                    }

                    break
                } else {
                    resultArray = append(resultArray, top)
                }
            }
        }
    }

    // 把resultArray转化为string
    result := ""
    for i:=0; i<len(stack); i++ {
        result = result + string(stack[i])
    }
    return result
}

// 对括号一个一个解码

// 遍历string，如果是数字、左括号、字符，就压栈
// 如果是右括号，就一直弹出，直到弹到左括号。中间弹出的字母，逆序，保存。然后下一个弹出的一定是数字。构造给定长度后，压栈
// 继续进行压栈

// 3[a2[c]]
// 压栈：3 [ a 2 [ c
// 弹栈: 3 [ a 2 逆序字符串为c
// 弹栈1个：3 [ a c c
// 继续扫描，又遇到右括号，弹栈：3 逆序字符串为acc
// 弹栈1个：accaccacc

// 如何处理数字：string。判断是否为数字：v, err := strconv.Atoi(s)

// *************************************************
// 优化：①可以把[]byte直接转化为string，就是string(array) ，不必遍历然后一个个加到string里，这样编译器底层可以优化 ②用倒序压栈，避免对数组的显式反转
// *************************************************


func decodeString(s string) string {
    stack := make([]byte, 0, 2*len(s)) // 只是给一个较大的初始容量，如果不够会自动扩容
    for i:=0; i<len(s); i++ { // byte类型，为什么可以和rune比较
    // 在 Go 中，'a' 这种单引号字面量是 untyped rune。当它与 byte (uint8) 比较时，如果其值在 0-255 范围内，编译器会自动处理，所以 s[i] != ']' 是合法且高效的。
        if s[i] != ']' {
            // 压栈
            stack = append(stack, s[i])
        } else { // 处理]
            // 1. 提取括号里的内容，逆序弹出
            temp := []byte{}
            for len(stack)>0 && stack[len(stack)-1] != '[' {
                // 弹栈
                top := stack[len(stack)-1]
                temp = append(temp, top)
                stack = stack[:len(stack)-1]
            }

            // 弹出原栈里的[
            stack = stack[:len(stack)-1]

            // 2.提取原栈里的数字
            countStr := ""
            for len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
                countStr = string(stack[len(stack)-1]) + countStr
                stack = stack[:len(stack)-1]
            }

            count, _ := strconv.Atoi(countStr)

            // 将内容再压回栈（temp是逆序的，所以倒着读)
            for j:=0; j<count; j++ {
                for l:=len(temp)-1; l>=0; l-- {
                    stack = append(stack, temp[l])
                }
            }

        }
    }

    return string(stack)
}

// 对括号一个一个解码

// 遍历string，如果是数字、左括号、字符，就压栈
// 如果是右括号，就一直弹出，直到弹到左括号。中间弹出的字母，逆序，保存。然后下一个弹出的一定是数字。构造给定长度后，压栈
// 继续进行压栈

// 3[a2[c]]
// 压栈：3 [ a 2 [ c
// 弹栈: 3 [ a 2 逆序字符串为c
// 弹栈1个：3 [ a c c
// 继续扫描，又遇到右括号，弹栈：3 逆序字符串为acc
// 弹栈1个：accaccacc

// 如何处理数字：string。判断是否为数字：v, err := strconv.Atoi(s)