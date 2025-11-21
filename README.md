# 100 Go Programming Questions & Solutions

A comprehensive collection of Go programming exercises ranging from basic to advanced level.

## Table of Contents
- [Basics (1-20)](#basics-1-20)
- [Data Structures (21-40)](#data-structures-21-40)
- [Functions & Methods (41-50)](#functions--methods-41-50)
- [Concurrency (51-65)](#concurrency-51-65)
- [Interfaces & Structs (66-75)](#interfaces--structs-66-75)
- [Error Handling (76-80)](#error-handling-76-80)
- [File I/O (81-85)](#file-io-81-85)
- [Advanced Topics (86-100)](#advanced-topics-86-100)

---

## Basics (1-20)

### 1. Hello World
**Question**: Print "Hello, World!" to console.
```go
package main
import "fmt"
func main() {
    fmt.Println("Hello, World!")
}
```

### 2. Variable Declaration
**Question**: Declare and initialize variables of different types.
```go
package main
import "fmt"
func main() {
    var name string = "Go"
    age := 13
    isAwesome := true
    fmt.Printf("Language: %s, Age: %d, Awesome: %t\n", name, age, isAwesome)
}
```

### 3. Check Even or Odd
**Question**: Check if a number is even or odd.
```go
package main
import "fmt"
func isEven(n int) bool {
    return n%2 == 0
}
func main() {
    num := 42
    if isEven(num) {
        fmt.Printf("%d is even\n", num)
    } else {
        fmt.Printf("%d is odd\n", num)
    }
}
```

### 4. Factorial
**Question**: Calculate factorial of a number.
```go
package main
import "fmt"
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}
func main() {
    fmt.Println("Factorial of 5:", factorial(5))
}
```

### 5. Fibonacci Sequence
**Question**: Generate Fibonacci sequence up to n terms.
```go
package main
import "fmt"
func fibonacci(n int) []int {
    fib := make([]int, n)
    fib[0], fib[1] = 0, 1
    for i := 2; i < n; i++ {
        fib[i] = fib[i-1] + fib[i-2]
    }
    return fib
}
func main() {
    fmt.Println(fibonacci(10))
}
```

### 6. Prime Number Check
**Question**: Check if a number is prime.
```go
package main
import "fmt"
func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
func main() {
    fmt.Println("Is 17 prime?", isPrime(17))
}
```

### 7. Sum of Array
**Question**: Calculate sum of array elements.
```go
package main
import "fmt"
func sumArray(arr []int) int {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return sum
}
func main() {
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println("Sum:", sumArray(nums))
}
```

### 8. Find Maximum
**Question**: Find maximum element in an array.
```go
package main
import "fmt"
func findMax(arr []int) int {
    max := arr[0]
    for _, v := range arr {
        if v > max {
            max = v
        }
    }
    return max
}
func main() {
    nums := []int{3, 7, 2, 9, 1}
    fmt.Println("Max:", findMax(nums))
}
```

### 9. Reverse String
**Question**: Reverse a string.
```go
package main
import "fmt"
func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
func main() {
    fmt.Println(reverseString("Hello"))
}
```

### 10. Palindrome Check
**Question**: Check if a string is palindrome.
```go
package main
import "fmt"
func isPalindrome(s string) bool {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        if runes[i] != runes[j] {
            return false
        }
    }
    return true
}
func main() {
    fmt.Println("Is 'radar' palindrome?", isPalindrome("radar"))
}
```

### 11. Count Vowels
**Question**: Count vowels in a string.
```go
package main
import "fmt"
func countVowels(s string) int {
    vowels := "aeiouAEIOU"
    count := 0
    for _, ch := range s {
        for _, v := range vowels {
            if ch == v {
                count++
                break
            }
        }
    }
    return count
}
func main() {
    fmt.Println("Vowels:", countVowels("Hello World"))
}
```

### 12. Swap Two Numbers
**Question**: Swap two numbers without temp variable.
```go
package main
import "fmt"
func swap(a, b int) (int, int) {
    return b, a
}
func main() {
    x, y := 5, 10
    x, y = swap(x, y)
    fmt.Printf("x=%d, y=%d\n", x, y)
}
```

### 13. Leap Year Check
**Question**: Check if a year is leap year.
```go
package main
import "fmt"
func isLeapYear(year int) bool {
    return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
func main() {
    fmt.Println("Is 2024 leap year?", isLeapYear(2024))
}
```

### 14. Armstrong Number
**Question**: Check if number is Armstrong number.
```go
package main
import ("fmt"; "math")
func isArmstrong(n int) bool {
    original, sum, digits := n, 0, 0
    temp := n
    for temp > 0 {
        digits++
        temp /= 10
    }
    temp = n
    for temp > 0 {
        digit := temp % 10
        sum += int(math.Pow(float64(digit), float64(digits)))
        temp /= 10
    }
    return sum == original
}
func main() {
    fmt.Println("Is 153 Armstrong?", isArmstrong(153))
}
```

### 15. GCD (Greatest Common Divisor)
**Question**: Find GCD of two numbers.
```go
package main
import "fmt"
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}
func main() {
    fmt.Println("GCD(48, 18):", gcd(48, 18))
}
```

### 16. LCM (Least Common Multiple)
**Question**: Find LCM of two numbers.
```go
package main
import "fmt"
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}
func lcm(a, b int) int {
    return (a * b) / gcd(a, b)
}
func main() {
    fmt.Println("LCM(12, 18):", lcm(12, 18))
}
```

### 17. Power Function
**Question**: Calculate power of a number.
```go
package main
import "fmt"
func power(base, exp int) int {
    result := 1
    for i := 0; i < exp; i++ {
        result *= base
    }
    return result
}
func main() {
    fmt.Println("2^8 =", power(2, 8))
}
```

### 18. Sum of Digits
**Question**: Calculate sum of digits.
```go
package main
import "fmt"
func sumDigits(n int) int {
    sum := 0
    for n > 0 {
        sum += n % 10
        n /= 10
    }
    return sum
}
func main() {
    fmt.Println("Sum of digits of 12345:", sumDigits(12345))
}
```

### 19. Reverse Number
**Question**: Reverse a number.
```go
package main
import "fmt"
func reverseNumber(n int) int {
    reversed := 0
    for n > 0 {
        reversed = reversed*10 + n%10
        n /= 10
    }
    return reversed
}
func main() {
    fmt.Println("Reverse of 12345:", reverseNumber(12345))
}
```

### 20. Multiplication Table
**Question**: Print multiplication table.
```go
package main
import "fmt"
func printTable(n int) {
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d x %d = %d\n", n, i, n*i)
    }
}
func main() {
    printTable(5)
}
```

---

## Data Structures (21-40)

### 21. Array Rotation
**Question**: Rotate array by k positions.
```go
package main
import "fmt"
func rotateArray(arr []int, k int) []int {
    n := len(arr)
    k = k % n
    result := make([]int, n)
    copy(result, arr[k:])
    copy(result[n-k:], arr[:k])
    return result
}
func main() {
    arr := []int{1, 2, 3, 4, 5}
    fmt.Println(rotateArray(arr, 2))
}
```

### 22. Remove Duplicates
**Question**: Remove duplicates from slice.
```go
package main
import "fmt"
func removeDuplicates(arr []int) []int {
    seen := make(map[int]bool)
    result := []int{}
    for _, v := range arr {
        if !seen[v] {
            seen[v] = true
            result = append(result, v)
        }
    }
    return result
}
func main() {
    arr := []int{1, 2, 2, 3, 4, 4, 5}
    fmt.Println(removeDuplicates(arr))
}
```

### 23. Merge Two Sorted Arrays
**Question**: Merge two sorted arrays.
```go
package main
import "fmt"
func mergeSorted(arr1, arr2 []int) []int {
    result := make([]int, 0, len(arr1)+len(arr2))
    i, j := 0, 0
    for i < len(arr1) && j < len(arr2) {
        if arr1[i] < arr2[j] {
            result = append(result, arr1[i])
            i++
        } else {
            result = append(result, arr2[j])
            j++
        }
    }
    result = append(result, arr1[i:]...)
    result = append(result, arr2[j:]...)
    return result
}
func main() {
    fmt.Println(mergeSorted([]int{1, 3, 5}, []int{2, 4, 6}))
}
```

### 24. Binary Search
**Question**: Implement binary search.
```go
package main
import "fmt"
func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1
    for left <= right {
        mid := left + (right-left)/2
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
func main() {
    arr := []int{1, 3, 5, 7, 9, 11}
    fmt.Println("Index:", binarySearch(arr, 7))
}
```

### 25. Bubble Sort
**Question**: Implement bubble sort.
```go
package main
import "fmt"
func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}
func main() {
    arr := []int{64, 34, 25, 12, 22}
    bubbleSort(arr)
    fmt.Println(arr)
}
```

### 26. Selection Sort
**Question**: Implement selection sort.
```go
package main
import "fmt"
func selectionSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
}
func main() {
    arr := []int{64, 25, 12, 22, 11}
    selectionSort(arr)
    fmt.Println(arr)
}
```

### 27. Insertion Sort
**Question**: Implement insertion sort.
```go
package main
import "fmt"
func insertionSort(arr []int) {
    for i := 1; i < len(arr); i++ {
        key := arr[i]
        j := i - 1
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = key
    }
}
func main() {
    arr := []int{12, 11, 13, 5, 6}
    insertionSort(arr)
    fmt.Println(arr)
}
```

### 28. Quick Sort
**Question**: Implement quick sort.
```go
package main
import "fmt"
func quickSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
    pivot := arr[len(arr)/2]
    var left, right []int
    for i, v := range arr {
        if i == len(arr)/2 {
            continue
        }
        if v <= pivot {
            left = append(left, v)
        } else {
            right = append(right, v)
        }
    }
    return append(append(quickSort(left), pivot), quickSort(right)...)
}
func main() {
    arr := []int{10, 7, 8, 9, 1, 5}
    fmt.Println(quickSort(arr))
}
```

### 29. Merge Sort
**Question**: Implement merge sort.
```go
package main
import "fmt"
func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    mid := len(arr) / 2
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])
    return merge(left, right)
}
func merge(left, right []int) []int {
    result := make([]int, 0, len(left)+len(right))
    i, j := 0, 0
    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }
    return append(append(result, left[i:]...), right[j:]...)
}
func main() {
    arr := []int{38, 27, 43, 3, 9, 82, 10}
    fmt.Println(mergeSort(arr))
}
```

### 30. Stack Implementation
**Question**: Implement a stack.
```go
package main
import "fmt"
type Stack struct {
    items []int
}
func (s *Stack) Push(item int) {
    s.items = append(s.items, item)
}
func (s *Stack) Pop() int {
    if len(s.items) == 0 {
        return -1
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}
func main() {
    stack := Stack{}
    stack.Push(1)
    stack.Push(2)
    fmt.Println(stack.Pop())
}
```

### 31. Queue Implementation
**Question**: Implement a queue.
```go
package main
import "fmt"
type Queue struct {
    items []int
}
func (q *Queue) Enqueue(item int) {
    q.items = append(q.items, item)
}
func (q *Queue) Dequeue() int {
    if len(q.items) == 0 {
        return -1
    }
    item := q.items[0]
    q.items = q.items[1:]
    return item
}
func main() {
    queue := Queue{}
    queue.Enqueue(1)
    queue.Enqueue(2)
    fmt.Println(queue.Dequeue())
}
```

### 32. Linked List
**Question**: Implement singly linked list.
```go
package main
import "fmt"
type Node struct {
    data int
    next *Node
}
type LinkedList struct {
    head *Node
}
func (ll *LinkedList) Insert(data int) {
    newNode := &Node{data: data}
    if ll.head == nil {
        ll.head = newNode
        return
    }
    current := ll.head
    for current.next != nil {
        current = current.next
    }
    current.next = newNode
}
func (ll *LinkedList) Print() {
    current := ll.head
    for current != nil {
        fmt.Print(current.data, " -> ")
        current = current.next
    }
    fmt.Println("nil")
}
func main() {
    ll := LinkedList{}
    ll.Insert(1)
    ll.Insert(2)
    ll.Insert(3)
    ll.Print()
}
```

### 33. Reverse Linked List
**Question**: Reverse a linked list.
```go
package main
import "fmt"
type Node struct {
    data int
    next *Node
}
func reverseList(head *Node) *Node {
    var prev *Node
    current := head
    for current != nil {
        next := current.next
        current.next = prev
        prev = current
        current = next
    }
    return prev
}
func printList(head *Node) {
    for head != nil {
        fmt.Print(head.data, " -> ")
        head = head.next
    }
    fmt.Println("nil")
}
func main() {
    head := &Node{1, &Node{2, &Node{3, nil}}}
    printList(head)
    reversed := reverseList(head)
    printList(reversed)
}
```

### 34. Detect Cycle in Linked List
**Question**: Detect cycle in linked list.
```go
package main
import "fmt"
type Node struct {
    data int
    next *Node
}
func hasCycle(head *Node) bool {
    slow, fast := head, head
    for fast != nil && fast.next != nil {
        slow = slow.next
        fast = fast.next.next
        if slow == fast {
            return true
        }
    }
    return false
}
func main() {
    node1 := &Node{data: 1}
    node2 := &Node{data: 2}
    node3 := &Node{data: 3}
    node1.next = node2
    node2.next = node3
    node3.next = node1 // Creates cycle
    fmt.Println("Has cycle:", hasCycle(node1))
}
```

### 35. Binary Tree Traversal
**Question**: Implement inorder traversal.
```go
package main
import "fmt"
type TreeNode struct {
    val   int
    left  *TreeNode
    right *TreeNode
}
func inorder(root *TreeNode) {
    if root == nil {
        return
    }
    inorder(root.left)
    fmt.Print(root.val, " ")
    inorder(root.right)
}
func main() {
    root := &TreeNode{val: 1}
    root.left = &TreeNode{val: 2}
    root.right = &TreeNode{val: 3}
    inorder(root)
}
```

### 36. Tree Height
**Question**: Calculate height of binary tree.
```go
package main
import "fmt"
type TreeNode struct {
    val   int
    left  *TreeNode
    right *TreeNode
}
func height(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftHeight := height(root.left)
    rightHeight := height(root.right)
    if leftHeight > rightHeight {
        return leftHeight + 1
    }
    return rightHeight + 1
}
func main() {
    root := &TreeNode{val: 1}
    root.left = &TreeNode{val: 2}
    root.right = &TreeNode{val: 3}
    root.left.left = &TreeNode{val: 4}
    fmt.Println("Height:", height(root))
}
```

### 37. Level Order Traversal
**Question**: Level order traversal of tree.
```go
package main
import "fmt"
type TreeNode struct {
    val   int
    left  *TreeNode
    right *TreeNode
}
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    result := [][]int{}
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        level := []int{}
        size := len(queue)
        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]
            level = append(level, node.val)
            if node.left != nil {
                queue = append(queue, node.left)
            }
            if node.right != nil {
                queue = append(queue, node.right)
            }
        }
        result = append(result, level)
    }
    return result
}
func main() {
    root := &TreeNode{val: 1}
    root.left = &TreeNode{val: 2}
    root.right = &TreeNode{val: 3}
    fmt.Println(levelOrder(root))
}
```

### 38. Hash Map Implementation
**Question**: Simple hash map.
```go
package main
import "fmt"
type HashMap struct {
    data map[string]int
}
func NewHashMap() *HashMap {
    return &HashMap{data: make(map[string]int)}
}
func (h *HashMap) Put(key string, value int) {
    h.data[key] = value
}
func (h *HashMap) Get(key string) (int, bool) {
    val, exists := h.data[key]
    return val, exists
}
func main() {
    hm := NewHashMap()
    hm.Put("age", 25)
    val, _ := hm.Get("age")
    fmt.Println("Age:", val)
}
```

### 39. Two Sum Problem
**Question**: Find two numbers that sum to target.
```go
package main
import "fmt"
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if j, ok := m[complement]; ok {
            return []int{j, i}
        }
        m[num] = i
    }
    return nil
}
func main() {
    nums := []int{2, 7, 11, 15}
    fmt.Println(twoSum(nums, 9))
}
```

### 40. Valid Parentheses
**Question**: Check if parentheses are valid.
```go
package main
import "fmt"
func isValid(s string) bool {
    stack := []rune{}
    pairs := map[rune]rune{')': '(', '}': '{', ']': '['}
    for _, ch := range s {
        if ch == '(' || ch == '{' || ch == '[' {
            stack = append(stack, ch)
        } else {
            if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}
func main() {
    fmt.Println(isValid("()[]{}"))
    fmt.Println(isValid("(]"))
}
```

---

## Functions & Methods (41-50)

### 41. Variadic Functions
**Question**: Sum using variadic function.
```go
package main
import "fmt"
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
func main() {
    fmt.Println(sum(1, 2, 3, 4, 5))
}
```

### 42. Closures
**Question**: Implement counter using closure.
```go
package main
import "fmt"
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
func main() {
    c := counter()
    fmt.Println(c())
    fmt.Println(c())
    fmt.Println(c())
}
```

### 43. Higher Order Functions
**Question**: Map function implementation.
```go
package main
import "fmt"
func mapFunc(arr []int, fn func(int) int) []int {
    result := make([]int, len(arr))
    for i, v := range arr {
        result[i] = fn(v)
    }
    return result
}
func main() {
    nums := []int{1, 2, 3, 4}
    squared := mapFunc(nums, func(n int) int { return n * n })
    fmt.Println(squared)
}
```

### 44. Filter Function
**Question**: Filter even numbers.
```go
package main
import "fmt"
func filter(arr []int, fn func(int) bool) []int {
    result := []int{}
    for _, v := range arr {
        if fn(v) {
            result = append(result, v)
        }
    }
    return result
}
func main() {
    nums := []int{1, 2, 3, 4, 5, 6}
    evens := filter(nums, func(n int) bool { return n%2 == 0 })
    fmt.Println(evens)
}
```

### 45. Reduce Function
**Question**: Reduce array to single value.
```go
package main
import "fmt"
func reduce(arr []int, fn func(int, int) int, initial int) int {
    result := initial
    for _, v := range arr {
        result = fn(result, v)
    }
    return result
}
func main() {
    nums := []int{1, 2, 3, 4, 5}
    sum := reduce(nums, func(a, b int) int { return a + b }, 0)
    fmt.Println(sum)
}
```

### 46. Method on Struct
**Question**: Rectangle area calculation.
```go
package main
import "fmt"
type Rectangle struct {
    width, height float64
}
func (r Rectangle) Area() float64 {
    return r.width * r.height
}
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.width + r.height)
}
func main() {
    rect := Rectangle{width: 10, height: 5}
    fmt.Println("Area:", rect.Area())
    fmt.Println("Perimeter:", rect.Perimeter())
}
```

### 47. Pointer Receivers
**Question**: Modify struct using pointer receiver.
```go
package main
import "fmt"
type Counter struct {
    count int
}
func (c *Counter) Increment() {
    c.count++
}
func (c Counter) Value() int {
    return c.count
}
func main() {
    c := Counter{}
    c.Increment()
    c.Increment()
    fmt.Println("Count:", c.Value())
}
```

### 48. Function as Return Value
**Question**: Return function from function.
```go
package main
import "fmt"
func multiplier(factor int) func(int) int {
    return func(n int) int {
        return n * factor
    }
}
func main() {
    double := multiplier(2)
    triple := multiplier(3)
    fmt.Println(double(5))
    fmt.Println(triple(5))
}
```

### 49. Defer Statement
**Question**: Demonstrate defer usage.
```go
package main
import "fmt"
func main() {
    defer fmt.Println("World")
    defer fmt.Println("Beautiful")
    fmt.Println("Hello")
}
```

### 50. Panic and Recover
**Question**: Handle panic with recover.
```go
package main
import "fmt"
func safeDivide(a, b int) (result int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
            result = 0
        }
    }()
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}
func main() {
    fmt.Println(safeDivide(10, 2))
    fmt.Println(safeDivide(10, 0))
}
```

---

## Concurrency (51-65)

### 51. Simple Goroutine
**Question**: Launch goroutine.
```go
package main
import ("fmt"; "time")
func sayHello() {
    fmt.Println("Hello from goroutine")
}
func main() {
    go sayHello()
    time.Sleep(time.Second)
    fmt.Println("Main function")
}
```

### 52. Multiple Goroutines
**Question**: Launch multiple goroutines.
```go
package main
import ("fmt"; "time")
func printNumber(n int) {
    fmt.Println(n)
}
func main() {
    for i := 1; i <= 5; i++ {
        go printNumber(i)
    }
    time.Sleep(time.Second)
}
```

### 53. Channels
**Question**: Send and receive via channel.
```go
package main
import "fmt"
func main() {
    ch := make(chan int)
    go func() {
        ch <- 42
    }()
    value := <-ch
    fmt.Println("Received:", value)
}
```

### 54. Buffered Channels
**Question**: Use buffered channel.
```go
package main
import "fmt"
func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

### 55. Channel Range
**Question**: Iterate over channel.
```go
package main
import "fmt"
func fibonacci(n int, ch chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        ch <- x
        x, y = y, x+y
    }
    close(ch)
}
func main() {
    ch := make(chan int)
    go fibonacci(10, ch)
    for num := range ch {
        fmt.Println(num)
    }
}
```

### 56. Select Statement
**Question**: Use select with multiple channels.
```go
package main
import ("fmt"; "time")
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "two"
    }()
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Received", msg1)
        case msg2 := <-ch2:
            fmt.Println("Received", msg2)
        }
    }
}
```

### 57. Worker Pool
**Question**: Implement worker pool pattern.
```go
package main
import ("fmt"; "time")
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}
func main() {
    jobs := make(chan int, 5)
    results := make(chan int, 5)
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    for a := 1; a <= 5; a++ {
        <-results
    }
}
```

### 58. Mutex
**Question**: Use mutex for synchronization.
```go
package main
import ("fmt"; "sync")
type SafeCounter struct {
    mu    sync.Mutex
    count int
}
func (c *SafeCounter) Inc() {
    c.mu.Lock()
    c.count++
    c.mu.Unlock()
}
func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}
func main() {
    counter := SafeCounter{}
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Inc()
        }()
    }
    wg.Wait()
    fmt.Println("Count:", counter.Value())
}
```

### 59. WaitGroup
**Question**: Use WaitGroup for synchronization.
```go
package main
import ("fmt"; "sync"; "time")
func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}
func main() {
    var wg sync.WaitGroup
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }
    wg.Wait()
    fmt.Println("All workers done")
}
```

### 60. Rate Limiting
**Question**: Implement rate limiter.
```go
package main
import ("fmt"; "time")
func main() {
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)
    limiter := time.Tick(200 * time.Millisecond)
    for req := range requests {
        <-limiter
        fmt.Println("Request", req, time.Now())
    }
}
```

### 61. Atomic Operations
**Question**: Use atomic operations.
```go
package main
import ("fmt"; "sync"; "sync/atomic")
func main() {
    var counter int64
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            atomic.AddInt64(&counter, 1)
        }()
    }
    wg.Wait()
    fmt.Println("Counter:", counter)
}
```

### 62. Context Usage
**Question**: Use context for cancellation.
```go
package main
import ("context"; "fmt"; "time")
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Worker cancelled")
            return
        default:
            fmt.Println("Working...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}
func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    go worker(ctx)
    time.Sleep(3 * time.Second)
}
```

### 63. Fan-Out Fan-In
**Question**: Implement fan-out fan-in pattern.
```go
package main
import ("fmt"; "sync")
func producer(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}
func merge(cs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    for _, c := range cs {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for n := range ch {
                out <- n
            }
        }(c)
    }
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}
func main() {
    in := producer(1, 2, 3, 4)
    c1 := square(in)
    c2 := square(in)
    for n := range merge(c1, c2) {
        fmt.Println(n)
    }
}
```

### 64. Pipeline Pattern
**Question**: Create processing pipeline.
```go
package main
import "fmt"
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}
func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}
func main() {
    for n := range sq(sq(gen(2, 3))) {
        fmt.Println(n)
    }
}
```

### 65. Semaphore Pattern
**Question**: Implement semaphore.
```go
package main
import ("fmt"; "time")
func main() {
    sem := make(chan struct{}, 3)
    for i := 1; i <= 10; i++ {
        sem <- struct{}{}
        go func(id int) {
            defer func() { <-sem }()
            fmt.Printf("Task %d running\n", id)
            time.Sleep(time.Second)
            fmt.Printf("Task %d done\n", id)
        }(i)
    }
    time.Sleep(5 * time.Second)
}
```

---

## Interfaces & Structs (66-75)

### 66. Interface Definition
**Question**: Define and implement interface.
```go
package main
import "fmt"
type Shape interface {
    Area() float64
}
type Circle struct {
    radius float64
}
func (c Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}
func main() {
    var s Shape = Circle{radius: 5}
    fmt.Println("Area:", s.Area())
}
```

### 67. Multiple Interfaces
**Question**: Implement multiple interfaces.
```go
package main
import "fmt"
type Reader interface {
    Read() string
}
type Writer interface {
    Write(string)
}
type File struct {
    content string
}
func (f *File) Read() string {
    return f.content
}
func (f *File) Write(s string) {
    f.content = s
}
func main() {
    f := &File{}
    f.Write("Hello")
    fmt.Println(f.Read())
}
```

### 68. Empty Interface
**Question**: Use empty interface.
```go
package main
import "fmt"
func describe(i interface{}) {
    fmt.Printf("Type: %T, Value: %v\n", i, i)
}
func main() {
    describe(42)
    describe("hello")
    describe(true)
}
```

### 69. Type Assertion
**Question**: Perform type assertion.
```go
package main
import "fmt"
func main() {
    var i interface{} = "hello"
    s, ok := i.(string)
    if ok {
        fmt.Println("String:", s)
    }
    n, ok := i.(int)
    if !ok {
        fmt.Println("Not an int:", n)
    }
}
```

### 70. Type Switch
**Question**: Use type switch.
```go
package main
import "fmt"
func do(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Int: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Bool: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
func main() {
    do(21)
    do("hello")
    do(true)
}
```

### 71. Embedded Structs
**Question**: Use struct embedding.
```go
package main
import "fmt"
type Person struct {
    name string
    age  int
}
type Employee struct {
    Person
    salary int
}
func main() {
    emp := Employee{
        Person: Person{name: "John", age: 30},
        salary: 50000,
    }
    fmt.Printf("%s is %d years old\n", emp.name, emp.age)
}
```

### 72. Interface Composition
**Question**: Compose interfaces.
```go
package main
import "fmt"
type Reader interface {
    Read() string
}
type Writer interface {
    Write(string)
}
type ReadWriter interface {
    Reader
    Writer
}
type Buffer struct {
    data string
}
func (b *Buffer) Read() string {
    return b.data
}
func (b *Buffer) Write(s string) {
    b.data = s
}
func main() {
    var rw ReadWriter = &Buffer{}
    rw.Write("Hello")
    fmt.Println(rw.Read())
}
```

### 73. Stringer Interface
**Question**: Implement Stringer interface.
```go
package main
import "fmt"
type Person struct {
    name string
    age  int
}
func (p Person) String() string {
    return fmt.Sprintf("%s (%d years)", p.name, p.age)
}
func main() {
    p := Person{name: "Alice", age: 25}
    fmt.Println(p)
}
```

### 74. Error Interface
**Question**: Implement custom error.
```go
package main
import "fmt"
type MyError struct {
    msg string
}
func (e *MyError) Error() string {
    return e.msg
}
func doSomething() error {
    return &MyError{msg: "something went wrong"}
}
func main() {
    if err := doSomething(); err != nil {
        fmt.Println("Error:", err)
    }
}
```

### 75. Interface Polymorphism
**Question**: Demonstrate polymorphism.
```go
package main
import "fmt"
type Animal interface {
    Speak() string
}
type Dog struct{}
func (d Dog) Speak() string {
    return "Woof!"
}
type Cat struct{}
func (c Cat) Speak() string {
    return "Meow!"
}
func main() {
    animals := []Animal{Dog{}, Cat{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
```

---

## Error Handling (76-80)

### 76. Basic Error Handling
**Question**: Handle errors properly.
```go
package main
import ("errors"; "fmt")
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```

### 77. Error Wrapping
**Question**: Wrap errors with context.
```go
package main
import ("fmt"; "errors")
func readFile() error {
    return errors.New("file not found")
}
func processFile() error {
    if err := readFile(); err != nil {
        return fmt.Errorf("processFile: %w", err)
    }
    return nil
}
func main() {
    if err := processFile(); err != nil {
        fmt.Println(err)
    }
}
```

### 78. Custom Error Types
**Question**: Create custom error type.
```go
package main
import "fmt"
type ValidationError struct {
    Field string
    Msg   string
}
func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}
func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{Field: "age", Msg: "must be positive"}
    }
    return nil
}
func main() {
    if err := validateAge(-5); err != nil {
        fmt.Println("Validation error:", err)
    }
}
```

### 79. Error Checking with errors.Is
**Question**: Use errors.Is for comparison.
```go
package main
import ("errors"; "fmt")
var ErrNotFound = errors.New("not found")
func findItem(id int) error {
    if id != 1 {
        return fmt.Errorf("item %d: %w", id, ErrNotFound)
    }
    return nil
}
func main() {
    err := findItem(2)
    if errors.Is(err, ErrNotFound) {
        fmt.Println("Item not found")
    }
}
```

### 80. Error Type Checking with errors.As
**Question**: Use errors.As for type assertion.
```go
package main
import ("errors"; "fmt")
type QueryError struct {
    Query string
    Err   error
}
func (e *QueryError) Error() string {
    return fmt.Sprintf("query error: %s", e.Query)
}
func executeQuery() error {
    return &QueryError{Query: "SELECT *", Err: errors.New("timeout")}
}
func main() {
    err := executeQuery()
    var qe *QueryError
    if errors.As(err, &qe) {
        fmt.Printf("Query failed: %s\n", qe.Query)
    }
}
```

---

## File I/O (81-85)

### 81. Read File
**Question**: Read entire file.
```go
package main
import ("fmt"; "os")
func main() {
    data, err := os.ReadFile("test.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(string(data))
}
```

### 82. Write File
**Question**: Write to file.
```go
package main
import ("fmt"; "os")
func main() {
    data := []byte("Hello, Go!")
    err := os.WriteFile("output.txt", data, 0644)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("File written successfully")
}
```

### 83. Append to File
**Question**: Append content to file.
```go
package main
import ("fmt"; "os")
func main() {
    f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer f.Close()
    if _, err := f.WriteString("New log entry\n"); err != nil {
        fmt.Println("Error:", err)
    }
}
```

### 84. Read File Line by Line
**Question**: Read file line by line.
```go
package main
import ("bufio"; "fmt"; "os")
func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error:", err)
    }
}
```

### 85. Check File Exists
**Question**: Check if file exists.
```go
package main
import ("fmt"; "os")
func fileExists(filename string) bool {
    _, err := os.Stat(filename)
    return !os.IsNotExist(err)
}
func main() {
    if fileExists("test.txt") {
        fmt.Println("File exists")
    } else {
        fmt.Println("File does not exist")
    }
}
```

---

## Advanced Topics (86-100)

### 86. JSON Marshaling
**Question**: Convert struct to JSON.
```go
package main
import ("encoding/json"; "fmt")
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
func main() {
    p := Person{Name: "Alice", Age: 25}
    jsonData, _ := json.Marshal(p)
    fmt.Println(string(jsonData))
}
```

### 87. JSON Unmarshaling
**Question**: Parse JSON to struct.
```go
package main
import ("encoding/json"; "fmt")
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
func main() {
    jsonStr := `{"name":"Bob","age":30}`
    var p Person
    json.Unmarshal([]byte(jsonStr), &p)
    fmt.Printf("%+v\n", p)
}
```

### 88. HTTP GET Request
**Question**: Make HTTP GET request.
```go
package main
import ("fmt"; "io"; "net/http")
func main() {
    resp, err := http.Get("https://api.github.com")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
}
```

### 89. HTTP Server
**Question**: Create simple HTTP server.
```go
package main
import ("fmt"; "net/http")
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

