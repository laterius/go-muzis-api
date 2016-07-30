package stringutil

import "fmt"

//замыкания
func mainEvenGenerator() func() uint{
	i := uint(0)
	return func()(ret uint){
		ret = i
		i+=2
		return
	}

}

func FuncLeaningPrint(){
	/*nextgen := mainEvenGenerator()
	fmt.Println(nextgen())
	fmt.Println(nextgen())
	fmt.Println(nextgen())*/

	/*defer  func(){
		str := recover()
		fmt.Println(str)
	}()

	panic("PANIC")*/

	fmt.Println(Task1(1,2,3))
}

func Task1(args ...int) int {
	var a int
	for _, i := range args{
		a = a+ i
	}
	return a
}

