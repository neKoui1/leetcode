//假设 Andy 和 Doris 想在晚餐时选择一家餐厅，并且他们都有一个表示最喜爱餐厅的列表，每个餐厅的名字用字符串表示。
//你需要帮助他们用最少的索引和找出他们共同喜爱的餐厅。 如果答案不止一个，则输出所有答案并且不考虑顺序。 你可以假设答案总是存在。

package main

func findRestaurant(list1 []string, list2 []string) []string {
	var result []string
	min := 65536
	for i1, v1 := range list1 {
		for i2, v2 := range list2 {
			if v1 == v2 {
				if i1+i2 < min {
					result = nil
					result = append(result, v1)
					min = i1 + i2
					continue
				} else if i1+i2 == min {
					result = append(result, v1)
					continue
				}
			}
		}
	}

	return result
}
