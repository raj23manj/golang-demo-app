package app

func MapTo(data []interface{}, f func(interface{}) interface{}) []interface{} {

	mapped := make([]interface{}, len(data))

	for i, e := range data {
		mapped[i] = f(e)
	}

	return mapped
}
