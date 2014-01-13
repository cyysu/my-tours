package main

import (
	"strings"
)

type Rule struct {
        rule_type int
        value string
}

const (
        TYPE_STRING_EQUAL = 1
)

func is_valid(output string, rule Rule) (bool, string){
	output = strings.TrimSpace(output)
        switch rule.rule_type {
        case TYPE_STRING_EQUAL:
		if output  ==  rule.value {
			return true, "赞！"
		} else {
			return false, "输出的东西不一样啊"
		}
	}

	return false, "Not implemented yet"	
}

func validate_python(num int, output string) (bool, string) {
        switch num {
        case 1:
                rule := Rule{TYPE_STRING_EQUAL, "hello python world"}
                return is_valid(output, rule)
        }

	return false, "Not implemented yet"
}
