package main

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 1. Print "Hello, World!"
func helloWorld() {
	fmt.Println("Hello, World!")
}

// 2. Sum of two numbers
func sum(a, b int) int {
	return a + b
}

// 3. Check if a number is even
func isEven(n int) bool {
	return n%2 == 0
}

// 4. Factorial of a number
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 5. Reverse a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 6. Check if a string is palindrome
func isPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return s == reverseString(s)
}

// 7. Find maximum element in an array
func maxElement(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

// 8. Fibonacci sequence nth number
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// 9. Swap two numbers without temp variable
func swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

// 10. Check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 11. Count vowels in a string
func countVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, c := range s {
		if strings.ContainsRune(vowels, c) {
			count++
		}
	}
	return count
}

// 12. Remove duplicates from an array
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

// 13. Merge two sorted arrays
func mergeSortedArrays(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

// 14. Find first non-repeating character
func firstNonRepeating(s string) rune {
	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}
	for _, c := range s {
		if count[c] == 1 {
			return c
		}
	}
	return -1
}

// 15. Check if two strings are anagrams
func areAnagrams(s1, s2 string) bool {
	s1 = strings.ToLower(strings.ReplaceAll(s1, " ", ""))
	s2 = strings.ToLower(strings.ReplaceAll(s2, " ", ""))
	if len(s1) != len(s2) {
		return false
	}
	count := make(map[rune]int)
	for _, c := range s1 {
		count[c]++
	}
	for _, c := range s2 {
		count[c]--
		if count[c] < 0 {
			return false
		}
	}
	return true
}

// 16. Find GCD of two numbers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 17. Check if a number is power of 2
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// 18. Rotate an array by k positions
func rotateArray(arr []int, k int) []int {
	k = k % len(arr)
	reverse := func(a []int, start, end int) {
		for start < end {
			a[start], a[end] = a[end], a[start]
			start++
			end--
		}
	}
	reverse(arr, 0, len(arr)-1)
	reverse(arr, 0, k-1)
	reverse(arr, k, len(arr)-1)
	return arr
}

// 19. Find missing number in array 1 to n
func findMissingNumber(arr []int, n int) int {
	expected := n * (n + 1) / 2
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return expected - sum
}

// 20. Implement binary search
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

// 21. Count occurrences of a number in sorted array
func countOccurrences(arr []int, target int) int {
	count := 0
	for _, v := range arr {
		if v == target {
			count++
		}
	}
	return count
}

// 22. Convert string to integer (atoi)
func atoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	sign := 1
	start := 0
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}
	result := 0
	for i := start; i < len(s); i++ {
		if !isDigit(s[i]) {
			break
		}
		result = result*10 + int(s[i]-'0')
	}
	return sign * result
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// 23. Generate all permutations of a string
func permutations(s string) []string {
	var result []string
	var backtrack func([]rune, int)
	backtrack = func(chars []rune, start int) {
		if start == len(chars) {
			result = append(result, string(chars))
			return
		}
		for i := start; i < len(chars); i++ {
			chars[start], chars[i] = chars[i], chars[start]
			backtrack(chars, start+1)
			chars[start], chars[i] = chars[i], chars[start]
		}
	}
	backtrack([]rune(s), 0)
	return result
}

// 24. Check if a string is valid parenthesis
func isValidParenthesis(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if (c == ')' && last != '(') || (c == '}' && last != '{') || (c == ']' && last != '[') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 25. Find longest common prefix
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

// 26. Implement a stack
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

// 27. Implement a queue
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

// 28. Reverse a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 29. Detect cycle in a linked list
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 30. Merge two sorted linked lists
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}
	return dummy.Next
}

// 31. Find intersection of two linked lists
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	for a != b {
		a = a.Next
		b = b.Next
		if a == nil && b == nil {
			return nil
		}
		if a == nil {
			a = headB
		}
		if b == nil {
			b = headA
		}
	}
	return a
}

// 32. Implement a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 33. Preorder traversal of binary tree
func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return result
}

// 34. Inorder traversal of binary tree
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return result
}

// 35. Postorder traversal of binary tree
func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		result = append(result, node.Val)
	}
	postorder(root)
	return result
}

// 36. Level order traversal of binary tree
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []int{}
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, currentLevel)
	}
	return result
}

