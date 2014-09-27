package main
import "sdlaudio"
import "time"
import "math"


func main(){
	sdlaudio.InitSDL()//only needed when not using gl or gl31
	sdlaudio.Init()
	sdlaudio.OutFunc(func(t int64)(int16, int16){
		return int16(math.Sin(float64(t)/50)*32000*math.Sin(float64(t)/50000))/10, int16(math.Sin(float64(t)/50)*32000*math.Cos(float64(t)/50000))/10
	})
	sdlaudio.Enable()
	time.Sleep(time.Second*50)
}
