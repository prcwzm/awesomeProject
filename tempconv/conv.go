package tempconv

func CTOF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func CTOK(c Celsius) Kelvins {
	return Kelvins(c + 273.15)
}
func KTOC(k Kelvins) Celsius {
	return Celsius(k - 273.15)
}
