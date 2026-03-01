package cdyleetcode

// *************************************************
// ❌❌❌ 只找股价最低点，然后找后面的最高点。
// 错误，因为全局最低点和之后的高点的差值，不一定有其他差值多。
// *************************************************
func maxProfit(prices []int) int {
    // 先找股价最低点
    lowest := prices[0]
    lowestDay := 0
    for i:=0; i<len(prices); i++ {
        if prices[i] < lowest {
            lowest = prices[i]
            lowestDay = i
        }
    }

    // 然后找之后的最高点
    hightest := lowest
    for j:=lowestDay+1; j<len(prices); j++ {
        if prices[j] > hightest {
            hightest = prices[j]
        }
    }

    return hightest-lowest

}
// *************************************************
// 灵神的思路
// *************************************************
func maxProfit(prices []int) int {
    lowest := prices[0]
    maxProfit := 0 // 到目前为止的最大利润

	// 遍历每一天的价格，把每一天都当作可能的“卖出日”
    for i:=1; i<len(prices); i++ {
        // 比较当天卖出和不卖出的利润
        maxProfit = max(prices[i]-lowest, maxProfit)
        // 当天也可以成为未来某个卖出日所对应的买入点
        lowest = min(lowest, prices[i])
    }

    return maxProfit

}

// 核心思想：如果我在今天卖出，那么为了获得最大利润，
// 我一定希望买入价是过去（包括今天）最低的那一天。
// 所以，用今天的价格(prices[i]) 减去 过去的最低价(lowest)，
// 就得到了“如果今天卖出”能获得的最大利润。
// 我们和之前记录的最大利润(maxProfit)比较，取较大的那个。


// 在考虑完“今天卖出”的可能性后，
// 更新“到当前为止的最低买入价”，为未来的卖出做准备。
// 因为未来的卖出日，可以选择在今天或之前任何一天买入。
