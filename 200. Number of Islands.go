func bfs(grid [][]byte, i,j int) {
    if i<0 || i>=len(grid) || j<0 || j>=len(grid[0]) || grid[i][j] != '1' {
        return
    }
    
    grid[i][j] = '2'
    
    bfs(grid, i+1, j)
    bfs(grid, i, j+1)
    bfs(grid, i-1, j)
    bfs(grid, i, j-1)
    
}

func numIslands(grid [][]byte) int {
    count := 0
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j] == '1' {
                count += 1
                bfs(grid, i, j)
            }
        }
    }
    return count
}
