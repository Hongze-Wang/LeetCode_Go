// 233. Number of Digit One
// int型实际大小依赖于机器能支持的最大位数 64位机器上就是64位相当于 int64

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func countDigitOne(n int) int {
    var count, i int
    for i=1; i<=n; i*=10 {
        divider := i*10
        count += (n / divider) * i + min(max(n%divider - i + 1, 0), i)
    }
    return count
}