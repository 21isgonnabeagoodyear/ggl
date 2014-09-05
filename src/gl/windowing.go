package gl
// #cgo CFLAGS: -g -I/usr/local/include/SDL2 -D_REENTRANT
// #cgo LDFLAGS: -L/usr/local/lib -Wl,-rpath,/usr/local/lib -lSDL2 -lpthread -lGL
// #define GL_GLEXT_PROTOTYPES
// #include "SDL.h"
// #include "glcorearb.h"
// #include <stdlib.h>
import "C"
import "errors"
import "fmt"
import "runtime"
import "unsafe"

func printerr( text string) error{
	_,file,line,_ := runtime.Caller(1)
	_,upperfile,upperline,_ := runtime.Caller(2)
	err := C.GoString(C.SDL_GetError())
	errtxt := err
	if len(err)>0{
		fmt.Println("SDL error: in ", file, " line ", line," after ", upperfile, " line ", upperline ,err, text)
	}
	glerr := C.glGetError()
	errtxt += map[int]string{0x0500:"GL_INVALID_ENUM",0x0501:"GL_INVALID_VALUE",0x0502:"GL_INVALID_OPERATION", 0x0505:"GL_OUT_OF_MEMORY"}[int(glerr)]
	if glerr != C.GL_NO_ERROR{
		fmt.Println("GL error: in ", file, " line ", line ," after ", upperfile, " line ", upperline , glerr, "("+map[int]string{0x0500:"GL_INVALID_ENUM",0x0501:"GL_INVALID_VALUE",0x0502:"GL_INVALID_OPERATION", 0x0505:"GL_OUT_OF_MEMORY"}[int(glerr)]+")", text)
	}
	if len(errtxt) >0 {return errors.New(errtxt)}
	return nil
}
func Printerr( text string) error{
	return printerr(text)
}


var win *C.struct_SDL_Window

func InitGL(width, height, msaa int, fullscreen bool){
	C.SDL_Init(C.SDL_INIT_VIDEO)
        C.SDL_VideoInit(/*nil*/C.SDL_GetVideoDriver(0))
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MAJOR_VERSION, 3)
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MINOR_VERSION, 2)

	if msaa != 0{
		C.SDL_GL_SetAttribute(C.SDL_GL_MULTISAMPLEBUFFERS, 1)
		C.SDL_GL_SetAttribute(C.SDL_GL_MULTISAMPLESAMPLES, C.int(msaa))
	}

C.SDL_GL_SetAttribute( C.SDL_GL_DEPTH_SIZE, 16 );
	//win = C.SDL_CreateWindow(nil, C.SDL_WINDOWPOS_CENTERED, C.SDL_WINDOWPOS_CENTERED,C.int(width), C.int(height), C.SDL_WINDOW_OPENGL)
	if fullscreen{
		win = C.SDL_CreateWindow(nil, C.SDL_WINDOWPOS_CENTERED, C.SDL_WINDOWPOS_CENTERED,C.int(width), C.int(height), C.SDL_WINDOW_OPENGL|C.SDL_WINDOW_FULLSCREEN_DESKTOP)
	}else{
		win = C.SDL_CreateWindow(nil, C.SDL_WINDOWPOS_CENTERED, C.SDL_WINDOWPOS_CENTERED,C.int(width), C.int(height), C.SDL_WINDOW_OPENGL)
	}

	C.SDL_ShowWindow(win)
	wat := C.SDL_GL_CreateContext(win)
	fmt.Println(C.GoString(C.SDL_GetVideoDriver(0)))

	C.SDL_GL_MakeCurrent(win, wat)
	//C.SDL_GL_SetSwapInterval(1)

	C.glEnable(C.GL_DEPTH_TEST);
	C.glDepthFunc(C.GL_LEQUAL)

	C.glClearColor(0.3,0.5,1,0)
	C.glClear(C.GL_COLOR_BUFFER_BIT|C.GL_DEPTH_BUFFER_BIT)
	printerr("failed to initialize openGL")

}
func Vsync(enable bool){
	if enable{
		C.SDL_GL_SetSwapInterval(1)
	}else{
		C.SDL_GL_SetSwapInterval(0)
	}
}
func Flip(){
	C.SDL_GL_SwapWindow(win)
}
func Q(){
        C.SDL_Quit()
}
func str2scancode(in string) int{
	if len(in) ==1{
		if in[0] >= 'a' && in[0] <= 'z'{
			return int(in[0])-'a'+C.SDL_SCANCODE_A
		}
		if in[0] >= '0' && in[0] <= '9'{
			return int(in[0])-'0'+C.SDL_SCANCODE_0
		}
	}
	if in == "rarrow"{return C.SDL_SCANCODE_RIGHT}
	if in == "larrow"{return C.SDL_SCANCODE_LEFT}
	if in == "uarrow"{return C.SDL_SCANCODE_UP}
	if in == "darrow"{return C.SDL_SCANCODE_DOWN}
	if in == "esc"{return C.SDL_SCANCODE_ESCAPE}
	if in == "ctrl"{return C.SDL_SCANCODE_LCTRL}
	if in == "shift"{return C.SDL_SCANCODE_LSHIFT}
	if in == " "{return C.SDL_SCANCODE_SPACE}
	if in == "f1"{return C.SDL_SCANCODE_F1}
	if in == "f2"{return C.SDL_SCANCODE_F2}
	if in == "f3"{return C.SDL_SCANCODE_F3}
	if in == "f4"{return C.SDL_SCANCODE_F4}
	if in == "f5"{return C.SDL_SCANCODE_F5}
	if in == "f6"{return C.SDL_SCANCODE_F6}
	if in == "f7"{return C.SDL_SCANCODE_F7}
	if in == "f8"{return C.SDL_SCANCODE_F8}
	if in == "f9"{return C.SDL_SCANCODE_F9}
	if in == "f10"{return C.SDL_SCANCODE_F10}
	if in == "f11"{return C.SDL_SCANCODE_F11}
	if in == "f12"{return C.SDL_SCANCODE_F12}


	return 0

}

