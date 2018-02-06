package main
import "fmt"
func main() {

	sayHello()

	fmt.Print("Enter Temperature in Farenehit: ");
	var input float64
	fmt.Scan("%f", &input)

	var celcius float64
	celcius =  (input-32)*5/9
	fmt.Println("Temperature in Celcius: %f", celcius)

}