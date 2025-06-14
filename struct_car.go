package main

import "fmt"

type Car struct {
	price      int
	color      string
	company    string
	isElectric bool
}

func changeColor(car *Car, color string) {
	car.color = color
}

func main() {

	defer fmt.Print("Chalo kaam khatam ab !")

	tesla_car := Car{45000, "red", "Tesla", true}
	bmw_car := Car{price: 50000, color: "black", company: "BMW"}

	fmt.Println(tesla_car.price, tesla_car.color)
	fmt.Println(bmw_car.price, bmw_car.color, bmw_car.isElectric)

	// Since we cannot create a fucntion inside a struct, we created a separate function to changeColor, now we can use this by passing Pointer to the Car
	changeColor(&tesla_car, "Grey")
	fmt.Printf("Color updated %v, %v\n", tesla_car.price, tesla_car.color)

}
