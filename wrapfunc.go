package main
//go run wrapfunc.go >src/gl/autogened.go
import "fmt"
import "strings"
import "os"
import "bufio"

//typemap = make(map[string] string)
var fnblacklist = map[string]bool{
"glTexPageCommitmentARB":true,
"glDispatchComputeGroupSizeARB":true,
"glIsImageHandleResidentARB":true,
"glIsTextureHandleResidentARB":true,
"glMakeImageHandleNonResidentARB":true,
"glMakeTextureHandleResidentARB":true,
"glMakeTextureHandleNonResidentARB":true,

}

var typemap map[string] [3]string
func init(){
typemap = make(map[string] [3]string)//type, convertto
typemap["GLenum"] =                   [3]string{"int", "\t$o := C.GLenum($i)", "\t$o := int($i)"}
typemap["GLint"] =                    [3]string{"int", "\t$o := C.GLint($i)", "\tpanic()"}
typemap["GLsizei"] =                  [3]string{"int", "\t$o := C.GLsizei($i)", "\tpanic()"}
typemap["const void *"] =             [3]string{"[]byte", "\tpanic()", "\tpanic()"}
typemap["GLfloat"] =                  [3]string{"float32", "\t$o := C.GLfloat($i)", "\tpanic()"}
typemap["GLboolean"] =                [3]string{"bool", "\t$o := C.GLboolean(1);if !$i{$o=C.GLboolean(0)}", "\tpanic()"}
typemap["GLuint"] =                   [3]string{"uint", "\t$o := C.GLuint($i)", "\tpanic()"}
typemap["void **"] =                  [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const void *"] =             [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const GLuint *"] =           [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLsizeiptr"] =               [3]string{"int", "\t$o := C.GLsizeiptr($i)", "\tpanic()"}
typemap["const GLint *"] =            [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLuint *"] =                 [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["void"] =                     [3]string{"", "", "\tpanic()"}
typemap["const GLfloat *"] =          [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLbitfield"] =               [3]string{"uint", "\t$o := C.GLbitfield($i)", "\t$o := uint($i)"}
typemap["GLdouble"] =                 [3]string{"float64", "\t$o := C.GLdouble($i)", "\tpanic()"}
typemap["GLboolean *"] =              [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLdouble *"] =               [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const GLdouble *"] =         [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLfloat *"] =                [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLint *"] =                  [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const_GLubyte_*"] =          [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["void *"] =                   [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["void*"] =                    [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const GLsizei *"] =          [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLsizei *"] =                [3]string{"int", "\t$o := C.GLsizei($i)", "\tpanic()"}
typemap["GLchar *"] =                 [3]string{"int", "\tpanic()", "\tpanic()"}
typemap["const void *const*"] =       [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLintptr"] =                 [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLshort"] =                  [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const GLshort *"] =          [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLenum *"] =                 [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const GLchar *"] =           [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const GLchar *const*"] =     [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const GLuint64EXT *"] =      [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["struct _cl_context *"] =     [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLint64 *"] =                [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLuint64"] =                 [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLuint64EXT"] =              [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const GLuint64 *"] =         [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["struct _cl_event *"] =       [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLuint64EXT *"] =            [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLDEBUGPROC"] =              [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLDEBUGPROCARB"] =           [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const GLintptr *"] =         [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const GLenum *"] =           [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLuint64 *"] =               [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLsync"] =                   [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const char *__restrict"] =   [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const GLushort *"] =         [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["GLubyte"] =                  [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const GLbyte *"] =           [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const GLubyte *"] =          [3]string{"???", "\tpanic()", "\tpanic()"}
typemap["const GLsizeiptr *"] =       [3]string{"???", "\tpanic()", "\tpanic()"}





typemap[""] =                         [3]string{"", "\tpanic()", "\tpanic()"}
}

func gotype(ctype string) string{
	if rv, ok := typemap[ctype]; ok{
		return rv[0]
	}
	return "(XXX:unknown type \""+ctype+"\")"
}
func goconverter(ctype string) string{
	if rv, ok := typemap[ctype]; ok{
		return rv[1]
	}
	return "(XXX:unknown type \""+ctype+"\")"
}
func goconverterfrom(ctype string) string{
	if rv, ok := typemap[ctype]; ok{
		return rv[2]
	}
	return "(XXX:unknown type \""+ctype+"\")"
}

func returntype(in string) string{
	//if strings.Split(in, " ")[1] == "void"{return ""}
	return strings.Split(in, " ")[1]
}
func fnname(in string) string{
	return strings.Split(in, " ")[2]
}
func params(in string) [][2]string{
	rv := make([][2]string, 0, 0)
	paramstring := strings.Split(in, "(")[1];
	paramstring = paramstring[0:len(paramstring)-2];
	pars := strings.Split(paramstring, ",");
	if len(pars) == 1 && strings.Trim(pars[0], " ") == "void"{return make([][2]string, 0, 0)}//void function
	for _, p := range pars{
		p = strings.Trim(p, " ")
		lastind := 0
		for i := range p{
			if p[i] == ' ' || p[i] == '*'{lastind = i+1}
		}
		//rv = append(rv, [2]string{"\""+strings.Trim(p[0:lastind], " ")+"\"", "\""+strings.Trim(p[lastind:], " ")+"\""})
		if strings.Trim(p[lastind:], " ") == "func"{
			rv = append(rv, [2]string{strings.Trim(p[0:lastind], " "), strings.Trim("function", " ")})
		}else if strings.Trim(p[lastind:], " ") == "type"{
			rv = append(rv, [2]string{strings.Trim(p[0:lastind], " "), strings.Trim("whichtype", " ")})
		}else{
			rv = append(rv, [2]string{strings.Trim(p[0:lastind], " "), strings.Trim(p[lastind:], " ")})
		}
	}
	return rv
}
func makeconvertto(proto string) string{
	rv := ""
	for _, param := range params(proto){
		rv = rv + strings.Replace(strings.Replace(goconverter(param[0]), "$o", "_"+param[1], -1), "$i", param[1], -1)+"\n"
	}
	return rv
}
func makecall(proto string) string{
	rv := ""
	for _, param := range params(proto){
		rv += "_"+param[1]+", "
	}
	if len(rv)>2{
		rv = rv[:len(rv)-2]
	}
	if returntype(proto) == "void"{
		return "\tC."+fnname(proto)+"("+rv+")\n"
	}else{
		return "\treturnvalue := C."+fnname(proto)+"("+rv+")\n"+strings.Replace(strings.Replace(goconverterfrom(returntype(proto)), "$o", "convreturnvalue", -1), "$i", "returnvalue", -1)+"\nreturn convreturnvalue\n"
	}
}
func main(){
//	str :="extern void glGetUniformSubroutineuiv (GLenum shadertype, GLint location, GLuint *params);"
	fmt.Println(`
package gl
// #cgo CFLAGS: -g -I/usr/local/include/SDL2 -D_REENTRANT
// #cgo LDFLAGS: -L/usr/local/lib -Wl,-rpath,/usr/local/lib -lSDL2 -lpthread -lGL
// #define GL_GLEXT_PROTOTYPES
// #include "SDL.h"
// #include "glcorearb.h"
// #include <stdlib.h>
import "C"
`)
	definesfile, _ := os.Open("defines.h")
	definescanner := bufio.NewScanner(definesfile)
	for definescanner.Scan(){
		fmt.Println("const "+strings.Replace(strings.Replace(definescanner.Text(), "#define ", "", 1)[3:], " ", "=", 1))//so sketch
	}


	inputfile, _ := os.Open("api.h")
	scanner := bufio.NewScanner(inputfile)
	failedcount := 0
	succeededcount := 0
	for scanner.Scan(){
		str := scanner.Text()
		//fmt.Println(params(str))
		newparams := ""
		for _, par := range params(str){
			newparams += par[1]+" "+gotype(par[0])+", "
		}
		if(len(newparams)>2){
			newparams = newparams[:len(newparams)-2]
		}
		funcdata := "func "+fnname(str)[2:]+"("+newparams+")"+gotype(returntype(str))+"{\n"+makeconvertto(str)+makecall(str)+"}"
		if _, ok := fnblacklist[fnname(str)]; ok || strings.Index(funcdata, "panic") != -1{fmt.Println("/*"+funcdata+"*/");failedcount ++;}else{fmt.Println(funcdata);succeededcount++}
	}



	fmt.Println("//failed: ", failedcount)
	fmt.Println("//succeeded: ", succeededcount)
}