// 37. Find maximum depth of binary tree
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return max(leftDepth, rightDepth) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 38. Check if binary tree is balanced
func isBalanced(root *TreeNode) bool {
	var checkHeight func(*TreeNode) int
	checkHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftHeight := checkHeight(node.Left)
		if leftHeight == -1 {
			return -1
		}
		rightHeight := checkHeight(node.Right)
		if rightHeight == -1 {
			return -1
		}
		if abs(leftHeight-rightHeight) > 1 {
			return -1
		}
		return max(leftHeight, rightHeight) + 1
	}
	return checkHeight(root) != -1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 39. Lowest common ancestor in binary tree
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

// 40. Check if binary tree is a BST
func isValidBST(root *TreeNode) bool {
	var isBST func(*TreeNode, *int, *int) bool
	isBST = func(node *TreeNode, min, max *int) bool {
		if node == nil {
			return true
		}
		if min != nil && node.Val <= *min {
			return false
		}
		if max != nil && node.Val >= *max {
			return false
		}
		return isBST(node.Left, min, &node.Val) && isBST(node.Right, &node.Val, max)
	}
	return isBST(root, nil, nil)
}

// 41. Implement a trie
type Trie struct {
	children map[rune]*Trie
	isEnd    bool
}

func NewTrie() *Trie {
	return &Trie{children: make(map[rune]*Trie)}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = NewTrie()
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}
	return node.isEnd
}

// 42. Find all subsets of a set
func subsets(nums []int) [][]int {
	result := [][]int{{}}
	for _, num := range nums {
		for _, subset := range result {
			newSubset := append([]int{}, subset...)
			newSubset = append(newSubset, num)
			result = append(result, newSubset)
		}
	}
	return result
}

// 43. Find kth largest element in an array
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 44. Implement LRU cache
type LRUCache struct {
	capacity int
	cache    map[int]*DListNode
	head     *DListNode
	tail     *DListNode
}

type DListNode struct {
	key, value int
	prev, next *DListNode
}

func NewLRUCache(capacity int) *LRUCache {
	lru := &LRUCache{
		cache:    make(map[int]*DListNode),
		capacity: capacity,
		head:     &DListNode{},
		tail:     &DListNode{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		lru.moveToHead(node)
		return node.value
	}
	return -1
}

func (lru *LRUCache) Put(key, value int) {
	if node, ok := lru.cache[key]; ok {
		node.value = value
		lru.moveToHead(node)
		return
	}
	newNode := &DListNode{key: key, value: value}
	lru.cache[key] = newNode
	lru.addToHead(newNode)
	if len(lru.cache) > lru.capacity {
		delete(lru.cache, lru.tail.prev.key)
		lru.removeNode(lru.tail.prev)
	}
}

func (lru *LRUCache) addToHead(node *DListNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) removeNode(node *DListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) moveToHead(node *DListNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

// 45. Implement a min heap
type MinHeap []int

func NewMinHeap() *MinHeap {
	h := &MinHeap{}
	return h
}

func (h *MinHeap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

func (h *MinHeap) Pop() int {
	if len(*h) == 0 {
		return -1
	}
	min := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.down(0)
	return min
}

func (h *MinHeap) up(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if (*h)[i] >= (*h)[parent] {
			break
		}
		(*h)[i], (*h)[parent] = (*h)[parent], (*h)[i]
		i = parent
	}
}

func (h *MinHeap) down(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i
		if left < len(*h) && (*h)[left] < (*h)[smallest] {
			smallest = left
		}
		if right < len(*h) && (*h)[right] < (*h)[smallest] {
			smallest = right
		}
		if smallest == i {
			break
		}
		(*h)[i], (*h)[smallest] = (*h)[smallest], (*h)[i]
		i = smallest
	}
}

// 46. Find median from data stream
type MedianFinder struct {
	minHeap MinHeap
	maxHeap MinHeap
}

func NewMedianFinder() *MedianFinder {
	return &MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	if len(mf.maxHeap) == 0 || num < mf.maxHeap[0] {
		mf.maxHeap.Push(-num)
	} else {
		mf.minHeap.Push(num)
	}
	if len(mf.maxHeap) > len(mf.minHeap)+1 {
		mf.minHeap.Push(-mf.maxHeap.Pop())
	} else if len(mf.minHeap) > len(mf.maxHeap) {
		mf.maxHeap.Push(-mf.minHeap.Pop())
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if len(mf.maxHeap) > len(mf.minHeap) {
		return float64(-mf.maxHeap[0])
	}
	return float64(-mf.maxHeap[0]+mf.minHeap[0]) / 2
}

// 47. Reverse words in a string
func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// 48. Implement a goroutine to print numbers
func printNumbers() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	for num := range ch {
		fmt.Println(num)
	}
}

// 49. Use sync.WaitGroup for goroutines
func waitGroupExample() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d\n", n)
		}(i)
	}
	wg.Wait()
}

