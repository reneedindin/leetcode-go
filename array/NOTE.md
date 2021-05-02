# Array Note

## Title

### 26. Remove Duplicates from Sorted Array

#### version1

TimeComplexity: O(n)

SpaceComplexity:

- 透過 golang slice 實作移除重複元素
- 使用 append (golang append 實作機制會影響 performance -> 印象中)
- nums 最後會移除重複的元素

#### version2

TimeComplexity: O(n)

SpaceComplexity:

- 只透過 for
- nums 並不會移除重複元素，而是計算不同的元素有幾個

### 35. Search Insert Position

#### version1

TimeComplexity: O(n)

SpaceComplexity:

#### version2

TimeComplexity: O(n)

SpaceComplexity:

- 減少 `idx == len(nums)-1 && nums[idx] < target` 的判斷