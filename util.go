package main

// Helper function for casting []interface{} to []string
func castStringList(inputItems []interface{}) []string {
	outputItems := make([]string, 0, len(value))
	for _, item := range inputItems {
		itemValue, ok := v.(string)
		if ok && itemValue != "" {
			outputItems = append(outputItems, itemValue)
		}
	}

	return outputItems
}
