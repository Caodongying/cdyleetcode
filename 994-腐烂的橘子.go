package cdyleetcode

// 使用切片构建FIFO队列

type coordinate struct {
    x int
    y int
}

func orangesRotting(grid [][]int) int {
    var queue []coordinate
    // 初始化队列
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j] == 2 {
                // 添加当前坐标到队列尾
                queue = append(queue, coordinate{i,j})
            }
        }
    }

    totalMin := 0
    // 开始处理
    for {
        // 备份队列当前腐烂橘子的个数
        currentRottedOrg := len(queue)
        for k:=0; k < currentRottedOrg; k++{
            // 腐烂邻居新鲜橘子并加入队列
            centerCoor := queue[k]
            rotNeighbour(grid, &queue, centerCoor)
        }

        //  将currentRottedOrg个橘子移除队列
        queue = queue[currentRottedOrg:]

        // 检查是否结束
        if len(queue) == 0 {
            break
        }

        totalMin++
    }

    // 遍历检查一遍是否还有新鲜橘子
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j] == 1 {
                totalMin = -1
            }
        }
    }
    return totalMin
}

func rotNeighbour(grid [][]int, queue *[]coordinate, center coordinate) {
    // 怎么避免数组浅拷贝
    x := center.x
    y := center.y
    // 检查上邻居
    if x - 1 >= 0 {
        if grid[x-1][y] == 1 {
            grid[x-1][y] = 2
            *queue = append(*queue, coordinate{x-1, y})
        }
    }
    // 检查右邻居
    if y+1 < len(grid[0]){
        if grid[x][y+1] == 1 {
            grid[x][y+1] = 2
            *queue = append(*queue, coordinate{x, y+1})
        }
    }
    // 检查下邻居
    if x + 1 < len(grid) {
        if grid[x+1][y] == 1 {
            grid[x+1][y] = 2
            *queue = append(*queue, coordinate{x+1, y})
        }
    }
    // 检查左邻居
    if y-1 >= 0 {
        if grid[x][y-1] == 1 {
            grid[x][y-1] = 2
            *queue = append(*queue, coordinate{x, y-1})
        }
    }
}
// 广度优先遍历
// 先初始化队列，也就是已经腐烂的橘子的坐标
// 如果此时二维数组中没有1，说明此时0分钟没有新鲜橘子了
// 只要队列不为空，就进入新一轮循环，并且保存进入循环时的副本
// 如果邻居有没访问过的新鲜橘子，就将之值设为2，并且坐标加入队列