package main

import "fmt"

func main() {
	mp := make(map[string]string)
	mp["etc"] = "Soumya"
	mp["etc2"] = "Soumyasw"

	delete(mp, "etc")
	fmt.Println(mp, len(mp))
	m := map[string]string{"deds": "wdff"}
	fmt.Print(m)

}
