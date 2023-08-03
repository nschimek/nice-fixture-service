package core

func OnlyPopulatedChildren[P any, C any](parents []P, getChildrenFunc func(parent P) []C) []P {
	o := []P{}
	for _, p := range parents {
		if len(getChildrenFunc(p)) > 0 {
			o = append(o, p)
		}
	}
	return o
}