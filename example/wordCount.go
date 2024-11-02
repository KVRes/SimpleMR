package main

import (
	"fmt"
	"github.com/KVRes/SimpleMR/node"
	"github.com/KVRes/SimpleMR/types"
	"strings"
)

func main() {
	master := node.NewMaster().
		WithMapWorker(2).
		WithReduceWorker(3)

	rst := master.Start(MapDocument, 2, 3, MapDocument, ReduceWordCount)

	for i, v := range rst {
		fmt.Printf("Reduce File #%d\n", i)
		vT := v.(map[string]int)
		for w, count := range vT {
			fmt.Printf("%s: %d\n", w, count)
		}
	}

}

func MapDocument(ctx *node.MapContext, document any) {
	doc := document.(string)

	ws := strings.Split(doc, " ")
	for _, w := range ws {
		ctx.Emit(strings.Trim(w, " "), 1)
	}
}

func ReduceWordCount(task types.ReduceTask) any {
	m := make(map[string]int)
	for k, list := range task {
		for _, elem := range list {
			m[k] = m[k] + elem.(int)
		}
	}
	return m
}
