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


var win *[0]byte

func InitGL(width, height, msaa int){
	C.SDL_Init(C.SDL_INIT_VIDEO)
        C.SDL_VideoInit(/*nil*/C.SDL_GetVideoDriver(0))
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MAJOR_VERSION, 3)
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MINOR_VERSION, 2)

	C.SDL_GL_SetAttribute(C.SDL_GL_MULTISAMPLEBUFFERS, 1)
        C.SDL_GL_SetAttribute(C.SDL_GL_MULTISAMPLESAMPLES, C.int(msaa))

C.SDL_GL_SetAttribute( C.SDL_GL_DEPTH_SIZE, 16 );
	win = C.SDL_CreateWindow(nil, C.SDL_WINDOWPOS_CENTERED, C.SDL_WINDOWPOS_CENTERED,C.int(width), C.int(height), C.SDL_WINDOW_OPENGL)

	C.SDL_ShowWindow(win)
	wat := C.SDL_GL_CreateContext(win)
	fmt.Println(C.GoString(C.SDL_GetVideoDriver(0)))

	C.SDL_GL_MakeCurrent(win, wat)
	C.SDL_GL_SetSwapInterval(1)

	C.glEnable(C.GL_DEPTH_TEST);
	C.glDepthFunc(C.GL_LEQUAL)

	C.glClearColor(0.3,0.5,1,0)
	C.glClear(C.GL_COLOR_BUFFER_BIT|C.GL_DEPTH_BUFFER_BIT)
	printerr("failed to initialize openGL")

}
func Flip(){
	C.SDL_GL_SwapWindow(win)
}
