// calculate
package mathutil

import (
	//"fmt"
	"math/rand"
	"math"
	"time"

)

//func main(){
//	fmt.Println(calculatePi(1000000))
//}

func CalculatePi(times int) float32{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sum:=float32(0)
	if times<0 {
		times = -times
	}
	for i:=0; i<times; i++{
		x:=r.Float64()
		y:=r.Float64()
		//sum+=r.Float32()
		//fmt.Println(math.Pow(x,2))
		if inSideCircle(x, y){
			sum++
		}
	}
	//fmt.Println(sum)
	//fmt.Println(float32(times))
	return sum/float32(times)*4
}

func inSideCircle(x float64, y float64)bool{
	xs:=math.Pow(x,2)
	ys:=math.Pow(y,2)
	if math.Sqrt(float64(xs+ys)) > float64(1){
		return false
	}
	return true
}	