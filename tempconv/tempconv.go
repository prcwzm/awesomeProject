package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvins float64

const (
	AbsoluteZero Celsius = -273.15
	FreezingC    Celsius = 0
	BoilingC     Celsius = 100
)

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CTOF(BoilingC)
	fmt.Printf("%g\n", boilingF-CTOF(FreezingC))
	c := FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
	k := CTOK(c)
	fmt.Printf("%g\n", k)
	fmt.Printf("%g\n", KTOC(k))
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g °C", c)
}