// 50. Implement a mutex for shared resource
func mutexExample() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}

// 51. Use channels for producer-consumer
func producerConsumer() {
	ch := make(chan int, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 3; i++ {
			ch <- i
			fmt.Println("Produced:", i)
		}
		close(ch)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Println("Consumed:", num)
		}
	}()
	wg.Wait()
}

// 52. Implement context with timeout
func contextTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Task cancelled:", ctx.Err())
	}
}

// 53. Generate random numbers
func randomNumbers(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = rand.Intn(100)
	}
	return result
}

// 54. Sort an array using sort package
func sortArray(arr []int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)
	return sorted
}

// 55. Check if two slices are equal
func areSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// 56. Find intersection of two arrays
func intersection(nums1, nums2 []int) []int {
	set := make(map[int]bool)
	for _, num := range nums1 {
		set[num] = true
	}
	result := []int{}
	for _, num := range nums2 {
		if set[num] {
			result = append(result, num)
			delete(set, num)
		}
	}
	return result
}

// 57. Implement a circular buffer
type CircularBuffer struct {
	items []int
	size  int
	head  int
	tail  int
	count int
}

func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{
		items: make([]int, size),
		size:  size,
	}
}

func (cb *CircularBuffer) Push(item int) bool {
	if cb.count == cb.size {
		return false
	}
	cb.items[cb.tail] = item
	cb.tail = (cb.tail + 1) % cb.size
	cb.count++
	return true
}

func (cb *CircularBuffer) Pop() (int, bool) {
	if cb.count == 0 {
		return 0, false
	}
	item := cb.items[cb.head]
	cb.head = (cb.head + 1) % cb.size
	cb.count--
	return item, true
}

// 58. Convert integer to Roman numeral
func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""
	for i := 0; i < len(values) && num > 0; i++ {
		for num >= values[i] {
			result += symbols[i]
			num -= values[i]
		}
	}
	return result
}

// 59. Find longest substring without repeating characters
func lengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int)
	maxLen, start := 0, 0
	for i, c := range s {
		if last, ok := seen[c]; ok && last >= start {
			start = last + 1
		} else {
			maxLen = max(maxLen, i-start+1)
		}
		seen[c] = i
	}
	return maxLen
}

// 60. Implement a bloom filter
type BloomFilter struct {
	bits []bool
	k    int
}

func NewBloomFilter(size, k int) *BloomFilter {
	return &BloomFilter{
		bits: make([]bool, size),
		k:    k,
	}
}

func (bf *BloomFilter) Add(item string) {
	for i := 0; i < bf.k; i++ {
		index := hash(item, i) % len(bf.bits)
		bf.bits[index] = true
	}
}

func (bf *BloomFilter) MightContain(item string) bool {
	for i := 0; i < bf.k; i++ {
		index := hash(item, i) % len(bf.bits)
		if !bf.bits[index] {
			return false
		}
	}
	return true
}

func hash(s string, seed int) int {
	h := 0
	for _, c := range s {
		h = h*31 + int(c) + seed
	}
	return h
}

// 61. Check if number is palindrome
func isNumberPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reversed := 0
	original := x
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return original == reversed
}

// 62. Find all prime numbers up to n (Sieve of Eratosthenes)
func sieveOfEratosthenes(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	n := arr[0]
	for _, v := range arr {
		if v > n {
			n = v
		}
	}
	sieve := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		sieve[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if sieve[i] {
			for j := i * i; j <= n; j += i {
				sieve[j] = false
			}
		}
	}
	primes := []int{}
	for _, v := range arr {
		if v >= 2 && sieve[v] {
			primes = append(primes, v)
		}
	}
	return primes
}

// 63. Implement quicksort
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	left := []int{}
	right := []int{}
	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

// 64. Find shortest path using BFS
func shortestPath(graph map[int][]int, start, end int) int {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true
	distance := make(map[int]int)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
				distance[neighbor] = distance[node] + 1
				if neighbor == end {
					return distance[neighbor]
				}
			}
		}
	}
	return -1
}

