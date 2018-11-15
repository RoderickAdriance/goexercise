package exercise

import "fmt"

func MapDemo() {
	countryCapitalMap := make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["France"]
	if ok {
		fmt.Println(capital)
	} else {
		fmt.Println("not exists!")
	}

}
