package main

import "fmt"

type BitInt struct {
	sign int
	num  string
}

func (a *BitInt) Add(b *BitInt) *BitInt {

}

func (a *BitInt) Subtract(b *BitInt) *BitInt {

}

func (a *BitInt) Multiply(b *BitInt) *BitInt {

}

func (a *BitInt) Divide(b *BitInt) *BitInt {

}

func main() {
	var a, b BitInt
	var oper byte
	fmt.Scanf("%s %c %s", &a.num, &oper, &b.num)
	switch oper {
	case '+':
		fmt.Print(a.Add(&b).num)
	case '-':
		fmt.Print(a.Subtract(&b).num)
	case '*':
		fmt.Print(a.Multiply(&b).num)
	case '/':
		fmt.Print(a.Divide(&b).num)
	default:
	}
}
