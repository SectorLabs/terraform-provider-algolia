package main

// Helper function for casting []interface{} to []string
func castStringList(inputItems []interface{}) []string {
	outputItems := make([]string, 0, len(inputItems))
	for _, item := range inputItems {
		itemValue, ok := item.(string)
		if ok && itemValue != "" {
			outputItems = append(outputItems, itemValue)
		}
	}

	return outputItems
}
