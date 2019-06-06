package main

import (
	"fmt"

	"github.com/gagliardetto/treeout"
)

func main() {
	t := treeout.New("DEPLOYMENTS")
	t.Child("dep1")
	t.Child("dep2")

	t2 := t.Child("dep3")
	t2.Child("pod1").Child("pod1").Child("pod1").Child("pod1").Child("pod1").Child("pod1").Child("pod1").Child("pod1").Child("pod1")
	t2.Child("pod2").Child("something")
	t2.Child("pod3").Child("something")
	t2.Child("pod4")

	t3 := t2.Child("pod5")
	t3.Child("err:", "something happened")
	t.Child("dep4").Child("child").Child("child").Child("child")

	t.Child("dep5").Child("child")

	fmt.Println(t)
}
