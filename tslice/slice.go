package tslice

// SliceRemoveDuplicates 切片去重
// Args: slice []string 切片
// Returns: []string 去重后的切片
func SliceRemoveDuplicates(slice []string) []string {
	uniqueMap := make(map[string]bool)
	var result []string
	for _, item := range slice {
		if _, exists := uniqueMap[item]; !exists {
			uniqueMap[item] = true
			result = append(result, item)
		}
	}
	return result
}
