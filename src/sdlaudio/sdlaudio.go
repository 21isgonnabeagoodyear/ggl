package sdlaudio

/*
#cgo CFLAGS: -g -I/usr/local/include/SDL2 -D_REENTRANT
#cgo LDFLAGS: -L/usr/local/lib -Wl,-rpath,/usr/local/lib -lSDL2 -lpthread
#include <SDL.h>
void audio_callback_go_cgo(void *userdata, void* stream, int length);
*/
import "C"


import "unsafe"
import "time"
import "math"

var dev C.SDL_AudioDeviceID
var acb func(apos int64) (int16, int16) = func(apos int64) (int16,int16){return 0, 0}

var audiopos int64
//export audio_callback_go
func audio_callback_go(userdata unsafe.Pointer, stream unsafe.Pointer, length C.int){
	for i :=0;i<int(length);i+=4{
		l, r := acb(audiopos)
		*((*int16)(unsafe.Pointer(uintptr(stream)+uintptr(i)))) = l
		*((*int16)(unsafe.Pointer(uintptr(stream)+uintptr(i)+2))) = r
		audiopos ++
	}
	//fmt.Println("called", stream, length, *(*int16)(stream))

}
func OutFunc(cb func(apos int64) (int16, int16)){
	acb = cb
}


func Test(){
	acb = func(apos int64) (int16,int16){return int16((math.Sin(float64(apos)/100))/5 * 64000/2), int16((math.Sin(float64(apos)/50))/5 * 64000/2)}
	Init()
	Enable()
	time.Sleep(time.Second*10)

}
func Enable(){
	C.SDL_PauseAudioDevice(dev, 0)
}
func Disable(){
	C.SDL_PauseAudioDevice(dev, 1)
}

func InitSDL(){
	if ok := C.SDL_Init(C.SDL_INIT_AUDIO); ok < 0{
		panic(ok)
	}
}
func Init(){
	if ok := C.SDL_InitSubSystem(C.SDL_INIT_AUDIO); ok < 0{
		panic(ok)
	}
	var want C.SDL_AudioSpec
	want.freq = 48000
	want.format = C.AUDIO_S16
	want.channels = 2
	want.callback=(C.SDL_AudioCallback)(unsafe.Pointer(C.audio_callback_go_cgo))
	var have C.SDL_AudioSpec
	if dev = C.SDL_OpenAudioDevice(nil, 0, &want, &have, 0); dev ==0{
		panic("dev=0")
	}
}
