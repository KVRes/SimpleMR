package node

import (
	"github.com/KVRes/SimpleMR/types"
	"hash/fnv"
	rnd "math/rand/v2"
)

func HashPartitioner(key string) int {
	algorithm := fnv.New32a()
	algorithm.Write([]byte(key))
	return int(algorithm.Sum32())
}

func RandomPartitioner(key string) int {
	return rnd.Int()
}

func KeyMapPartitionerGenerator(keyMapper func(string) string, partitioner types.Partitioner) types.Partitioner {
	return func(key string) int {
		key = keyMapper(key)
		return partitioner(key)
	}
}
