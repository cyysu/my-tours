package main

import (
        "strings"
        "strconv"
        "log"
)

type Rule struct {
        rule_type int
        value string
}

const (
        TYPE_STRING_EQUAL = 1
        TYPE_NUM_EQUAL = 2
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

func valid_int(output string, num int) (bool, string){
        istr, err := strconv.Atoi(output)
        if err != nil {
                return false, "需要一个数字类型"
        }

        if istr == num {
                return true, "cool"
        } else if istr < num {
                return false, "数字太小"
        } else {
                return false, "数字太大"
        }
}

func custom_valid_str(output string, target string, right_resp string, wrong_resp string) (bool, string){
        if output == target {
                return true, right_resp
        } else  {
                return false, wrong_resp
        }
}

func dumb_validation() (bool, string){
        return true, ""
}


func validate_python(num int, output string) (bool, string) {
        output = strings.TrimSpace(output)
        switch num {
        case 1:
                var rule Rule
                rule = Rule{TYPE_STRING_EQUAL, "hello python world"}
                return is_valid(output, rule)
        case 2:
                return valid_int(output, 6)
        case 3, 5, 6, 7:
                return dumb_validation()
        case 4:
                return custom_valid_str(output, "python is awesome", "没错", "python才不糟糕呢")
        }
        log.Fatal("Not implemented yet")
        return false, ""
}

func validate_ruby_variable(output string) (bool, string){
        lines := strings.Fields(output)
        if lines[0] != "25" {
                return false, "my_num的值错误"
        }

        if lines[1] != "true" {
                return false, "my_boolean的值错误"
        }
        if lines[2] != "Ruby" {
                return false, "my_string的值错误"
        }

        return true, "正确"
}

func validate_ruby_case(output string) (bool, string){
        lines := strings.Fields(output)
        if lines[0] != "UPCASE" {
		return false, "upcase没有大写"
	} else if lines[1] != "downcase" {
		return false, "downcase没有大写"
	}

	return true, "厉害"
}

func validate_ruby(num int, output string) (bool, string) {
        output = strings.TrimSpace(output)
        switch num {
        case 2:
                return validate_ruby_variable(output)
        case 1, 3, 4, 5, 6:
                return dumb_validation()
        case 7:
                return validate_ruby_case(output)
        }
        log.Fatal("Not implemented yet")
        return false, "Not implemented yet"
}
