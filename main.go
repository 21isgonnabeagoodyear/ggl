package main
import "gl"
import "time"
import "fmt"

var vshadsource = []byte(`
#version 330
void main(){
	if(gl_VertexID == 0)
		gl_Position = vec4(-1,-1,0,1);
	if(gl_VertexID == 1)
		gl_Position = vec4(-1,1,0,1);
	if(gl_VertexID == 2)
		gl_Position = vec4(1,1,0,1);
	if(gl_VertexID == 3)
		gl_Position = vec4(1,-1,0,1);

}

`)
var fshadsource = []byte(`
#version 330
out vec4 col;
void main(){
	//col = vec4(int(gl_FragCoord.x/4)&int(gl_FragCoord.y/4));
	col = vec4(0,0,1,1);
	col.a = float(int(gl_FragCoord.x/4)&int(gl_FragCoord.y/4));
}
`)

func printshaderinfolog(shaderid uint32){
	var status int32
	gl.GetShaderiv(shaderid, gl.COMPILE_STATUS, &status)
	if status != gl.TRUE{
		var log [1000]byte
		gl.GetShaderInfoLog(shaderid, 1000, nil, &log[0])
		fmt.Println(string(log[:]))
	}
}

func main(){
	gl.InitGL(800,600, 1)
	gl.ClearColor(1,0.5,0,1)
	gl.Clear(gl.COLOR_BUFFER_BIT|gl.DEPTH_BUFFER_BIT)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	//gl.Vert

	var ptrtobyte *byte
	var length int32

	prog := gl.CreateProgram()
	vertshad := gl.CreateShader(gl.VERTEX_SHADER)
	ptrtobyte = &vshadsource[0]
	length = int32(len(vshadsource))
	gl.ShaderSource(vertshad, 1, &ptrtobyte, &length)
	gl.CompileShader(vertshad)
	printshaderinfolog(vertshad)
	gl.AttachShader(prog, vertshad)

	fragshad := gl.CreateShader(gl.FRAGMENT_SHADER)
	ptrtobyte = &fshadsource[0]
	length = int32(len(vshadsource))
	gl.ShaderSource(fragshad, 1, &ptrtobyte, &length)
	gl.CompileShader(fragshad)
	printshaderinfolog(fragshad)
	gl.AttachShader(prog, fragshad)

	gl.LinkProgram(prog)
	var log [1000]byte
	gl.GetProgramInfoLog(prog, 1000, nil, &log[0])
	fmt.Println(string(log[:]))

	gl.UseProgram(prog)

	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4)
	gl.Flip()
	time.Sleep(1*time.Second)

	gl.Q()

}
