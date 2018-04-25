package distribution

import "log"

type Name struct {
    NamingRecords map[string]CalculatorProxy
    ServiceName string
    CalculatorProxy CalculatorProxy
    Search string
}


func (nr Name) Bind() {

    log.Println("bind name service")
    nr.NamingRecords[nr.ServiceName] = nr.CalculatorProxy
    log.Println("new map:", nr.NamingRecords)

}

func (nr Name) Lookup() CalculatorProxy {

    return nr.NamingRecords[nr.Search]

}


type Calculator struct {
    x, y float64
}


func (calc Calculator) Sum() float64{
    return calc.x + calc.y
}

func (calc Calculator) Sub() float64{
    return calc.x - calc.y
}

func (calc Calculator) Div() float64{
    return calc.x / calc.y
}

func (calc Calculator) Mul() float64{
    return calc.x * calc.y
}