// 65. Implement a priority queue using heap
type PriorityQueue struct {
	items []struct{ val, priority int }
}

func (pq *PriorityQueue) Push(val, priority int) {
	pq.items = append(pq.items, struct{ val, priority int }{val, priority})
	heapifyUp(pq, len(pq.items)-1)
}

func (pq *PriorityQueue) Pop() (int, int) {
	if len(pq.items) == 0 {
		return 0, -1
	}
	item := pq.items[0]
	pq.items[0] = pq.items[len(pq.items)-1]
	pq.items = pq.items[:len(pq.items)-1]
	heapifyDown(pq, 0)
	return item.val, item.priority
}

func heapifyUp(pq *PriorityQueue, i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if pq.items[i].priority <= pq.items[parent].priority {
			break
		}
		pq.items[i], pq.items[parent] = pq.items[parent], pq.items[i]
		i = parent
	}
}

func heapifyDown(pq *PriorityQueue, i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i
		if left < len(pq.items) && pq.items[left].priority > pq.items[largest].priority {
			largest = left
		}
		if right < len(pq.items) && pq.items[right].priority > pq.items[largest].priority {
			largest = right
		}
		if largest == i {
			break
		}
		pq.items[i], pq.items[largest] = pq.items[largest], pq.items[i]
		i = largest
	}
}

// 66. Find longest increasing path in matrix
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	rows := len(matrix)
	cols := len(matrix[0])
	cache := make(map[[2]int]int)
	var dfs func(int, int, int) int
	dfs = func(row, col, prev int) int {
		if row < 0 || row >= rows || col >= 0 || col >= cols || matrix[row][col] <= prev {
			return 0
		}
		if val, ok := cache[[2]int{row, col}]; ok {
			return val
		}
		path := 1
		dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, dir := range dirs {
			path = max(path, 1+dfs(row+dir[0], col+dir[0]+dir[1], matrix[row][col]))
		}
		cache[[2]int{row, col}] = path
		return path
	}
	maxPath := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			maxPath = max(maxPath, dfs(i, j, math.MinInt))
		}
	}
	return maxPath
}

// 67. Implement a concurrent counter
type ConcurrentCounter struct {
	count int
	mutex sync.Mutex
}

func (c *ConcurrentCounter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
}

func (c *ConcurrentCounter) Count() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.count
}

// 68. Find all anagrams in a string
func findAnagrams(s, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	result := []int{}
	countP := make(map[rune]int)
	countS := make(map[rune]int)
	for _, c := range p {
		countP[c]++
	}
	for _, c := range s[:len(p)] {
		countS[c]++
	}
	if reflect.DeepEqual(countP, countS) == true {
		result = append(result, 0)
	}
	for i := len(p); i < len(s); i++ {
		countS[rune(s[i-len(p)])]--
		if countS[rune(s[i-len(p)])] == 0 {
			delete(countS, rune(s[i-len(p)]))
		}
		countS[rune(s[i])]++
		if reflect.DeepEqual(countP, countS) {
			result = append(result, i-len(p)+1)
		}
	}
	return result
}

// 69. Implement a rate limiter
func rateLimiter() {
	limiter := make(chan struct{}, 2) // 2 requests/sec)
	go func() {
		ticker := time.NewTicker(time.Second / 2)
		defer ticker.Stop()
		for {
			<-ticker.C
			limiter <- struct{}{}
		}
	}()
	for i := 0; i < 5; i++ {
		<-limiter
		fmt.Printf("Request %d processed at %s\n", i+1, time.Now())
	}
}

// 70. Find shortest palindrome
func shortestPalindrome(s string) string {
	rev := reverseString(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[:i+1] == rev[len(s)-i-1:] {
			return rev[:len(s)-i-1] + s
		}
	}
	return s
}

// 71. Implement a number to binary string
func numberToBinary(n int) string {
	if n == 0 {
		return "0"
	}
	binary := ""
	for n > 0 {
		binary = strconv.Itoa(n%2) + binary
		n /= 2
	}
	return binary
}

