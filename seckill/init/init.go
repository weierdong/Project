package initt

import "seckill/api"

func JoinInit(name string, id string, email string, tel string, pass1 string, pass2 string) api.Join { //初始化join
	var j api.Join
	if pass1 == pass2 {

		j.Name = name
		j.Id = id
		j.Tel = tel
		j.Pass1 = pass1
		return j
	} else {
		return j
	}
}
