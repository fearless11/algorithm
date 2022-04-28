## 如何衡量一个算法？

- 算法执行效率：时间
- 算法的内存消耗：空间
- 算法的稳定性：随着数据规模增长的变化，执行效率和内存消耗的稳定性

## 时间复杂度

Time complexity：描述算法在运行过程中消耗的时间。

大O符合表示法，即 T(n) = O(f(n))

```go
// 假设每行代码执行时间为t1，计算执行时间

// 常数阶O(1)
// T(n)=(1+1)*t1，不会随者变量增长而增长，T(n)=1
i, j := 1, 2 // 执行1次
z := i + y    // 执行1次

// 线性阶O(n)
// T(n)=(2n)*t1，当n趋近无限大时，T(n)=O(n)
for i = 1; i <= n; i++ { 
    j := i // 执行n次
    j++   // 执行n次
}

// 对数阶O(logN)
// T(n)=log(2n)*t1，当n趋近无限大时，T(n)=O(logN) 
for i := 1; i < 2*n; { 
    i = i * 10 // 执行log(2n)次，因为10^x > 2n，x=log(2n)
}

// 平方阶O(n^2)
// T(n)=m*n*t1，当m，n趋近无限大时，T(n)=O(n^2) 
for i = 1; i <= m; i++ { 
    for j = 1; j <= n; j++ { 
        k := i + j // 执行m*n次
    }
}
```

## 空间复杂度

Space complexity：描述算法在运行过程中临时占用存储空间大小。

大O符合表示法，即 S(n) = O(f(n))

```go
// 常数阶O(1)
// i，j，z所分配的空间不会随着数据量变化，S(n)=O(1)
i, j := 1, 2 
z := i + y    

// 线性阶O(n)
// S(n)=O(2n)，当n趋近无限大时，S(n)=O(n)
m := make([]int, 2*n)  // S(n)=O(2n)
for i := 0; i < n; i++ {
    j := m[i]
}
```

## leetcode
- https://leetcode-cn.com/study-plan/algorithms/?progress=nufcu6d