// 72. Find missing ranges in array
func findMissingRanges(nums []int, lower, upper int) []string {
	result := []string{}
	if len(nums) == 0 {
		return []string{fmt.Sprintf("%d->%d", lower, upper)}
	}
	if nums[0] > lower {
		result = append(result, fmt.Sprintf("%d->%d", lower, nums[0]-1))
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			result = append(result, fmt.Sprintf("%d->%d", nums[i-1]+1, nums[i]-1))
		}
	}
	if nums[len(nums)-1] < upper {
		result = append(result, fmt.Sprintf("%d->%d", nums[len(nums)-1]+1, upper))
	}
	return result
}

// 73. Implement a thread-safe map
type SafeMap struct {
	m     map[string]int
	mutex sync.Mutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	v, ok := sm.m[key]
	return v, ok
}

// 74. Find next greater element
func nextGreaterElement(nums1, nums2 []int) []int {
	result := make([]int, len(nums1))
	stack := []int{}
	m := make(map[int]int)
	for _, num := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			m[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}
	for i, num := range nums1 {
		if next, ok := m[num]; ok {
			result[i] = next
		} else {
			result[i] = -1
		}
	}
	return result
}

// 75. Implement a JSON encoder
func simpleJSONEncoder(data map[string]interface{}) string {
	var builder strings.Builder
	builder.WriteString("{")
	first := true
	for k, v := range data {
		if !first {
			builder.WriteString(",")
		}
		first = false
		builder.WriteString(fmt.Sprintf("\"%s\":", k))
		switch val := v.(type) {
		case string:
			builder.WriteString(fmt.Sprintf("\"%s\"", val))
		case int:
			builder.WriteString(strconv.Itoa(val))
		default:
			builder.WriteString("null")
		}
	}
	builder.WriteString("}")
	return builder.String()
}

// 76. Find longest valid parentheses
func longestValidParentheses(s string) int {
	stack := []int{-1}
	maxLen := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLen = max(maxLen, i-stack[len(stack)-1])
			}
		}
	}
	return maxLen
}

// 77. Implement a sliding window maximum
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	result := []int{}
	deque := []int{}
	for i := 0; i < len(nums); i++ {
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}

// 78. Find minimum window substring
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	count := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		count[t[i]]++
	}
	needed := len(count)
	formed := 0
	left, right := 0, 0
	minLen := math.MaxInt32
	minLeft := 0
	windowCounts := make(map[byte]int)
	for right < len(s) {
		c := s[right]
		windowCounts[c]++
		if _, ok := count[c]; ok && windowCounts[c] == count[c] {
			formed++
		}
		for left <= right && formed == needed {
			if right-left+1 < minLen {
				minLen = right - left + 1
				minLeft = left
			}
			windowCounts[s[left]]--
			if _, ok := count[s[left]]; ok && windowCounts[s[left]] < count[s[left]] {
				formed--
			}
			left++
		}
		right++
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[minLeft : minLeft+minLen]
}

// 79. Implement a task scheduler
type TaskScheduler struct {
	tasks chan func()
	wg    sync.WaitGroup
}

func newTaskScheduler(numWorkers int) *TaskScheduler {
	ts := &TaskScheduler{
		tasks: make(chan func(), 100),
	}
	for i := 0; i < numWorkers; i++ {
		go ts.worker()
	}
	return ts
}

func (ts *TaskScheduler) Submit(task func()) {
	ts.tasks <- task
}

func (ts *TaskScheduler) worker() {
	for task := range ts.tasks {
		task()
	}
}

// 80. Find first missing positive integer
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

// 81. Implement a regular expression matcher
func isRegularMatch(s, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	firstMatch := len(s) > 0 && (p[0] == s[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return isRegularMatch(p[2:], s) || firstMatch && isRegularMatch(s[1:], s[1:])
	}
	return firstMatch && isRegularMatch(p[1:], s[1:])
}

// 82. Find maximum profit from stock prices
func maxStockProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for _, price := range prices[1:] {
		if price < minPrice {
			minPrice = price
		} else {
			if price-minPrice > maxProfit {
				maxProfit = price - minPrice
			}
		}
	}
	return maxProfit
}

// 83. Implement a LFU cache
type LFUCache struct {
	cache     map[int]*NodeLFU
	freqLists map[int]*DListLFU
	minFreq   int
	capacity  int
}

type NodeLFU struct {
	key, value, freq int
	next             *NodeLFU
}

type DListLFU struct {
	head, tail *NodeLFU
}

