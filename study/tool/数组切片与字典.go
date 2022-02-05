package tool

import "fmt"

func ArrayandSlice() {
	//用索引对数组初始化
	array_2 := []string{0: "sdf", 5: "trt"}
	fmt.Printf("len(array_2): %v\n", len(array_2))
	fmt.Printf("array_2: %v\n", array_2)

	//用数组对切片进行初始化
	slice_1 := array_2[:] //前后都没有数表示取所有元素，前闭后开
	fmt.Printf("slice_1: %T\n", slice_1)

	//s
	slice_1 = append(slice_1, "100")
	fmt.Printf("slice_1: %v\n", slice_1)
	//切片删除索引为i的元素
	i := 6
	slice_1 = append(slice_1[:i])
	fmt.Printf("slice_1: %v\n", slice_1)

}

func Mapdic() {
	var map_1 map[string]string //前键后值
	map_1 = make(map[string]string)
	map_1["name"] = "guang"
	map_1["age"] = "22"
	fmt.Printf("map_1: %v\n", map_1)

	map_2 := map[string]string{"name": "guang", "age": "22"}
	fmt.Printf("map_2: %v\n", map_2)
}
