package lessoncodebasic

/*Реализуйте методы структуры Counter, представляющую собой счётчик,
хранящий неотрицательное целочисленное значение и позволяющий это значение изменять:

метод Inc(delta int) должен увеличивать текущее значение на delta единиц (на 1 по умолчанию),
метод Dec(delta int) должен уменьшать текущее значение на delta единиц.*/

type Counter struct {
	Value int
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (c *Counter) Inc(delta int) {
	if delta <1 {
		delta = 1
	}
	c.Value += delta
}

func (c *Counter) Dec(delta int) {
		if delta < 1 {
		delta = 1
	}
	c.Value -= delta
	if c.Value < 0 {
		c.Value = 0
	}
}


