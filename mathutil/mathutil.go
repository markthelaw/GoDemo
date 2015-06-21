// calculate
package mathutil

import (
	"fmt"
	"math/rand"
	"math"
	"time"
	"runtime"
	//"sync"

)

//func main(){
//	fmt.Println(calculatePi(1000000))
//}
//we make this run in parallel
func CalculatePi(times int) float64{
	numCPU := MaxParallelism()
	//fmt.Println(numCPU)
	//set max cpu to use
	runtime.GOMAXPROCS(numCPU)
	//channel to store the return result from go routine
	inside := make(chan float64, numCPU)
		
	eachCoreTrials := times/numCPU
	//need to add this due to division round down above.
	lastCoreTrials := eachCoreTrials + (eachCoreTrials*numCPU-times)
	//fmt.Println(lastCoreTrials)
	sum:=float64(0)
	fmt.Println(numCPU)
	for i:=0;i<numCPU-1;i++ {
		go taskPi(eachCoreTrials, inside)
		fmt.Println("taskPi %d is running", i)
	}
	
	//for last core run
	go taskPi(lastCoreTrials, inside)
	
	fmt.Println("finished 1st for loop")
	for i:=0; i<numCPU; i++ {
		j:= <- inside
		fmt.Println("%d", j)
		sum += j
	}
	fmt.Println(sum)

	return sum/float64(times)*4
}

func taskPi(times int, inside chan float64){
	
	//defer wg.Done()
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sum:=float64(0)
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
	inside <- sum
	//fmt.Println(sum)
	//fmt.Println(float32(times))
	//return sum, float64(times)
	fmt.Println("pi finished")
}

func inSideCircle(x float64, y float64)bool{
	xs:=math.Pow(x,2)
	ys:=math.Pow(y,2)
	if math.Sqrt(float64(xs+ys)) > float64(1){
		return false
	}
	return true
}
//return max core system has
func MaxParallelism() int {
    maxProcs := runtime.GOMAXPROCS(0)
    numCPU := runtime.NumCPU()
    if maxProcs > numCPU {
        return maxProcs
    }
    return numCPU
}	