package main

import (
	"fmt"
	"github.com/KVRes/SimpleMR/node"
	"github.com/KVRes/SimpleMR/splitter"
	"github.com/KVRes/SimpleMR/types"
	"strings"
)

func main() {
	master := node.NewMaster().
		WithMapWorker(2).
		WithReduceWorker(3).
		WithSplitter(splitter.TextSplit)

	Raw := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
	data := strings.Replace(Raw, ". ", ".\n", -1)
	rst := master.Start(data, 2, 3, MapDocument, ReduceWordCount)

	for i, v := range rst {
		fmt.Printf("---------------\n")
		fmt.Printf("Reduced File #%d\n", i)
		fmt.Printf("---------------\n")
		vT := v.(map[string]int)
		for w, count := range vT {
			fmt.Printf("%s: %d\n", w, count)
		}
	}

}

func cleanWord(w string) string {
	return strings.Trim(w, ",.!? ")
}

func MapDocument(ctx *node.MapContext, document any) {
	doc := document.(string)

	ws := strings.Split(doc, " ")
	for _, w := range ws {
		w = strings.ToLower(w)
		w = cleanWord(w)

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
