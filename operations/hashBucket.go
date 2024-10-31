package operations

import "hash/fnv"

func HashBucket(key string, r int) int {
	algorithm := fnv.New32a()
	algorithm.Write([]byte(key))
	return int(algorithm.Sum32()) % r
}
