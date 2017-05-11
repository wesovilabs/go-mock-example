package matches

import (
	"github.com/wesovilabs.com/go-mock-example/utils"
)

type Result struct {
}

func GetResultForMatch(local string, visitor string) string {
	result := utils.Result()
	switch result {
	case 0:
		return local + " empata con " + visitor
	case 1:
		return local + " vence a " + visitor
	case 2:
		return local + " pierde contra " + visitor
	}
	return "Unexpected reuslt."

}
