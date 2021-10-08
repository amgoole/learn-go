package iexten

import str "oop/struction"

type Man str.Person

func (m *Man) GetAge() int {
	age := m.Age
	return age
}

// Man实例实现了调用Person的方法

