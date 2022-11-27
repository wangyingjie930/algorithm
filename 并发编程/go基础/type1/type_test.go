package type1

import (
	"fmt"
	"testing"
)

func TestDog_Type(t *testing.T) {
	// 示例1。
	dog := Dog{"little pig"}
	//接口类型断言
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	ret, ok := interface{}(&dog).(Pet)
	fmt.Printf("类型断言, T为接口类型: %+v\n, type: %T\n", ret, ret)
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)
	fmt.Println()

	//&dog这里是动态值, *Dog这里是动态类型
	//Pet这里是静态类型
	var pet Pet = &dog
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
	fmt.Printf("%+v %T", pet, pet)

	//dog本身的实现的方法多, 覆盖了smallPet, Pet接口定义好的那些函数
	//变量smallPet的动态类型仍为*Dog, 但是smallPet能调用的就只有smallPet接口定义好的那些
	var smallPet smallPet = &dog
	//此时将smallPet类型断言, 转化为Pet接口, 只能使用Pet接口定义好的那些函数(这些函数也必须是*Dog实现过的)
	//动态类型仍为*Dog
	smallPetToPet, ok := interface{}(smallPet).(Pet)
	fmt.Printf("%T %T", smallPet, smallPetToPet)
}