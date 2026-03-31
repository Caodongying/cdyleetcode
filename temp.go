package main

import "fmt"

// import "fmt"


var digitCharMap map[int][]string

func buildMap() {
	char := 'a'
	for i := 2; i <= 9; i++ {
		value := []string{}
		if i == 7 || i == 9 {
			value = append(value, string(char), string(char+1), string(char+2), string(char+3))
			char = char + 4
		} else {
			value = append(value, string(char), string(char+1), string(char+2))
			char = char + 3
		}

		digitCharMap[i] = value
	}

    printMap()

}

func printMap() {
    for k, v := range digitCharMap {
        fmt.Printf("key: %v, value: %v\n", k, v)
    }
}

func letterCombinations(digits string) []string {
	// digitCharMap := make(map[int][]string, 0, 8) 错误 map只能直接指定容量
    digitCharMap = make(map[int][]string, 8) // 不能是digitCharMap :=
	buildMap()
	result := make([]string, 0)
    combination := ""
    n := len(digits)
    fmt.Println("n是", n)

    nums := make([]int, 0, n)
    // 把digits转化为[]int
    for _, v := range digits {
        nums = append(nums, int(v-'0'))
    }

    var dfs func(idx int, combination string) // 在comnination后，分别加上当前idx对应的集合元素
    dfs = func(idx int, combination string) {
        fmt.Println("进入dfs")
        // 最终的组合字符串长度一定是len(digits)
        if len(combination) == n {
            fmt.Println("combination是", combination)
            result = append(result, combination)
            return
        }

        // charSet := digitCharMap[idx] // 字符可以从chatSet数组中选取
        charSet := digitCharMap[nums[idx]]
        fmt.Printf("%d对应的charSet是%v", idx, charSet)

        for i:=0; i<len(charSet); i++ {
            combination += charSet[i]
            dfs(idx+1, combination)
            combination = combination[:len(combination)-1]
        }
    }

    dfs(0, combination)

    fmt.Println(result)

    return result

}

// 每一个digit都是一个字符集合
// 按字符集合顺序，从每一个集合里取出一个字符，顺序拼接在一起

func main() {
   letterCombinations("23")
}