var joy *C.SDL_Joystick
func makestickopened(){
	if joy == nil{
		C.SDL_InitSubSystem(C.SDL_INIT_JOYSTICK)
		fmt.Println(int(C.SDL_NumJoysticks()))
		joy = C.SDL_JoystickOpen(0)
		fmt.Println(joy)
	}
}

func Keystate(key string) int{
	if key == "joy1stick1x"{
		makestickopened()
		return int(C.SDL_JoystickGetAxis(joy, 0))
	}
	if key == "joy1stick1y"{
		makestickopened()
		return int(C.SDL_JoystickGetAxis(joy, 1))
	}
	if key == "joy1stick2x"{
		makestickopened()
		return int(C.SDL_JoystickGetAxis(joy, 3))
	}
	if key == "joy1stick2y"{
		makestickopened()
		return int(C.SDL_JoystickGetAxis(joy, 4))
	}
	if key == "joy1a"{
		makestickopened()
		return int(C.SDL_JoystickGetButton(joy, 0))
	}
	var length C.int
	keys := C.SDL_GetKeyboardState(&length)
	return int(*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(keys))+uintptr(str2scancode(key)))))//letter scancodes start at 4

}
func Updatekeys(){
	C.SDL_PumpEvents()
}
func Mouse() (int, int, uint32){
	var x C.int
	var y C.int
	state := C.SDL_GetMouseState(&x,&y)
	return int(x),int(y),uint32(state)
}
func Mouserelative() (int, int, uint32){
	var x C.int
	var y C.int
	state := C.SDL_GetRelativeMouseState(&x,&y)
	C.SDL_SetRelativeMouseMode(C.SDL_TRUE)
	return int(x),int(y),uint32(state)
/*
	state := C.SDL_GetMouseState(&x,&y)
	C.SDL_WarpMouseInWindow(win, 100,100)
	return int(x)-100,int(y)-100,uint32(state)
*/
}
func Fullscreen(fs bool){
	if fs{
		C.SDL_SetWindowFullscreen(win, C.SDL_WINDOW_FULLSCREEN_DESKTOP)
	}else{
		C.SDL_SetWindowFullscreen(win, 0)
	}
}
func Isfocussed() bool{
	return C.SDL_GetWindowFlags(win)&C.SDL_WINDOW_INPUT_FOCUS != 0
}
func TrapMouse(trap bool){
	if trap{
		C.SDL_SetWindowGrab(win, C.SDL_TRUE)
	}else{
		C.SDL_SetWindowGrab(win, C.SDL_FALSE)
	}
}
func ShowCursor(show bool){
	if show{
		C.SDL_ShowCursor(C.SDL_TRUE)
	}else{
		C.SDL_ShowCursor(C.SDL_FALSE)
	}
}

func GetWindowSize() (int, int){
	var w, h C.int
	C.SDL_GetWindowSize(win, &w, &h)
	return int(w), int(h)
}
