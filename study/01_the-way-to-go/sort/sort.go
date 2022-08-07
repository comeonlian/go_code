package sort

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

type IntArray []int // 别名类型

func (p IntArray) Len() int {
	return len(p)
}

func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type StringArray []string

func (p StringArray) Len() int {
	return len(p)
}

func (p StringArray) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p StringArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func SortInts(a []int) {
	ia := IntArray(a) // 数组类型转换为别名类型
	Sort(ia)
}

func SortStrings(a []string) {
	sa := StringArray(a)
	Sort(sa)
}

func IntsAreSorted(a []int) bool {
	ia := IntArray(a)
	return IsSorted(ia)
}

func StringsAreSorted(a []string) bool {
	sa := StringArray(a)
	return IsSorted(sa)
}