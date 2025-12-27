package homework01

import (
	"sort"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	var resultMap map[int]int = make(map[int]int)

	for i:=0; i<len(nums); i++ {
		if _, key := resultMap[nums[i]]; key {
			resultMap[nums[i]] += 1
		} else {
			resultMap[nums[i]] = 1
		}
	}

	for key, value := range resultMap {
		if value == 1 {
			return key
		}
	}

	return 0;
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
    	if x<0 || (x%10==0 && x!=0) {
		return false;
	}

	revertedNumber := 0

	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x/10
		x = x/10
	}

	return x==revertedNumber || x==(revertedNumber/10)
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	pair := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []byte{}

	for i:=0; i<n; i++ {
		if pair[s[i]] > 0 {
			if len(stack)==0 || stack[len(stack)-1]!=pair[s[i]] {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}		
	}

	return len(stack) == 0
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minStr := strs[0];
	for _, str := range strs {
		minStr = min(minStr, str)
	}

	for i:=0; i<len(strs); i++ {
		current := minStr[i];
		for _, buf := range strs {
			if(current != buf[i]) {
				return minStr[:i]
			}
		}
	}

	return minStr
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
    	n:= len(digits)

	for i:=n-1; i>=0; i-- {
		if digits[i]!=9 {
			digits[i]++;
			for j:=i+1; j<n; j++ {
				digits[j] = 0;
			}

			return digits;
		}
	}

	digits = make([]int, n+1)
	digits[0] = 1

	return digits
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0;
	}

	slow := 1
	for fast:=1; fast<n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	var result [][]int

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result = append(result, []int{intervals[0][0], intervals[0][1]})

	for i, interval := range intervals {
		lastLen := len(result) - 1
		lastItem := result[lastLen]

		if interval[0] <= lastItem[1] {
			if interval[1] > lastItem[1] {
				lastItem[1] = intervals[i][1]
			}
		} else {
			result = append(result, []int{interval[0], interval[1]})
		}
	}

	return result
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i+1; j<len(nums); j++ {
			if x + nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil;
}
