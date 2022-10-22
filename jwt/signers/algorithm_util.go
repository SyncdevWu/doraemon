package signers

import (
	"github.com/syncdevwu/doraemon/core/hashmap"
	"strings"
)

var algs = hashmap.NewHashMap[string, string](&hashmap.Options[string, string]{})

func init() {
	algs.Put("HS256", "")
}

// GetAlgorithm 获取算法，用户传入算法ID返回算法名，传入算法名返回本身
func GetAlgorithm(idOrAlgorithm string) string {
	if res := getAlgorithmById(idOrAlgorithm); res != "" {
		return res
	}
	return idOrAlgorithm
}

// GetId 获取算法ID，用户传入算法名返回算法ID，传入算法ID返回本身
func GetId(idOrAlgorithm string) string {
	if res := getIdByAlgorithm(idOrAlgorithm); res != "" {
		return res
	}
	return idOrAlgorithm
}

func getAlgorithmById(id string) string {
	res, _ := algs.Get(strings.ToUpper(id))
	return res
}

func getIdByAlgorithm(algorithm string) string {
	keys, vals := algs.Entries()
	if len(keys) != len(vals) {
		return ""
	}
	for i := 0; i < len(vals); i++ {
		if vals[i] == algorithm {
			return keys[i]
		}
	}
	return ""
}
