package main

import "fmt"

type sigZodiacal int
type meses int

const (
	ACUARIO sigZodiacal = iota
	PISCIS
	ARIES
	TAURO
	GEMINIS
	CANCER
	LEO
	VIRGO
	LIBRA
	ESCORPIO
	SAGITARIO
	CAPRICORNIO
)

const (
	ENERO meses = 1 + iota
	FEBRERO
	MARZO
	ABRIL
	MAYO
	JUNIO
	JULIO
	AGOSTO
	SEPTIEMBRE
	OCTUBRE
	NOVIEMBRE
	DICIEMBRE
)

func (s sigZodiacal) String() string {
	return [12]string{"acuario", "piscis", "aries", "tauro", "geminis", "cancer", "leo", "virgo", "libra", "escorpio", "sagitario", "capricornio"}[s]
}

func main() {
	var mes meses
	var dia int
	var signo sigZodiacal
	fmt.Scan(&dia)
	fmt.Scan(&mes)
	switch mes {
	case ENERO:
		if dia < 20 {
			signo = CAPRICORNIO
		} else {
			signo = ACUARIO
		}
	case FEBRERO:
		if dia < 19 {
			signo = ACUARIO
		} else {
			signo = PISCIS
		}
	case MARZO:
		if dia < 21 {
			signo = PISCIS
		} else {
			signo = ARIES
		}
	case ABRIL:
		if dia < 20 {
			signo = ARIES
		} else {
			signo = TAURO
		}
	case MAYO:
		if dia < 21 {
			signo = TAURO
		} else {
			signo = GEMINIS
		}
	case JUNIO:
		if dia < 21 {
			signo = GEMINIS
		} else {
			signo = CANCER
		}
	case JULIO:
		if dia < 23 {
			signo = CANCER
		} else {
			signo = LEO
		}
	case AGOSTO:
		if dia < 23 {
			signo = LEO
		} else {
			signo = VIRGO
		}
	case SEPTIEMBRE:
		if dia < 23 {
			signo = VIRGO
		} else {
			signo = LIBRA
		}
	case OCTUBRE:
		if dia < 23 {
			signo = LIBRA
		} else {
			signo = ESCORPIO
		}
	case NOVIEMBRE:
		if dia < 22 {
			signo = ESCORPIO
		} else {
			signo = SAGITARIO
		}
	case DICIEMBRE:
		if dia < 22 {
			signo = SAGITARIO
		} else {
			signo = CAPRICORNIO
		}
	}
	fmt.Print(signo)
}