// Remove removes and returns the first node in the list (for LFU eviction)
func (dl *DListLFU) remove() *NodeLFU {
	if dl.head == nil {
		return nil
	}
	node := dl.head
	dl.head = dl.head.next
	if dl.head == nil {
		dl.tail = nil
	}
	return node
}

// Add adds a node to the end of the list
func (dl *DListLFU) add(node *NodeLFU) {
	node.next = nil
	if dl.tail == nil {
		dl.head = node
		dl.tail = node
	} else {
		dl.tail.next = node
		dl.tail = node
	}
}

func newLFU(capacity int) *LFUCache {
	return &LFUCache{
		cache:     make(map[int]*NodeLFU),
		freqLists: make(map[int]*DListLFU),
		capacity:  capacity,
	}
}

// Simplified: Implement LFU, update
func (lfu *LFUCache) GetLFU(key int) int {
	if node, ok := lfu.cache[key]; ok {
		lfu.updateFreq(node)
		return node.value
	}
	return -1
}

// updateFreq increases the frequency of a node and moves it to the correct frequency list
func (lfu *LFUCache) updateFreq(node *NodeLFU) {
	oldFreq := node.freq
	if len(lfu.cache) >= lfu.capacity {
		// Remove the least frequently used node
		lfuList := lfu.freqLists[lfu.minFreq]
		evicted := lfuList.remove()
		if evicted != nil {
			delete(lfu.cache, evicted.key)
		}
	}
	newNode := &NodeLFU{key: node.key, value: node.value, freq: 1}
	if lfu.freqLists[1] == nil {
		lfu.freqLists[1] = &DListLFU{}
	}
	lfu.freqLists[1].add(newNode)
	lfu.cache[node.key] = newNode
	lfu.minFreq = 1
	// Update minFreq if needed
	if lfu.freqLists[oldFreq] == nil || /* list is empty */ false {
		if lfu.minFreq == oldFreq {
			lfu.minFreq = node.freq
		}
	}
}

func (lfu *LFUCache) PutLFU(key, value int) {
	if node, ok := lfu.cache[key]; ok {
		node.value = value
		lfu.updateFreq(node)
		return
	}
	if len(lfu.cache) >= lfu.capacity {
		// Remove the least frequently used node
		lfuList := lfu.freqLists[lfu.minFreq]
		// Assuming remove returns the node removed
		evicted := lfuList.remove()
		delete(lfu.cache, evicted.key)
	}
	newNode := &NodeLFU{key: key, value: value, freq: 1}
	lfu.cache[key] = newNode
	lfu.freqLists[1].add(newNode)
	lfu.minFreq = 1
}

// 84. Find all triplets sum to zero
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

// 85. Implement a concurrent map
type ConcurrentMap struct {
	m map[string]int
	sync.RWMutex
}

func newConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[string]int),
	}
}

func (cm *ConcurrentMap) Set(key string, value int) {
	cm.Lock()
	defer cm.Unlock()
	cm.m[key] = value
}

func (cm *ConcurrentMap) Get(key string) (int, bool) {
	cm.RLock()
	defer cm.RUnlock()
	v, ok := cm.m[key]
	return v, ok
}

// 86. Find maximum subarray sum
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]
	for _, num := range nums[1:] {
		currentSum = max(num, currentSum+num)
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

// 87. Implement a binary search tree
type BST struct {
	root *TreeNode
}

func (bst *BST) Insert(val int) {
	bst.root = bst.insertRec(bst.root, val)
}

func (bst *BST) insertRec(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val}
	}
	if val < node.Val {
		node.Left = bst.insertRec(node.Left, val)
	} else {
		node.Right = bst.insertRec(node.Right, val)
	}
	return node
}

// 88. Find k closest points to origin
func kClosest(points [][]int, k int) [][]int {
	heap := make([][]int, 0)
	for _, p := range points {
		heap = append(heap, p)
	}
	heapSort(heap, func(a, b []int) bool {
		return a[0]*a[0]+a[1]*a[1] < b[0]*b[0]+b[1]*b[1]
	})
	return heap[:k]
}

func heapSort(heap [][]int, less func(a, b []int) bool) {
	for i := len(heap)/2 - 1; i >= 0; i-- {
		heapify(heap, i, len(heap), less)
	}
	for i := len(heap) - 1; i > 0; i-- {
		heap[0], heap[i] = heap[i], heap[0]
		heapify(heap, 0, i, less)
	}
}

