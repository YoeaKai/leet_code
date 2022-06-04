package expression_add_operators

import (
	"fmt"
	"strconv"
)

type ResultsBucket struct {
	Target  int
	Results []string
}

func AddOperators(num string, target int) (ret []string) {
	res := &ResultsBucket{Target: target}
	addOperatorsRecursive(num, "", 0, 0, res)
	return res.Results
}

func addOperatorsRecursive(num string, exp string, prev int, current int, res *ResultsBucket) {
	if len(num) == 0 {
		if res.Target == current {
			res.Results = append(res.Results, exp)
		}
		return
	}

	for i := 1; i < len(num)+1; i++ {
		if i > 1 && num[0] == '0' {
			break
		}
		newArg, _ := strconv.Atoi(num[:i])
		if len(exp) == 0 {
			addOperatorsRecursive(num[i:], num[:i], newArg, newArg, res)
		} else {
			addOperatorsRecursive(num[i:], fmt.Sprintf("%v+%v", exp, num[:i]), newArg, current+newArg, res)
			addOperatorsRecursive(num[i:], fmt.Sprintf("%v-%v", exp, num[:i]), -newArg, current-newArg, res)
			addOperatorsRecursive(num[i:], fmt.Sprintf("%v*%v", exp, num[:i]), newArg*prev, current-prev+prev*newArg, res)
		}
	}
}
