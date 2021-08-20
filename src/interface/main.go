package main
import "fmt"

////////////////////////////
type Usb interface{
	Start()
	Stop()
}
////////////////////////////
type Phone struct{}

func(p Phone) Start(){
	fmt.Println("phone is Start()")
}

func(p Phone) Stop(){
	fmt.Println("phone is stop()")
}
////////////////////////////
type Camera struct{}

func(p Camera) Start(){
	fmt.Println("Camera is Start()")
}

func(p Camera) Stop(){
	fmt.Println("Camera is stop()")
}
////////////////////////////
type PC struct{}

func(c PC) InsertPC(usb Usb){
	usb.Start()
	usb.Stop()
}

////////////////////////////

func main()  {
	var pc PC
	var phone Phone
	var camera Camera

	pc.InsertPC(phone)
	pc.InsertPC(camera)

}