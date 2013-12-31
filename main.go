package main
import "gl"
import "time"

func main(){
	gl.InitGL(800,600, 1)
	gl.ClearColor(1,0.5,0,1)
	gl.Clear(gl.COLOR_BUFFER_BIT|gl.DEPTH_BUFFER_BIT)
	gl.Flip()
	time.Sleep(time.Second)

}
