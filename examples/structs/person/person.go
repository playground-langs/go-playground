package person

type Person struct {
	Name       string
	Age        uint8
	unexported int
}

type A struct {
	B1 *B
}

type B struct {
	A1 *A
}

type Employee struct {
	Name string
	Age  int
	//匿名嵌入struct,可以有多个
	Address
}

type Address struct {
	Street string
	City   string
}