### 90. Regular Expressions
**Question**: Use regex for pattern matching.
```go
package main
import ("fmt"; "regexp")
func main() {
    re := regexp.MustCompile(`\d+`)
    fmt.Println(re.FindString("Age: 25"))
    fmt.Println(re.FindAllString("1 2 3 4", -1))
}
```

### 91. Time Formatting
**Question**: Format and parse time.
```go
package main
import ("fmt"; "time")
func main() {
    now := time.Now()
    fmt.Println(now.Format("2006-01-02 15:04:05"))
    t, _ := time.Parse("2006-01-02", "2024-01-15")
    fmt.Println(t)
}
```

### 92. Timers
**Question**: Use timers.
```go
package main
import ("fmt"; "time")
func main() {
    timer := time.NewTimer(2 * time.Second)
    fmt.Println("Waiting...")
    <-timer.C
    fmt.Println("Timer expired")
}
```

### 93. Tickers
**Question**: Use tickers for periodic tasks.
```go
package main
import ("fmt"; "time")
func main() {
    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)
    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()
    time.Sleep(2 * time.Second)
    ticker.Stop()
    done <- true
}
```

### 94. Command Line Arguments
**Question**: Parse command line args.
```go
package main
import ("fmt"; "os")
func main() {
    args := os.Args[1:]
    for i, arg := range args {
        fmt.Printf("Arg %d: %s\n", i, arg)
    }
}
```

