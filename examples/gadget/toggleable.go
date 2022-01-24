package gadget

type Toggleable interface {
	Toggle()
}

type Switch string

// Toggle 指针类型接收者 需要将指针传递给接口
func (s *Switch) Toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
}

type Switch1 string

func (s Switch1) Toggle() {
	if s == "on" {
		s = "off"
	} else {
		s = "on"
	}
}