func heapify(heap [][]int, i, n int, less func(a, b []int) bool) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i
		if left < n && less(heap[left], heap[largest]) {
			largest = left
		}
		if right < n && less(heap[right], heap[largest]) {
			largest = right
		}
		if largest == i {
			break
		}
		heap[i], heap[largest] = heap[largest], heap[i]
		i = largest
	}
}

// 89. Implement a topological sort
func topologicalSort(graph map[string][]string) []string {
	visited := make(map[string]bool)
	stack := []string{}
	var dfs func(node *string)
	dfs = func(node *string) {
		visited[*node] = true
		for _, neighbor := range graph[*node] {
			if !visited[neighbor] {
				dfs(&neighbor)
			}
		}
		stack = append([]string{*node}, stack...)
	}
	for node := range graph {
		if !visited[node] {
			dfs(&node)
		}
	}
	return stack
}

// 90. Find minimum spanning tree (Prim's)
type Edge struct {
	weight int
	node   int
}

type EdgeMinHeap []Edge

func (h *EdgeMinHeap) Push(x Edge) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

func (h *EdgeMinHeap) Pop() Edge {
	if len(*h) == 0 {
		return Edge{}
	}
	min := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.down(0)
	return min
}

func (h *EdgeMinHeap) up(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if (*h)[i].weight >= (*h)[parent].weight {
			break
		}
		(*h)[i], (*h)[parent] = (*h)[parent], (*h)[i]
		i = parent
	}
}

func (h *EdgeMinHeap) down(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i
		if left < len(*h) && (*h)[left].weight < (*h)[smallest].weight {
			smallest = left
		}
		if right < len(*h) && (*h)[right].weight < (*h)[smallest].weight {
			smallest = right
		}
		if smallest == i {
			break
		}
		(*h)[i], (*h)[smallest] = (*h)[smallest], (*h)[i]
		i = smallest
	}
}

func primMST(graph [][]int) int {
	n := len(graph)
	visited := make([]bool, n)
	minHeap := &EdgeMinHeap{}
	minHeap.Push(Edge{weight: 0, node: 0}) // {weight, node}
	totalWeight := 0
	for len(*minHeap) > 0 {
		edge := minHeap.Pop()
		weight, node := edge.weight, edge.node
		if visited[node] {
			continue
		}
		visited[node] = true
		totalWeight += weight
		for i := 0; i < n; i++ {
			if !visited[i] && graph[node][i] != 0 {
				minHeap.Push(Edge{weight: graph[node][i], node: i})
			}
		}
	}
	return totalWeight
}

// 91. Implement Dijkstra’s algorithm
func dijkstra(graph [][]int, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0
	type pair struct{ dist, node int }
	minHeap := &MinHeap{}
	minHeap.Push(0)
	minHeap.Push(start)
	for len(*minHeap) > 0 {
		u := 0
		if len(*minHeap) >= 2 {
			u = (*minHeap)[1]
			*minHeap = (*minHeap)[2:]
		} else {
			break
		}
		if visited[u] {
			continue
		}
		visited[u] = true
		for v := 0; v < n; v++ {
			if graph[u][v] != 0 && dist[v] > dist[u]+graph[u][v] {
				dist[v] = dist[u] + graph[u][v]
				minHeap.Push(dist[v])
				minHeap.Push(v)
			}
		}
	}
	return dist
}

// 92. Find longest common substring
func longestCommonSubstring(s1, s2 string) string {
	rows := len(s1)
	cols := len(s2)
	dp := make([][]int, rows+1)
	for i := range dp {
		dp[i] = make([]int, cols+1)
	}
	maxLen := 0
	endIndex := 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
					endIndex = i
				}
			}
		}
	}
	return s1[endIndex-maxLen : endIndex]
}

// 93. Implement a union-find with rank
type UnionFind struct {
	parent []int
	rank   []int
}

func newUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := range uf.parent {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if x != uf.parent[x] {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	px := uf.Find(x)
	py := uf.Find(y)
	if px == py {
		return
	}
	if uf.rank[px] < uf.rank[py] {
		uf.parent[px] = py
	} else if uf.rank[px] > uf.rank[py] {
		uf.parent[py] = px
	} else {
		uf.parent[py] = px
		uf.rank[px]++
	}
}

// 94. Find word break problem
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// 95. Implement a consistent hashing
type ConsistentHash struct {
	ring  map[int]string
	nodes []int
}

func newConsistentHash() *ConsistentHash {
	return &ConsistentHash{
		ring: make(map[int]string),
	}
}

func (ch *ConsistentHash) Add(node string) {
	h := hash(node, 0) % 360
	ch.ring[h] = node
	ch.nodes = append(ch.nodes, h)
	sort.Ints(ch.nodes)
}

func (ch *ConsistentHash) Get(key string) string {
	hash := hash(key, 0) % 360
	for _, h := range ch.nodes {
		if h >= hash {
			return ch.ring[h]
		}
	}
	return ch.ring[ch.nodes[0]]
}

// 96. Find maximum rectangle in histogram
func largestRectangleArea(heights []int) int {
	maxArea := 0
	heights = append(heights, 0)
	stack := []int{}
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}
			maxArea = max(maxArea, height*width)
		}
		stack = append(stack, i)
	}
	return maxArea
}

// 97. Implement a skip list
type SkipListNode struct {
	val     *int
	forward []*SkipListNode
}

type SkipList struct {
	head  *SkipListNode
	level int
}

func newSkipList() *SkipList {
	return &SkipList{
		head:  &SkipListNode{forward: make([]*SkipListNode, 16)},
		level: 0,
	}
}

// Simplified insert
func (sl *SkipList) Insert(val *int) {
	update := make([]*SkipListNode, sl.level+1)
	current := sl.head
	for i := sl.level; i >= 0; i-- {
		for current.forward[i] != nil && *current.forward[i].val <= *val {
			current = current.forward[i]
		}
		update[i] = current
	}
	level := randomLevel()
	if level > sl.level {
		for i := sl.level + 1; i <= level; i++ {
			update = append(update, sl.head)
		}
		sl.level = level
	}
	newNode := &SkipListNode{val: val, forward: make([]*SkipListNode, level+1)}
	for i := 0; i <= level; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}
}

func randomLevel() int {
	level := 0
	for rand.Float64() < 0.5 && level < 15 {
		level++
	}
	return level
}

// 98. Find longest palindromic substring
func longestPalindromicSubstring(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, maxLen := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		maxCurrent := max(len1, len2)
		if maxCurrent > maxLen {
			maxLen = maxCurrent
			start = i - (maxLen-1)/2
		}
	}
	return s[start : start+maxLen]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// 99. Implement a concurrent queue
type ConcurrentQueue struct {
	items []int
	mutex sync.Mutex
}

func newConcurrentQueue() *ConcurrentQueue {
	return &ConcurrentQueue{
		items: make([]int, 0),
	}
}

func (cq *ConcurrentQueue) Enqueue(item int) {
	cq.mutex.Lock()
	defer cq.mutex.Unlock()
	cq.items = append(cq.items, item)
}

func (cq *ConcurrentQueue) Dequeue() (int, bool) {
	cq.mutex.Lock()
	defer cq.mutex.Unlock()
	if len(cq.items) == 0 {
		return 0, false
	}
	item := cq.items[0]
	cq.items = cq.items[1:]
	return item, true
}

// 100. Find maximum product subarray
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSoFar := nums[0]
	minSoFar := nums[0]
	maxProduct := nums[0]
	for i := 1; i < len(nums); i++ {
		tempMax := max(nums[i], max(maxSoFar*nums[i], minSoFar*nums[i]))
		minSoFar = min(nums[i], min(maxSoFar*nums[i], minSoFar*nums[i]))
		maxSoFar = tempMax
		maxProduct = max(maxProduct, maxSoFar)
	}
	return maxProduct
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BasicMain() {
	// Testing some of the function as required by program as upto user to test others.
	helloWorld()
	fmt.Println("Sum:", sum(3, 4))
	fmt.Println("Is even:", isEven(6))
	fmt.Println("Factorial:", factorial(5))
	fmt.Println("Reverse string:", reverseString("hello"))
	fmt.Println("Is palindrome:", isPalindrome("racecar"))
	fmt.Println("Max element:", maxElement([]int{3, 1, 4, 5, 2}))
	fmt.Println("Fibonacci:", fibonacci(6))
	fmt.Println("Is prime:", isPrime(17))
	fmt.Println("Count vowels:", countVowels("hello"))
}

// main runs the demo entry points of this package
func main() {
	BasicMain()
	DemoMain()
	HelloMain()
}
