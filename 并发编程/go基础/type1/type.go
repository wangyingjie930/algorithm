package type1

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}
//*Dog实现了3 个方法, Dog只实现了2个
type Dog struct {
	name string // 名字。
}
func (dog *Dog) SetName(name string) {
	dog.name = name
}
func (dog Dog) Name() string {
	return dog.name
}
func (dog Dog) Category() string {
	return "dog"
}

type smallPet interface {
	Name() string
	Category() string
}