### 95. Flag Package
**Question**: Use flag package for CLI.
```go
package main
import ("flag"; "fmt")
func main() {
    name := flag.String("name", "World", "a name")
    age := flag.Int("age", 0, "an age")
    flag.Parse()
    fmt.Printf("Hello, %s! Age: %d\n", *name, *age)
}
```

### 96. Sorting
**Question**: Sort slices.
```go
package main
import ("fmt"; "sort")
func main() {
    nums := []int{4, 2, 7, 1, 9}
    sort.Ints(nums)
    fmt.Println(nums)
    strs := []string{"banana", "apple", "cherry"}
    sort.Strings(strs)
    fmt.Println(strs)
}
```

### 97. Custom Sorting
**Question**: Sort with custom comparator.
```go
package main
import ("fmt"; "sort")
type Person struct {
    Name string
    Age  int
}
type ByAge []Person
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func main() {
    people := []Person{{"Bob", 30}, {"Alice", 25}, {"Charlie", 35}}
    sort.Sort(ByAge(people))
    fmt.Println(people)
}
```

### 98. Reflection
**Question**: Use reflection to inspect types.
```go
package main
import ("fmt"; "reflect")
func main() {
    var x float64 = 3.4
    v := reflect.ValueOf(x)
    fmt.Println("Type:", v.Type())
    fmt.Println("Kind:", v.Kind())
    fmt.Println("Value:", v.Float())
}
```

### 99. Testing
**Question**: Write unit test.
```go
// math.go
package main
func Add(a, b int) int {
    return a + b
}

// math_test.go
package main
import "testing"
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, got %d", result)
    }
}
```

### 100. Benchmarking
**Question**: Write benchmark test.
```go
package main
import "testing"
func Fibonacci(n int) int {
    if n < 2 {
        return n
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}
func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fibonacci(10)
    }
}
```

---

## Contributing
Feel free to add more questions or improve existing solutions!

## License
MIT License
