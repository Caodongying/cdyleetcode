package cdyleetcode

// *************************************************
// ❌❌ 自己写的，直接把digit下标作为map的key了
// *************************************************
var digitCharMap map[int][]string

func buildMap() {
	char := 'a'
	for i := 2; i <= 9; i++ {
		value := []string{}
		if i == 7 || i == 9 {
			value = append(value, char, string(char+1), string(char+2), string(char+3))
			char = string(char + 4)
		} else {
			value = append(value, char, string(char+1), string(char+2))
			char = string(char + 3)
		}

		digitCharMap[i] = value
	}
}

func letterCombinations(digits string) []string {
	// digitCharMap := make(map[int][]string, 0, 8) 错误 map只能直接指定容量
    digitCharMap = make(map[int][]string, 8) // 不能是digitCharMap :=
	buildMap()
	result := make([]string, 0)
    combination = ""
    n := len(digits)

    nums := make([]int, 0, n)
    // 把digits转化为[]int
    for _, v := range digits {
        nums = append(nums, int(v-'0'))
    }
2 3
    var dfs func(idx int) // 计算组合字符串第idx个字符
    dfs = func(idx int) {
        // 最终的组合字符串长度一定是len(digits)
        if len(combination) == n {
            result = append(result, combination)
            return
        }

        charSet := digitCharMap[idx] // 字符可以从chatSet数组中选取

        for i:=0; i<len(charSet); i++ {
            combination += charSet[i]
            dfs(idx+1)
            combination = combination[:len(combination)]
        }
    }

    dfs(0)

    return result

}

// 每一个digit都是一个字符集合
// 按字符集合顺序，从每一个集合里取出一个字符，顺序拼接在一起

// *************************************************
// 修正了map的key使用错误的问题
// *************************************************
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
}

func letterCombinations(digits string) []string {
	// digitCharMap := make(map[int][]string, 0, 8) 错误 map只能直接指定容量
    digitCharMap = make(map[int][]string, 8) // 不能是digitCharMap :=
	buildMap()
	result := make([]string, 0)
    combination := ""
    n := len(digits)

    nums := make([]int, 0, n)
    // 把digits转化为[]int
    for _, v := range digits {
        nums = append(nums, int(v-'0'))
    }

    var dfs func(idx int, combination string) // 在comnination后，分别加上当前idx对应的集合元素
    dfs = func(idx int, combination string) {
        // 最终的组合字符串长度一定是len(digits)
        if len(combination) == n {
            result = append(result, combination)
            return
        }

        // charSet := digitCharMap[idx] // 字符可以从chatSet数组中选取
        charSet := digitCharMap[nums[idx]] // 字符可以从chatSet数组中选取

        for i:=0; i<len(charSet); i++ {
            combination += charSet[i]
            dfs(idx+1, combination)
            combination = combination[:len(combination)-1]
        }
    }

    dfs(0, combination)

    return result

}

// 每一个digit都是一个字符集合
// 按字符集合顺序，从每一个集合里取出一个字符，顺序拼接在一起

// *************************************************
// 优化，combination不用作为参数。不用函数构建map，直接用[]string。用数组替代map。
// *************************************************
func letterCombinations(digits string) []string {
    digitCharMap := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	result := make([]string, 0)
    combination := ""
    n := len(digits)

    var dfs func(idx int) // 在comnination后，分别加上当前idx对应的集合元素
    dfs = func(idx int) {
        // 最终的组合字符串长度一定是len(digits)
        if len(combination) == n {
            result = append(result, combination)
            return
        }

        // charSet := digitCharMap[idx] // 字符可以从chatSet数组中选取
        charSet := digitCharMap[int(digits[idx]-'0')-2] // 字符可以从chatSet数组中选取

        for i:=0; i<len(charSet); i++ {
            combination += string(charSet[i])
            dfs(idx+1)
            combination = combination[:len(combination)-1]
        }
    }

    dfs(0)

    return result

}

// 每一个digit都是一个字符集合
// 按字符集合顺序，从每一个集合里取出一个字符，顺序拼接在一起
