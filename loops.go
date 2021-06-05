package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}
}
 {i := 0
for i < 5 {
 fmt.Print(i)
 i++
 }
}

arr :=[3]string{"a", "b", "c"}

for index, value := range arr {
	fmt.Print("index:", index, "value:", value)
}
m := make(map[string]string)
m["a"] = "alpha"
m["b"] = "beta"

for key, value := range m {
	fmt.Print("key:", key, "value", value)
}
