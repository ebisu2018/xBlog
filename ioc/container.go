package ioc

func Container() *Ioc {
	return container
}

var container = &Ioc{
	container: make(map[string]IocInf),
}


func ApiHandler() *Ioc {
	return apiContainer
}

var apiContainer = &Ioc{
	container: make(map[string]IocInf),
}