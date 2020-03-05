package setup

import (
	"github.com/astaxie/beego"
)

func functionsForView() {
	_ = beego.AddFuncMap("calculate", calculate)
	_ = beego.AddFuncMap("dynamicMap", dynamicMap)
}

// calculate used for calculation logic in views
func calculate(x, y int, option string) int {

	switch option {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	default:
		return 0
	}

}

// dynamicMap functionused to append the args to map in key value format
func dynamicMap(args ...interface{}) (out map[interface{}]interface{}) {
	maps := make(map[interface{}]interface{})

	for i := 0; i < len(args); i++ {
		maps[args[i]] = args[i+1]
		i = i + 1
	}

	return maps
}
