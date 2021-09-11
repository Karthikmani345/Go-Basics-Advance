package main

// Any type which provides definition for all the methods of an interface
// is said to implicitly implement that interface.

// Base type
// interface can contain only function signature
// interface cant hold properties
type human interface {
	getgender() string
	fullname() string
}

// -----------------------------------
// male class and its function
type male struct {
	firstName string
	gender    string
}

func (p male) fullname() string {
	return p.firstName
}

func (p male) getgender() string {
	return p.gender
}

// -----------------------------------
// male class and its function
type female struct {
	firstName string
	gender    string
}

func (p female) fullname() string {
	return p.firstName
}

func (p female) getgender() string {
	return p.gender
}

func main() {
	maleEmp := male{"karthik", "male"}
	femaleEmp := female{"aishu", "female"}

	// fmt.Println(maleEmp)
	// fmt.Println(maleEmp.fullname())

	// fmt.Println("--------------------")

	// fmt.Println(femaleEmp.fullname())
	// fmt.Println(femaleEmp.getgender())

	getHumangender(maleEmp)
	getHumangender(femaleEmp)

}

func getHumangender(arg human) {
	println(arg.fullname())
	println(arg.getgender())
}
