package setup

import "github.com/astaxie/beego"

func functionsForView() {
	beego.AddFuncMap("incrementByOne", incrementByOne)
}

// incrementByOne function will increment by one to the given integer
func incrementByOne(in int) (out int) {
	out = in + 1
	return
}
