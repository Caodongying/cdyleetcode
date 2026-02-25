package cdyleetcode

type coordinate struct{
    x int
    y int
}

func numIslands(grid [][]byte) int {
  num := 0
  for i, row := range(grid){
    for j, _ := range(row){
        if grid[i][j] != '1' {
            continue
        }
        c := coordinate{x:i, y:j}
        dfs(c, grid)
        num++
    }
  }
  return num
}

func dfs(c coordinate, grid [][]byte) {
    if c.x < 0 || c.x > len(grid)-1 {
        return
    }
    if c.y < 0 || c.y > len(grid[0])-1 {
        return
    }
    if grid[c.x][c.y] != '1' {
        return
    }

    grid[c.x][c.y] = '2'

    dfs(coordinate{c.x-1, c.y}, grid)
    dfs(coordinate{c.x, c.y-1}, grid)
    dfs(coordinate{c.x+1, c.y}, grid)
    dfs(coordinate{c.x, c.y+1}, grid)
}

// 方法一：直接在二维数组上改
// 每遇到一个1，就进行深度遍历，将所有访问到的岛屿都设置为2（其实就是标记visited）
// 深度遍历完后，将岛屿数量+1。这是因为dfs已经可以遍历到当前连通子图上的所有节点了。

// --  初始思路，错错的
    // 设置一个map，key为value里的最小值，value是多个标记，每个标记表示一个连通子图，整个value表示一个连通图
    // 遍历矩阵
    // 遇到第一个1，标记为2，其实是表示第一个连通图，然后继续遍历
    // 如果当前值为1，检查上下左右，如果有大于1的值，就把当前值设为上下左右的非零最小值。如果上下左右非0非1值多于一个，就把多的加入map中（key是最小的），并且检查map里是否有所有的要合并的keys