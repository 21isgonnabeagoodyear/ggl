
package gl
// #cgo CFLAGS: -g -I/usr/local/include/SDL2 -D_REENTRANT
// #cgo LDFLAGS: -L/usr/local/lib -Wl,-rpath,/usr/local/lib -lSDL2 -lpthread -lGL
// #define GL_GLEXT_PROTOTYPES
// #include "SDL.h"
// #include "glcorearb.h"
// #include <stdlib.h>
import "C"

const VERSION_1_0=1
const VERSION_1_1=1
const DEPTH_BUFFER_BIT=              0x00000100
const STENCIL_BUFFER_BIT=            0x00000400
const COLOR_BUFFER_BIT=              0x00004000
const FALSE=                         0
const TRUE=                          1
const POINTS=                        0x0000
const LINES=                         0x0001
const LINE_LOOP=                     0x0002
const LINE_STRIP=                    0x0003
const TRIANGLES=                     0x0004
const TRIANGLE_STRIP=                0x0005
const TRIANGLE_FAN=                  0x0006
const QUADS=                         0x0007
const NEVER=                         0x0200
const LESS=                          0x0201
const EQUAL=                         0x0202
const LEQUAL=                        0x0203
const GREATER=                       0x0204
const NOTEQUAL=                      0x0205
const GEQUAL=                        0x0206
const ALWAYS=                        0x0207
const ZERO=                          0
const ONE=                           1
const SRC_COLOR=                     0x0300
const ONE_MINUS_SRC_COLOR=           0x0301
const SRC_ALPHA=                     0x0302
const ONE_MINUS_SRC_ALPHA=           0x0303
const DST_ALPHA=                     0x0304
const ONE_MINUS_DST_ALPHA=           0x0305
const DST_COLOR=                     0x0306
const ONE_MINUS_DST_COLOR=           0x0307
const SRC_ALPHA_SATURATE=            0x0308
const NONE=                          0
const FRONT_LEFT=                    0x0400
const FRONT_RIGHT=                   0x0401
const BACK_LEFT=                     0x0402
const BACK_RIGHT=                    0x0403
const FRONT=                         0x0404
const BACK=                          0x0405
const LEFT=                          0x0406
const RIGHT=                         0x0407
const FRONT_AND_BACK=                0x0408
const NO_ERROR=                      0
const INVALID_ENUM=                  0x0500
const INVALID_VALUE=                 0x0501
const INVALID_OPERATION=             0x0502
const OUT_OF_MEMORY=                 0x0505
const CW=                            0x0900
const CCW=                           0x0901
const POINT_SIZE=                    0x0B11
const POINT_SIZE_RANGE=              0x0B12
const POINT_SIZE_GRANULARITY=        0x0B13
const LINE_SMOOTH=                   0x0B20
const LINE_WIDTH=                    0x0B21
const LINE_WIDTH_RANGE=              0x0B22
const LINE_WIDTH_GRANULARITY=        0x0B23
const POLYGON_MODE=                  0x0B40
const POLYGON_SMOOTH=                0x0B41
const CULL_FACE=                     0x0B44
const CULL_FACE_MODE=                0x0B45
const FRONT_FACE=                    0x0B46
const DEPTH_RANGE=                   0x0B70
const DEPTH_TEST=                    0x0B71
const DEPTH_WRITEMASK=               0x0B72
const DEPTH_CLEAR_VALUE=             0x0B73
const DEPTH_FUNC=                    0x0B74
const STENCIL_TEST=                  0x0B90
const STENCIL_CLEAR_VALUE=           0x0B91
const STENCIL_FUNC=                  0x0B92
const STENCIL_VALUE_MASK=            0x0B93
const STENCIL_FAIL=                  0x0B94
const STENCIL_PASS_DEPTH_FAIL=       0x0B95
const STENCIL_PASS_DEPTH_PASS=       0x0B96
const STENCIL_REF=                   0x0B97
const STENCIL_WRITEMASK=             0x0B98
const VIEWPORT=                      0x0BA2
const DITHER=                        0x0BD0
const BLEND_DST=                     0x0BE0
const BLEND_SRC=                     0x0BE1
const BLEND=                         0x0BE2
const LOGIC_OP_MODE=                 0x0BF0
const COLOR_LOGIC_OP=                0x0BF2
const DRAW_BUFFER=                   0x0C01
const READ_BUFFER=                   0x0C02
const SCISSOR_BOX=                   0x0C10
const SCISSOR_TEST=                  0x0C11
const COLOR_CLEAR_VALUE=             0x0C22
const COLOR_WRITEMASK=               0x0C23
const DOUBLEBUFFER=                  0x0C32
const STEREO=                        0x0C33
const LINE_SMOOTH_HINT=              0x0C52
const POLYGON_SMOOTH_HINT=           0x0C53
const UNPACK_SWAP_BYTES=             0x0CF0
const UNPACK_LSB_FIRST=              0x0CF1
const UNPACK_ROW_LENGTH=             0x0CF2
const UNPACK_SKIP_ROWS=              0x0CF3
const UNPACK_SKIP_PIXELS=            0x0CF4
const UNPACK_ALIGNMENT=              0x0CF5
const PACK_SWAP_BYTES=               0x0D00
const PACK_LSB_FIRST=                0x0D01
const PACK_ROW_LENGTH=               0x0D02
const PACK_SKIP_ROWS=                0x0D03
const PACK_SKIP_PIXELS=              0x0D04
const PACK_ALIGNMENT=                0x0D05
const MAX_TEXTURE_SIZE=              0x0D33
const MAX_VIEWPORT_DIMS=             0x0D3A
const SUBPIXEL_BITS=                 0x0D50
const TEXTURE_1D=                    0x0DE0
const TEXTURE_2D=                    0x0DE1
const POLYGON_OFFSET_UNITS=          0x2A00
const POLYGON_OFFSET_POINT=          0x2A01
const POLYGON_OFFSET_LINE=           0x2A02
const POLYGON_OFFSET_FILL=           0x8037
const POLYGON_OFFSET_FACTOR=         0x8038
const TEXTURE_BINDING_1D=            0x8068
const TEXTURE_BINDING_2D=            0x8069
const TEXTURE_WIDTH=                 0x1000
const TEXTURE_HEIGHT=                0x1001
const TEXTURE_INTERNAL_FORMAT=       0x1003
const TEXTURE_BORDER_COLOR=          0x1004
const TEXTURE_RED_SIZE=              0x805C
const TEXTURE_GREEN_SIZE=            0x805D
const TEXTURE_BLUE_SIZE=             0x805E
const TEXTURE_ALPHA_SIZE=            0x805F
const DONT_CARE=                     0x1100
const FASTEST=                       0x1101
const NICEST=                        0x1102
const BYTE=                          0x1400
const UNSIGNED_BYTE=                 0x1401
const SHORT=                         0x1402
const UNSIGNED_SHORT=                0x1403
const INT=                           0x1404
const UNSIGNED_INT=                  0x1405
const FLOAT=                         0x1406
const DOUBLE=                        0x140A
const STACK_OVERFLOW=                0x0503
const STACK_UNDERFLOW=               0x0504
const CLEAR=                         0x1500
const AND=                           0x1501
const AND_REVERSE=                   0x1502
const COPY=                          0x1503
const AND_INVERTED=                  0x1504
const NOOP=                          0x1505
const XOR=                           0x1506
const OR=                            0x1507
const NOR=                           0x1508
const EQUIV=                         0x1509
const INVERT=                        0x150A
const OR_REVERSE=                    0x150B
const COPY_INVERTED=                 0x150C
const OR_INVERTED=                   0x150D
const NAND=                          0x150E
const SET=                           0x150F
const TEXTURE=                       0x1702
const COLOR=                         0x1800
const DEPTH=                         0x1801
const STENCIL=                       0x1802
const STENCIL_INDEX=                 0x1901
const DEPTH_COMPONENT=               0x1902
const RED=                           0x1903
const GREEN=                         0x1904
const BLUE=                          0x1905
const ALPHA=                         0x1906
const RGB=                           0x1907
const RGBA=                          0x1908
const POINT=                         0x1B00
const LINE=                          0x1B01
const FILL=                          0x1B02
const KEEP=                          0x1E00
const REPLACE=                       0x1E01
const INCR=                          0x1E02
const DECR=                          0x1E03
const VENDOR=                        0x1F00
const RENDERER=                      0x1F01
const VERSION=                       0x1F02
const EXTENSIONS=                    0x1F03
const NEAREST=                       0x2600
const LINEAR=                        0x2601
const NEAREST_MIPMAP_NEAREST=        0x2700
const LINEAR_MIPMAP_NEAREST=         0x2701
const NEAREST_MIPMAP_LINEAR=         0x2702
const LINEAR_MIPMAP_LINEAR=          0x2703
const TEXTURE_MAG_FILTER=            0x2800
const TEXTURE_MIN_FILTER=            0x2801
const TEXTURE_WRAP_S=                0x2802
const TEXTURE_WRAP_T=                0x2803
const PROXY_TEXTURE_1D=              0x8063
const PROXY_TEXTURE_2D=              0x8064
const REPEAT=                        0x2901
const R3_G3_B2=                      0x2A10
const RGB4=                          0x804F
const RGB5=                          0x8050
const RGB8=                          0x8051
const RGB10=                         0x8052
const RGB12=                         0x8053
const RGB16=                         0x8054
const RGBA2=                         0x8055
const RGBA4=                         0x8056
const RGB5_A1=                       0x8057
const RGBA8=                         0x8058
const RGB10_A2=                      0x8059
const RGBA12=                        0x805A
const RGBA16=                        0x805B
const VERTEX_ARRAY=                  0x8074
const VERSION_1_2=1
const UNSIGNED_BYTE_3_3_2=           0x8032
const UNSIGNED_SHORT_4_4_4_4=        0x8033
const UNSIGNED_SHORT_5_5_5_1=        0x8034
const UNSIGNED_INT_8_8_8_8=          0x8035
const UNSIGNED_INT_10_10_10_2=       0x8036
const TEXTURE_BINDING_3D=            0x806A
const PACK_SKIP_IMAGES=              0x806B
const PACK_IMAGE_HEIGHT=             0x806C
const UNPACK_SKIP_IMAGES=            0x806D
const UNPACK_IMAGE_HEIGHT=           0x806E
const TEXTURE_3D=                    0x806F
const PROXY_TEXTURE_3D=              0x8070
const TEXTURE_DEPTH=                 0x8071
const TEXTURE_WRAP_R=                0x8072
const MAX_3D_TEXTURE_SIZE=           0x8073
const UNSIGNED_BYTE_2_3_3_REV=       0x8362
const UNSIGNED_SHORT_5_6_5=          0x8363
const UNSIGNED_SHORT_5_6_5_REV=      0x8364
const UNSIGNED_SHORT_4_4_4_4_REV=    0x8365
const UNSIGNED_SHORT_1_5_5_5_REV=    0x8366
const UNSIGNED_INT_8_8_8_8_REV=      0x8367
const UNSIGNED_INT_2_10_10_10_REV=   0x8368
const BGR=                           0x80E0
const BGRA=                          0x80E1
const MAX_ELEMENTS_VERTICES=         0x80E8
const MAX_ELEMENTS_INDICES=          0x80E9
const CLAMP_TO_EDGE=                 0x812F
const TEXTURE_MIN_LOD=               0x813A
const TEXTURE_MAX_LOD=               0x813B
const TEXTURE_BASE_LEVEL=            0x813C
const TEXTURE_MAX_LEVEL=             0x813D
const SMOOTH_POINT_SIZE_RANGE=       0x0B12
const SMOOTH_POINT_SIZE_GRANULARITY= 0x0B13
const SMOOTH_LINE_WIDTH_RANGE=       0x0B22
const SMOOTH_LINE_WIDTH_GRANULARITY= 0x0B23
const ALIASED_LINE_WIDTH_RANGE=      0x846E
const VERSION_1_3=1
const TEXTURE0=                      0x84C0
const TEXTURE1=                      0x84C1
const TEXTURE2=                      0x84C2
const TEXTURE3=                      0x84C3
const TEXTURE4=                      0x84C4
const TEXTURE5=                      0x84C5
const TEXTURE6=                      0x84C6
const TEXTURE7=                      0x84C7
const TEXTURE8=                      0x84C8
const TEXTURE9=                      0x84C9
const TEXTURE10=                     0x84CA
const TEXTURE11=                     0x84CB
const TEXTURE12=                     0x84CC
const TEXTURE13=                     0x84CD
const TEXTURE14=                     0x84CE
const TEXTURE15=                     0x84CF
const TEXTURE16=                     0x84D0
const TEXTURE17=                     0x84D1
const TEXTURE18=                     0x84D2
const TEXTURE19=                     0x84D3
const TEXTURE20=                     0x84D4
const TEXTURE21=                     0x84D5
const TEXTURE22=                     0x84D6
const TEXTURE23=                     0x84D7
const TEXTURE24=                     0x84D8
const TEXTURE25=                     0x84D9
const TEXTURE26=                     0x84DA
const TEXTURE27=                     0x84DB
const TEXTURE28=                     0x84DC
const TEXTURE29=                     0x84DD
const TEXTURE30=                     0x84DE
const TEXTURE31=                     0x84DF
const ACTIVE_TEXTURE=                0x84E0
const MULTISAMPLE=                   0x809D
const SAMPLE_ALPHA_TO_COVERAGE=      0x809E
const SAMPLE_ALPHA_TO_ONE=           0x809F
const SAMPLE_COVERAGE=               0x80A0
const SAMPLE_BUFFERS=                0x80A8
const SAMPLES=                       0x80A9
const SAMPLE_COVERAGE_VALUE=         0x80AA
const SAMPLE_COVERAGE_INVERT=        0x80AB
const TEXTURE_CUBE_MAP=              0x8513
const TEXTURE_BINDING_CUBE_MAP=      0x8514
const TEXTURE_CUBE_MAP_POSITIVE_X=   0x8515
const TEXTURE_CUBE_MAP_NEGATIVE_X=   0x8516
const TEXTURE_CUBE_MAP_POSITIVE_Y=   0x8517
const TEXTURE_CUBE_MAP_NEGATIVE_Y=   0x8518
const TEXTURE_CUBE_MAP_POSITIVE_Z=   0x8519
const TEXTURE_CUBE_MAP_NEGATIVE_Z=   0x851A
const PROXY_TEXTURE_CUBE_MAP=        0x851B
const MAX_CUBE_MAP_TEXTURE_SIZE=     0x851C
const COMPRESSED_RGB=                0x84ED
const COMPRESSED_RGBA=               0x84EE
const TEXTURE_COMPRESSION_HINT=      0x84EF
const TEXTURE_COMPRESSED_IMAGE_SIZE= 0x86A0
const TEXTURE_COMPRESSED=            0x86A1
const NUM_COMPRESSED_TEXTURE_FORMATS=0x86A2
const COMPRESSED_TEXTURE_FORMATS=    0x86A3
const CLAMP_TO_BORDER=               0x812D
const VERSION_1_4=1
const BLEND_DST_RGB=                 0x80C8
const BLEND_SRC_RGB=                 0x80C9
const BLEND_DST_ALPHA=               0x80CA
const BLEND_SRC_ALPHA=               0x80CB
const POINT_FADE_THRESHOLD_SIZE=     0x8128
const DEPTH_COMPONENT16=             0x81A5
const DEPTH_COMPONENT24=             0x81A6
const DEPTH_COMPONENT32=             0x81A7
const MIRRORED_REPEAT=               0x8370
const MAX_TEXTURE_LOD_BIAS=          0x84FD
const TEXTURE_LOD_BIAS=              0x8501
const INCR_WRAP=                     0x8507
const DECR_WRAP=                     0x8508
const TEXTURE_DEPTH_SIZE=            0x884A
const TEXTURE_COMPARE_MODE=          0x884C
const TEXTURE_COMPARE_FUNC=          0x884D
const FUNC_ADD=                      0x8006
const FUNC_SUBTRACT=                 0x800A
const FUNC_REVERSE_SUBTRACT=         0x800B
const MIN=                           0x8007
const MAX=                           0x8008
const CONSTANT_COLOR=                0x8001
const ONE_MINUS_CONSTANT_COLOR=      0x8002
const CONSTANT_ALPHA=                0x8003
const ONE_MINUS_CONSTANT_ALPHA=      0x8004
const VERSION_1_5=1
const BUFFER_SIZE=                   0x8764
const BUFFER_USAGE=                  0x8765
const QUERY_COUNTER_BITS=            0x8864
const CURRENT_QUERY=                 0x8865
const QUERY_RESULT=                  0x8866
const QUERY_RESULT_AVAILABLE=        0x8867
const ARRAY_BUFFER=                  0x8892
const ELEMENT_ARRAY_BUFFER=          0x8893
const ARRAY_BUFFER_BINDING=          0x8894
const ELEMENT_ARRAY_BUFFER_BINDING=  0x8895
const VERTEX_ATTRIB_ARRAY_BUFFER_BINDING=0x889F
const READ_ONLY=                     0x88B8
const WRITE_ONLY=                    0x88B9
const READ_WRITE=                    0x88BA
const BUFFER_ACCESS=                 0x88BB
const BUFFER_MAPPED=                 0x88BC
const BUFFER_MAP_POINTER=            0x88BD
const STREAM_DRAW=                   0x88E0
const STREAM_READ=                   0x88E1
const STREAM_COPY=                   0x88E2
const STATIC_DRAW=                   0x88E4
const STATIC_READ=                   0x88E5
const STATIC_COPY=                   0x88E6
const DYNAMIC_DRAW=                  0x88E8
const DYNAMIC_READ=                  0x88E9
const DYNAMIC_COPY=                  0x88EA
const SAMPLES_PASSED=                0x8914
const SRC1_ALPHA=                    0x8589
const VERSION_2_0=1
const BLEND_EQUATION_RGB=            0x8009
const VERTEX_ATTRIB_ARRAY_ENABLED=   0x8622
const VERTEX_ATTRIB_ARRAY_SIZE=      0x8623
const VERTEX_ATTRIB_ARRAY_STRIDE=    0x8624
const VERTEX_ATTRIB_ARRAY_TYPE=      0x8625
const CURRENT_VERTEX_ATTRIB=         0x8626
const VERTEX_PROGRAM_POINT_SIZE=     0x8642
const VERTEX_ATTRIB_ARRAY_POINTER=   0x8645
const STENCIL_BACK_FUNC=             0x8800
const STENCIL_BACK_FAIL=             0x8801
const STENCIL_BACK_PASS_DEPTH_FAIL=  0x8802
const STENCIL_BACK_PASS_DEPTH_PASS=  0x8803
const MAX_DRAW_BUFFERS=              0x8824
const DRAW_BUFFER0=                  0x8825
const DRAW_BUFFER1=                  0x8826
const DRAW_BUFFER2=                  0x8827
const DRAW_BUFFER3=                  0x8828
const DRAW_BUFFER4=                  0x8829
const DRAW_BUFFER5=                  0x882A
const DRAW_BUFFER6=                  0x882B
const DRAW_BUFFER7=                  0x882C
const DRAW_BUFFER8=                  0x882D
const DRAW_BUFFER9=                  0x882E
const DRAW_BUFFER10=                 0x882F
const DRAW_BUFFER11=                 0x8830
const DRAW_BUFFER12=                 0x8831
const DRAW_BUFFER13=                 0x8832
const DRAW_BUFFER14=                 0x8833
const DRAW_BUFFER15=                 0x8834
const BLEND_EQUATION_ALPHA=          0x883D
const MAX_VERTEX_ATTRIBS=            0x8869
const VERTEX_ATTRIB_ARRAY_NORMALIZED=0x886A
const MAX_TEXTURE_IMAGE_UNITS=       0x8872
const FRAGMENT_SHADER=               0x8B30
const VERTEX_SHADER=                 0x8B31
const MAX_FRAGMENT_UNIFORM_COMPONENTS=0x8B49
const MAX_VERTEX_UNIFORM_COMPONENTS= 0x8B4A
const MAX_VARYING_FLOATS=            0x8B4B
const MAX_VERTEX_TEXTURE_IMAGE_UNITS=0x8B4C
const MAX_COMBINED_TEXTURE_IMAGE_UNITS=0x8B4D
const SHADER_TYPE=                   0x8B4F
const FLOAT_VEC2=                    0x8B50
const FLOAT_VEC3=                    0x8B51
const FLOAT_VEC4=                    0x8B52
const INT_VEC2=                      0x8B53
const INT_VEC3=                      0x8B54
const INT_VEC4=                      0x8B55
const BOOL=                          0x8B56
const BOOL_VEC2=                     0x8B57
const BOOL_VEC3=                     0x8B58
const BOOL_VEC4=                     0x8B59
const FLOAT_MAT2=                    0x8B5A
const FLOAT_MAT3=                    0x8B5B
const FLOAT_MAT4=                    0x8B5C
const SAMPLER_1D=                    0x8B5D
const SAMPLER_2D=                    0x8B5E
const SAMPLER_3D=                    0x8B5F
const SAMPLER_CUBE=                  0x8B60
const SAMPLER_1D_SHADOW=             0x8B61
const SAMPLER_2D_SHADOW=             0x8B62
const DELETE_STATUS=                 0x8B80
const COMPILE_STATUS=                0x8B81
const LINK_STATUS=                   0x8B82
const VALIDATE_STATUS=               0x8B83
const INFO_LOG_LENGTH=               0x8B84
const ATTACHED_SHADERS=              0x8B85
const ACTIVE_UNIFORMS=               0x8B86
const ACTIVE_UNIFORM_MAX_LENGTH=     0x8B87
const SHADER_SOURCE_LENGTH=          0x8B88
const ACTIVE_ATTRIBUTES=             0x8B89
const ACTIVE_ATTRIBUTE_MAX_LENGTH=   0x8B8A
const FRAGMENT_SHADER_DERIVATIVE_HINT=0x8B8B
const SHADING_LANGUAGE_VERSION=      0x8B8C
const CURRENT_PROGRAM=               0x8B8D
const POINT_SPRITE_COORD_ORIGIN=     0x8CA0
const LOWER_LEFT=                    0x8CA1
const UPPER_LEFT=                    0x8CA2
const STENCIL_BACK_REF=              0x8CA3
const STENCIL_BACK_VALUE_MASK=       0x8CA4
const STENCIL_BACK_WRITEMASK=        0x8CA5
const VERSION_2_1=1
const PIXEL_PACK_BUFFER=             0x88EB
const PIXEL_UNPACK_BUFFER=           0x88EC
const PIXEL_PACK_BUFFER_BINDING=     0x88ED
const PIXEL_UNPACK_BUFFER_BINDING=   0x88EF
const FLOAT_MAT2x3=                  0x8B65
const FLOAT_MAT2x4=                  0x8B66
const FLOAT_MAT3x2=                  0x8B67
const FLOAT_MAT3x4=                  0x8B68
const FLOAT_MAT4x2=                  0x8B69
const FLOAT_MAT4x3=                  0x8B6A
const SRGB=                          0x8C40
const SRGB8=                         0x8C41
const SRGB_ALPHA=                    0x8C42
const SRGB8_ALPHA8=                  0x8C43
const COMPRESSED_SRGB=               0x8C48
const COMPRESSED_SRGB_ALPHA=         0x8C49
const VERSION_3_0=1
const COMPARE_REF_TO_TEXTURE=        0x884E
const CLIP_DISTANCE0=                0x3000
const CLIP_DISTANCE1=                0x3001
const CLIP_DISTANCE2=                0x3002
const CLIP_DISTANCE3=                0x3003
const CLIP_DISTANCE4=                0x3004
const CLIP_DISTANCE5=                0x3005
const CLIP_DISTANCE6=                0x3006
const CLIP_DISTANCE7=                0x3007
const MAX_CLIP_DISTANCES=            0x0D32
const MAJOR_VERSION=                 0x821B
const MINOR_VERSION=                 0x821C
const NUM_EXTENSIONS=                0x821D
const CONTEXT_FLAGS=                 0x821E
const COMPRESSED_RED=                0x8225
const COMPRESSED_RG=                 0x8226
const CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT=0x00000001
const RGBA32F=                       0x8814
const RGB32F=                        0x8815
const RGBA16F=                       0x881A
const RGB16F=                        0x881B
const VERTEX_ATTRIB_ARRAY_INTEGER=   0x88FD
const MAX_ARRAY_TEXTURE_LAYERS=      0x88FF
const MIN_PROGRAM_TEXEL_OFFSET=      0x8904
const MAX_PROGRAM_TEXEL_OFFSET=      0x8905
const CLAMP_READ_COLOR=              0x891C
const FIXED_ONLY=                    0x891D
const MAX_VARYING_COMPONENTS=        0x8B4B
const TEXTURE_1D_ARRAY=              0x8C18
const PROXY_TEXTURE_1D_ARRAY=        0x8C19
const TEXTURE_2D_ARRAY=              0x8C1A
const PROXY_TEXTURE_2D_ARRAY=        0x8C1B
const TEXTURE_BINDING_1D_ARRAY=      0x8C1C
const TEXTURE_BINDING_2D_ARRAY=      0x8C1D
const R11F_G11F_B10F=                0x8C3A
const UNSIGNED_INT_10F_11F_11F_REV=  0x8C3B
const RGB9_E5=                       0x8C3D
const UNSIGNED_INT_5_9_9_9_REV=      0x8C3E
const TEXTURE_SHARED_SIZE=           0x8C3F
const TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH=0x8C76
const TRANSFORM_FEEDBACK_BUFFER_MODE=0x8C7F
const MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS=0x8C80
const TRANSFORM_FEEDBACK_VARYINGS=   0x8C83
const TRANSFORM_FEEDBACK_BUFFER_START=0x8C84
const TRANSFORM_FEEDBACK_BUFFER_SIZE=0x8C85
const PRIMITIVES_GENERATED=          0x8C87
const TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN=0x8C88
const RASTERIZER_DISCARD=            0x8C89
const MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS=0x8C8A
const MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS=0x8C8B
const INTERLEAVED_ATTRIBS=           0x8C8C
const SEPARATE_ATTRIBS=              0x8C8D
const TRANSFORM_FEEDBACK_BUFFER=     0x8C8E
const TRANSFORM_FEEDBACK_BUFFER_BINDING=0x8C8F
const RGBA32UI=                      0x8D70
const RGB32UI=                       0x8D71
const RGBA16UI=                      0x8D76
const RGB16UI=                       0x8D77
const RGBA8UI=                       0x8D7C
const RGB8UI=                        0x8D7D
const RGBA32I=                       0x8D82
const RGB32I=                        0x8D83
const RGBA16I=                       0x8D88
const RGB16I=                        0x8D89
const RGBA8I=                        0x8D8E
const RGB8I=                         0x8D8F
const RED_INTEGER=                   0x8D94
const GREEN_INTEGER=                 0x8D95
const BLUE_INTEGER=                  0x8D96
const RGB_INTEGER=                   0x8D98
const RGBA_INTEGER=                  0x8D99
const BGR_INTEGER=                   0x8D9A
const BGRA_INTEGER=                  0x8D9B
const SAMPLER_1D_ARRAY=              0x8DC0
const SAMPLER_2D_ARRAY=              0x8DC1
const SAMPLER_1D_ARRAY_SHADOW=       0x8DC3
const SAMPLER_2D_ARRAY_SHADOW=       0x8DC4
const SAMPLER_CUBE_SHADOW=           0x8DC5
const UNSIGNED_INT_VEC2=             0x8DC6
const UNSIGNED_INT_VEC3=             0x8DC7
const UNSIGNED_INT_VEC4=             0x8DC8
const INT_SAMPLER_1D=                0x8DC9
const INT_SAMPLER_2D=                0x8DCA
const INT_SAMPLER_3D=                0x8DCB
const INT_SAMPLER_CUBE=              0x8DCC
const INT_SAMPLER_1D_ARRAY=          0x8DCE
const INT_SAMPLER_2D_ARRAY=          0x8DCF
const UNSIGNED_INT_SAMPLER_1D=       0x8DD1
const UNSIGNED_INT_SAMPLER_2D=       0x8DD2
const UNSIGNED_INT_SAMPLER_3D=       0x8DD3
const UNSIGNED_INT_SAMPLER_CUBE=     0x8DD4
const UNSIGNED_INT_SAMPLER_1D_ARRAY= 0x8DD6
const UNSIGNED_INT_SAMPLER_2D_ARRAY= 0x8DD7
const QUERY_WAIT=                    0x8E13
const QUERY_NO_WAIT=                 0x8E14
const QUERY_BY_REGION_WAIT=          0x8E15
const QUERY_BY_REGION_NO_WAIT=       0x8E16
const BUFFER_ACCESS_FLAGS=           0x911F
const BUFFER_MAP_LENGTH=             0x9120
const BUFFER_MAP_OFFSET=             0x9121
const DEPTH_COMPONENT32F=            0x8CAC
const DEPTH32F_STENCIL8=             0x8CAD
const FLOAT_32_UNSIGNED_INT_24_8_REV=0x8DAD
const INVALID_FRAMEBUFFER_OPERATION= 0x0506
const FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING=0x8210
const FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE=0x8211
const FRAMEBUFFER_ATTACHMENT_RED_SIZE=0x8212
const FRAMEBUFFER_ATTACHMENT_GREEN_SIZE=0x8213
const FRAMEBUFFER_ATTACHMENT_BLUE_SIZE=0x8214
const FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE=0x8215
const FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE=0x8216
const FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE=0x8217
const FRAMEBUFFER_DEFAULT=           0x8218
const FRAMEBUFFER_UNDEFINED=         0x8219
const DEPTH_STENCIL_ATTACHMENT=      0x821A
const MAX_RENDERBUFFER_SIZE=         0x84E8
const DEPTH_STENCIL=                 0x84F9
const UNSIGNED_INT_24_8=             0x84FA
const DEPTH24_STENCIL8=              0x88F0
const TEXTURE_STENCIL_SIZE=          0x88F1
const TEXTURE_RED_TYPE=              0x8C10
const TEXTURE_GREEN_TYPE=            0x8C11
const TEXTURE_BLUE_TYPE=             0x8C12
const TEXTURE_ALPHA_TYPE=            0x8C13
const TEXTURE_DEPTH_TYPE=            0x8C16
const UNSIGNED_NORMALIZED=           0x8C17
const FRAMEBUFFER_BINDING=           0x8CA6
const DRAW_FRAMEBUFFER_BINDING=      0x8CA6
const RENDERBUFFER_BINDING=          0x8CA7
const READ_FRAMEBUFFER=              0x8CA8
const DRAW_FRAMEBUFFER=              0x8CA9
const READ_FRAMEBUFFER_BINDING=      0x8CAA
const RENDERBUFFER_SAMPLES=          0x8CAB
const FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE=0x8CD0
const FRAMEBUFFER_ATTACHMENT_OBJECT_NAME=0x8CD1
const FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL=0x8CD2
const FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE=0x8CD3
const FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER=0x8CD4
const FRAMEBUFFER_COMPLETE=          0x8CD5
const FRAMEBUFFER_INCOMPLETE_ATTACHMENT=0x8CD6
const FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT=0x8CD7
const FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER=0x8CDB
const FRAMEBUFFER_INCOMPLETE_READ_BUFFER=0x8CDC
const FRAMEBUFFER_UNSUPPORTED=       0x8CDD
const MAX_COLOR_ATTACHMENTS=         0x8CDF
const COLOR_ATTACHMENT0=             0x8CE0
const COLOR_ATTACHMENT1=             0x8CE1
const COLOR_ATTACHMENT2=             0x8CE2
const COLOR_ATTACHMENT3=             0x8CE3
const COLOR_ATTACHMENT4=             0x8CE4
const COLOR_ATTACHMENT5=             0x8CE5
const COLOR_ATTACHMENT6=             0x8CE6
const COLOR_ATTACHMENT7=             0x8CE7
const COLOR_ATTACHMENT8=             0x8CE8
const COLOR_ATTACHMENT9=             0x8CE9
const COLOR_ATTACHMENT10=            0x8CEA
const COLOR_ATTACHMENT11=            0x8CEB
const COLOR_ATTACHMENT12=            0x8CEC
const COLOR_ATTACHMENT13=            0x8CED
const COLOR_ATTACHMENT14=            0x8CEE
const COLOR_ATTACHMENT15=            0x8CEF
const DEPTH_ATTACHMENT=              0x8D00
const STENCIL_ATTACHMENT=            0x8D20
const FRAMEBUFFER=                   0x8D40
const RENDERBUFFER=                  0x8D41
const RENDERBUFFER_WIDTH=            0x8D42
const RENDERBUFFER_HEIGHT=           0x8D43
const RENDERBUFFER_INTERNAL_FORMAT=  0x8D44
const STENCIL_INDEX1=                0x8D46
const STENCIL_INDEX4=                0x8D47
const STENCIL_INDEX8=                0x8D48
const STENCIL_INDEX16=               0x8D49
const RENDERBUFFER_RED_SIZE=         0x8D50
const RENDERBUFFER_GREEN_SIZE=       0x8D51
const RENDERBUFFER_BLUE_SIZE=        0x8D52
const RENDERBUFFER_ALPHA_SIZE=       0x8D53
const RENDERBUFFER_DEPTH_SIZE=       0x8D54
const RENDERBUFFER_STENCIL_SIZE=     0x8D55
const FRAMEBUFFER_INCOMPLETE_MULTISAMPLE=0x8D56
const MAX_SAMPLES=                   0x8D57
const FRAMEBUFFER_SRGB=              0x8DB9
const HALF_FLOAT=                    0x140B
const MAP_READ_BIT=                  0x0001
const MAP_WRITE_BIT=                 0x0002
const MAP_INVALIDATE_RANGE_BIT=      0x0004
const MAP_INVALIDATE_BUFFER_BIT=     0x0008
const MAP_FLUSH_EXPLICIT_BIT=        0x0010
const MAP_UNSYNCHRONIZED_BIT=        0x0020
const COMPRESSED_RED_RGTC1=          0x8DBB
const COMPRESSED_SIGNED_RED_RGTC1=   0x8DBC
const COMPRESSED_RG_RGTC2=           0x8DBD
const COMPRESSED_SIGNED_RG_RGTC2=    0x8DBE
const RG=                            0x8227
const RG_INTEGER=                    0x8228
const R8=                            0x8229
const R16=                           0x822A
const RG8=                           0x822B
const RG16=                          0x822C
const R16F=                          0x822D
const R32F=                          0x822E
const RG16F=                         0x822F
const RG32F=                         0x8230
const R8I=                           0x8231
const R8UI=                          0x8232
const R16I=                          0x8233
const R16UI=                         0x8234
const R32I=                          0x8235
const R32UI=                         0x8236
const RG8I=                          0x8237
const RG8UI=                         0x8238
const RG16I=                         0x8239
const RG16UI=                        0x823A
const RG32I=                         0x823B
const RG32UI=                        0x823C
const VERTEX_ARRAY_BINDING=          0x85B5
const VERSION_3_1=1
const SAMPLER_2D_RECT=               0x8B63
const SAMPLER_2D_RECT_SHADOW=        0x8B64
const SAMPLER_BUFFER=                0x8DC2
const INT_SAMPLER_2D_RECT=           0x8DCD
const INT_SAMPLER_BUFFER=            0x8DD0
const UNSIGNED_INT_SAMPLER_2D_RECT=  0x8DD5
const UNSIGNED_INT_SAMPLER_BUFFER=   0x8DD8
const TEXTURE_BUFFER=                0x8C2A
const MAX_TEXTURE_BUFFER_SIZE=       0x8C2B
const TEXTURE_BINDING_BUFFER=        0x8C2C
const TEXTURE_BUFFER_DATA_STORE_BINDING=0x8C2D
const TEXTURE_RECTANGLE=             0x84F5
const TEXTURE_BINDING_RECTANGLE=     0x84F6
const PROXY_TEXTURE_RECTANGLE=       0x84F7
const MAX_RECTANGLE_TEXTURE_SIZE=    0x84F8
const R8_SNORM=                      0x8F94
const RG8_SNORM=                     0x8F95
const RGB8_SNORM=                    0x8F96
const RGBA8_SNORM=                   0x8F97
const R16_SNORM=                     0x8F98
const RG16_SNORM=                    0x8F99
const RGB16_SNORM=                   0x8F9A
const RGBA16_SNORM=                  0x8F9B
const SIGNED_NORMALIZED=             0x8F9C
const PRIMITIVE_RESTART=             0x8F9D
const PRIMITIVE_RESTART_INDEX=       0x8F9E
const COPY_READ_BUFFER=              0x8F36
const COPY_WRITE_BUFFER=             0x8F37
const UNIFORM_BUFFER=                0x8A11
const UNIFORM_BUFFER_BINDING=        0x8A28
const UNIFORM_BUFFER_START=          0x8A29
const UNIFORM_BUFFER_SIZE=           0x8A2A
const MAX_VERTEX_UNIFORM_BLOCKS=     0x8A2B
const MAX_FRAGMENT_UNIFORM_BLOCKS=   0x8A2D
const MAX_COMBINED_UNIFORM_BLOCKS=   0x8A2E
const MAX_UNIFORM_BUFFER_BINDINGS=   0x8A2F
const MAX_UNIFORM_BLOCK_SIZE=        0x8A30
const MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS=0x8A31
const MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS=0x8A33
const UNIFORM_BUFFER_OFFSET_ALIGNMENT=0x8A34
const ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH=0x8A35
const ACTIVE_UNIFORM_BLOCKS=         0x8A36
const UNIFORM_TYPE=                  0x8A37
const UNIFORM_SIZE=                  0x8A38
const UNIFORM_NAME_LENGTH=           0x8A39
const UNIFORM_BLOCK_INDEX=           0x8A3A
const UNIFORM_OFFSET=                0x8A3B
const UNIFORM_ARRAY_STRIDE=          0x8A3C
const UNIFORM_MATRIX_STRIDE=         0x8A3D
const UNIFORM_IS_ROW_MAJOR=          0x8A3E
const UNIFORM_BLOCK_BINDING=         0x8A3F
const UNIFORM_BLOCK_DATA_SIZE=       0x8A40
const UNIFORM_BLOCK_NAME_LENGTH=     0x8A41
const UNIFORM_BLOCK_ACTIVE_UNIFORMS= 0x8A42
const UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES=0x8A43
const UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER=0x8A44
const UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER=0x8A46
const INVALID_INDEX=                 0xFFFFFFFF
const VERSION_3_2=1
const CONTEXT_CORE_PROFILE_BIT=      0x00000001
const CONTEXT_COMPATIBILITY_PROFILE_BIT=0x00000002
const LINES_ADJACENCY=               0x000A
const LINE_STRIP_ADJACENCY=          0x000B
const TRIANGLES_ADJACENCY=           0x000C
const TRIANGLE_STRIP_ADJACENCY=      0x000D
const PROGRAM_POINT_SIZE=            0x8642
const MAX_GEOMETRY_TEXTURE_IMAGE_UNITS=0x8C29
const FRAMEBUFFER_ATTACHMENT_LAYERED=0x8DA7
const FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS=0x8DA8
const GEOMETRY_SHADER=               0x8DD9
const GEOMETRY_VERTICES_OUT=         0x8916
const GEOMETRY_INPUT_TYPE=           0x8917
const GEOMETRY_OUTPUT_TYPE=          0x8918
const MAX_GEOMETRY_UNIFORM_COMPONENTS=0x8DDF
const MAX_GEOMETRY_OUTPUT_VERTICES=  0x8DE0
const MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS=0x8DE1
const MAX_VERTEX_OUTPUT_COMPONENTS=  0x9122
const MAX_GEOMETRY_INPUT_COMPONENTS= 0x9123
const MAX_GEOMETRY_OUTPUT_COMPONENTS=0x9124
const MAX_FRAGMENT_INPUT_COMPONENTS= 0x9125
const CONTEXT_PROFILE_MASK=          0x9126
const DEPTH_CLAMP=                   0x864F
const QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION=0x8E4C
const FIRST_VERTEX_CONVENTION=       0x8E4D
const LAST_VERTEX_CONVENTION=        0x8E4E
const PROVOKING_VERTEX=              0x8E4F
const TEXTURE_CUBE_MAP_SEAMLESS=     0x884F
const MAX_SERVER_WAIT_TIMEOUT=       0x9111
const OBJECT_TYPE=                   0x9112
const SYNC_CONDITION=                0x9113
const SYNC_STATUS=                   0x9114
const SYNC_FLAGS=                    0x9115
const SYNC_FENCE=                    0x9116
const SYNC_GPU_COMMANDS_COMPLETE=    0x9117
const UNSIGNALED=                    0x9118
const SIGNALED=                      0x9119
const ALREADY_SIGNALED=              0x911A
const TIMEOUT_EXPIRED=               0x911B
const CONDITION_SATISFIED=           0x911C
const WAIT_FAILED=                   0x911D
const TIMEOUT_IGNORED=               0xFFFFFFFFFFFFFFFF
const SYNC_FLUSH_COMMANDS_BIT=       0x00000001
const SAMPLE_POSITION=               0x8E50
const SAMPLE_MASK=                   0x8E51
const SAMPLE_MASK_VALUE=             0x8E52
const MAX_SAMPLE_MASK_WORDS=         0x8E59
const TEXTURE_2D_MULTISAMPLE=        0x9100
const PROXY_TEXTURE_2D_MULTISAMPLE=  0x9101
const TEXTURE_2D_MULTISAMPLE_ARRAY=  0x9102
const PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY=0x9103
const TEXTURE_BINDING_2D_MULTISAMPLE=0x9104
const TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY=0x9105
const TEXTURE_SAMPLES=               0x9106
const TEXTURE_FIXED_SAMPLE_LOCATIONS=0x9107
const SAMPLER_2D_MULTISAMPLE=        0x9108
const INT_SAMPLER_2D_MULTISAMPLE=    0x9109
const UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE=0x910A
const SAMPLER_2D_MULTISAMPLE_ARRAY=  0x910B
const INT_SAMPLER_2D_MULTISAMPLE_ARRAY=0x910C
const UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY=0x910D
const MAX_COLOR_TEXTURE_SAMPLES=     0x910E
const MAX_DEPTH_TEXTURE_SAMPLES=     0x910F
const MAX_INTEGER_SAMPLES=           0x9110
const VERSION_3_3=1
const VERTEX_ATTRIB_ARRAY_DIVISOR=   0x88FE
const SRC1_COLOR=                    0x88F9
const ONE_MINUS_SRC1_COLOR=          0x88FA
const ONE_MINUS_SRC1_ALPHA=          0x88FB
const MAX_DUAL_SOURCE_DRAW_BUFFERS=  0x88FC
const ANY_SAMPLES_PASSED=            0x8C2F
const SAMPLER_BINDING=               0x8919
const RGB10_A2UI=                    0x906F
const TEXTURE_SWIZZLE_R=             0x8E42
const TEXTURE_SWIZZLE_G=             0x8E43
const TEXTURE_SWIZZLE_B=             0x8E44
const TEXTURE_SWIZZLE_A=             0x8E45
const TEXTURE_SWIZZLE_RGBA=          0x8E46
const TIME_ELAPSED=                  0x88BF
const TIMESTAMP=                     0x8E28
const INT_2_10_10_10_REV=            0x8D9F
const VERSION_4_0=1
const SAMPLE_SHADING=                0x8C36
const MIN_SAMPLE_SHADING_VALUE=      0x8C37
const MIN_PROGRAM_TEXTURE_GATHER_OFFSET=0x8E5E
const MAX_PROGRAM_TEXTURE_GATHER_OFFSET=0x8E5F
const TEXTURE_CUBE_MAP_ARRAY=        0x9009
const TEXTURE_BINDING_CUBE_MAP_ARRAY=0x900A
const PROXY_TEXTURE_CUBE_MAP_ARRAY=  0x900B
const SAMPLER_CUBE_MAP_ARRAY=        0x900C
const SAMPLER_CUBE_MAP_ARRAY_SHADOW= 0x900D
const INT_SAMPLER_CUBE_MAP_ARRAY=    0x900E
const UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY=0x900F
const DRAW_INDIRECT_BUFFER=          0x8F3F
const DRAW_INDIRECT_BUFFER_BINDING=  0x8F43
const GEOMETRY_SHADER_INVOCATIONS=   0x887F
const MAX_GEOMETRY_SHADER_INVOCATIONS=0x8E5A
const MIN_FRAGMENT_INTERPOLATION_OFFSET=0x8E5B
const MAX_FRAGMENT_INTERPOLATION_OFFSET=0x8E5C
const FRAGMENT_INTERPOLATION_OFFSET_BITS=0x8E5D
const MAX_VERTEX_STREAMS=            0x8E71
const DOUBLE_VEC2=                   0x8FFC
const DOUBLE_VEC3=                   0x8FFD
const DOUBLE_VEC4=                   0x8FFE
const DOUBLE_MAT2=                   0x8F46
const DOUBLE_MAT3=                   0x8F47
const DOUBLE_MAT4=                   0x8F48
const DOUBLE_MAT2x3=                 0x8F49
const DOUBLE_MAT2x4=                 0x8F4A
const DOUBLE_MAT3x2=                 0x8F4B
const DOUBLE_MAT3x4=                 0x8F4C
const DOUBLE_MAT4x2=                 0x8F4D
const DOUBLE_MAT4x3=                 0x8F4E
const ACTIVE_SUBROUTINES=            0x8DE5
const ACTIVE_SUBROUTINE_UNIFORMS=    0x8DE6
const ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS=0x8E47
const ACTIVE_SUBROUTINE_MAX_LENGTH=  0x8E48
const ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH=0x8E49
const MAX_SUBROUTINES=               0x8DE7
const MAX_SUBROUTINE_UNIFORM_LOCATIONS=0x8DE8
const NUM_COMPATIBLE_SUBROUTINES=    0x8E4A
const COMPATIBLE_SUBROUTINES=        0x8E4B
const PATCHES=                       0x000E
const PATCH_VERTICES=                0x8E72
const PATCH_DEFAULT_INNER_LEVEL=     0x8E73
const PATCH_DEFAULT_OUTER_LEVEL=     0x8E74
const TESS_CONTROL_OUTPUT_VERTICES=  0x8E75
const TESS_GEN_MODE=                 0x8E76
const TESS_GEN_SPACING=              0x8E77
const TESS_GEN_VERTEX_ORDER=         0x8E78
const TESS_GEN_POINT_MODE=           0x8E79
const ISOLINES=                      0x8E7A
const FRACTIONAL_ODD=                0x8E7B
const FRACTIONAL_EVEN=               0x8E7C
const MAX_PATCH_VERTICES=            0x8E7D
const MAX_TESS_GEN_LEVEL=            0x8E7E
const MAX_TESS_CONTROL_UNIFORM_COMPONENTS=0x8E7F
const MAX_TESS_EVALUATION_UNIFORM_COMPONENTS=0x8E80
const MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS=0x8E81
const MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS=0x8E82
const MAX_TESS_CONTROL_OUTPUT_COMPONENTS=0x8E83
const MAX_TESS_PATCH_COMPONENTS=     0x8E84
const MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS=0x8E85
const MAX_TESS_EVALUATION_OUTPUT_COMPONENTS=0x8E86
const MAX_TESS_CONTROL_UNIFORM_BLOCKS=0x8E89
const MAX_TESS_EVALUATION_UNIFORM_BLOCKS=0x8E8A
const MAX_TESS_CONTROL_INPUT_COMPONENTS=0x886C
const MAX_TESS_EVALUATION_INPUT_COMPONENTS=0x886D
const MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS=0x8E1E
const MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS=0x8E1F
const UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER=0x84F0
const UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER=0x84F1
const TESS_EVALUATION_SHADER=        0x8E87
const TESS_CONTROL_SHADER=           0x8E88
const TRANSFORM_FEEDBACK=            0x8E22
const TRANSFORM_FEEDBACK_BUFFER_PAUSED=0x8E23
const TRANSFORM_FEEDBACK_BUFFER_ACTIVE=0x8E24
const TRANSFORM_FEEDBACK_BINDING=    0x8E25
const MAX_TRANSFORM_FEEDBACK_BUFFERS=0x8E70
const VERSION_4_1=1
const FIXED=                         0x140C
const IMPLEMENTATION_COLOR_READ_TYPE=0x8B9A
const IMPLEMENTATION_COLOR_READ_FORMAT=0x8B9B
const LOW_FLOAT=                     0x8DF0
const MEDIUM_FLOAT=                  0x8DF1
const HIGH_FLOAT=                    0x8DF2
const LOW_INT=                       0x8DF3
const MEDIUM_INT=                    0x8DF4
const HIGH_INT=                      0x8DF5
const SHADER_COMPILER=               0x8DFA
const SHADER_BINARY_FORMATS=         0x8DF8
const NUM_SHADER_BINARY_FORMATS=     0x8DF9
const MAX_VERTEX_UNIFORM_VECTORS=    0x8DFB
const MAX_VARYING_VECTORS=           0x8DFC
const MAX_FRAGMENT_UNIFORM_VECTORS=  0x8DFD
const RGB565=                        0x8D62
const PROGRAM_BINARY_RETRIEVABLE_HINT=0x8257
const PROGRAM_BINARY_LENGTH=         0x8741
const NUM_PROGRAM_BINARY_FORMATS=    0x87FE
const PROGRAM_BINARY_FORMATS=        0x87FF
const VERTEX_SHADER_BIT=             0x00000001
const FRAGMENT_SHADER_BIT=           0x00000002
const GEOMETRY_SHADER_BIT=           0x00000004
const TESS_CONTROL_SHADER_BIT=       0x00000008
const TESS_EVALUATION_SHADER_BIT=    0x00000010
const ALL_SHADER_BITS=               0xFFFFFFFF
const PROGRAM_SEPARABLE=             0x8258
const ACTIVE_PROGRAM=                0x8259
const PROGRAM_PIPELINE_BINDING=      0x825A
const MAX_VIEWPORTS=                 0x825B
const VIEWPORT_SUBPIXEL_BITS=        0x825C
const VIEWPORT_BOUNDS_RANGE=         0x825D
const LAYER_PROVOKING_VERTEX=        0x825E
const VIEWPORT_INDEX_PROVOKING_VERTEX=0x825F
const UNDEFINED_VERTEX=              0x8260
const VERSION_4_2=1
const UNPACK_COMPRESSED_BLOCK_WIDTH= 0x9127
const UNPACK_COMPRESSED_BLOCK_HEIGHT=0x9128
const UNPACK_COMPRESSED_BLOCK_DEPTH= 0x9129
const UNPACK_COMPRESSED_BLOCK_SIZE=  0x912A
const PACK_COMPRESSED_BLOCK_WIDTH=   0x912B
const PACK_COMPRESSED_BLOCK_HEIGHT=  0x912C
const PACK_COMPRESSED_BLOCK_DEPTH=   0x912D
const PACK_COMPRESSED_BLOCK_SIZE=    0x912E
const NUM_SAMPLE_COUNTS=             0x9380
const MIN_MAP_BUFFER_ALIGNMENT=      0x90BC
const ATOMIC_COUNTER_BUFFER=         0x92C0
const ATOMIC_COUNTER_BUFFER_BINDING= 0x92C1
const ATOMIC_COUNTER_BUFFER_START=   0x92C2
const ATOMIC_COUNTER_BUFFER_SIZE=    0x92C3
const ATOMIC_COUNTER_BUFFER_DATA_SIZE=0x92C4
const ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS=0x92C5
const ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES=0x92C6
const ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER=0x92C7
const ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER=0x92C8
const ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER=0x92C9
const ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER=0x92CA
const ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER=0x92CB
const MAX_VERTEX_ATOMIC_COUNTER_BUFFERS=0x92CC
const MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS=0x92CD
const MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS=0x92CE
const MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS=0x92CF
const MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS=0x92D0
const MAX_COMBINED_ATOMIC_COUNTER_BUFFERS=0x92D1
const MAX_VERTEX_ATOMIC_COUNTERS=    0x92D2
const MAX_TESS_CONTROL_ATOMIC_COUNTERS=0x92D3
const MAX_TESS_EVALUATION_ATOMIC_COUNTERS=0x92D4
const MAX_GEOMETRY_ATOMIC_COUNTERS=  0x92D5
const MAX_FRAGMENT_ATOMIC_COUNTERS=  0x92D6
const MAX_COMBINED_ATOMIC_COUNTERS=  0x92D7
const MAX_ATOMIC_COUNTER_BUFFER_SIZE=0x92D8
const MAX_ATOMIC_COUNTER_BUFFER_BINDINGS=0x92DC
const ACTIVE_ATOMIC_COUNTER_BUFFERS= 0x92D9
const UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX=0x92DA
const UNSIGNED_INT_ATOMIC_COUNTER=   0x92DB
const VERTEX_ATTRIB_ARRAY_BARRIER_BIT=0x00000001
const ELEMENT_ARRAY_BARRIER_BIT=     0x00000002
const UNIFORM_BARRIER_BIT=           0x00000004
const TEXTURE_FETCH_BARRIER_BIT=     0x00000008
const SHADER_IMAGE_ACCESS_BARRIER_BIT=0x00000020
const COMMAND_BARRIER_BIT=           0x00000040
const PIXEL_BUFFER_BARRIER_BIT=      0x00000080
const TEXTURE_UPDATE_BARRIER_BIT=    0x00000100
const BUFFER_UPDATE_BARRIER_BIT=     0x00000200
const FRAMEBUFFER_BARRIER_BIT=       0x00000400
const TRANSFORM_FEEDBACK_BARRIER_BIT=0x00000800
const ATOMIC_COUNTER_BARRIER_BIT=    0x00001000
const ALL_BARRIER_BITS=              0xFFFFFFFF
const MAX_IMAGE_UNITS=               0x8F38
const MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS=0x8F39
const IMAGE_BINDING_NAME=            0x8F3A
const IMAGE_BINDING_LEVEL=           0x8F3B
const IMAGE_BINDING_LAYERED=         0x8F3C
const IMAGE_BINDING_LAYER=           0x8F3D
const IMAGE_BINDING_ACCESS=          0x8F3E
const IMAGE_1D=                      0x904C
const IMAGE_2D=                      0x904D
const IMAGE_3D=                      0x904E
const IMAGE_2D_RECT=                 0x904F
const IMAGE_CUBE=                    0x9050
const IMAGE_BUFFER=                  0x9051
const IMAGE_1D_ARRAY=                0x9052
const IMAGE_2D_ARRAY=                0x9053
const IMAGE_CUBE_MAP_ARRAY=          0x9054
const IMAGE_2D_MULTISAMPLE=          0x9055
const IMAGE_2D_MULTISAMPLE_ARRAY=    0x9056
const INT_IMAGE_1D=                  0x9057
const INT_IMAGE_2D=                  0x9058
const INT_IMAGE_3D=                  0x9059
const INT_IMAGE_2D_RECT=             0x905A
const INT_IMAGE_CUBE=                0x905B
const INT_IMAGE_BUFFER=              0x905C
const INT_IMAGE_1D_ARRAY=            0x905D
const INT_IMAGE_2D_ARRAY=            0x905E
const INT_IMAGE_CUBE_MAP_ARRAY=      0x905F
const INT_IMAGE_2D_MULTISAMPLE=      0x9060
const INT_IMAGE_2D_MULTISAMPLE_ARRAY=0x9061
const UNSIGNED_INT_IMAGE_1D=         0x9062
const UNSIGNED_INT_IMAGE_2D=         0x9063
const UNSIGNED_INT_IMAGE_3D=         0x9064
const UNSIGNED_INT_IMAGE_2D_RECT=    0x9065
const UNSIGNED_INT_IMAGE_CUBE=       0x9066
const UNSIGNED_INT_IMAGE_BUFFER=     0x9067
const UNSIGNED_INT_IMAGE_1D_ARRAY=   0x9068
const UNSIGNED_INT_IMAGE_2D_ARRAY=   0x9069
const UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY=0x906A
const UNSIGNED_INT_IMAGE_2D_MULTISAMPLE=0x906B
const UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY=0x906C
const MAX_IMAGE_SAMPLES=             0x906D
const IMAGE_BINDING_FORMAT=          0x906E
const IMAGE_FORMAT_COMPATIBILITY_TYPE=0x90C7
const IMAGE_FORMAT_COMPATIBILITY_BY_SIZE=0x90C8
const IMAGE_FORMAT_COMPATIBILITY_BY_CLASS=0x90C9
const MAX_VERTEX_IMAGE_UNIFORMS=     0x90CA
const MAX_TESS_CONTROL_IMAGE_UNIFORMS=0x90CB
const MAX_TESS_EVALUATION_IMAGE_UNIFORMS=0x90CC
const MAX_GEOMETRY_IMAGE_UNIFORMS=   0x90CD
const MAX_FRAGMENT_IMAGE_UNIFORMS=   0x90CE
const MAX_COMBINED_IMAGE_UNIFORMS=   0x90CF
const COMPRESSED_RGBA_BPTC_UNORM=    0x8E8C
const COMPRESSED_SRGB_ALPHA_BPTC_UNORM=0x8E8D
const COMPRESSED_RGB_BPTC_SIGNED_FLOAT=0x8E8E
const COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT=0x8E8F
const TEXTURE_IMMUTABLE_FORMAT=      0x912F
const VERSION_4_3=1
const NUM_SHADING_LANGUAGE_VERSIONS= 0x82E9
const VERTEX_ATTRIB_ARRAY_LONG=      0x874E
const COMPRESSED_RGB8_ETC2=          0x9274
const COMPRESSED_SRGB8_ETC2=         0x9275
const COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2=0x9276
const COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2=0x9277
const COMPRESSED_RGBA8_ETC2_EAC=     0x9278
const COMPRESSED_SRGB8_ALPHA8_ETC2_EAC=0x9279
const COMPRESSED_R11_EAC=            0x9270
const COMPRESSED_SIGNED_R11_EAC=     0x9271
const COMPRESSED_RG11_EAC=           0x9272
const COMPRESSED_SIGNED_RG11_EAC=    0x9273
const PRIMITIVE_RESTART_FIXED_INDEX= 0x8D69
const ANY_SAMPLES_PASSED_CONSERVATIVE=0x8D6A
const MAX_ELEMENT_INDEX=             0x8D6B
const COMPUTE_SHADER=                0x91B9
const MAX_COMPUTE_UNIFORM_BLOCKS=    0x91BB
const MAX_COMPUTE_TEXTURE_IMAGE_UNITS=0x91BC
const MAX_COMPUTE_IMAGE_UNIFORMS=    0x91BD
const MAX_COMPUTE_SHARED_MEMORY_SIZE=0x8262
const MAX_COMPUTE_UNIFORM_COMPONENTS=0x8263
const MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS=0x8264
const MAX_COMPUTE_ATOMIC_COUNTERS=   0x8265
const MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS=0x8266
const MAX_COMPUTE_WORK_GROUP_INVOCATIONS=0x90EB
const MAX_COMPUTE_WORK_GROUP_COUNT=  0x91BE
const MAX_COMPUTE_WORK_GROUP_SIZE=   0x91BF
const COMPUTE_WORK_GROUP_SIZE=       0x8267
const UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER=0x90EC
const ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER=0x90ED
const DISPATCH_INDIRECT_BUFFER=      0x90EE
const DISPATCH_INDIRECT_BUFFER_BINDING=0x90EF
const DEBUG_OUTPUT_SYNCHRONOUS=      0x8242
const DEBUG_NEXT_LOGGED_MESSAGE_LENGTH=0x8243
const DEBUG_CALLBACK_FUNCTION=       0x8244
const DEBUG_CALLBACK_USER_PARAM=     0x8245
const DEBUG_SOURCE_API=              0x8246
const DEBUG_SOURCE_WINDOW_SYSTEM=    0x8247
const DEBUG_SOURCE_SHADER_COMPILER=  0x8248
const DEBUG_SOURCE_THIRD_PARTY=      0x8249
const DEBUG_SOURCE_APPLICATION=      0x824A
const DEBUG_SOURCE_OTHER=            0x824B
const DEBUG_TYPE_ERROR=              0x824C
const DEBUG_TYPE_DEPRECATED_BEHAVIOR=0x824D
const DEBUG_TYPE_UNDEFINED_BEHAVIOR= 0x824E
const DEBUG_TYPE_PORTABILITY=        0x824F
const DEBUG_TYPE_PERFORMANCE=        0x8250
const DEBUG_TYPE_OTHER=              0x8251
const MAX_DEBUG_MESSAGE_LENGTH=      0x9143
const MAX_DEBUG_LOGGED_MESSAGES=     0x9144
const DEBUG_LOGGED_MESSAGES=         0x9145
const DEBUG_SEVERITY_HIGH=           0x9146
const DEBUG_SEVERITY_MEDIUM=         0x9147
const DEBUG_SEVERITY_LOW=            0x9148
const DEBUG_TYPE_MARKER=             0x8268
const DEBUG_TYPE_PUSH_GROUP=         0x8269
const DEBUG_TYPE_POP_GROUP=          0x826A
const DEBUG_SEVERITY_NOTIFICATION=   0x826B
const MAX_DEBUG_GROUP_STACK_DEPTH=   0x826C
const DEBUG_GROUP_STACK_DEPTH=       0x826D
const BUFFER=                        0x82E0
const SHADER=                        0x82E1
const PROGRAM=                       0x82E2
const QUERY=                         0x82E3
const PROGRAM_PIPELINE=              0x82E4
const SAMPLER=                       0x82E6
const MAX_LABEL_LENGTH=              0x82E8
const DEBUG_OUTPUT=                  0x92E0
const CONTEXT_FLAG_DEBUG_BIT=        0x00000002
const MAX_UNIFORM_LOCATIONS=         0x826E
const FRAMEBUFFER_DEFAULT_WIDTH=     0x9310
const FRAMEBUFFER_DEFAULT_HEIGHT=    0x9311
const FRAMEBUFFER_DEFAULT_LAYERS=    0x9312
const FRAMEBUFFER_DEFAULT_SAMPLES=   0x9313
const FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS=0x9314
const MAX_FRAMEBUFFER_WIDTH=         0x9315
const MAX_FRAMEBUFFER_HEIGHT=        0x9316
const MAX_FRAMEBUFFER_LAYERS=        0x9317
const MAX_FRAMEBUFFER_SAMPLES=       0x9318
const INTERNALFORMAT_SUPPORTED=      0x826F
const INTERNALFORMAT_PREFERRED=      0x8270
const INTERNALFORMAT_RED_SIZE=       0x8271
const INTERNALFORMAT_GREEN_SIZE=     0x8272
const INTERNALFORMAT_BLUE_SIZE=      0x8273
const INTERNALFORMAT_ALPHA_SIZE=     0x8274
const INTERNALFORMAT_DEPTH_SIZE=     0x8275
const INTERNALFORMAT_STENCIL_SIZE=   0x8276
const INTERNALFORMAT_SHARED_SIZE=    0x8277
const INTERNALFORMAT_RED_TYPE=       0x8278
const INTERNALFORMAT_GREEN_TYPE=     0x8279
const INTERNALFORMAT_BLUE_TYPE=      0x827A
const INTERNALFORMAT_ALPHA_TYPE=     0x827B
const INTERNALFORMAT_DEPTH_TYPE=     0x827C
const INTERNALFORMAT_STENCIL_TYPE=   0x827D
const MAX_WIDTH=                     0x827E
const MAX_HEIGHT=                    0x827F
const MAX_DEPTH=                     0x8280
const MAX_LAYERS=                    0x8281
const MAX_COMBINED_DIMENSIONS=       0x8282
const COLOR_COMPONENTS=              0x8283
const DEPTH_COMPONENTS=              0x8284
const STENCIL_COMPONENTS=            0x8285
const COLOR_RENDERABLE=              0x8286
const DEPTH_RENDERABLE=              0x8287
const STENCIL_RENDERABLE=            0x8288
const FRAMEBUFFER_RENDERABLE=        0x8289
const FRAMEBUFFER_RENDERABLE_LAYERED=0x828A
const FRAMEBUFFER_BLEND=             0x828B
const READ_PIXELS=                   0x828C
const READ_PIXELS_FORMAT=            0x828D
const READ_PIXELS_TYPE=              0x828E
const TEXTURE_IMAGE_FORMAT=          0x828F
const TEXTURE_IMAGE_TYPE=            0x8290
const GET_TEXTURE_IMAGE_FORMAT=      0x8291
const GET_TEXTURE_IMAGE_TYPE=        0x8292
const MIPMAP=                        0x8293
const MANUAL_GENERATE_MIPMAP=        0x8294
const AUTO_GENERATE_MIPMAP=          0x8295
const COLOR_ENCODING=                0x8296
const SRGB_READ=                     0x8297
const SRGB_WRITE=                    0x8298
const FILTER=                        0x829A
const VERTEX_TEXTURE=                0x829B
const TESS_CONTROL_TEXTURE=          0x829C
const TESS_EVALUATION_TEXTURE=       0x829D
const GEOMETRY_TEXTURE=              0x829E
const FRAGMENT_TEXTURE=              0x829F
const COMPUTE_TEXTURE=               0x82A0
const TEXTURE_SHADOW=                0x82A1
const TEXTURE_GATHER=                0x82A2
const TEXTURE_GATHER_SHADOW=         0x82A3
const SHADER_IMAGE_LOAD=             0x82A4
const SHADER_IMAGE_STORE=            0x82A5
const SHADER_IMAGE_ATOMIC=           0x82A6
const IMAGE_TEXEL_SIZE=              0x82A7
const IMAGE_COMPATIBILITY_CLASS=     0x82A8
const IMAGE_PIXEL_FORMAT=            0x82A9
const IMAGE_PIXEL_TYPE=              0x82AA
const SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST=0x82AC
const SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST=0x82AD
const SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE=0x82AE
const SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE=0x82AF
const TEXTURE_COMPRESSED_BLOCK_WIDTH=0x82B1
const TEXTURE_COMPRESSED_BLOCK_HEIGHT=0x82B2
const TEXTURE_COMPRESSED_BLOCK_SIZE= 0x82B3
const CLEAR_BUFFER=                  0x82B4
const TEXTURE_VIEW=                  0x82B5
const VIEW_COMPATIBILITY_CLASS=      0x82B6
const FULL_SUPPORT=                  0x82B7
const CAVEAT_SUPPORT=                0x82B8
const IMAGE_CLASS_4_X_32=            0x82B9
const IMAGE_CLASS_2_X_32=            0x82BA
const IMAGE_CLASS_1_X_32=            0x82BB
const IMAGE_CLASS_4_X_16=            0x82BC
const IMAGE_CLASS_2_X_16=            0x82BD
const IMAGE_CLASS_1_X_16=            0x82BE
const IMAGE_CLASS_4_X_8=             0x82BF
const IMAGE_CLASS_2_X_8=             0x82C0
const IMAGE_CLASS_1_X_8=             0x82C1
const IMAGE_CLASS_11_11_10=          0x82C2
const IMAGE_CLASS_10_10_10_2=        0x82C3
const VIEW_CLASS_128_BITS=           0x82C4
const VIEW_CLASS_96_BITS=            0x82C5
const VIEW_CLASS_64_BITS=            0x82C6
const VIEW_CLASS_48_BITS=            0x82C7
const VIEW_CLASS_32_BITS=            0x82C8
const VIEW_CLASS_24_BITS=            0x82C9
const VIEW_CLASS_16_BITS=            0x82CA
const VIEW_CLASS_8_BITS=             0x82CB
const VIEW_CLASS_S3TC_DXT1_RGB=      0x82CC
const VIEW_CLASS_S3TC_DXT1_RGBA=     0x82CD
const VIEW_CLASS_S3TC_DXT3_RGBA=     0x82CE
const VIEW_CLASS_S3TC_DXT5_RGBA=     0x82CF
const VIEW_CLASS_RGTC1_RED=          0x82D0
const VIEW_CLASS_RGTC2_RG=           0x82D1
const VIEW_CLASS_BPTC_UNORM=         0x82D2
const VIEW_CLASS_BPTC_FLOAT=         0x82D3
const UNIFORM=                       0x92E1
const UNIFORM_BLOCK=                 0x92E2
const PROGRAM_INPUT=                 0x92E3
const PROGRAM_OUTPUT=                0x92E4
const BUFFER_VARIABLE=               0x92E5
const SHADER_STORAGE_BLOCK=          0x92E6
const VERTEX_SUBROUTINE=             0x92E8
const TESS_CONTROL_SUBROUTINE=       0x92E9
const TESS_EVALUATION_SUBROUTINE=    0x92EA
const GEOMETRY_SUBROUTINE=           0x92EB
const FRAGMENT_SUBROUTINE=           0x92EC
const COMPUTE_SUBROUTINE=            0x92ED
const VERTEX_SUBROUTINE_UNIFORM=     0x92EE
const TESS_CONTROL_SUBROUTINE_UNIFORM=0x92EF
const TESS_EVALUATION_SUBROUTINE_UNIFORM=0x92F0
const GEOMETRY_SUBROUTINE_UNIFORM=   0x92F1
const FRAGMENT_SUBROUTINE_UNIFORM=   0x92F2
const COMPUTE_SUBROUTINE_UNIFORM=    0x92F3
const TRANSFORM_FEEDBACK_VARYING=    0x92F4
const ACTIVE_RESOURCES=              0x92F5
const MAX_NAME_LENGTH=               0x92F6
const MAX_NUM_ACTIVE_VARIABLES=      0x92F7
const MAX_NUM_COMPATIBLE_SUBROUTINES=0x92F8
const NAME_LENGTH=                   0x92F9
const TYPE=                          0x92FA
const ARRAY_SIZE=                    0x92FB
const OFFSET=                        0x92FC
const BLOCK_INDEX=                   0x92FD
const ARRAY_STRIDE=                  0x92FE
const MATRIX_STRIDE=                 0x92FF
const IS_ROW_MAJOR=                  0x9300
const ATOMIC_COUNTER_BUFFER_INDEX=   0x9301
const BUFFER_BINDING=                0x9302
const BUFFER_DATA_SIZE=              0x9303
const NUM_ACTIVE_VARIABLES=          0x9304
const ACTIVE_VARIABLES=              0x9305
const REFERENCED_BY_VERTEX_SHADER=   0x9306
const REFERENCED_BY_TESS_CONTROL_SHADER=0x9307
const REFERENCED_BY_TESS_EVALUATION_SHADER=0x9308
const REFERENCED_BY_GEOMETRY_SHADER= 0x9309
const REFERENCED_BY_FRAGMENT_SHADER= 0x930A
const REFERENCED_BY_COMPUTE_SHADER=  0x930B
const TOP_LEVEL_ARRAY_SIZE=          0x930C
const TOP_LEVEL_ARRAY_STRIDE=        0x930D
const LOCATION=                      0x930E
const LOCATION_INDEX=                0x930F
const IS_PER_PATCH=                  0x92E7
const SHADER_STORAGE_BUFFER=         0x90D2
const SHADER_STORAGE_BUFFER_BINDING= 0x90D3
const SHADER_STORAGE_BUFFER_START=   0x90D4
const SHADER_STORAGE_BUFFER_SIZE=    0x90D5
const MAX_VERTEX_SHADER_STORAGE_BLOCKS=0x90D6
const MAX_GEOMETRY_SHADER_STORAGE_BLOCKS=0x90D7
const MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS=0x90D8
const MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS=0x90D9
const MAX_FRAGMENT_SHADER_STORAGE_BLOCKS=0x90DA
const MAX_COMPUTE_SHADER_STORAGE_BLOCKS=0x90DB
const MAX_COMBINED_SHADER_STORAGE_BLOCKS=0x90DC
const MAX_SHADER_STORAGE_BUFFER_BINDINGS=0x90DD
const MAX_SHADER_STORAGE_BLOCK_SIZE= 0x90DE
const SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT=0x90DF
const SHADER_STORAGE_BARRIER_BIT=    0x00002000
const MAX_COMBINED_SHADER_OUTPUT_RESOURCES=0x8F39
const DEPTH_STENCIL_TEXTURE_MODE=    0x90EA
const TEXTURE_BUFFER_OFFSET=         0x919D
const TEXTURE_BUFFER_SIZE=           0x919E
const TEXTURE_BUFFER_OFFSET_ALIGNMENT=0x919F
const TEXTURE_VIEW_MIN_LEVEL=        0x82DB
const TEXTURE_VIEW_NUM_LEVELS=       0x82DC
const TEXTURE_VIEW_MIN_LAYER=        0x82DD
const TEXTURE_VIEW_NUM_LAYERS=       0x82DE
const TEXTURE_IMMUTABLE_LEVELS=      0x82DF
const VERTEX_ATTRIB_BINDING=         0x82D4
const VERTEX_ATTRIB_RELATIVE_OFFSET= 0x82D5
const VERTEX_BINDING_DIVISOR=        0x82D6
const VERTEX_BINDING_OFFSET=         0x82D7
const VERTEX_BINDING_STRIDE=         0x82D8
const MAX_VERTEX_ATTRIB_RELATIVE_OFFSET=0x82D9
const MAX_VERTEX_ATTRIB_BINDINGS=    0x82DA
const VERSION_4_4=1
const MAX_VERTEX_ATTRIB_STRIDE=      0x82E5
const PRIMITIVE_RESTART_FOR_PATCHES_SUPPORTED=0x8221
const TEXTURE_BUFFER_BINDING=        0x8C2A
const MAP_PERSISTENT_BIT=            0x0040
const MAP_COHERENT_BIT=              0x0080
const DYNAMIC_STORAGE_BIT=           0x0100
const CLIENT_STORAGE_BIT=            0x0200
const CLIENT_MAPPED_BUFFER_BARRIER_BIT=0x00004000
const BUFFER_IMMUTABLE_STORAGE=      0x821F
const BUFFER_STORAGE_FLAGS=          0x8220
const CLEAR_TEXTURE=                 0x9365
const LOCATION_COMPONENT=            0x934A
const TRANSFORM_FEEDBACK_BUFFER_INDEX=0x934B
const TRANSFORM_FEEDBACK_BUFFER_STRIDE=0x934C
const QUERY_BUFFER=                  0x9192
const QUERY_BUFFER_BARRIER_BIT=      0x00008000
const QUERY_BUFFER_BINDING=          0x9193
const QUERY_RESULT_NO_WAIT=          0x9194
const MIRROR_CLAMP_TO_EDGE=          0x8743
const ARB_ES2_compatibility=1
const ARB_ES3_compatibility=1
const ARB_arrays_of_arrays=1
const ARB_base_instance=1
const ARB_bindless_texture=1
const UNSIGNED_INT64_ARB=            0x140F
const ARB_blend_func_extended=1
const ARB_buffer_storage=1
const ARB_cl_event=1
const SYNC_CL_EVENT_ARB=             0x8240
const SYNC_CL_EVENT_COMPLETE_ARB=    0x8241
const ARB_clear_buffer_object=1
const ARB_clear_texture=1
const ARB_compressed_texture_pixel_storage=1
const ARB_compute_shader=1
const COMPUTE_SHADER_BIT=            0x00000020
const ARB_compute_variable_group_size=1
const MAX_COMPUTE_VARIABLE_GROUP_INVOCATIONS_ARB=0x9344
const MAX_COMPUTE_FIXED_GROUP_INVOCATIONS_ARB=0x90EB
const MAX_COMPUTE_VARIABLE_GROUP_SIZE_ARB=0x9345
const MAX_COMPUTE_FIXED_GROUP_SIZE_ARB=0x91BF
const ARB_conservative_depth=1
const ARB_copy_buffer=1
const COPY_READ_BUFFER_BINDING=      0x8F36
const COPY_WRITE_BUFFER_BINDING=     0x8F37
const ARB_copy_image=1
const ARB_debug_output=1
const DEBUG_OUTPUT_SYNCHRONOUS_ARB=  0x8242
const DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB=0x8243
const DEBUG_CALLBACK_FUNCTION_ARB=   0x8244
const DEBUG_CALLBACK_USER_PARAM_ARB= 0x8245
const DEBUG_SOURCE_API_ARB=          0x8246
const DEBUG_SOURCE_WINDOW_SYSTEM_ARB=0x8247
const DEBUG_SOURCE_SHADER_COMPILER_ARB=0x8248
const DEBUG_SOURCE_THIRD_PARTY_ARB=  0x8249
const DEBUG_SOURCE_APPLICATION_ARB=  0x824A
const DEBUG_SOURCE_OTHER_ARB=        0x824B
const DEBUG_TYPE_ERROR_ARB=          0x824C
const DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB=0x824D
const DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB=0x824E
const DEBUG_TYPE_PORTABILITY_ARB=    0x824F
const DEBUG_TYPE_PERFORMANCE_ARB=    0x8250
const DEBUG_TYPE_OTHER_ARB=          0x8251
const MAX_DEBUG_MESSAGE_LENGTH_ARB=  0x9143
const MAX_DEBUG_LOGGED_MESSAGES_ARB= 0x9144
const DEBUG_LOGGED_MESSAGES_ARB=     0x9145
const DEBUG_SEVERITY_HIGH_ARB=       0x9146
const DEBUG_SEVERITY_MEDIUM_ARB=     0x9147
const DEBUG_SEVERITY_LOW_ARB=        0x9148
const ARB_depth_buffer_float=1
const ARB_depth_clamp=1
const ARB_draw_buffers_blend=1
const ARB_draw_elements_base_vertex=1
const ARB_draw_indirect=1
const ARB_enhanced_layouts=1
const ARB_explicit_attrib_location=1
const ARB_explicit_uniform_location=1
const ARB_fragment_coord_conventions=1
const ARB_fragment_layer_viewport=1
const ARB_framebuffer_no_attachments=1
const ARB_framebuffer_object=1
const ARB_framebuffer_sRGB=1
const ARB_get_program_binary=1
const ARB_gpu_shader5=1
const ARB_gpu_shader_fp64=1
const ARB_half_float_vertex=1
const ARB_imaging=1
const BLEND_COLOR=                   0x8005
const BLEND_EQUATION=                0x8009
const ARB_indirect_parameters=1
const PARAMETER_BUFFER_ARB=          0x80EE
const PARAMETER_BUFFER_BINDING_ARB=  0x80EF
const ARB_internalformat_query=1
const ARB_internalformat_query2=1
const SRGB_DECODE_ARB=               0x8299
const ARB_invalidate_subdata=1
const ARB_map_buffer_alignment=1
const ARB_map_buffer_range=1
const ARB_multi_bind=1
const ARB_multi_draw_indirect=1
const ARB_occlusion_query2=1
const ARB_program_interface_query=1
const ARB_provoking_vertex=1
const ARB_query_buffer_object=1
const ARB_robust_buffer_access_behavior=1
const ARB_robustness=1
const CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB=0x00000004
const LOSE_CONTEXT_ON_RESET_ARB=     0x8252
const GUILTY_CONTEXT_RESET_ARB=      0x8253
const INNOCENT_CONTEXT_RESET_ARB=    0x8254
const UNKNOWN_CONTEXT_RESET_ARB=     0x8255
const RESET_NOTIFICATION_STRATEGY_ARB=0x8256
const NO_RESET_NOTIFICATION_ARB=     0x8261
const ARB_robustness_isolation=1
const ARB_sample_shading=1
const SAMPLE_SHADING_ARB=            0x8C36
const MIN_SAMPLE_SHADING_VALUE_ARB=  0x8C37
const ARB_sampler_objects=1
const ARB_seamless_cube_map=1
const ARB_seamless_cubemap_per_texture=1
const ARB_separate_shader_objects=1
const ARB_shader_atomic_counters=1
const ARB_shader_bit_encoding=1
const ARB_shader_draw_parameters=1
const ARB_shader_group_vote=1
const ARB_shader_image_load_store=1
const ARB_shader_image_size=1
const ARB_shader_precision=1
const ARB_shader_stencil_export=1
const ARB_shader_storage_buffer_object=1
const ARB_shader_subroutine=1
const ARB_shading_language_420pack=1
const ARB_shading_language_include=1
const SHADER_INCLUDE_ARB=            0x8DAE
const NAMED_STRING_LENGTH_ARB=       0x8DE9
const NAMED_STRING_TYPE_ARB=         0x8DEA
const ARB_shading_language_packing=1
const ARB_sparse_texture=1
const TEXTURE_SPARSE_ARB=            0x91A6
const VIRTUAL_PAGE_SIZE_INDEX_ARB=   0x91A7
const MIN_SPARSE_LEVEL_ARB=          0x919B
const NUM_VIRTUAL_PAGE_SIZES_ARB=    0x91A8
const VIRTUAL_PAGE_SIZE_X_ARB=       0x9195
const VIRTUAL_PAGE_SIZE_Y_ARB=       0x9196
const VIRTUAL_PAGE_SIZE_Z_ARB=       0x9197
const MAX_SPARSE_TEXTURE_SIZE_ARB=   0x9198
const MAX_SPARSE_3D_TEXTURE_SIZE_ARB=0x9199
const MAX_SPARSE_ARRAY_TEXTURE_LAYERS_ARB=0x919A
const SPARSE_TEXTURE_FULL_ARRAY_CUBE_MIPMAPS_ARB=0x91A9
const ARB_stencil_texturing=1
const ARB_sync=1
const ARB_tessellation_shader=1
const ARB_texture_buffer_object_rgb32=1
const ARB_texture_buffer_range=1
const ARB_texture_compression_bptc=1
const COMPRESSED_RGBA_BPTC_UNORM_ARB=0x8E8C
const COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB=0x8E8D
const COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB=0x8E8E
const COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB=0x8E8F
const ARB_texture_compression_rgtc=1
const ARB_texture_cube_map_array=1
const TEXTURE_CUBE_MAP_ARRAY_ARB=    0x9009
const TEXTURE_BINDING_CUBE_MAP_ARRAY_ARB=0x900A
const PROXY_TEXTURE_CUBE_MAP_ARRAY_ARB=0x900B
const SAMPLER_CUBE_MAP_ARRAY_ARB=    0x900C
const SAMPLER_CUBE_MAP_ARRAY_SHADOW_ARB=0x900D
const INT_SAMPLER_CUBE_MAP_ARRAY_ARB=0x900E
const UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY_ARB=0x900F
const ARB_texture_gather=1
const MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB=0x8E5E
const MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB=0x8E5F
const MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB=0x8F9F
const ARB_texture_mirror_clamp_to_edge=1
const ARB_texture_multisample=1
const ARB_texture_query_levels=1
const ARB_texture_query_lod=1
const ARB_texture_rg=1
const ARB_texture_rgb10_a2ui=1
const ARB_texture_stencil8=1
const ARB_texture_storage=1
const ARB_texture_storage_multisample=1
const ARB_texture_swizzle=1
const ARB_texture_view=1
const ARB_timer_query=1
const ARB_transform_feedback2=1
const TRANSFORM_FEEDBACK_PAUSED=     0x8E23
const TRANSFORM_FEEDBACK_ACTIVE=     0x8E24
const ARB_transform_feedback3=1
const ARB_transform_feedback_instanced=1
const ARB_uniform_buffer_object=1
const MAX_GEOMETRY_UNIFORM_BLOCKS=   0x8A2C
const MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS=0x8A32
const UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER=0x8A45
const ARB_vertex_array_bgra=1
const ARB_vertex_array_object=1
const ARB_vertex_attrib_64bit=1
const ARB_vertex_attrib_binding=1
const ARB_vertex_type_10f_11f_11f_rev=1
const ARB_vertex_type_2_10_10_10_rev=1
const ARB_viewport_array=1
const KHR_debug=1
const KHR_texture_compression_astc_hdr=1
const COMPRESSED_RGBA_ASTC_4x4_KHR=  0x93B0
const COMPRESSED_RGBA_ASTC_5x4_KHR=  0x93B1
const COMPRESSED_RGBA_ASTC_5x5_KHR=  0x93B2
const COMPRESSED_RGBA_ASTC_6x5_KHR=  0x93B3
const COMPRESSED_RGBA_ASTC_6x6_KHR=  0x93B4
const COMPRESSED_RGBA_ASTC_8x5_KHR=  0x93B5
const COMPRESSED_RGBA_ASTC_8x6_KHR=  0x93B6
const COMPRESSED_RGBA_ASTC_8x8_KHR=  0x93B7
const COMPRESSED_RGBA_ASTC_10x5_KHR= 0x93B8
const COMPRESSED_RGBA_ASTC_10x6_KHR= 0x93B9
const COMPRESSED_RGBA_ASTC_10x8_KHR= 0x93BA
const COMPRESSED_RGBA_ASTC_10x10_KHR=0x93BB
const COMPRESSED_RGBA_ASTC_12x10_KHR=0x93BC
const COMPRESSED_RGBA_ASTC_12x12_KHR=0x93BD
const COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR=0x93D0
const COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR=0x93D1
const COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR=0x93D2
const COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR=0x93D3
const COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR=0x93D4
const COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR=0x93D5
const COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR=0x93D6
const COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR=0x93D7
const COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR=0x93D8
const COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR=0x93D9
const COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR=0x93DA
const COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR=0x93DB
const COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR=0x93DC
const COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR=0x93DD
const KHR_texture_compression_astc_ldr=1
func CullFace(mode int){
	_mode := C.GLenum(mode)
	C.glCullFace(_mode)
}
func FrontFace(mode int){
	_mode := C.GLenum(mode)
	C.glFrontFace(_mode)
}
func Hint(target int, mode int){
	_target := C.GLenum(target)
	_mode := C.GLenum(mode)
	C.glHint(_target, _mode)
}
func LineWidth(width float32){
	_width := C.GLfloat(width)
	C.glLineWidth(_width)
}
func PointSize(size float32){
	_size := C.GLfloat(size)
	C.glPointSize(_size)
}
func PolygonMode(face int, mode int){
	_face := C.GLenum(face)
	_mode := C.GLenum(mode)
	C.glPolygonMode(_face, _mode)
}
func Scissor(x int, y int, width int, height int){
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	C.glScissor(_x, _y, _width, _height)
}
func TexParameterf(target int, pname int, param float32){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	_param := C.GLfloat(param)
	C.glTexParameterf(_target, _pname, _param)
}
/*func TexParameterfv(target int, pname int, params ???){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	panic()
	C.glTexParameterfv(_target, _pname, _params)
}*/
func TexParameteri(target int, pname int, param int){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	_param := C.GLint(param)
	C.glTexParameteri(_target, _pname, _param)
}
/*func TexParameteriv(target int, pname int, params ???){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	panic()
	C.glTexParameteriv(_target, _pname, _params)
}*/
/*func TexImage1D(target int, level int, internalformat int, width int, border int, format int, whichtype int, pixels ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_internalformat := C.GLint(internalformat)
	_width := C.GLsizei(width)
	_border := C.GLint(border)
	_format := C.GLenum(format)
	_whichtype := C.GLenum(whichtype)
	panic()
	C.glTexImage1D(_target, _level, _internalformat, _width, _border, _format, _whichtype, _pixels)
}*/
/*func TexImage2D(target int, level int, internalformat int, width int, height int, border int, format int, whichtype int, pixels ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_internalformat := C.GLint(internalformat)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_border := C.GLint(border)
	_format := C.GLenum(format)
	_whichtype := C.GLenum(whichtype)
	panic()
	C.glTexImage2D(_target, _level, _internalformat, _width, _height, _border, _format, _whichtype, _pixels)
}*/
func DrawBuffer(mode int){
	_mode := C.GLenum(mode)
	C.glDrawBuffer(_mode)
}
func Clear(mask uint){
	_mask := C.GLbitfield(mask)
	C.glClear(_mask)
}
func ClearColor(red float32, green float32, blue float32, alpha float32){
	_red := C.GLfloat(red)
	_green := C.GLfloat(green)
	_blue := C.GLfloat(blue)
	_alpha := C.GLfloat(alpha)
	C.glClearColor(_red, _green, _blue, _alpha)
}
func ClearStencil(s int){
	_s := C.GLint(s)
	C.glClearStencil(_s)
}
func ClearDepth(depth float64){
	_depth := C.GLdouble(depth)
	C.glClearDepth(_depth)
}
func StencilMask(mask uint){
	_mask := C.GLuint(mask)
	C.glStencilMask(_mask)
}
func ColorMask(red bool, green bool, blue bool, alpha bool){
	_red := C.GLboolean(1);if !red{_red=C.GLboolean(0)}
	_green := C.GLboolean(1);if !green{_green=C.GLboolean(0)}
	_blue := C.GLboolean(1);if !blue{_blue=C.GLboolean(0)}
	_alpha := C.GLboolean(1);if !alpha{_alpha=C.GLboolean(0)}
	C.glColorMask(_red, _green, _blue, _alpha)
}
func DepthMask(flag bool){
	_flag := C.GLboolean(1);if !flag{_flag=C.GLboolean(0)}
	C.glDepthMask(_flag)
}
func Disable(cap int){
	_cap := C.GLenum(cap)
	C.glDisable(_cap)
}
func Enable(cap int){
	_cap := C.GLenum(cap)
	C.glEnable(_cap)
}
func Finish(){
	C.glFinish()
}
func Flush(){
	C.glFlush()
}
func BlendFunc(sfactor int, dfactor int){
	_sfactor := C.GLenum(sfactor)
	_dfactor := C.GLenum(dfactor)
	C.glBlendFunc(_sfactor, _dfactor)
}
func LogicOp(opcode int){
	_opcode := C.GLenum(opcode)
	C.glLogicOp(_opcode)
}
func StencilFunc(function int, ref int, mask uint){
	_function := C.GLenum(function)
	_ref := C.GLint(ref)
	_mask := C.GLuint(mask)
	C.glStencilFunc(_function, _ref, _mask)
}
func StencilOp(fail int, zfail int, zpass int){
	_fail := C.GLenum(fail)
	_zfail := C.GLenum(zfail)
	_zpass := C.GLenum(zpass)
	C.glStencilOp(_fail, _zfail, _zpass)
}
func DepthFunc(function int){
	_function := C.GLenum(function)
	C.glDepthFunc(_function)
}
func PixelStoref(pname int, param float32){
	_pname := C.GLenum(pname)
	_param := C.GLfloat(param)
	C.glPixelStoref(_pname, _param)
}
func PixelStorei(pname int, param int){
	_pname := C.GLenum(pname)
	_param := C.GLint(param)
	C.glPixelStorei(_pname, _param)
}
func ReadBuffer(mode int){
	_mode := C.GLenum(mode)
	C.glReadBuffer(_mode)
}
/*func ReadPixels(x int, y int, width int, height int, format int, whichtype int, pixels ???){
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_format := C.GLenum(format)
	_whichtype := C.GLenum(whichtype)
	panic()
	C.glReadPixels(_x, _y, _width, _height, _format, _whichtype, _pixels)
}*/
/*func GetBooleanv(pname int, data ???){
	_pname := C.GLenum(pname)
	panic()
	C.glGetBooleanv(_pname, _data)
}*/
/*func GetDoublev(pname int, data ???){
	_pname := C.GLenum(pname)
	panic()
	C.glGetDoublev(_pname, _data)
}*/
func GetError()int{
	returnvalue := C.glGetError()
	convreturnvalue := int(returnvalue)
return convreturnvalue
}
/*func GetFloatv(pname int, data ???){
	_pname := C.GLenum(pname)
	panic()
	C.glGetFloatv(_pname, _data)
}*/
/*func GetIntegerv(pname int, data ???){
	_pname := C.GLenum(pname)
	panic()
	C.glGetIntegerv(_pname, _data)
}*/
/*func GetString(name int)???{
	_name := C.GLenum(name)
	returnvalue := C.glGetString(_name)
	panic()
return convreturnvalue
}*/
/*func GetTexImage(target int, level int, format int, whichtype int, pixels ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_format := C.GLenum(format)
	_whichtype := C.GLenum(whichtype)
	panic()
	C.glGetTexImage(_target, _level, _format, _whichtype, _pixels)
}*/
/*func GetTexParameterfv(target int, pname int, params ???){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	panic()
	C.glGetTexParameterfv(_target, _pname, _params)
}*/
/*func GetTexParameteriv(target int, pname int, params ???){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	panic()
	C.glGetTexParameteriv(_target, _pname, _params)
}*/
/*func GetTexLevelParameterfv(target int, level int, pname int, params ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_pname := C.GLenum(pname)
	panic()
	C.glGetTexLevelParameterfv(_target, _level, _pname, _params)
}*/
/*func GetTexLevelParameteriv(target int, level int, pname int, params ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_pname := C.GLenum(pname)
	panic()
	C.glGetTexLevelParameteriv(_target, _level, _pname, _params)
}*/
/*func IsEnabled(cap int)bool{
	_cap := C.GLenum(cap)
	returnvalue := C.glIsEnabled(_cap)
	panic()
return convreturnvalue
}*/
func DepthRange(near float64, far float64){
	_near := C.GLdouble(near)
	_far := C.GLdouble(far)
	C.glDepthRange(_near, _far)
}
func Viewport(x int, y int, width int, height int){
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	C.glViewport(_x, _y, _width, _height)
}
func DrawArrays(mode int, first int, count int){
	_mode := C.GLenum(mode)
	_first := C.GLint(first)
	_count := C.GLsizei(count)
	C.glDrawArrays(_mode, _first, _count)
}
/*func DrawElements(mode int, count int, whichtype int, indices ???){
	_mode := C.GLenum(mode)
	_count := C.GLsizei(count)
	_whichtype := C.GLenum(whichtype)
	panic()
	C.glDrawElements(_mode, _count, _whichtype, _indices)
}*/
/*func GetPointerv(pname int, params ???){
	_pname := C.GLenum(pname)
	panic()
	C.glGetPointerv(_pname, _params)
}*/
func PolygonOffset(factor float32, units float32){
	_factor := C.GLfloat(factor)
	_units := C.GLfloat(units)
	C.glPolygonOffset(_factor, _units)
}
func CopyTexImage1D(target int, level int, internalformat int, x int, y int, width int, border int){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_internalformat := C.GLenum(internalformat)
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	_border := C.GLint(border)
	C.glCopyTexImage1D(_target, _level, _internalformat, _x, _y, _width, _border)
}
func CopyTexImage2D(target int, level int, internalformat int, x int, y int, width int, height int, border int){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_internalformat := C.GLenum(internalformat)
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_border := C.GLint(border)
	C.glCopyTexImage2D(_target, _level, _internalformat, _x, _y, _width, _height, _border)
}
func CopyTexSubImage1D(target int, level int, xoffset int, x int, y int, width int){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_xoffset := C.GLint(xoffset)
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	C.glCopyTexSubImage1D(_target, _level, _xoffset, _x, _y, _width)
}
func CopyTexSubImage2D(target int, level int, xoffset int, yoffset int, x int, y int, width int, height int){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_xoffset := C.GLint(xoffset)
	_yoffset := C.GLint(yoffset)
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	C.glCopyTexSubImage2D(_target, _level, _xoffset, _yoffset, _x, _y, _width, _height)
}
/*func TexSubImage1D(target int, level int, xoffset int, width int, format int, whichtype int, pixels ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_xoffset := C.GLint(xoffset)
	_width := C.GLsizei(width)
	_format := C.GLenum(format)
	_whichtype := C.GLenum(whichtype)
	panic()
	C.glTexSubImage1D(_target, _level, _xoffset, _width, _format, _whichtype, _pixels)
}*/
/*func TexSubImage2D(target int, level int, xoffset int, yoffset int, width int, height int, format int, whichtype int, pixels ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_xoffset := C.GLint(xoffset)
	_yoffset := C.GLint(yoffset)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_format := C.GLenum(format)
	_whichtype := C.GLenum(whichtype)
	panic()
	C.glTexSubImage2D(_target, _level, _xoffset, _yoffset, _width, _height, _format, _whichtype, _pixels)
}*/
func BindTexture(target int, texture uint){
	_target := C.GLenum(target)
	_texture := C.GLuint(texture)
	C.glBindTexture(_target, _texture)
}
/*func DeleteTextures(n int, textures ???){
	_n := C.GLsizei(n)
	panic()
	C.glDeleteTextures(_n, _textures)
}*/
/*func GenTextures(n int, textures ???){
	_n := C.GLsizei(n)
	panic()
	C.glGenTextures(_n, _textures)
}*/
/*func IsTexture(texture uint)bool{
	_texture := C.GLuint(texture)
	returnvalue := C.glIsTexture(_texture)
	panic()
return convreturnvalue
}*/
/*func DrawRangeElements(mode int, start uint, end uint, count int, whichtype int, indices ???){
	_mode := C.GLenum(mode)
	_start := C.GLuint(start)
	_end := C.GLuint(end)
	_count := C.GLsizei(count)
	_whichtype := C.GLenum(whichtype)
	panic()
	C.glDrawRangeElements(_mode, _start, _end, _count, _whichtype, _indices)
}*/
/*func TexImage3D(target int, level int, internalformat int, width int, height int, depth int, border int, format int, whichtype int, pixels ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_internalformat := C.GLint(internalformat)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_depth := C.GLsizei(depth)
	_border := C.GLint(border)
	_format := C.GLenum(format)
	_whichtype := C.GLenum(whichtype)
	panic()
	C.glTexImage3D(_target, _level, _internalformat, _width, _height, _depth, _border, _format, _whichtype, _pixels)
}*/
/*func TexSubImage3D(target int, level int, xoffset int, yoffset int, zoffset int, width int, height int, depth int, format int, whichtype int, pixels ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_xoffset := C.GLint(xoffset)
	_yoffset := C.GLint(yoffset)
	_zoffset := C.GLint(zoffset)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_depth := C.GLsizei(depth)
	_format := C.GLenum(format)
	_whichtype := C.GLenum(whichtype)
	panic()
	C.glTexSubImage3D(_target, _level, _xoffset, _yoffset, _zoffset, _width, _height, _depth, _format, _whichtype, _pixels)
}*/
func CopyTexSubImage3D(target int, level int, xoffset int, yoffset int, zoffset int, x int, y int, width int, height int){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_xoffset := C.GLint(xoffset)
	_yoffset := C.GLint(yoffset)
	_zoffset := C.GLint(zoffset)
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	C.glCopyTexSubImage3D(_target, _level, _xoffset, _yoffset, _zoffset, _x, _y, _width, _height)
}
func ActiveTexture(texture int){
	_texture := C.GLenum(texture)
	C.glActiveTexture(_texture)
}
func SampleCoverage(value float32, invert bool){
	_value := C.GLfloat(value)
	_invert := C.GLboolean(1);if !invert{_invert=C.GLboolean(0)}
	C.glSampleCoverage(_value, _invert)
}
/*func CompressedTexImage3D(target int, level int, internalformat int, width int, height int, depth int, border int, imageSize int, data ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_internalformat := C.GLenum(internalformat)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_depth := C.GLsizei(depth)
	_border := C.GLint(border)
	_imageSize := C.GLsizei(imageSize)
	panic()
	C.glCompressedTexImage3D(_target, _level, _internalformat, _width, _height, _depth, _border, _imageSize, _data)
}*/
/*func CompressedTexImage2D(target int, level int, internalformat int, width int, height int, border int, imageSize int, data ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_internalformat := C.GLenum(internalformat)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_border := C.GLint(border)
	_imageSize := C.GLsizei(imageSize)
	panic()
	C.glCompressedTexImage2D(_target, _level, _internalformat, _width, _height, _border, _imageSize, _data)
}*/
/*func CompressedTexImage1D(target int, level int, internalformat int, width int, border int, imageSize int, data ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_internalformat := C.GLenum(internalformat)
	_width := C.GLsizei(width)
	_border := C.GLint(border)
	_imageSize := C.GLsizei(imageSize)
	panic()
	C.glCompressedTexImage1D(_target, _level, _internalformat, _width, _border, _imageSize, _data)
}*/
/*func CompressedTexSubImage3D(target int, level int, xoffset int, yoffset int, zoffset int, width int, height int, depth int, format int, imageSize int, data ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_xoffset := C.GLint(xoffset)
	_yoffset := C.GLint(yoffset)
	_zoffset := C.GLint(zoffset)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_depth := C.GLsizei(depth)
	_format := C.GLenum(format)
	_imageSize := C.GLsizei(imageSize)
	panic()
	C.glCompressedTexSubImage3D(_target, _level, _xoffset, _yoffset, _zoffset, _width, _height, _depth, _format, _imageSize, _data)
}*/
/*func CompressedTexSubImage2D(target int, level int, xoffset int, yoffset int, width int, height int, format int, imageSize int, data ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_xoffset := C.GLint(xoffset)
	_yoffset := C.GLint(yoffset)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_format := C.GLenum(format)
	_imageSize := C.GLsizei(imageSize)
	panic()
	C.glCompressedTexSubImage2D(_target, _level, _xoffset, _yoffset, _width, _height, _format, _imageSize, _data)
}*/
/*func CompressedTexSubImage1D(target int, level int, xoffset int, width int, format int, imageSize int, data ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_xoffset := C.GLint(xoffset)
	_width := C.GLsizei(width)
	_format := C.GLenum(format)
	_imageSize := C.GLsizei(imageSize)
	panic()
	C.glCompressedTexSubImage1D(_target, _level, _xoffset, _width, _format, _imageSize, _data)
}*/
/*func GetCompressedTexImage(target int, level int, img ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	panic()
	C.glGetCompressedTexImage(_target, _level, _img)
}*/
func BlendFuncSeparate(sfactorRGB int, dfactorRGB int, sfactorAlpha int, dfactorAlpha int){
	_sfactorRGB := C.GLenum(sfactorRGB)
	_dfactorRGB := C.GLenum(dfactorRGB)
	_sfactorAlpha := C.GLenum(sfactorAlpha)
	_dfactorAlpha := C.GLenum(dfactorAlpha)
	C.glBlendFuncSeparate(_sfactorRGB, _dfactorRGB, _sfactorAlpha, _dfactorAlpha)
}
/*func MultiDrawArrays(mode int, first ???, count ???, drawcount int){
	_mode := C.GLenum(mode)
	panic()
	panic()
	_drawcount := C.GLsizei(drawcount)
	C.glMultiDrawArrays(_mode, _first, _count, _drawcount)
}*/
/*func MultiDrawElements(mode int, count ???, whichtype int, indices ???, drawcount int){
	_mode := C.GLenum(mode)
	panic()
	_whichtype := C.GLenum(whichtype)
	panic()
	_drawcount := C.GLsizei(drawcount)
	C.glMultiDrawElements(_mode, _count, _whichtype, _indices, _drawcount)
}*/
func PointParameterf(pname int, param float32){
	_pname := C.GLenum(pname)
	_param := C.GLfloat(param)
	C.glPointParameterf(_pname, _param)
}
/*func PointParameterfv(pname int, params ???){
	_pname := C.GLenum(pname)
	panic()
	C.glPointParameterfv(_pname, _params)
}*/
func PointParameteri(pname int, param int){
	_pname := C.GLenum(pname)
	_param := C.GLint(param)
	C.glPointParameteri(_pname, _param)
}
/*func PointParameteriv(pname int, params ???){
	_pname := C.GLenum(pname)
	panic()
	C.glPointParameteriv(_pname, _params)
}*/
func BlendColor(red float32, green float32, blue float32, alpha float32){
	_red := C.GLfloat(red)
	_green := C.GLfloat(green)
	_blue := C.GLfloat(blue)
	_alpha := C.GLfloat(alpha)
	C.glBlendColor(_red, _green, _blue, _alpha)
}
func BlendEquation(mode int){
	_mode := C.GLenum(mode)
	C.glBlendEquation(_mode)
}
/*func GenQueries(n int, ids ???){
	_n := C.GLsizei(n)
	panic()
	C.glGenQueries(_n, _ids)
}*/
/*func DeleteQueries(n int, ids ???){
	_n := C.GLsizei(n)
	panic()
	C.glDeleteQueries(_n, _ids)
}*/
/*func IsQuery(id uint)bool{
	_id := C.GLuint(id)
	returnvalue := C.glIsQuery(_id)
	panic()
return convreturnvalue
}*/
func BeginQuery(target int, id uint){
	_target := C.GLenum(target)
	_id := C.GLuint(id)
	C.glBeginQuery(_target, _id)
}
func EndQuery(target int){
	_target := C.GLenum(target)
	C.glEndQuery(_target)
}
/*func GetQueryiv(target int, pname int, params ???){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	panic()
	C.glGetQueryiv(_target, _pname, _params)
}*/
/*func GetQueryObjectiv(id uint, pname int, params ???){
	_id := C.GLuint(id)
	_pname := C.GLenum(pname)
	panic()
	C.glGetQueryObjectiv(_id, _pname, _params)
}*/
/*func GetQueryObjectuiv(id uint, pname int, params ???){
	_id := C.GLuint(id)
	_pname := C.GLenum(pname)
	panic()
	C.glGetQueryObjectuiv(_id, _pname, _params)
}*/
func BindBuffer(target int, buffer uint){
	_target := C.GLenum(target)
	_buffer := C.GLuint(buffer)
	C.glBindBuffer(_target, _buffer)
}
/*func DeleteBuffers(n int, buffers ???){
	_n := C.GLsizei(n)
	panic()
	C.glDeleteBuffers(_n, _buffers)
}*/
/*func GenBuffers(n int, buffers ???){
	_n := C.GLsizei(n)
	panic()
	C.glGenBuffers(_n, _buffers)
}*/
/*func IsBuffer(buffer uint)bool{
	_buffer := C.GLuint(buffer)
	returnvalue := C.glIsBuffer(_buffer)
	panic()
return convreturnvalue
}*/
/*func BufferData(target int, size int, data ???, usage int){
	_target := C.GLenum(target)
	_size := C.GLsizeiptr(size)
	panic()
	_usage := C.GLenum(usage)
	C.glBufferData(_target, _size, _data, _usage)
}*/
/*func BufferSubData(target int, offset ???, size int, data ???){
	_target := C.GLenum(target)
	panic()
	_size := C.GLsizeiptr(size)
	panic()
	C.glBufferSubData(_target, _offset, _size, _data)
}*/
/*func GetBufferSubData(target int, offset ???, size int, data ???){
	_target := C.GLenum(target)
	panic()
	_size := C.GLsizeiptr(size)
	panic()
	C.glGetBufferSubData(_target, _offset, _size, _data)
}*/
/*func MapBuffer(target int, access int)???{
	_target := C.GLenum(target)
	_access := C.GLenum(access)
	returnvalue := C.glMapBuffer(_target, _access)
	panic()
return convreturnvalue
}*/
/*func UnmapBuffer(target int)bool{
	_target := C.GLenum(target)
	returnvalue := C.glUnmapBuffer(_target)
	panic()
return convreturnvalue
}*/
/*func GetBufferParameteriv(target int, pname int, params ???){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	panic()
	C.glGetBufferParameteriv(_target, _pname, _params)
}*/
/*func GetBufferPointerv(target int, pname int, params ???){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	panic()
	C.glGetBufferPointerv(_target, _pname, _params)
}*/
func BlendEquationSeparate(modeRGB int, modeAlpha int){
	_modeRGB := C.GLenum(modeRGB)
	_modeAlpha := C.GLenum(modeAlpha)
	C.glBlendEquationSeparate(_modeRGB, _modeAlpha)
}
/*func DrawBuffers(n int, bufs ???){
	_n := C.GLsizei(n)
	panic()
	C.glDrawBuffers(_n, _bufs)
}*/
func StencilOpSeparate(face int, sfail int, dpfail int, dppass int){
	_face := C.GLenum(face)
	_sfail := C.GLenum(sfail)
	_dpfail := C.GLenum(dpfail)
	_dppass := C.GLenum(dppass)
	C.glStencilOpSeparate(_face, _sfail, _dpfail, _dppass)
}
func StencilFuncSeparate(face int, function int, ref int, mask uint){
	_face := C.GLenum(face)
	_function := C.GLenum(function)
	_ref := C.GLint(ref)
	_mask := C.GLuint(mask)
	C.glStencilFuncSeparate(_face, _function, _ref, _mask)
}
func StencilMaskSeparate(face int, mask uint){
	_face := C.GLenum(face)
	_mask := C.GLuint(mask)
	C.glStencilMaskSeparate(_face, _mask)
}
func AttachShader(program uint, shader uint){
	_program := C.GLuint(program)
	_shader := C.GLuint(shader)
	C.glAttachShader(_program, _shader)
}
/*func BindAttribLocation(program uint, index uint, name ???){
	_program := C.GLuint(program)
	_index := C.GLuint(index)
	panic()
	C.glBindAttribLocation(_program, _index, _name)
}*/
func CompileShader(shader uint){
	_shader := C.GLuint(shader)
	C.glCompileShader(_shader)
}
/*func CreateProgram()uint{
	returnvalue := C.glCreateProgram()
	panic()
return convreturnvalue
}*/
/*func CreateShader(whichtype int)uint{
	_whichtype := C.GLenum(whichtype)
	returnvalue := C.glCreateShader(_whichtype)
	panic()
return convreturnvalue
}*/
func DeleteProgram(program uint){
	_program := C.GLuint(program)
	C.glDeleteProgram(_program)
}
func DeleteShader(shader uint){
	_shader := C.GLuint(shader)
	C.glDeleteShader(_shader)
}
func DetachShader(program uint, shader uint){
	_program := C.GLuint(program)
	_shader := C.GLuint(shader)
	C.glDetachShader(_program, _shader)
}
func DisableVertexAttribArray(index uint){
	_index := C.GLuint(index)
	C.glDisableVertexAttribArray(_index)
}
func EnableVertexAttribArray(index uint){
	_index := C.GLuint(index)
	C.glEnableVertexAttribArray(_index)
}
/*func GetActiveAttrib(program uint, index uint, bufSize int, length int, size ???, whichtype ???, name int){
	_program := C.GLuint(program)
	_index := C.GLuint(index)
	_bufSize := C.GLsizei(bufSize)
	_length := C.GLsizei(length)
	panic()
	panic()
	panic()
	C.glGetActiveAttrib(_program, _index, _bufSize, _length, _size, _whichtype, _name)
}*/
/*func GetActiveUniform(program uint, index uint, bufSize int, length int, size ???, whichtype ???, name int){
	_program := C.GLuint(program)
	_index := C.GLuint(index)
	_bufSize := C.GLsizei(bufSize)
	_length := C.GLsizei(length)
	panic()
	panic()
	panic()
	C.glGetActiveUniform(_program, _index, _bufSize, _length, _size, _whichtype, _name)
}*/
/*func GetAttachedShaders(program uint, maxCount int, count int, shaders ???){
	_program := C.GLuint(program)
	_maxCount := C.GLsizei(maxCount)
	_count := C.GLsizei(count)
	panic()
	C.glGetAttachedShaders(_program, _maxCount, _count, _shaders)
}*/
/*func GetAttribLocation(program uint, name ???)int{
	_program := C.GLuint(program)
	panic()
	returnvalue := C.glGetAttribLocation(_program, _name)
	panic()
return convreturnvalue
}*/
/*func GetProgramiv(program uint, pname int, params ???){
	_program := C.GLuint(program)
	_pname := C.GLenum(pname)
	panic()
	C.glGetProgramiv(_program, _pname, _params)
}*/
/*func GetProgramInfoLog(program uint, bufSize int, length int, infoLog int){
	_program := C.GLuint(program)
	_bufSize := C.GLsizei(bufSize)
	_length := C.GLsizei(length)
	panic()
	C.glGetProgramInfoLog(_program, _bufSize, _length, _infoLog)
}*/
/*func GetShaderiv(shader uint, pname int, params ???){
	_shader := C.GLuint(shader)
	_pname := C.GLenum(pname)
	panic()
	C.glGetShaderiv(_shader, _pname, _params)
}*/
/*func GetShaderInfoLog(shader uint, bufSize int, length int, infoLog int){
	_shader := C.GLuint(shader)
	_bufSize := C.GLsizei(bufSize)
	_length := C.GLsizei(length)
	panic()
	C.glGetShaderInfoLog(_shader, _bufSize, _length, _infoLog)
}*/
/*func GetShaderSource(shader uint, bufSize int, length int, source int){
	_shader := C.GLuint(shader)
	_bufSize := C.GLsizei(bufSize)
	_length := C.GLsizei(length)
	panic()
	C.glGetShaderSource(_shader, _bufSize, _length, _source)
}*/
/*func GetUniformLocation(program uint, name ???)int{
	_program := C.GLuint(program)
	panic()
	returnvalue := C.glGetUniformLocation(_program, _name)
	panic()
return convreturnvalue
}*/
/*func GetUniformfv(program uint, location int, params ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	panic()
	C.glGetUniformfv(_program, _location, _params)
}*/
/*func GetUniformiv(program uint, location int, params ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	panic()
	C.glGetUniformiv(_program, _location, _params)
}*/
/*func GetVertexAttribdv(index uint, pname int, params ???){
	_index := C.GLuint(index)
	_pname := C.GLenum(pname)
	panic()
	C.glGetVertexAttribdv(_index, _pname, _params)
}*/
/*func GetVertexAttribfv(index uint, pname int, params ???){
	_index := C.GLuint(index)
	_pname := C.GLenum(pname)
	panic()
	C.glGetVertexAttribfv(_index, _pname, _params)
}*/
/*func GetVertexAttribiv(index uint, pname int, params ???){
	_index := C.GLuint(index)
	_pname := C.GLenum(pname)
	panic()
	C.glGetVertexAttribiv(_index, _pname, _params)
}*/
/*func GetVertexAttribPointerv(index uint, pname int, pointer ???){
	_index := C.GLuint(index)
	_pname := C.GLenum(pname)
	panic()
	C.glGetVertexAttribPointerv(_index, _pname, _pointer)
}*/
/*func IsProgram(program uint)bool{
	_program := C.GLuint(program)
	returnvalue := C.glIsProgram(_program)
	panic()
return convreturnvalue
}*/
/*func IsShader(shader uint)bool{
	_shader := C.GLuint(shader)
	returnvalue := C.glIsShader(_shader)
	panic()
return convreturnvalue
}*/
func LinkProgram(program uint){
	_program := C.GLuint(program)
	C.glLinkProgram(_program)
}
/*func ShaderSource(shader uint, count int, string ???, length ???){
	_shader := C.GLuint(shader)
	_count := C.GLsizei(count)
	panic()
	panic()
	C.glShaderSource(_shader, _count, _string, _length)
}*/
func UseProgram(program uint){
	_program := C.GLuint(program)
	C.glUseProgram(_program)
}
func Uniform1f(location int, v0 float32){
	_location := C.GLint(location)
	_v0 := C.GLfloat(v0)
	C.glUniform1f(_location, _v0)
}
func Uniform2f(location int, v0 float32, v1 float32){
	_location := C.GLint(location)
	_v0 := C.GLfloat(v0)
	_v1 := C.GLfloat(v1)
	C.glUniform2f(_location, _v0, _v1)
}
func Uniform3f(location int, v0 float32, v1 float32, v2 float32){
	_location := C.GLint(location)
	_v0 := C.GLfloat(v0)
	_v1 := C.GLfloat(v1)
	_v2 := C.GLfloat(v2)
	C.glUniform3f(_location, _v0, _v1, _v2)
}
func Uniform4f(location int, v0 float32, v1 float32, v2 float32, v3 float32){
	_location := C.GLint(location)
	_v0 := C.GLfloat(v0)
	_v1 := C.GLfloat(v1)
	_v2 := C.GLfloat(v2)
	_v3 := C.GLfloat(v3)
	C.glUniform4f(_location, _v0, _v1, _v2, _v3)
}
func Uniform1i(location int, v0 int){
	_location := C.GLint(location)
	_v0 := C.GLint(v0)
	C.glUniform1i(_location, _v0)
}
func Uniform2i(location int, v0 int, v1 int){
	_location := C.GLint(location)
	_v0 := C.GLint(v0)
	_v1 := C.GLint(v1)
	C.glUniform2i(_location, _v0, _v1)
}
func Uniform3i(location int, v0 int, v1 int, v2 int){
	_location := C.GLint(location)
	_v0 := C.GLint(v0)
	_v1 := C.GLint(v1)
	_v2 := C.GLint(v2)
	C.glUniform3i(_location, _v0, _v1, _v2)
}
func Uniform4i(location int, v0 int, v1 int, v2 int, v3 int){
	_location := C.GLint(location)
	_v0 := C.GLint(v0)
	_v1 := C.GLint(v1)
	_v2 := C.GLint(v2)
	_v3 := C.GLint(v3)
	C.glUniform4i(_location, _v0, _v1, _v2, _v3)
}
/*func Uniform1fv(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniform1fv(_location, _count, _value)
}*/
/*func Uniform2fv(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniform2fv(_location, _count, _value)
}*/
/*func Uniform3fv(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniform3fv(_location, _count, _value)
}*/
/*func Uniform4fv(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniform4fv(_location, _count, _value)
}*/
/*func Uniform1iv(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniform1iv(_location, _count, _value)
}*/
/*func Uniform2iv(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniform2iv(_location, _count, _value)
}*/
/*func Uniform3iv(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniform3iv(_location, _count, _value)
}*/
/*func Uniform4iv(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniform4iv(_location, _count, _value)
}*/
/*func UniformMatrix2fv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix2fv(_location, _count, _transpose, _value)
}*/
/*func UniformMatrix3fv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix3fv(_location, _count, _transpose, _value)
}*/
/*func UniformMatrix4fv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix4fv(_location, _count, _transpose, _value)
}*/
func ValidateProgram(program uint){
	_program := C.GLuint(program)
	C.glValidateProgram(_program)
}
func VertexAttrib1d(index uint, x float64){
	_index := C.GLuint(index)
	_x := C.GLdouble(x)
	C.glVertexAttrib1d(_index, _x)
}
/*func VertexAttrib1dv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib1dv(_index, _v)
}*/
func VertexAttrib1f(index uint, x float32){
	_index := C.GLuint(index)
	_x := C.GLfloat(x)
	C.glVertexAttrib1f(_index, _x)
}
/*func VertexAttrib1fv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib1fv(_index, _v)
}*/
/*func VertexAttrib1s(index uint, x ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib1s(_index, _x)
}*/
/*func VertexAttrib1sv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib1sv(_index, _v)
}*/
func VertexAttrib2d(index uint, x float64, y float64){
	_index := C.GLuint(index)
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	C.glVertexAttrib2d(_index, _x, _y)
}
/*func VertexAttrib2dv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib2dv(_index, _v)
}*/
func VertexAttrib2f(index uint, x float32, y float32){
	_index := C.GLuint(index)
	_x := C.GLfloat(x)
	_y := C.GLfloat(y)
	C.glVertexAttrib2f(_index, _x, _y)
}
/*func VertexAttrib2fv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib2fv(_index, _v)
}*/
/*func VertexAttrib2s(index uint, x ???, y ???){
	_index := C.GLuint(index)
	panic()
	panic()
	C.glVertexAttrib2s(_index, _x, _y)
}*/
/*func VertexAttrib2sv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib2sv(_index, _v)
}*/
func VertexAttrib3d(index uint, x float64, y float64, z float64){
	_index := C.GLuint(index)
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	_z := C.GLdouble(z)
	C.glVertexAttrib3d(_index, _x, _y, _z)
}
/*func VertexAttrib3dv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib3dv(_index, _v)
}*/
func VertexAttrib3f(index uint, x float32, y float32, z float32){
	_index := C.GLuint(index)
	_x := C.GLfloat(x)
	_y := C.GLfloat(y)
	_z := C.GLfloat(z)
	C.glVertexAttrib3f(_index, _x, _y, _z)
}
/*func VertexAttrib3fv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib3fv(_index, _v)
}*/
/*func VertexAttrib3s(index uint, x ???, y ???, z ???){
	_index := C.GLuint(index)
	panic()
	panic()
	panic()
	C.glVertexAttrib3s(_index, _x, _y, _z)
}*/
/*func VertexAttrib3sv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib3sv(_index, _v)
}*/
/*func VertexAttrib4Nbv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib4Nbv(_index, _v)
}*/
/*func VertexAttrib4Niv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib4Niv(_index, _v)
}*/
/*func VertexAttrib4Nsv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib4Nsv(_index, _v)
}*/
/*func VertexAttrib4Nub(index uint, x ???, y ???, z ???, w ???){
	_index := C.GLuint(index)
	panic()
	panic()
	panic()
	panic()
	C.glVertexAttrib4Nub(_index, _x, _y, _z, _w)
}*/
/*func VertexAttrib4Nubv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib4Nubv(_index, _v)
}*/
/*func VertexAttrib4Nuiv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib4Nuiv(_index, _v)
}*/
/*func VertexAttrib4Nusv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib4Nusv(_index, _v)
}*/
/*func VertexAttrib4bv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib4bv(_index, _v)
}*/
func VertexAttrib4d(index uint, x float64, y float64, z float64, w float64){
	_index := C.GLuint(index)
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	_z := C.GLdouble(z)
	_w := C.GLdouble(w)
	C.glVertexAttrib4d(_index, _x, _y, _z, _w)
}
/*func VertexAttrib4dv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib4dv(_index, _v)
}*/
func VertexAttrib4f(index uint, x float32, y float32, z float32, w float32){
	_index := C.GLuint(index)
	_x := C.GLfloat(x)
	_y := C.GLfloat(y)
	_z := C.GLfloat(z)
	_w := C.GLfloat(w)
	C.glVertexAttrib4f(_index, _x, _y, _z, _w)
}
/*func VertexAttrib4fv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib4fv(_index, _v)
}*/
/*func VertexAttrib4iv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib4iv(_index, _v)
}*/
/*func VertexAttrib4s(index uint, x ???, y ???, z ???, w ???){
	_index := C.GLuint(index)
	panic()
	panic()
	panic()
	panic()
	C.glVertexAttrib4s(_index, _x, _y, _z, _w)
}*/
/*func VertexAttrib4sv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib4sv(_index, _v)
}*/
/*func VertexAttrib4ubv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib4ubv(_index, _v)
}*/
/*func VertexAttrib4uiv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib4uiv(_index, _v)
}*/
/*func VertexAttrib4usv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttrib4usv(_index, _v)
}*/
/*func VertexAttribPointer(index uint, size int, whichtype int, normalized bool, stride int, pointer ???){
	_index := C.GLuint(index)
	_size := C.GLint(size)
	_whichtype := C.GLenum(whichtype)
	_normalized := C.GLboolean(1);if !normalized{_normalized=C.GLboolean(0)}
	_stride := C.GLsizei(stride)
	panic()
	C.glVertexAttribPointer(_index, _size, _whichtype, _normalized, _stride, _pointer)
}*/
/*func UniformMatrix2x3fv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix2x3fv(_location, _count, _transpose, _value)
}*/
/*func UniformMatrix3x2fv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix3x2fv(_location, _count, _transpose, _value)
}*/
/*func UniformMatrix2x4fv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix2x4fv(_location, _count, _transpose, _value)
}*/
/*func UniformMatrix4x2fv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix4x2fv(_location, _count, _transpose, _value)
}*/
/*func UniformMatrix3x4fv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix3x4fv(_location, _count, _transpose, _value)
}*/
/*func UniformMatrix4x3fv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix4x3fv(_location, _count, _transpose, _value)
}*/
func ColorMaski(index uint, r bool, g bool, b bool, a bool){
	_index := C.GLuint(index)
	_r := C.GLboolean(1);if !r{_r=C.GLboolean(0)}
	_g := C.GLboolean(1);if !g{_g=C.GLboolean(0)}
	_b := C.GLboolean(1);if !b{_b=C.GLboolean(0)}
	_a := C.GLboolean(1);if !a{_a=C.GLboolean(0)}
	C.glColorMaski(_index, _r, _g, _b, _a)
}
/*func GetBooleani_v(target int, index uint, data ???){
	_target := C.GLenum(target)
	_index := C.GLuint(index)
	panic()
	C.glGetBooleani_v(_target, _index, _data)
}*/
/*func GetIntegeri_v(target int, index uint, data ???){
	_target := C.GLenum(target)
	_index := C.GLuint(index)
	panic()
	C.glGetIntegeri_v(_target, _index, _data)
}*/
func Enablei(target int, index uint){
	_target := C.GLenum(target)
	_index := C.GLuint(index)
	C.glEnablei(_target, _index)
}
func Disablei(target int, index uint){
	_target := C.GLenum(target)
	_index := C.GLuint(index)
	C.glDisablei(_target, _index)
}
/*func IsEnabledi(target int, index uint)bool{
	_target := C.GLenum(target)
	_index := C.GLuint(index)
	returnvalue := C.glIsEnabledi(_target, _index)
	panic()
return convreturnvalue
}*/
func BeginTransformFeedback(primitiveMode int){
	_primitiveMode := C.GLenum(primitiveMode)
	C.glBeginTransformFeedback(_primitiveMode)
}
func EndTransformFeedback(){
	C.glEndTransformFeedback()
}
/*func BindBufferRange(target int, index uint, buffer uint, offset ???, size int){
	_target := C.GLenum(target)
	_index := C.GLuint(index)
	_buffer := C.GLuint(buffer)
	panic()
	_size := C.GLsizeiptr(size)
	C.glBindBufferRange(_target, _index, _buffer, _offset, _size)
}*/
func BindBufferBase(target int, index uint, buffer uint){
	_target := C.GLenum(target)
	_index := C.GLuint(index)
	_buffer := C.GLuint(buffer)
	C.glBindBufferBase(_target, _index, _buffer)
}
/*func TransformFeedbackVaryings(program uint, count int, varyings ???, bufferMode int){
	_program := C.GLuint(program)
	_count := C.GLsizei(count)
	panic()
	_bufferMode := C.GLenum(bufferMode)
	C.glTransformFeedbackVaryings(_program, _count, _varyings, _bufferMode)
}*/
/*func GetTransformFeedbackVarying(program uint, index uint, bufSize int, length int, size int, whichtype ???, name int){
	_program := C.GLuint(program)
	_index := C.GLuint(index)
	_bufSize := C.GLsizei(bufSize)
	_length := C.GLsizei(length)
	_size := C.GLsizei(size)
	panic()
	panic()
	C.glGetTransformFeedbackVarying(_program, _index, _bufSize, _length, _size, _whichtype, _name)
}*/
func ClampColor(target int, clamp int){
	_target := C.GLenum(target)
	_clamp := C.GLenum(clamp)
	C.glClampColor(_target, _clamp)
}
func BeginConditionalRender(id uint, mode int){
	_id := C.GLuint(id)
	_mode := C.GLenum(mode)
	C.glBeginConditionalRender(_id, _mode)
}
func EndConditionalRender(){
	C.glEndConditionalRender()
}
/*func VertexAttribIPointer(index uint, size int, whichtype int, stride int, pointer ???){
	_index := C.GLuint(index)
	_size := C.GLint(size)
	_whichtype := C.GLenum(whichtype)
	_stride := C.GLsizei(stride)
	panic()
	C.glVertexAttribIPointer(_index, _size, _whichtype, _stride, _pointer)
}*/
/*func GetVertexAttribIiv(index uint, pname int, params ???){
	_index := C.GLuint(index)
	_pname := C.GLenum(pname)
	panic()
	C.glGetVertexAttribIiv(_index, _pname, _params)
}*/
/*func GetVertexAttribIuiv(index uint, pname int, params ???){
	_index := C.GLuint(index)
	_pname := C.GLenum(pname)
	panic()
	C.glGetVertexAttribIuiv(_index, _pname, _params)
}*/
func VertexAttribI1i(index uint, x int){
	_index := C.GLuint(index)
	_x := C.GLint(x)
	C.glVertexAttribI1i(_index, _x)
}
func VertexAttribI2i(index uint, x int, y int){
	_index := C.GLuint(index)
	_x := C.GLint(x)
	_y := C.GLint(y)
	C.glVertexAttribI2i(_index, _x, _y)
}
func VertexAttribI3i(index uint, x int, y int, z int){
	_index := C.GLuint(index)
	_x := C.GLint(x)
	_y := C.GLint(y)
	_z := C.GLint(z)
	C.glVertexAttribI3i(_index, _x, _y, _z)
}
func VertexAttribI4i(index uint, x int, y int, z int, w int){
	_index := C.GLuint(index)
	_x := C.GLint(x)
	_y := C.GLint(y)
	_z := C.GLint(z)
	_w := C.GLint(w)
	C.glVertexAttribI4i(_index, _x, _y, _z, _w)
}
func VertexAttribI1ui(index uint, x uint){
	_index := C.GLuint(index)
	_x := C.GLuint(x)
	C.glVertexAttribI1ui(_index, _x)
}
func VertexAttribI2ui(index uint, x uint, y uint){
	_index := C.GLuint(index)
	_x := C.GLuint(x)
	_y := C.GLuint(y)
	C.glVertexAttribI2ui(_index, _x, _y)
}
func VertexAttribI3ui(index uint, x uint, y uint, z uint){
	_index := C.GLuint(index)
	_x := C.GLuint(x)
	_y := C.GLuint(y)
	_z := C.GLuint(z)
	C.glVertexAttribI3ui(_index, _x, _y, _z)
}
func VertexAttribI4ui(index uint, x uint, y uint, z uint, w uint){
	_index := C.GLuint(index)
	_x := C.GLuint(x)
	_y := C.GLuint(y)
	_z := C.GLuint(z)
	_w := C.GLuint(w)
	C.glVertexAttribI4ui(_index, _x, _y, _z, _w)
}
/*func VertexAttribI1iv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribI1iv(_index, _v)
}*/
/*func VertexAttribI2iv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribI2iv(_index, _v)
}*/
/*func VertexAttribI3iv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribI3iv(_index, _v)
}*/
/*func VertexAttribI4iv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribI4iv(_index, _v)
}*/
/*func VertexAttribI1uiv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribI1uiv(_index, _v)
}*/
/*func VertexAttribI2uiv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribI2uiv(_index, _v)
}*/
/*func VertexAttribI3uiv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribI3uiv(_index, _v)
}*/
/*func VertexAttribI4uiv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribI4uiv(_index, _v)
}*/
/*func VertexAttribI4bv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribI4bv(_index, _v)
}*/
/*func VertexAttribI4sv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribI4sv(_index, _v)
}*/
/*func VertexAttribI4ubv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribI4ubv(_index, _v)
}*/
/*func VertexAttribI4usv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribI4usv(_index, _v)
}*/
/*func GetUniformuiv(program uint, location int, params ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	panic()
	C.glGetUniformuiv(_program, _location, _params)
}*/
/*func BindFragDataLocation(program uint, color uint, name ???){
	_program := C.GLuint(program)
	_color := C.GLuint(color)
	panic()
	C.glBindFragDataLocation(_program, _color, _name)
}*/
/*func GetFragDataLocation(program uint, name ???)int{
	_program := C.GLuint(program)
	panic()
	returnvalue := C.glGetFragDataLocation(_program, _name)
	panic()
return convreturnvalue
}*/
func Uniform1ui(location int, v0 uint){
	_location := C.GLint(location)
	_v0 := C.GLuint(v0)
	C.glUniform1ui(_location, _v0)
}
func Uniform2ui(location int, v0 uint, v1 uint){
	_location := C.GLint(location)
	_v0 := C.GLuint(v0)
	_v1 := C.GLuint(v1)
	C.glUniform2ui(_location, _v0, _v1)
}
func Uniform3ui(location int, v0 uint, v1 uint, v2 uint){
	_location := C.GLint(location)
	_v0 := C.GLuint(v0)
	_v1 := C.GLuint(v1)
	_v2 := C.GLuint(v2)
	C.glUniform3ui(_location, _v0, _v1, _v2)
}
func Uniform4ui(location int, v0 uint, v1 uint, v2 uint, v3 uint){
	_location := C.GLint(location)
	_v0 := C.GLuint(v0)
	_v1 := C.GLuint(v1)
	_v2 := C.GLuint(v2)
	_v3 := C.GLuint(v3)
	C.glUniform4ui(_location, _v0, _v1, _v2, _v3)
}
/*func Uniform1uiv(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniform1uiv(_location, _count, _value)
}*/
/*func Uniform2uiv(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniform2uiv(_location, _count, _value)
}*/
/*func Uniform3uiv(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniform3uiv(_location, _count, _value)
}*/
/*func Uniform4uiv(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniform4uiv(_location, _count, _value)
}*/
/*func TexParameterIiv(target int, pname int, params ???){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	panic()
	C.glTexParameterIiv(_target, _pname, _params)
}*/
/*func TexParameterIuiv(target int, pname int, params ???){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	panic()
	C.glTexParameterIuiv(_target, _pname, _params)
}*/
/*func GetTexParameterIiv(target int, pname int, params ???){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	panic()
	C.glGetTexParameterIiv(_target, _pname, _params)
}*/
/*func GetTexParameterIuiv(target int, pname int, params ???){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	panic()
	C.glGetTexParameterIuiv(_target, _pname, _params)
}*/
/*func ClearBufferiv(buffer int, drawbuffer int, value ???){
	_buffer := C.GLenum(buffer)
	_drawbuffer := C.GLint(drawbuffer)
	panic()
	C.glClearBufferiv(_buffer, _drawbuffer, _value)
}*/
/*func ClearBufferuiv(buffer int, drawbuffer int, value ???){
	_buffer := C.GLenum(buffer)
	_drawbuffer := C.GLint(drawbuffer)
	panic()
	C.glClearBufferuiv(_buffer, _drawbuffer, _value)
}*/
/*func ClearBufferfv(buffer int, drawbuffer int, value ???){
	_buffer := C.GLenum(buffer)
	_drawbuffer := C.GLint(drawbuffer)
	panic()
	C.glClearBufferfv(_buffer, _drawbuffer, _value)
}*/
func ClearBufferfi(buffer int, drawbuffer int, depth float32, stencil int){
	_buffer := C.GLenum(buffer)
	_drawbuffer := C.GLint(drawbuffer)
	_depth := C.GLfloat(depth)
	_stencil := C.GLint(stencil)
	C.glClearBufferfi(_buffer, _drawbuffer, _depth, _stencil)
}
/*func GetStringi(name int, index uint)???{
	_name := C.GLenum(name)
	_index := C.GLuint(index)
	returnvalue := C.glGetStringi(_name, _index)
	panic()
return convreturnvalue
}*/
/*func IsRenderbuffer(renderbuffer uint)bool{
	_renderbuffer := C.GLuint(renderbuffer)
	returnvalue := C.glIsRenderbuffer(_renderbuffer)
	panic()
return convreturnvalue
}*/
func BindRenderbuffer(target int, renderbuffer uint){
	_target := C.GLenum(target)
	_renderbuffer := C.GLuint(renderbuffer)
	C.glBindRenderbuffer(_target, _renderbuffer)
}
/*func DeleteRenderbuffers(n int, renderbuffers ???){
	_n := C.GLsizei(n)
	panic()
	C.glDeleteRenderbuffers(_n, _renderbuffers)
}*/
/*func GenRenderbuffers(n int, renderbuffers ???){
	_n := C.GLsizei(n)
	panic()
	C.glGenRenderbuffers(_n, _renderbuffers)
}*/
func RenderbufferStorage(target int, internalformat int, width int, height int){
	_target := C.GLenum(target)
	_internalformat := C.GLenum(internalformat)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	C.glRenderbufferStorage(_target, _internalformat, _width, _height)
}
/*func GetRenderbufferParameteriv(target int, pname int, params ???){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	panic()
	C.glGetRenderbufferParameteriv(_target, _pname, _params)
}*/
/*func IsFramebuffer(framebuffer uint)bool{
	_framebuffer := C.GLuint(framebuffer)
	returnvalue := C.glIsFramebuffer(_framebuffer)
	panic()
return convreturnvalue
}*/
func BindFramebuffer(target int, framebuffer uint){
	_target := C.GLenum(target)
	_framebuffer := C.GLuint(framebuffer)
	C.glBindFramebuffer(_target, _framebuffer)
}
/*func DeleteFramebuffers(n int, framebuffers ???){
	_n := C.GLsizei(n)
	panic()
	C.glDeleteFramebuffers(_n, _framebuffers)
}*/
/*func GenFramebuffers(n int, framebuffers ???){
	_n := C.GLsizei(n)
	panic()
	C.glGenFramebuffers(_n, _framebuffers)
}*/
func CheckFramebufferStatus(target int)int{
	_target := C.GLenum(target)
	returnvalue := C.glCheckFramebufferStatus(_target)
	convreturnvalue := int(returnvalue)
return convreturnvalue
}
func FramebufferTexture1D(target int, attachment int, textarget int, texture uint, level int){
	_target := C.GLenum(target)
	_attachment := C.GLenum(attachment)
	_textarget := C.GLenum(textarget)
	_texture := C.GLuint(texture)
	_level := C.GLint(level)
	C.glFramebufferTexture1D(_target, _attachment, _textarget, _texture, _level)
}
func FramebufferTexture2D(target int, attachment int, textarget int, texture uint, level int){
	_target := C.GLenum(target)
	_attachment := C.GLenum(attachment)
	_textarget := C.GLenum(textarget)
	_texture := C.GLuint(texture)
	_level := C.GLint(level)
	C.glFramebufferTexture2D(_target, _attachment, _textarget, _texture, _level)
}
func FramebufferTexture3D(target int, attachment int, textarget int, texture uint, level int, zoffset int){
	_target := C.GLenum(target)
	_attachment := C.GLenum(attachment)
	_textarget := C.GLenum(textarget)
	_texture := C.GLuint(texture)
	_level := C.GLint(level)
	_zoffset := C.GLint(zoffset)
	C.glFramebufferTexture3D(_target, _attachment, _textarget, _texture, _level, _zoffset)
}
func FramebufferRenderbuffer(target int, attachment int, renderbuffertarget int, renderbuffer uint){
	_target := C.GLenum(target)
	_attachment := C.GLenum(attachment)
	_renderbuffertarget := C.GLenum(renderbuffertarget)
	_renderbuffer := C.GLuint(renderbuffer)
	C.glFramebufferRenderbuffer(_target, _attachment, _renderbuffertarget, _renderbuffer)
}
/*func GetFramebufferAttachmentParameteriv(target int, attachment int, pname int, params ???){
	_target := C.GLenum(target)
	_attachment := C.GLenum(attachment)
	_pname := C.GLenum(pname)
	panic()
	C.glGetFramebufferAttachmentParameteriv(_target, _attachment, _pname, _params)
}*/
func GenerateMipmap(target int){
	_target := C.GLenum(target)
	C.glGenerateMipmap(_target)
}
func BlitFramebuffer(srcX0 int, srcY0 int, srcX1 int, srcY1 int, dstX0 int, dstY0 int, dstX1 int, dstY1 int, mask uint, filter int){
	_srcX0 := C.GLint(srcX0)
	_srcY0 := C.GLint(srcY0)
	_srcX1 := C.GLint(srcX1)
	_srcY1 := C.GLint(srcY1)
	_dstX0 := C.GLint(dstX0)
	_dstY0 := C.GLint(dstY0)
	_dstX1 := C.GLint(dstX1)
	_dstY1 := C.GLint(dstY1)
	_mask := C.GLbitfield(mask)
	_filter := C.GLenum(filter)
	C.glBlitFramebuffer(_srcX0, _srcY0, _srcX1, _srcY1, _dstX0, _dstY0, _dstX1, _dstY1, _mask, _filter)
}
func RenderbufferStorageMultisample(target int, samples int, internalformat int, width int, height int){
	_target := C.GLenum(target)
	_samples := C.GLsizei(samples)
	_internalformat := C.GLenum(internalformat)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	C.glRenderbufferStorageMultisample(_target, _samples, _internalformat, _width, _height)
}
func FramebufferTextureLayer(target int, attachment int, texture uint, level int, layer int){
	_target := C.GLenum(target)
	_attachment := C.GLenum(attachment)
	_texture := C.GLuint(texture)
	_level := C.GLint(level)
	_layer := C.GLint(layer)
	C.glFramebufferTextureLayer(_target, _attachment, _texture, _level, _layer)
}
/*func MapBufferRange(target int, offset ???, length int, access uint)???{
	_target := C.GLenum(target)
	panic()
	_length := C.GLsizeiptr(length)
	_access := C.GLbitfield(access)
	returnvalue := C.glMapBufferRange(_target, _offset, _length, _access)
	panic()
return convreturnvalue
}*/
/*func FlushMappedBufferRange(target int, offset ???, length int){
	_target := C.GLenum(target)
	panic()
	_length := C.GLsizeiptr(length)
	C.glFlushMappedBufferRange(_target, _offset, _length)
}*/
func BindVertexArray(array uint){
	_array := C.GLuint(array)
	C.glBindVertexArray(_array)
}
/*func DeleteVertexArrays(n int, arrays ???){
	_n := C.GLsizei(n)
	panic()
	C.glDeleteVertexArrays(_n, _arrays)
}*/
/*func GenVertexArrays(n int, arrays ???){
	_n := C.GLsizei(n)
	panic()
	C.glGenVertexArrays(_n, _arrays)
}*/
/*func IsVertexArray(array uint)bool{
	_array := C.GLuint(array)
	returnvalue := C.glIsVertexArray(_array)
	panic()
return convreturnvalue
}*/
func DrawArraysInstanced(mode int, first int, count int, instancecount int){
	_mode := C.GLenum(mode)
	_first := C.GLint(first)
	_count := C.GLsizei(count)
	_instancecount := C.GLsizei(instancecount)
	C.glDrawArraysInstanced(_mode, _first, _count, _instancecount)
}
/*func DrawElementsInstanced(mode int, count int, whichtype int, indices ???, instancecount int){
	_mode := C.GLenum(mode)
	_count := C.GLsizei(count)
	_whichtype := C.GLenum(whichtype)
	panic()
	_instancecount := C.GLsizei(instancecount)
	C.glDrawElementsInstanced(_mode, _count, _whichtype, _indices, _instancecount)
}*/
func TexBuffer(target int, internalformat int, buffer uint){
	_target := C.GLenum(target)
	_internalformat := C.GLenum(internalformat)
	_buffer := C.GLuint(buffer)
	C.glTexBuffer(_target, _internalformat, _buffer)
}
func PrimitiveRestartIndex(index uint){
	_index := C.GLuint(index)
	C.glPrimitiveRestartIndex(_index)
}
/*func CopyBufferSubData(readTarget int, writeTarget int, readOffset ???, writeOffset ???, size int){
	_readTarget := C.GLenum(readTarget)
	_writeTarget := C.GLenum(writeTarget)
	panic()
	panic()
	_size := C.GLsizeiptr(size)
	C.glCopyBufferSubData(_readTarget, _writeTarget, _readOffset, _writeOffset, _size)
}*/
/*func GetUniformIndices(program uint, uniformCount int, uniformNames ???, uniformIndices ???){
	_program := C.GLuint(program)
	_uniformCount := C.GLsizei(uniformCount)
	panic()
	panic()
	C.glGetUniformIndices(_program, _uniformCount, _uniformNames, _uniformIndices)
}*/
/*func GetActiveUniformsiv(program uint, uniformCount int, uniformIndices ???, pname int, params ???){
	_program := C.GLuint(program)
	_uniformCount := C.GLsizei(uniformCount)
	panic()
	_pname := C.GLenum(pname)
	panic()
	C.glGetActiveUniformsiv(_program, _uniformCount, _uniformIndices, _pname, _params)
}*/
/*func GetActiveUniformName(program uint, uniformIndex uint, bufSize int, length int, uniformName int){
	_program := C.GLuint(program)
	_uniformIndex := C.GLuint(uniformIndex)
	_bufSize := C.GLsizei(bufSize)
	_length := C.GLsizei(length)
	panic()
	C.glGetActiveUniformName(_program, _uniformIndex, _bufSize, _length, _uniformName)
}*/
/*func GetUniformBlockIndex(program uint, uniformBlockName ???)uint{
	_program := C.GLuint(program)
	panic()
	returnvalue := C.glGetUniformBlockIndex(_program, _uniformBlockName)
	panic()
return convreturnvalue
}*/
/*func GetActiveUniformBlockiv(program uint, uniformBlockIndex uint, pname int, params ???){
	_program := C.GLuint(program)
	_uniformBlockIndex := C.GLuint(uniformBlockIndex)
	_pname := C.GLenum(pname)
	panic()
	C.glGetActiveUniformBlockiv(_program, _uniformBlockIndex, _pname, _params)
}*/
/*func GetActiveUniformBlockName(program uint, uniformBlockIndex uint, bufSize int, length int, uniformBlockName int){
	_program := C.GLuint(program)
	_uniformBlockIndex := C.GLuint(uniformBlockIndex)
	_bufSize := C.GLsizei(bufSize)
	_length := C.GLsizei(length)
	panic()
	C.glGetActiveUniformBlockName(_program, _uniformBlockIndex, _bufSize, _length, _uniformBlockName)
}*/
func UniformBlockBinding(program uint, uniformBlockIndex uint, uniformBlockBinding uint){
	_program := C.GLuint(program)
	_uniformBlockIndex := C.GLuint(uniformBlockIndex)
	_uniformBlockBinding := C.GLuint(uniformBlockBinding)
	C.glUniformBlockBinding(_program, _uniformBlockIndex, _uniformBlockBinding)
}
/*func DrawElementsBaseVertex(mode int, count int, whichtype int, indices ???, basevertex int){
	_mode := C.GLenum(mode)
	_count := C.GLsizei(count)
	_whichtype := C.GLenum(whichtype)
	panic()
	_basevertex := C.GLint(basevertex)
	C.glDrawElementsBaseVertex(_mode, _count, _whichtype, _indices, _basevertex)
}*/
/*func DrawRangeElementsBaseVertex(mode int, start uint, end uint, count int, whichtype int, indices ???, basevertex int){
	_mode := C.GLenum(mode)
	_start := C.GLuint(start)
	_end := C.GLuint(end)
	_count := C.GLsizei(count)
	_whichtype := C.GLenum(whichtype)
	panic()
	_basevertex := C.GLint(basevertex)
	C.glDrawRangeElementsBaseVertex(_mode, _start, _end, _count, _whichtype, _indices, _basevertex)
}*/
/*func DrawElementsInstancedBaseVertex(mode int, count int, whichtype int, indices ???, instancecount int, basevertex int){
	_mode := C.GLenum(mode)
	_count := C.GLsizei(count)
	_whichtype := C.GLenum(whichtype)
	panic()
	_instancecount := C.GLsizei(instancecount)
	_basevertex := C.GLint(basevertex)
	C.glDrawElementsInstancedBaseVertex(_mode, _count, _whichtype, _indices, _instancecount, _basevertex)
}*/
/*func MultiDrawElementsBaseVertex(mode int, count ???, whichtype int, indices ???, drawcount int, basevertex ???){
	_mode := C.GLenum(mode)
	panic()
	_whichtype := C.GLenum(whichtype)
	panic()
	_drawcount := C.GLsizei(drawcount)
	panic()
	C.glMultiDrawElementsBaseVertex(_mode, _count, _whichtype, _indices, _drawcount, _basevertex)
}*/
func ProvokingVertex(mode int){
	_mode := C.GLenum(mode)
	C.glProvokingVertex(_mode)
}
/*func FenceSync(condition int, flags uint)???{
	_condition := C.GLenum(condition)
	_flags := C.GLbitfield(flags)
	returnvalue := C.glFenceSync(_condition, _flags)
	panic()
return convreturnvalue
}*/
/*func IsSync(sync ???)bool{
	panic()
	returnvalue := C.glIsSync(_sync)
	panic()
return convreturnvalue
}*/
/*func DeleteSync(sync ???){
	panic()
	C.glDeleteSync(_sync)
}*/
/*func ClientWaitSync(sync ???, flags uint, timeout ???)int{
	panic()
	_flags := C.GLbitfield(flags)
	panic()
	returnvalue := C.glClientWaitSync(_sync, _flags, _timeout)
	convreturnvalue := int(returnvalue)
return convreturnvalue
}*/
/*func WaitSync(sync ???, flags uint, timeout ???){
	panic()
	_flags := C.GLbitfield(flags)
	panic()
	C.glWaitSync(_sync, _flags, _timeout)
}*/
/*func GetInteger64v(pname int, data ???){
	_pname := C.GLenum(pname)
	panic()
	C.glGetInteger64v(_pname, _data)
}*/
/*func GetSynciv(sync ???, pname int, bufSize int, length int, values ???){
	panic()
	_pname := C.GLenum(pname)
	_bufSize := C.GLsizei(bufSize)
	_length := C.GLsizei(length)
	panic()
	C.glGetSynciv(_sync, _pname, _bufSize, _length, _values)
}*/
/*func GetInteger64i_v(target int, index uint, data ???){
	_target := C.GLenum(target)
	_index := C.GLuint(index)
	panic()
	C.glGetInteger64i_v(_target, _index, _data)
}*/
/*func GetBufferParameteri64v(target int, pname int, params ???){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	panic()
	C.glGetBufferParameteri64v(_target, _pname, _params)
}*/
func FramebufferTexture(target int, attachment int, texture uint, level int){
	_target := C.GLenum(target)
	_attachment := C.GLenum(attachment)
	_texture := C.GLuint(texture)
	_level := C.GLint(level)
	C.glFramebufferTexture(_target, _attachment, _texture, _level)
}
func TexImage2DMultisample(target int, samples int, internalformat int, width int, height int, fixedsamplelocations bool){
	_target := C.GLenum(target)
	_samples := C.GLsizei(samples)
	_internalformat := C.GLenum(internalformat)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_fixedsamplelocations := C.GLboolean(1);if !fixedsamplelocations{_fixedsamplelocations=C.GLboolean(0)}
	C.glTexImage2DMultisample(_target, _samples, _internalformat, _width, _height, _fixedsamplelocations)
}
func TexImage3DMultisample(target int, samples int, internalformat int, width int, height int, depth int, fixedsamplelocations bool){
	_target := C.GLenum(target)
	_samples := C.GLsizei(samples)
	_internalformat := C.GLenum(internalformat)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_depth := C.GLsizei(depth)
	_fixedsamplelocations := C.GLboolean(1);if !fixedsamplelocations{_fixedsamplelocations=C.GLboolean(0)}
	C.glTexImage3DMultisample(_target, _samples, _internalformat, _width, _height, _depth, _fixedsamplelocations)
}
/*func GetMultisamplefv(pname int, index uint, val ???){
	_pname := C.GLenum(pname)
	_index := C.GLuint(index)
	panic()
	C.glGetMultisamplefv(_pname, _index, _val)
}*/
func SampleMaski(maskNumber uint, mask uint){
	_maskNumber := C.GLuint(maskNumber)
	_mask := C.GLbitfield(mask)
	C.glSampleMaski(_maskNumber, _mask)
}
/*func BindFragDataLocationIndexed(program uint, colorNumber uint, index uint, name ???){
	_program := C.GLuint(program)
	_colorNumber := C.GLuint(colorNumber)
	_index := C.GLuint(index)
	panic()
	C.glBindFragDataLocationIndexed(_program, _colorNumber, _index, _name)
}*/
/*func GetFragDataIndex(program uint, name ???)int{
	_program := C.GLuint(program)
	panic()
	returnvalue := C.glGetFragDataIndex(_program, _name)
	panic()
return convreturnvalue
}*/
/*func GenSamplers(count int, samplers ???){
	_count := C.GLsizei(count)
	panic()
	C.glGenSamplers(_count, _samplers)
}*/
/*func DeleteSamplers(count int, samplers ???){
	_count := C.GLsizei(count)
	panic()
	C.glDeleteSamplers(_count, _samplers)
}*/
/*func IsSampler(sampler uint)bool{
	_sampler := C.GLuint(sampler)
	returnvalue := C.glIsSampler(_sampler)
	panic()
return convreturnvalue
}*/
func BindSampler(unit uint, sampler uint){
	_unit := C.GLuint(unit)
	_sampler := C.GLuint(sampler)
	C.glBindSampler(_unit, _sampler)
}
func SamplerParameteri(sampler uint, pname int, param int){
	_sampler := C.GLuint(sampler)
	_pname := C.GLenum(pname)
	_param := C.GLint(param)
	C.glSamplerParameteri(_sampler, _pname, _param)
}
/*func SamplerParameteriv(sampler uint, pname int, param ???){
	_sampler := C.GLuint(sampler)
	_pname := C.GLenum(pname)
	panic()
	C.glSamplerParameteriv(_sampler, _pname, _param)
}*/
func SamplerParameterf(sampler uint, pname int, param float32){
	_sampler := C.GLuint(sampler)
	_pname := C.GLenum(pname)
	_param := C.GLfloat(param)
	C.glSamplerParameterf(_sampler, _pname, _param)
}
/*func SamplerParameterfv(sampler uint, pname int, param ???){
	_sampler := C.GLuint(sampler)
	_pname := C.GLenum(pname)
	panic()
	C.glSamplerParameterfv(_sampler, _pname, _param)
}*/
/*func SamplerParameterIiv(sampler uint, pname int, param ???){
	_sampler := C.GLuint(sampler)
	_pname := C.GLenum(pname)
	panic()
	C.glSamplerParameterIiv(_sampler, _pname, _param)
}*/
/*func SamplerParameterIuiv(sampler uint, pname int, param ???){
	_sampler := C.GLuint(sampler)
	_pname := C.GLenum(pname)
	panic()
	C.glSamplerParameterIuiv(_sampler, _pname, _param)
}*/
/*func GetSamplerParameteriv(sampler uint, pname int, params ???){
	_sampler := C.GLuint(sampler)
	_pname := C.GLenum(pname)
	panic()
	C.glGetSamplerParameteriv(_sampler, _pname, _params)
}*/
/*func GetSamplerParameterIiv(sampler uint, pname int, params ???){
	_sampler := C.GLuint(sampler)
	_pname := C.GLenum(pname)
	panic()
	C.glGetSamplerParameterIiv(_sampler, _pname, _params)
}*/
/*func GetSamplerParameterfv(sampler uint, pname int, params ???){
	_sampler := C.GLuint(sampler)
	_pname := C.GLenum(pname)
	panic()
	C.glGetSamplerParameterfv(_sampler, _pname, _params)
}*/
/*func GetSamplerParameterIuiv(sampler uint, pname int, params ???){
	_sampler := C.GLuint(sampler)
	_pname := C.GLenum(pname)
	panic()
	C.glGetSamplerParameterIuiv(_sampler, _pname, _params)
}*/
func QueryCounter(id uint, target int){
	_id := C.GLuint(id)
	_target := C.GLenum(target)
	C.glQueryCounter(_id, _target)
}
/*func GetQueryObjecti64v(id uint, pname int, params ???){
	_id := C.GLuint(id)
	_pname := C.GLenum(pname)
	panic()
	C.glGetQueryObjecti64v(_id, _pname, _params)
}*/
/*func GetQueryObjectui64v(id uint, pname int, params ???){
	_id := C.GLuint(id)
	_pname := C.GLenum(pname)
	panic()
	C.glGetQueryObjectui64v(_id, _pname, _params)
}*/
func VertexAttribDivisor(index uint, divisor uint){
	_index := C.GLuint(index)
	_divisor := C.GLuint(divisor)
	C.glVertexAttribDivisor(_index, _divisor)
}
func VertexAttribP1ui(index uint, whichtype int, normalized bool, value uint){
	_index := C.GLuint(index)
	_whichtype := C.GLenum(whichtype)
	_normalized := C.GLboolean(1);if !normalized{_normalized=C.GLboolean(0)}
	_value := C.GLuint(value)
	C.glVertexAttribP1ui(_index, _whichtype, _normalized, _value)
}
/*func VertexAttribP1uiv(index uint, whichtype int, normalized bool, value ???){
	_index := C.GLuint(index)
	_whichtype := C.GLenum(whichtype)
	_normalized := C.GLboolean(1);if !normalized{_normalized=C.GLboolean(0)}
	panic()
	C.glVertexAttribP1uiv(_index, _whichtype, _normalized, _value)
}*/
func VertexAttribP2ui(index uint, whichtype int, normalized bool, value uint){
	_index := C.GLuint(index)
	_whichtype := C.GLenum(whichtype)
	_normalized := C.GLboolean(1);if !normalized{_normalized=C.GLboolean(0)}
	_value := C.GLuint(value)
	C.glVertexAttribP2ui(_index, _whichtype, _normalized, _value)
}
/*func VertexAttribP2uiv(index uint, whichtype int, normalized bool, value ???){
	_index := C.GLuint(index)
	_whichtype := C.GLenum(whichtype)
	_normalized := C.GLboolean(1);if !normalized{_normalized=C.GLboolean(0)}
	panic()
	C.glVertexAttribP2uiv(_index, _whichtype, _normalized, _value)
}*/
func VertexAttribP3ui(index uint, whichtype int, normalized bool, value uint){
	_index := C.GLuint(index)
	_whichtype := C.GLenum(whichtype)
	_normalized := C.GLboolean(1);if !normalized{_normalized=C.GLboolean(0)}
	_value := C.GLuint(value)
	C.glVertexAttribP3ui(_index, _whichtype, _normalized, _value)
}
/*func VertexAttribP3uiv(index uint, whichtype int, normalized bool, value ???){
	_index := C.GLuint(index)
	_whichtype := C.GLenum(whichtype)
	_normalized := C.GLboolean(1);if !normalized{_normalized=C.GLboolean(0)}
	panic()
	C.glVertexAttribP3uiv(_index, _whichtype, _normalized, _value)
}*/
func VertexAttribP4ui(index uint, whichtype int, normalized bool, value uint){
	_index := C.GLuint(index)
	_whichtype := C.GLenum(whichtype)
	_normalized := C.GLboolean(1);if !normalized{_normalized=C.GLboolean(0)}
	_value := C.GLuint(value)
	C.glVertexAttribP4ui(_index, _whichtype, _normalized, _value)
}
/*func VertexAttribP4uiv(index uint, whichtype int, normalized bool, value ???){
	_index := C.GLuint(index)
	_whichtype := C.GLenum(whichtype)
	_normalized := C.GLboolean(1);if !normalized{_normalized=C.GLboolean(0)}
	panic()
	C.glVertexAttribP4uiv(_index, _whichtype, _normalized, _value)
}*/
func MinSampleShading(value float32){
	_value := C.GLfloat(value)
	C.glMinSampleShading(_value)
}
func BlendEquationi(buf uint, mode int){
	_buf := C.GLuint(buf)
	_mode := C.GLenum(mode)
	C.glBlendEquationi(_buf, _mode)
}
func BlendEquationSeparatei(buf uint, modeRGB int, modeAlpha int){
	_buf := C.GLuint(buf)
	_modeRGB := C.GLenum(modeRGB)
	_modeAlpha := C.GLenum(modeAlpha)
	C.glBlendEquationSeparatei(_buf, _modeRGB, _modeAlpha)
}
func BlendFunci(buf uint, src int, dst int){
	_buf := C.GLuint(buf)
	_src := C.GLenum(src)
	_dst := C.GLenum(dst)
	C.glBlendFunci(_buf, _src, _dst)
}
func BlendFuncSeparatei(buf uint, srcRGB int, dstRGB int, srcAlpha int, dstAlpha int){
	_buf := C.GLuint(buf)
	_srcRGB := C.GLenum(srcRGB)
	_dstRGB := C.GLenum(dstRGB)
	_srcAlpha := C.GLenum(srcAlpha)
	_dstAlpha := C.GLenum(dstAlpha)
	C.glBlendFuncSeparatei(_buf, _srcRGB, _dstRGB, _srcAlpha, _dstAlpha)
}
/*func DrawArraysIndirect(mode int, indirect ???){
	_mode := C.GLenum(mode)
	panic()
	C.glDrawArraysIndirect(_mode, _indirect)
}*/
/*func DrawElementsIndirect(mode int, whichtype int, indirect ???){
	_mode := C.GLenum(mode)
	_whichtype := C.GLenum(whichtype)
	panic()
	C.glDrawElementsIndirect(_mode, _whichtype, _indirect)
}*/
func Uniform1d(location int, x float64){
	_location := C.GLint(location)
	_x := C.GLdouble(x)
	C.glUniform1d(_location, _x)
}
func Uniform2d(location int, x float64, y float64){
	_location := C.GLint(location)
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	C.glUniform2d(_location, _x, _y)
}
func Uniform3d(location int, x float64, y float64, z float64){
	_location := C.GLint(location)
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	_z := C.GLdouble(z)
	C.glUniform3d(_location, _x, _y, _z)
}
func Uniform4d(location int, x float64, y float64, z float64, w float64){
	_location := C.GLint(location)
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	_z := C.GLdouble(z)
	_w := C.GLdouble(w)
	C.glUniform4d(_location, _x, _y, _z, _w)
}
/*func Uniform1dv(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniform1dv(_location, _count, _value)
}*/
/*func Uniform2dv(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniform2dv(_location, _count, _value)
}*/
/*func Uniform3dv(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniform3dv(_location, _count, _value)
}*/
/*func Uniform4dv(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniform4dv(_location, _count, _value)
}*/
/*func UniformMatrix2dv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix2dv(_location, _count, _transpose, _value)
}*/
/*func UniformMatrix3dv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix3dv(_location, _count, _transpose, _value)
}*/
/*func UniformMatrix4dv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix4dv(_location, _count, _transpose, _value)
}*/
/*func UniformMatrix2x3dv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix2x3dv(_location, _count, _transpose, _value)
}*/
/*func UniformMatrix2x4dv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix2x4dv(_location, _count, _transpose, _value)
}*/
/*func UniformMatrix3x2dv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix3x2dv(_location, _count, _transpose, _value)
}*/
/*func UniformMatrix3x4dv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix3x4dv(_location, _count, _transpose, _value)
}*/
/*func UniformMatrix4x2dv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix4x2dv(_location, _count, _transpose, _value)
}*/
/*func UniformMatrix4x3dv(location int, count int, transpose bool, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glUniformMatrix4x3dv(_location, _count, _transpose, _value)
}*/
/*func GetUniformdv(program uint, location int, params ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	panic()
	C.glGetUniformdv(_program, _location, _params)
}*/
/*func GetSubroutineUniformLocation(program uint, shadertype int, name ???)int{
	_program := C.GLuint(program)
	_shadertype := C.GLenum(shadertype)
	panic()
	returnvalue := C.glGetSubroutineUniformLocation(_program, _shadertype, _name)
	panic()
return convreturnvalue
}*/
/*func GetSubroutineIndex(program uint, shadertype int, name ???)uint{
	_program := C.GLuint(program)
	_shadertype := C.GLenum(shadertype)
	panic()
	returnvalue := C.glGetSubroutineIndex(_program, _shadertype, _name)
	panic()
return convreturnvalue
}*/
/*func GetActiveSubroutineUniformiv(program uint, shadertype int, index uint, pname int, values ???){
	_program := C.GLuint(program)
	_shadertype := C.GLenum(shadertype)
	_index := C.GLuint(index)
	_pname := C.GLenum(pname)
	panic()
	C.glGetActiveSubroutineUniformiv(_program, _shadertype, _index, _pname, _values)
}*/
/*func GetActiveSubroutineUniformName(program uint, shadertype int, index uint, bufsize int, length int, name int){
	_program := C.GLuint(program)
	_shadertype := C.GLenum(shadertype)
	_index := C.GLuint(index)
	_bufsize := C.GLsizei(bufsize)
	_length := C.GLsizei(length)
	panic()
	C.glGetActiveSubroutineUniformName(_program, _shadertype, _index, _bufsize, _length, _name)
}*/
/*func GetActiveSubroutineName(program uint, shadertype int, index uint, bufsize int, length int, name int){
	_program := C.GLuint(program)
	_shadertype := C.GLenum(shadertype)
	_index := C.GLuint(index)
	_bufsize := C.GLsizei(bufsize)
	_length := C.GLsizei(length)
	panic()
	C.glGetActiveSubroutineName(_program, _shadertype, _index, _bufsize, _length, _name)
}*/
/*func UniformSubroutinesuiv(shadertype int, count int, indices ???){
	_shadertype := C.GLenum(shadertype)
	_count := C.GLsizei(count)
	panic()
	C.glUniformSubroutinesuiv(_shadertype, _count, _indices)
}*/
/*func GetUniformSubroutineuiv(shadertype int, location int, params ???){
	_shadertype := C.GLenum(shadertype)
	_location := C.GLint(location)
	panic()
	C.glGetUniformSubroutineuiv(_shadertype, _location, _params)
}*/
/*func GetProgramStageiv(program uint, shadertype int, pname int, values ???){
	_program := C.GLuint(program)
	_shadertype := C.GLenum(shadertype)
	_pname := C.GLenum(pname)
	panic()
	C.glGetProgramStageiv(_program, _shadertype, _pname, _values)
}*/
func PatchParameteri(pname int, value int){
	_pname := C.GLenum(pname)
	_value := C.GLint(value)
	C.glPatchParameteri(_pname, _value)
}
/*func PatchParameterfv(pname int, values ???){
	_pname := C.GLenum(pname)
	panic()
	C.glPatchParameterfv(_pname, _values)
}*/
func BindTransformFeedback(target int, id uint){
	_target := C.GLenum(target)
	_id := C.GLuint(id)
	C.glBindTransformFeedback(_target, _id)
}
/*func DeleteTransformFeedbacks(n int, ids ???){
	_n := C.GLsizei(n)
	panic()
	C.glDeleteTransformFeedbacks(_n, _ids)
}*/
/*func GenTransformFeedbacks(n int, ids ???){
	_n := C.GLsizei(n)
	panic()
	C.glGenTransformFeedbacks(_n, _ids)
}*/
/*func IsTransformFeedback(id uint)bool{
	_id := C.GLuint(id)
	returnvalue := C.glIsTransformFeedback(_id)
	panic()
return convreturnvalue
}*/
func PauseTransformFeedback(){
	C.glPauseTransformFeedback()
}
func ResumeTransformFeedback(){
	C.glResumeTransformFeedback()
}
func DrawTransformFeedback(mode int, id uint){
	_mode := C.GLenum(mode)
	_id := C.GLuint(id)
	C.glDrawTransformFeedback(_mode, _id)
}
func DrawTransformFeedbackStream(mode int, id uint, stream uint){
	_mode := C.GLenum(mode)
	_id := C.GLuint(id)
	_stream := C.GLuint(stream)
	C.glDrawTransformFeedbackStream(_mode, _id, _stream)
}
func BeginQueryIndexed(target int, index uint, id uint){
	_target := C.GLenum(target)
	_index := C.GLuint(index)
	_id := C.GLuint(id)
	C.glBeginQueryIndexed(_target, _index, _id)
}
func EndQueryIndexed(target int, index uint){
	_target := C.GLenum(target)
	_index := C.GLuint(index)
	C.glEndQueryIndexed(_target, _index)
}
/*func GetQueryIndexediv(target int, index uint, pname int, params ???){
	_target := C.GLenum(target)
	_index := C.GLuint(index)
	_pname := C.GLenum(pname)
	panic()
	C.glGetQueryIndexediv(_target, _index, _pname, _params)
}*/
func ReleaseShaderCompiler(){
	C.glReleaseShaderCompiler()
}
/*func ShaderBinary(count int, shaders ???, binaryformat int, binary ???, length int){
	_count := C.GLsizei(count)
	panic()
	_binaryformat := C.GLenum(binaryformat)
	panic()
	_length := C.GLsizei(length)
	C.glShaderBinary(_count, _shaders, _binaryformat, _binary, _length)
}*/
/*func GetShaderPrecisionFormat(shadertype int, precisiontype int, range ???, precision ???){
	_shadertype := C.GLenum(shadertype)
	_precisiontype := C.GLenum(precisiontype)
	panic()
	panic()
	C.glGetShaderPrecisionFormat(_shadertype, _precisiontype, _range, _precision)
}*/
func DepthRangef(n float32, f float32){
	_n := C.GLfloat(n)
	_f := C.GLfloat(f)
	C.glDepthRangef(_n, _f)
}
func ClearDepthf(d float32){
	_d := C.GLfloat(d)
	C.glClearDepthf(_d)
}
/*func GetProgramBinary(program uint, bufSize int, length int, binaryFormat ???, binary ???){
	_program := C.GLuint(program)
	_bufSize := C.GLsizei(bufSize)
	_length := C.GLsizei(length)
	panic()
	panic()
	C.glGetProgramBinary(_program, _bufSize, _length, _binaryFormat, _binary)
}*/
/*func ProgramBinary(program uint, binaryFormat int, binary ???, length int){
	_program := C.GLuint(program)
	_binaryFormat := C.GLenum(binaryFormat)
	panic()
	_length := C.GLsizei(length)
	C.glProgramBinary(_program, _binaryFormat, _binary, _length)
}*/
func ProgramParameteri(program uint, pname int, value int){
	_program := C.GLuint(program)
	_pname := C.GLenum(pname)
	_value := C.GLint(value)
	C.glProgramParameteri(_program, _pname, _value)
}
func UseProgramStages(pipeline uint, stages uint, program uint){
	_pipeline := C.GLuint(pipeline)
	_stages := C.GLbitfield(stages)
	_program := C.GLuint(program)
	C.glUseProgramStages(_pipeline, _stages, _program)
}
func ActiveShaderProgram(pipeline uint, program uint){
	_pipeline := C.GLuint(pipeline)
	_program := C.GLuint(program)
	C.glActiveShaderProgram(_pipeline, _program)
}
/*func CreateShaderProgramv(whichtype int, count int, strings ???)uint{
	_whichtype := C.GLenum(whichtype)
	_count := C.GLsizei(count)
	panic()
	returnvalue := C.glCreateShaderProgramv(_whichtype, _count, _strings)
	panic()
return convreturnvalue
}*/
func BindProgramPipeline(pipeline uint){
	_pipeline := C.GLuint(pipeline)
	C.glBindProgramPipeline(_pipeline)
}
/*func DeleteProgramPipelines(n int, pipelines ???){
	_n := C.GLsizei(n)
	panic()
	C.glDeleteProgramPipelines(_n, _pipelines)
}*/
/*func GenProgramPipelines(n int, pipelines ???){
	_n := C.GLsizei(n)
	panic()
	C.glGenProgramPipelines(_n, _pipelines)
}*/
/*func IsProgramPipeline(pipeline uint)bool{
	_pipeline := C.GLuint(pipeline)
	returnvalue := C.glIsProgramPipeline(_pipeline)
	panic()
return convreturnvalue
}*/
/*func GetProgramPipelineiv(pipeline uint, pname int, params ???){
	_pipeline := C.GLuint(pipeline)
	_pname := C.GLenum(pname)
	panic()
	C.glGetProgramPipelineiv(_pipeline, _pname, _params)
}*/
func ProgramUniform1i(program uint, location int, v0 int){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_v0 := C.GLint(v0)
	C.glProgramUniform1i(_program, _location, _v0)
}
/*func ProgramUniform1iv(program uint, location int, count int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniform1iv(_program, _location, _count, _value)
}*/
func ProgramUniform1f(program uint, location int, v0 float32){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_v0 := C.GLfloat(v0)
	C.glProgramUniform1f(_program, _location, _v0)
}
/*func ProgramUniform1fv(program uint, location int, count int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniform1fv(_program, _location, _count, _value)
}*/
func ProgramUniform1d(program uint, location int, v0 float64){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_v0 := C.GLdouble(v0)
	C.glProgramUniform1d(_program, _location, _v0)
}
/*func ProgramUniform1dv(program uint, location int, count int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniform1dv(_program, _location, _count, _value)
}*/
func ProgramUniform1ui(program uint, location int, v0 uint){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_v0 := C.GLuint(v0)
	C.glProgramUniform1ui(_program, _location, _v0)
}
/*func ProgramUniform1uiv(program uint, location int, count int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniform1uiv(_program, _location, _count, _value)
}*/
func ProgramUniform2i(program uint, location int, v0 int, v1 int){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_v0 := C.GLint(v0)
	_v1 := C.GLint(v1)
	C.glProgramUniform2i(_program, _location, _v0, _v1)
}
/*func ProgramUniform2iv(program uint, location int, count int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniform2iv(_program, _location, _count, _value)
}*/
func ProgramUniform2f(program uint, location int, v0 float32, v1 float32){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_v0 := C.GLfloat(v0)
	_v1 := C.GLfloat(v1)
	C.glProgramUniform2f(_program, _location, _v0, _v1)
}
/*func ProgramUniform2fv(program uint, location int, count int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniform2fv(_program, _location, _count, _value)
}*/
func ProgramUniform2d(program uint, location int, v0 float64, v1 float64){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_v0 := C.GLdouble(v0)
	_v1 := C.GLdouble(v1)
	C.glProgramUniform2d(_program, _location, _v0, _v1)
}
/*func ProgramUniform2dv(program uint, location int, count int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniform2dv(_program, _location, _count, _value)
}*/
func ProgramUniform2ui(program uint, location int, v0 uint, v1 uint){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_v0 := C.GLuint(v0)
	_v1 := C.GLuint(v1)
	C.glProgramUniform2ui(_program, _location, _v0, _v1)
}
/*func ProgramUniform2uiv(program uint, location int, count int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniform2uiv(_program, _location, _count, _value)
}*/
func ProgramUniform3i(program uint, location int, v0 int, v1 int, v2 int){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_v0 := C.GLint(v0)
	_v1 := C.GLint(v1)
	_v2 := C.GLint(v2)
	C.glProgramUniform3i(_program, _location, _v0, _v1, _v2)
}
/*func ProgramUniform3iv(program uint, location int, count int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniform3iv(_program, _location, _count, _value)
}*/
func ProgramUniform3f(program uint, location int, v0 float32, v1 float32, v2 float32){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_v0 := C.GLfloat(v0)
	_v1 := C.GLfloat(v1)
	_v2 := C.GLfloat(v2)
	C.glProgramUniform3f(_program, _location, _v0, _v1, _v2)
}
/*func ProgramUniform3fv(program uint, location int, count int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniform3fv(_program, _location, _count, _value)
}*/
func ProgramUniform3d(program uint, location int, v0 float64, v1 float64, v2 float64){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_v0 := C.GLdouble(v0)
	_v1 := C.GLdouble(v1)
	_v2 := C.GLdouble(v2)
	C.glProgramUniform3d(_program, _location, _v0, _v1, _v2)
}
/*func ProgramUniform3dv(program uint, location int, count int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniform3dv(_program, _location, _count, _value)
}*/
func ProgramUniform3ui(program uint, location int, v0 uint, v1 uint, v2 uint){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_v0 := C.GLuint(v0)
	_v1 := C.GLuint(v1)
	_v2 := C.GLuint(v2)
	C.glProgramUniform3ui(_program, _location, _v0, _v1, _v2)
}
/*func ProgramUniform3uiv(program uint, location int, count int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniform3uiv(_program, _location, _count, _value)
}*/
func ProgramUniform4i(program uint, location int, v0 int, v1 int, v2 int, v3 int){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_v0 := C.GLint(v0)
	_v1 := C.GLint(v1)
	_v2 := C.GLint(v2)
	_v3 := C.GLint(v3)
	C.glProgramUniform4i(_program, _location, _v0, _v1, _v2, _v3)
}
/*func ProgramUniform4iv(program uint, location int, count int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniform4iv(_program, _location, _count, _value)
}*/
func ProgramUniform4f(program uint, location int, v0 float32, v1 float32, v2 float32, v3 float32){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_v0 := C.GLfloat(v0)
	_v1 := C.GLfloat(v1)
	_v2 := C.GLfloat(v2)
	_v3 := C.GLfloat(v3)
	C.glProgramUniform4f(_program, _location, _v0, _v1, _v2, _v3)
}
/*func ProgramUniform4fv(program uint, location int, count int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniform4fv(_program, _location, _count, _value)
}*/
func ProgramUniform4d(program uint, location int, v0 float64, v1 float64, v2 float64, v3 float64){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_v0 := C.GLdouble(v0)
	_v1 := C.GLdouble(v1)
	_v2 := C.GLdouble(v2)
	_v3 := C.GLdouble(v3)
	C.glProgramUniform4d(_program, _location, _v0, _v1, _v2, _v3)
}
/*func ProgramUniform4dv(program uint, location int, count int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniform4dv(_program, _location, _count, _value)
}*/
func ProgramUniform4ui(program uint, location int, v0 uint, v1 uint, v2 uint, v3 uint){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_v0 := C.GLuint(v0)
	_v1 := C.GLuint(v1)
	_v2 := C.GLuint(v2)
	_v3 := C.GLuint(v3)
	C.glProgramUniform4ui(_program, _location, _v0, _v1, _v2, _v3)
}
/*func ProgramUniform4uiv(program uint, location int, count int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniform4uiv(_program, _location, _count, _value)
}*/
/*func ProgramUniformMatrix2fv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix2fv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix3fv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix3fv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix4fv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix4fv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix2dv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix2dv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix3dv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix3dv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix4dv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix4dv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix2x3fv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix2x3fv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix3x2fv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix3x2fv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix2x4fv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix2x4fv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix4x2fv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix4x2fv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix3x4fv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix3x4fv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix4x3fv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix4x3fv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix2x3dv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix2x3dv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix3x2dv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix3x2dv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix2x4dv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix2x4dv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix4x2dv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix4x2dv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix3x4dv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix3x4dv(_program, _location, _count, _transpose, _value)
}*/
/*func ProgramUniformMatrix4x3dv(program uint, location int, count int, transpose bool, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	_transpose := C.GLboolean(1);if !transpose{_transpose=C.GLboolean(0)}
	panic()
	C.glProgramUniformMatrix4x3dv(_program, _location, _count, _transpose, _value)
}*/
func ValidateProgramPipeline(pipeline uint){
	_pipeline := C.GLuint(pipeline)
	C.glValidateProgramPipeline(_pipeline)
}
/*func GetProgramPipelineInfoLog(pipeline uint, bufSize int, length int, infoLog int){
	_pipeline := C.GLuint(pipeline)
	_bufSize := C.GLsizei(bufSize)
	_length := C.GLsizei(length)
	panic()
	C.glGetProgramPipelineInfoLog(_pipeline, _bufSize, _length, _infoLog)
}*/
func VertexAttribL1d(index uint, x float64){
	_index := C.GLuint(index)
	_x := C.GLdouble(x)
	C.glVertexAttribL1d(_index, _x)
}
func VertexAttribL2d(index uint, x float64, y float64){
	_index := C.GLuint(index)
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	C.glVertexAttribL2d(_index, _x, _y)
}
func VertexAttribL3d(index uint, x float64, y float64, z float64){
	_index := C.GLuint(index)
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	_z := C.GLdouble(z)
	C.glVertexAttribL3d(_index, _x, _y, _z)
}
func VertexAttribL4d(index uint, x float64, y float64, z float64, w float64){
	_index := C.GLuint(index)
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	_z := C.GLdouble(z)
	_w := C.GLdouble(w)
	C.glVertexAttribL4d(_index, _x, _y, _z, _w)
}
/*func VertexAttribL1dv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribL1dv(_index, _v)
}*/
/*func VertexAttribL2dv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribL2dv(_index, _v)
}*/
/*func VertexAttribL3dv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribL3dv(_index, _v)
}*/
/*func VertexAttribL4dv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribL4dv(_index, _v)
}*/
/*func VertexAttribLPointer(index uint, size int, whichtype int, stride int, pointer ???){
	_index := C.GLuint(index)
	_size := C.GLint(size)
	_whichtype := C.GLenum(whichtype)
	_stride := C.GLsizei(stride)
	panic()
	C.glVertexAttribLPointer(_index, _size, _whichtype, _stride, _pointer)
}*/
/*func GetVertexAttribLdv(index uint, pname int, params ???){
	_index := C.GLuint(index)
	_pname := C.GLenum(pname)
	panic()
	C.glGetVertexAttribLdv(_index, _pname, _params)
}*/
/*func ViewportArrayv(first uint, count int, v ???){
	_first := C.GLuint(first)
	_count := C.GLsizei(count)
	panic()
	C.glViewportArrayv(_first, _count, _v)
}*/
func ViewportIndexedf(index uint, x float32, y float32, w float32, h float32){
	_index := C.GLuint(index)
	_x := C.GLfloat(x)
	_y := C.GLfloat(y)
	_w := C.GLfloat(w)
	_h := C.GLfloat(h)
	C.glViewportIndexedf(_index, _x, _y, _w, _h)
}
/*func ViewportIndexedfv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glViewportIndexedfv(_index, _v)
}*/
/*func ScissorArrayv(first uint, count int, v ???){
	_first := C.GLuint(first)
	_count := C.GLsizei(count)
	panic()
	C.glScissorArrayv(_first, _count, _v)
}*/
func ScissorIndexed(index uint, left int, bottom int, width int, height int){
	_index := C.GLuint(index)
	_left := C.GLint(left)
	_bottom := C.GLint(bottom)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	C.glScissorIndexed(_index, _left, _bottom, _width, _height)
}
/*func ScissorIndexedv(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glScissorIndexedv(_index, _v)
}*/
/*func DepthRangeArrayv(first uint, count int, v ???){
	_first := C.GLuint(first)
	_count := C.GLsizei(count)
	panic()
	C.glDepthRangeArrayv(_first, _count, _v)
}*/
func DepthRangeIndexed(index uint, n float64, f float64){
	_index := C.GLuint(index)
	_n := C.GLdouble(n)
	_f := C.GLdouble(f)
	C.glDepthRangeIndexed(_index, _n, _f)
}
/*func GetFloati_v(target int, index uint, data ???){
	_target := C.GLenum(target)
	_index := C.GLuint(index)
	panic()
	C.glGetFloati_v(_target, _index, _data)
}*/
/*func GetDoublei_v(target int, index uint, data ???){
	_target := C.GLenum(target)
	_index := C.GLuint(index)
	panic()
	C.glGetDoublei_v(_target, _index, _data)
}*/
func DrawArraysInstancedBaseInstance(mode int, first int, count int, instancecount int, baseinstance uint){
	_mode := C.GLenum(mode)
	_first := C.GLint(first)
	_count := C.GLsizei(count)
	_instancecount := C.GLsizei(instancecount)
	_baseinstance := C.GLuint(baseinstance)
	C.glDrawArraysInstancedBaseInstance(_mode, _first, _count, _instancecount, _baseinstance)
}
/*func DrawElementsInstancedBaseInstance(mode int, count int, whichtype int, indices ???, instancecount int, baseinstance uint){
	_mode := C.GLenum(mode)
	_count := C.GLsizei(count)
	_whichtype := C.GLenum(whichtype)
	panic()
	_instancecount := C.GLsizei(instancecount)
	_baseinstance := C.GLuint(baseinstance)
	C.glDrawElementsInstancedBaseInstance(_mode, _count, _whichtype, _indices, _instancecount, _baseinstance)
}*/
/*func DrawElementsInstancedBaseVertexBaseInstance(mode int, count int, whichtype int, indices ???, instancecount int, basevertex int, baseinstance uint){
	_mode := C.GLenum(mode)
	_count := C.GLsizei(count)
	_whichtype := C.GLenum(whichtype)
	panic()
	_instancecount := C.GLsizei(instancecount)
	_basevertex := C.GLint(basevertex)
	_baseinstance := C.GLuint(baseinstance)
	C.glDrawElementsInstancedBaseVertexBaseInstance(_mode, _count, _whichtype, _indices, _instancecount, _basevertex, _baseinstance)
}*/
/*func GetInternalformativ(target int, internalformat int, pname int, bufSize int, params ???){
	_target := C.GLenum(target)
	_internalformat := C.GLenum(internalformat)
	_pname := C.GLenum(pname)
	_bufSize := C.GLsizei(bufSize)
	panic()
	C.glGetInternalformativ(_target, _internalformat, _pname, _bufSize, _params)
}*/
/*func GetActiveAtomicCounterBufferiv(program uint, bufferIndex uint, pname int, params ???){
	_program := C.GLuint(program)
	_bufferIndex := C.GLuint(bufferIndex)
	_pname := C.GLenum(pname)
	panic()
	C.glGetActiveAtomicCounterBufferiv(_program, _bufferIndex, _pname, _params)
}*/
func BindImageTexture(unit uint, texture uint, level int, layered bool, layer int, access int, format int){
	_unit := C.GLuint(unit)
	_texture := C.GLuint(texture)
	_level := C.GLint(level)
	_layered := C.GLboolean(1);if !layered{_layered=C.GLboolean(0)}
	_layer := C.GLint(layer)
	_access := C.GLenum(access)
	_format := C.GLenum(format)
	C.glBindImageTexture(_unit, _texture, _level, _layered, _layer, _access, _format)
}
func MemoryBarrier(barriers uint){
	_barriers := C.GLbitfield(barriers)
	C.glMemoryBarrier(_barriers)
}
func TexStorage1D(target int, levels int, internalformat int, width int){
	_target := C.GLenum(target)
	_levels := C.GLsizei(levels)
	_internalformat := C.GLenum(internalformat)
	_width := C.GLsizei(width)
	C.glTexStorage1D(_target, _levels, _internalformat, _width)
}
func TexStorage2D(target int, levels int, internalformat int, width int, height int){
	_target := C.GLenum(target)
	_levels := C.GLsizei(levels)
	_internalformat := C.GLenum(internalformat)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	C.glTexStorage2D(_target, _levels, _internalformat, _width, _height)
}
func TexStorage3D(target int, levels int, internalformat int, width int, height int, depth int){
	_target := C.GLenum(target)
	_levels := C.GLsizei(levels)
	_internalformat := C.GLenum(internalformat)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_depth := C.GLsizei(depth)
	C.glTexStorage3D(_target, _levels, _internalformat, _width, _height, _depth)
}
func DrawTransformFeedbackInstanced(mode int, id uint, instancecount int){
	_mode := C.GLenum(mode)
	_id := C.GLuint(id)
	_instancecount := C.GLsizei(instancecount)
	C.glDrawTransformFeedbackInstanced(_mode, _id, _instancecount)
}
func DrawTransformFeedbackStreamInstanced(mode int, id uint, stream uint, instancecount int){
	_mode := C.GLenum(mode)
	_id := C.GLuint(id)
	_stream := C.GLuint(stream)
	_instancecount := C.GLsizei(instancecount)
	C.glDrawTransformFeedbackStreamInstanced(_mode, _id, _stream, _instancecount)
}
/*func ClearBufferData(target int, internalformat int, format int, whichtype int, data ???){
	_target := C.GLenum(target)
	_internalformat := C.GLenum(internalformat)
	_format := C.GLenum(format)
	_whichtype := C.GLenum(whichtype)
	panic()
	C.glClearBufferData(_target, _internalformat, _format, _whichtype, _data)
}*/
/*func ClearBufferSubData(target int, internalformat int, offset ???, size int, format int, whichtype int, data ???){
	_target := C.GLenum(target)
	_internalformat := C.GLenum(internalformat)
	panic()
	_size := C.GLsizeiptr(size)
	_format := C.GLenum(format)
	_whichtype := C.GLenum(whichtype)
	panic()
	C.glClearBufferSubData(_target, _internalformat, _offset, _size, _format, _whichtype, _data)
}*/
func DispatchCompute(num_groups_x uint, num_groups_y uint, num_groups_z uint){
	_num_groups_x := C.GLuint(num_groups_x)
	_num_groups_y := C.GLuint(num_groups_y)
	_num_groups_z := C.GLuint(num_groups_z)
	C.glDispatchCompute(_num_groups_x, _num_groups_y, _num_groups_z)
}
/*func DispatchComputeIndirect(indirect ???){
	panic()
	C.glDispatchComputeIndirect(_indirect)
}*/
func CopyImageSubData(srcName uint, srcTarget int, srcLevel int, srcX int, srcY int, srcZ int, dstName uint, dstTarget int, dstLevel int, dstX int, dstY int, dstZ int, srcWidth int, srcHeight int, srcDepth int){
	_srcName := C.GLuint(srcName)
	_srcTarget := C.GLenum(srcTarget)
	_srcLevel := C.GLint(srcLevel)
	_srcX := C.GLint(srcX)
	_srcY := C.GLint(srcY)
	_srcZ := C.GLint(srcZ)
	_dstName := C.GLuint(dstName)
	_dstTarget := C.GLenum(dstTarget)
	_dstLevel := C.GLint(dstLevel)
	_dstX := C.GLint(dstX)
	_dstY := C.GLint(dstY)
	_dstZ := C.GLint(dstZ)
	_srcWidth := C.GLsizei(srcWidth)
	_srcHeight := C.GLsizei(srcHeight)
	_srcDepth := C.GLsizei(srcDepth)
	C.glCopyImageSubData(_srcName, _srcTarget, _srcLevel, _srcX, _srcY, _srcZ, _dstName, _dstTarget, _dstLevel, _dstX, _dstY, _dstZ, _srcWidth, _srcHeight, _srcDepth)
}
func FramebufferParameteri(target int, pname int, param int){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	_param := C.GLint(param)
	C.glFramebufferParameteri(_target, _pname, _param)
}
/*func GetFramebufferParameteriv(target int, pname int, params ???){
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	panic()
	C.glGetFramebufferParameteriv(_target, _pname, _params)
}*/
/*func GetInternalformati64v(target int, internalformat int, pname int, bufSize int, params ???){
	_target := C.GLenum(target)
	_internalformat := C.GLenum(internalformat)
	_pname := C.GLenum(pname)
	_bufSize := C.GLsizei(bufSize)
	panic()
	C.glGetInternalformati64v(_target, _internalformat, _pname, _bufSize, _params)
}*/
func InvalidateTexSubImage(texture uint, level int, xoffset int, yoffset int, zoffset int, width int, height int, depth int){
	_texture := C.GLuint(texture)
	_level := C.GLint(level)
	_xoffset := C.GLint(xoffset)
	_yoffset := C.GLint(yoffset)
	_zoffset := C.GLint(zoffset)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_depth := C.GLsizei(depth)
	C.glInvalidateTexSubImage(_texture, _level, _xoffset, _yoffset, _zoffset, _width, _height, _depth)
}
func InvalidateTexImage(texture uint, level int){
	_texture := C.GLuint(texture)
	_level := C.GLint(level)
	C.glInvalidateTexImage(_texture, _level)
}
/*func InvalidateBufferSubData(buffer uint, offset ???, length int){
	_buffer := C.GLuint(buffer)
	panic()
	_length := C.GLsizeiptr(length)
	C.glInvalidateBufferSubData(_buffer, _offset, _length)
}*/
func InvalidateBufferData(buffer uint){
	_buffer := C.GLuint(buffer)
	C.glInvalidateBufferData(_buffer)
}
/*func InvalidateFramebuffer(target int, numAttachments int, attachments ???){
	_target := C.GLenum(target)
	_numAttachments := C.GLsizei(numAttachments)
	panic()
	C.glInvalidateFramebuffer(_target, _numAttachments, _attachments)
}*/
/*func InvalidateSubFramebuffer(target int, numAttachments int, attachments ???, x int, y int, width int, height int){
	_target := C.GLenum(target)
	_numAttachments := C.GLsizei(numAttachments)
	panic()
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	C.glInvalidateSubFramebuffer(_target, _numAttachments, _attachments, _x, _y, _width, _height)
}*/
/*func MultiDrawArraysIndirect(mode int, indirect ???, drawcount int, stride int){
	_mode := C.GLenum(mode)
	panic()
	_drawcount := C.GLsizei(drawcount)
	_stride := C.GLsizei(stride)
	C.glMultiDrawArraysIndirect(_mode, _indirect, _drawcount, _stride)
}*/
/*func MultiDrawElementsIndirect(mode int, whichtype int, indirect ???, drawcount int, stride int){
	_mode := C.GLenum(mode)
	_whichtype := C.GLenum(whichtype)
	panic()
	_drawcount := C.GLsizei(drawcount)
	_stride := C.GLsizei(stride)
	C.glMultiDrawElementsIndirect(_mode, _whichtype, _indirect, _drawcount, _stride)
}*/
/*func GetProgramInterfaceiv(program uint, programInterface int, pname int, params ???){
	_program := C.GLuint(program)
	_programInterface := C.GLenum(programInterface)
	_pname := C.GLenum(pname)
	panic()
	C.glGetProgramInterfaceiv(_program, _programInterface, _pname, _params)
}*/
/*func GetProgramResourceIndex(program uint, programInterface int, name ???)uint{
	_program := C.GLuint(program)
	_programInterface := C.GLenum(programInterface)
	panic()
	returnvalue := C.glGetProgramResourceIndex(_program, _programInterface, _name)
	panic()
return convreturnvalue
}*/
/*func GetProgramResourceName(program uint, programInterface int, index uint, bufSize int, length int, name int){
	_program := C.GLuint(program)
	_programInterface := C.GLenum(programInterface)
	_index := C.GLuint(index)
	_bufSize := C.GLsizei(bufSize)
	_length := C.GLsizei(length)
	panic()
	C.glGetProgramResourceName(_program, _programInterface, _index, _bufSize, _length, _name)
}*/
/*func GetProgramResourceiv(program uint, programInterface int, index uint, propCount int, props ???, bufSize int, length int, params ???){
	_program := C.GLuint(program)
	_programInterface := C.GLenum(programInterface)
	_index := C.GLuint(index)
	_propCount := C.GLsizei(propCount)
	panic()
	_bufSize := C.GLsizei(bufSize)
	_length := C.GLsizei(length)
	panic()
	C.glGetProgramResourceiv(_program, _programInterface, _index, _propCount, _props, _bufSize, _length, _params)
}*/
/*func GetProgramResourceLocation(program uint, programInterface int, name ???)int{
	_program := C.GLuint(program)
	_programInterface := C.GLenum(programInterface)
	panic()
	returnvalue := C.glGetProgramResourceLocation(_program, _programInterface, _name)
	panic()
return convreturnvalue
}*/
/*func GetProgramResourceLocationIndex(program uint, programInterface int, name ???)int{
	_program := C.GLuint(program)
	_programInterface := C.GLenum(programInterface)
	panic()
	returnvalue := C.glGetProgramResourceLocationIndex(_program, _programInterface, _name)
	panic()
return convreturnvalue
}*/
func ShaderStorageBlockBinding(program uint, storageBlockIndex uint, storageBlockBinding uint){
	_program := C.GLuint(program)
	_storageBlockIndex := C.GLuint(storageBlockIndex)
	_storageBlockBinding := C.GLuint(storageBlockBinding)
	C.glShaderStorageBlockBinding(_program, _storageBlockIndex, _storageBlockBinding)
}
/*func TexBufferRange(target int, internalformat int, buffer uint, offset ???, size int){
	_target := C.GLenum(target)
	_internalformat := C.GLenum(internalformat)
	_buffer := C.GLuint(buffer)
	panic()
	_size := C.GLsizeiptr(size)
	C.glTexBufferRange(_target, _internalformat, _buffer, _offset, _size)
}*/
func TexStorage2DMultisample(target int, samples int, internalformat int, width int, height int, fixedsamplelocations bool){
	_target := C.GLenum(target)
	_samples := C.GLsizei(samples)
	_internalformat := C.GLenum(internalformat)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_fixedsamplelocations := C.GLboolean(1);if !fixedsamplelocations{_fixedsamplelocations=C.GLboolean(0)}
	C.glTexStorage2DMultisample(_target, _samples, _internalformat, _width, _height, _fixedsamplelocations)
}
func TexStorage3DMultisample(target int, samples int, internalformat int, width int, height int, depth int, fixedsamplelocations bool){
	_target := C.GLenum(target)
	_samples := C.GLsizei(samples)
	_internalformat := C.GLenum(internalformat)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_depth := C.GLsizei(depth)
	_fixedsamplelocations := C.GLboolean(1);if !fixedsamplelocations{_fixedsamplelocations=C.GLboolean(0)}
	C.glTexStorage3DMultisample(_target, _samples, _internalformat, _width, _height, _depth, _fixedsamplelocations)
}
func TextureView(texture uint, target int, origtexture uint, internalformat int, minlevel uint, numlevels uint, minlayer uint, numlayers uint){
	_texture := C.GLuint(texture)
	_target := C.GLenum(target)
	_origtexture := C.GLuint(origtexture)
	_internalformat := C.GLenum(internalformat)
	_minlevel := C.GLuint(minlevel)
	_numlevels := C.GLuint(numlevels)
	_minlayer := C.GLuint(minlayer)
	_numlayers := C.GLuint(numlayers)
	C.glTextureView(_texture, _target, _origtexture, _internalformat, _minlevel, _numlevels, _minlayer, _numlayers)
}
/*func BindVertexBuffer(bindingindex uint, buffer uint, offset ???, stride int){
	_bindingindex := C.GLuint(bindingindex)
	_buffer := C.GLuint(buffer)
	panic()
	_stride := C.GLsizei(stride)
	C.glBindVertexBuffer(_bindingindex, _buffer, _offset, _stride)
}*/
func VertexAttribFormat(attribindex uint, size int, whichtype int, normalized bool, relativeoffset uint){
	_attribindex := C.GLuint(attribindex)
	_size := C.GLint(size)
	_whichtype := C.GLenum(whichtype)
	_normalized := C.GLboolean(1);if !normalized{_normalized=C.GLboolean(0)}
	_relativeoffset := C.GLuint(relativeoffset)
	C.glVertexAttribFormat(_attribindex, _size, _whichtype, _normalized, _relativeoffset)
}
func VertexAttribIFormat(attribindex uint, size int, whichtype int, relativeoffset uint){
	_attribindex := C.GLuint(attribindex)
	_size := C.GLint(size)
	_whichtype := C.GLenum(whichtype)
	_relativeoffset := C.GLuint(relativeoffset)
	C.glVertexAttribIFormat(_attribindex, _size, _whichtype, _relativeoffset)
}
func VertexAttribLFormat(attribindex uint, size int, whichtype int, relativeoffset uint){
	_attribindex := C.GLuint(attribindex)
	_size := C.GLint(size)
	_whichtype := C.GLenum(whichtype)
	_relativeoffset := C.GLuint(relativeoffset)
	C.glVertexAttribLFormat(_attribindex, _size, _whichtype, _relativeoffset)
}
func VertexAttribBinding(attribindex uint, bindingindex uint){
	_attribindex := C.GLuint(attribindex)
	_bindingindex := C.GLuint(bindingindex)
	C.glVertexAttribBinding(_attribindex, _bindingindex)
}
func VertexBindingDivisor(bindingindex uint, divisor uint){
	_bindingindex := C.GLuint(bindingindex)
	_divisor := C.GLuint(divisor)
	C.glVertexBindingDivisor(_bindingindex, _divisor)
}
/*func DebugMessageControl(source int, whichtype int, severity int, count int, ids ???, enabled bool){
	_source := C.GLenum(source)
	_whichtype := C.GLenum(whichtype)
	_severity := C.GLenum(severity)
	_count := C.GLsizei(count)
	panic()
	_enabled := C.GLboolean(1);if !enabled{_enabled=C.GLboolean(0)}
	C.glDebugMessageControl(_source, _whichtype, _severity, _count, _ids, _enabled)
}*/
/*func DebugMessageInsert(source int, whichtype int, id uint, severity int, length int, buf ???){
	_source := C.GLenum(source)
	_whichtype := C.GLenum(whichtype)
	_id := C.GLuint(id)
	_severity := C.GLenum(severity)
	_length := C.GLsizei(length)
	panic()
	C.glDebugMessageInsert(_source, _whichtype, _id, _severity, _length, _buf)
}*/
/*func DebugMessageCallback(callback ???, userParam ???){
	panic()
	panic()
	C.glDebugMessageCallback(_callback, _userParam)
}*/
/*func GetDebugMessageLog(count uint, bufSize int, sources ???, types ???, ids ???, severities ???, lengths int, messageLog int)uint{
	_count := C.GLuint(count)
	_bufSize := C.GLsizei(bufSize)
	panic()
	panic()
	panic()
	panic()
	_lengths := C.GLsizei(lengths)
	panic()
	returnvalue := C.glGetDebugMessageLog(_count, _bufSize, _sources, _types, _ids, _severities, _lengths, _messageLog)
	panic()
return convreturnvalue
}*/
/*func PushDebugGroup(source int, id uint, length int, message ???){
	_source := C.GLenum(source)
	_id := C.GLuint(id)
	_length := C.GLsizei(length)
	panic()
	C.glPushDebugGroup(_source, _id, _length, _message)
}*/
func PopDebugGroup(){
	C.glPopDebugGroup()
}
/*func ObjectLabel(identifier int, name uint, length int, label ???){
	_identifier := C.GLenum(identifier)
	_name := C.GLuint(name)
	_length := C.GLsizei(length)
	panic()
	C.glObjectLabel(_identifier, _name, _length, _label)
}*/
/*func GetObjectLabel(identifier int, name uint, bufSize int, length int, label int){
	_identifier := C.GLenum(identifier)
	_name := C.GLuint(name)
	_bufSize := C.GLsizei(bufSize)
	_length := C.GLsizei(length)
	panic()
	C.glGetObjectLabel(_identifier, _name, _bufSize, _length, _label)
}*/
/*func ObjectPtrLabel(ptr ???, length int, label ???){
	panic()
	_length := C.GLsizei(length)
	panic()
	C.glObjectPtrLabel(_ptr, _length, _label)
}*/
/*func GetObjectPtrLabel(ptr ???, bufSize int, length int, label int){
	panic()
	_bufSize := C.GLsizei(bufSize)
	_length := C.GLsizei(length)
	panic()
	C.glGetObjectPtrLabel(_ptr, _bufSize, _length, _label)
}*/
/*func BufferStorage(target int, size int, data ???, flags uint){
	_target := C.GLenum(target)
	_size := C.GLsizeiptr(size)
	panic()
	_flags := C.GLbitfield(flags)
	C.glBufferStorage(_target, _size, _data, _flags)
}*/
/*func ClearTexImage(texture uint, level int, format int, whichtype int, data ???){
	_texture := C.GLuint(texture)
	_level := C.GLint(level)
	_format := C.GLenum(format)
	_whichtype := C.GLenum(whichtype)
	panic()
	C.glClearTexImage(_texture, _level, _format, _whichtype, _data)
}*/
/*func ClearTexSubImage(texture uint, level int, xoffset int, yoffset int, zoffset int, width int, height int, depth int, format int, whichtype int, data ???){
	_texture := C.GLuint(texture)
	_level := C.GLint(level)
	_xoffset := C.GLint(xoffset)
	_yoffset := C.GLint(yoffset)
	_zoffset := C.GLint(zoffset)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_depth := C.GLsizei(depth)
	_format := C.GLenum(format)
	_whichtype := C.GLenum(whichtype)
	panic()
	C.glClearTexSubImage(_texture, _level, _xoffset, _yoffset, _zoffset, _width, _height, _depth, _format, _whichtype, _data)
}*/
/*func BindBuffersBase(target int, first uint, count int, buffers ???){
	_target := C.GLenum(target)
	_first := C.GLuint(first)
	_count := C.GLsizei(count)
	panic()
	C.glBindBuffersBase(_target, _first, _count, _buffers)
}*/
/*func BindBuffersRange(target int, first uint, count int, buffers ???, offsets ???, sizes ???){
	_target := C.GLenum(target)
	_first := C.GLuint(first)
	_count := C.GLsizei(count)
	panic()
	panic()
	panic()
	C.glBindBuffersRange(_target, _first, _count, _buffers, _offsets, _sizes)
}*/
/*func BindTextures(first uint, count int, textures ???){
	_first := C.GLuint(first)
	_count := C.GLsizei(count)
	panic()
	C.glBindTextures(_first, _count, _textures)
}*/
/*func BindSamplers(first uint, count int, samplers ???){
	_first := C.GLuint(first)
	_count := C.GLsizei(count)
	panic()
	C.glBindSamplers(_first, _count, _samplers)
}*/
/*func BindImageTextures(first uint, count int, textures ???){
	_first := C.GLuint(first)
	_count := C.GLsizei(count)
	panic()
	C.glBindImageTextures(_first, _count, _textures)
}*/
/*func BindVertexBuffers(first uint, count int, buffers ???, offsets ???, strides ???){
	_first := C.GLuint(first)
	_count := C.GLsizei(count)
	panic()
	panic()
	panic()
	C.glBindVertexBuffers(_first, _count, _buffers, _offsets, _strides)
}*/
/*func GetTextureHandleARB(texture uint)???{
	_texture := C.GLuint(texture)
	returnvalue := C.glGetTextureHandleARB(_texture)
	panic()
return convreturnvalue
}*/
/*func GetTextureSamplerHandleARB(texture uint, sampler uint)???{
	_texture := C.GLuint(texture)
	_sampler := C.GLuint(sampler)
	returnvalue := C.glGetTextureSamplerHandleARB(_texture, _sampler)
	panic()
return convreturnvalue
}*/
/*func MakeTextureHandleResidentARB(handle ???){
	panic()
	C.glMakeTextureHandleResidentARB(_handle)
}*/
/*func MakeTextureHandleNonResidentARB(handle ???){
	panic()
	C.glMakeTextureHandleNonResidentARB(_handle)
}*/
/*func GetImageHandleARB(texture uint, level int, layered bool, layer int, format int)???{
	_texture := C.GLuint(texture)
	_level := C.GLint(level)
	_layered := C.GLboolean(1);if !layered{_layered=C.GLboolean(0)}
	_layer := C.GLint(layer)
	_format := C.GLenum(format)
	returnvalue := C.glGetImageHandleARB(_texture, _level, _layered, _layer, _format)
	panic()
return convreturnvalue
}*/
/*func MakeImageHandleResidentARB(handle ???, access int){
	panic()
	_access := C.GLenum(access)
	C.glMakeImageHandleResidentARB(_handle, _access)
}*/
/*func MakeImageHandleNonResidentARB(handle ???){
	panic()
	C.glMakeImageHandleNonResidentARB(_handle)
}*/
/*func UniformHandleui64ARB(location int, value ???){
	_location := C.GLint(location)
	panic()
	C.glUniformHandleui64ARB(_location, _value)
}*/
/*func UniformHandleui64vARB(location int, count int, value ???){
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glUniformHandleui64vARB(_location, _count, _value)
}*/
/*func ProgramUniformHandleui64ARB(program uint, location int, value ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	panic()
	C.glProgramUniformHandleui64ARB(_program, _location, _value)
}*/
/*func ProgramUniformHandleui64vARB(program uint, location int, count int, values ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_count := C.GLsizei(count)
	panic()
	C.glProgramUniformHandleui64vARB(_program, _location, _count, _values)
}*/
/*func IsTextureHandleResidentARB(handle ???)bool{
	panic()
	returnvalue := C.glIsTextureHandleResidentARB(_handle)
	panic()
return convreturnvalue
}*/
/*func IsImageHandleResidentARB(handle ???)bool{
	panic()
	returnvalue := C.glIsImageHandleResidentARB(_handle)
	panic()
return convreturnvalue
}*/
/*func VertexAttribL1ui64ARB(index uint, x ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribL1ui64ARB(_index, _x)
}*/
/*func VertexAttribL1ui64vARB(index uint, v ???){
	_index := C.GLuint(index)
	panic()
	C.glVertexAttribL1ui64vARB(_index, _v)
}*/
/*func GetVertexAttribLui64vARB(index uint, pname int, params ???){
	_index := C.GLuint(index)
	_pname := C.GLenum(pname)
	panic()
	C.glGetVertexAttribLui64vARB(_index, _pname, _params)
}*/
/*func CreateSyncFromCLeventARB(context ???, event ???, flags uint)???{
	panic()
	panic()
	_flags := C.GLbitfield(flags)
	returnvalue := C.glCreateSyncFromCLeventARB(_context, _event, _flags)
	panic()
return convreturnvalue
}*/
/*func DispatchComputeGroupSizeARB(num_groups_x uint, num_groups_y uint, num_groups_z uint, group_size_x uint, group_size_y uint, group_size_z uint){
	_num_groups_x := C.GLuint(num_groups_x)
	_num_groups_y := C.GLuint(num_groups_y)
	_num_groups_z := C.GLuint(num_groups_z)
	_group_size_x := C.GLuint(group_size_x)
	_group_size_y := C.GLuint(group_size_y)
	_group_size_z := C.GLuint(group_size_z)
	C.glDispatchComputeGroupSizeARB(_num_groups_x, _num_groups_y, _num_groups_z, _group_size_x, _group_size_y, _group_size_z)
}*/
/*func DebugMessageControlARB(source int, whichtype int, severity int, count int, ids ???, enabled bool){
	_source := C.GLenum(source)
	_whichtype := C.GLenum(whichtype)
	_severity := C.GLenum(severity)
	_count := C.GLsizei(count)
	panic()
	_enabled := C.GLboolean(1);if !enabled{_enabled=C.GLboolean(0)}
	C.glDebugMessageControlARB(_source, _whichtype, _severity, _count, _ids, _enabled)
}*/
/*func DebugMessageInsertARB(source int, whichtype int, id uint, severity int, length int, buf ???){
	_source := C.GLenum(source)
	_whichtype := C.GLenum(whichtype)
	_id := C.GLuint(id)
	_severity := C.GLenum(severity)
	_length := C.GLsizei(length)
	panic()
	C.glDebugMessageInsertARB(_source, _whichtype, _id, _severity, _length, _buf)
}*/
/*func DebugMessageCallbackARB(callback ???, userParam ???){
	panic()
	panic()
	C.glDebugMessageCallbackARB(_callback, _userParam)
}*/
/*func GetDebugMessageLogARB(count uint, bufSize int, sources ???, types ???, ids ???, severities ???, lengths int, messageLog int)uint{
	_count := C.GLuint(count)
	_bufSize := C.GLsizei(bufSize)
	panic()
	panic()
	panic()
	panic()
	_lengths := C.GLsizei(lengths)
	panic()
	returnvalue := C.glGetDebugMessageLogARB(_count, _bufSize, _sources, _types, _ids, _severities, _lengths, _messageLog)
	panic()
return convreturnvalue
}*/
func BlendEquationiARB(buf uint, mode int){
	_buf := C.GLuint(buf)
	_mode := C.GLenum(mode)
	C.glBlendEquationiARB(_buf, _mode)
}
func BlendEquationSeparateiARB(buf uint, modeRGB int, modeAlpha int){
	_buf := C.GLuint(buf)
	_modeRGB := C.GLenum(modeRGB)
	_modeAlpha := C.GLenum(modeAlpha)
	C.glBlendEquationSeparateiARB(_buf, _modeRGB, _modeAlpha)
}
func BlendFunciARB(buf uint, src int, dst int){
	_buf := C.GLuint(buf)
	_src := C.GLenum(src)
	_dst := C.GLenum(dst)
	C.glBlendFunciARB(_buf, _src, _dst)
}
func BlendFuncSeparateiARB(buf uint, srcRGB int, dstRGB int, srcAlpha int, dstAlpha int){
	_buf := C.GLuint(buf)
	_srcRGB := C.GLenum(srcRGB)
	_dstRGB := C.GLenum(dstRGB)
	_srcAlpha := C.GLenum(srcAlpha)
	_dstAlpha := C.GLenum(dstAlpha)
	C.glBlendFuncSeparateiARB(_buf, _srcRGB, _dstRGB, _srcAlpha, _dstAlpha)
}
/*func MultiDrawArraysIndirectCountARB(mode int, indirect ???, drawcount ???, maxdrawcount int, stride int){
	_mode := C.GLenum(mode)
	panic()
	panic()
	_maxdrawcount := C.GLsizei(maxdrawcount)
	_stride := C.GLsizei(stride)
	C.glMultiDrawArraysIndirectCountARB(_mode, _indirect, _drawcount, _maxdrawcount, _stride)
}*/
/*func MultiDrawElementsIndirectCountARB(mode int, whichtype int, indirect ???, drawcount ???, maxdrawcount int, stride int){
	_mode := C.GLenum(mode)
	_whichtype := C.GLenum(whichtype)
	panic()
	panic()
	_maxdrawcount := C.GLsizei(maxdrawcount)
	_stride := C.GLsizei(stride)
	C.glMultiDrawElementsIndirectCountARB(_mode, _whichtype, _indirect, _drawcount, _maxdrawcount, _stride)
}*/
func GetGraphicsResetStatusARB()int{
	returnvalue := C.glGetGraphicsResetStatusARB()
	convreturnvalue := int(returnvalue)
return convreturnvalue
}
/*func GetnTexImageARB(target int, level int, format int, whichtype int, bufSize int, img ???){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_format := C.GLenum(format)
	_whichtype := C.GLenum(whichtype)
	_bufSize := C.GLsizei(bufSize)
	panic()
	C.glGetnTexImageARB(_target, _level, _format, _whichtype, _bufSize, _img)
}*/
/*func ReadnPixelsARB(x int, y int, width int, height int, format int, whichtype int, bufSize int, data ???){
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_format := C.GLenum(format)
	_whichtype := C.GLenum(whichtype)
	_bufSize := C.GLsizei(bufSize)
	panic()
	C.glReadnPixelsARB(_x, _y, _width, _height, _format, _whichtype, _bufSize, _data)
}*/
/*func GetnCompressedTexImageARB(target int, lod int, bufSize int, img ???){
	_target := C.GLenum(target)
	_lod := C.GLint(lod)
	_bufSize := C.GLsizei(bufSize)
	panic()
	C.glGetnCompressedTexImageARB(_target, _lod, _bufSize, _img)
}*/
/*func GetnUniformfvARB(program uint, location int, bufSize int, params ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_bufSize := C.GLsizei(bufSize)
	panic()
	C.glGetnUniformfvARB(_program, _location, _bufSize, _params)
}*/
/*func GetnUniformivARB(program uint, location int, bufSize int, params ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_bufSize := C.GLsizei(bufSize)
	panic()
	C.glGetnUniformivARB(_program, _location, _bufSize, _params)
}*/
/*func GetnUniformuivARB(program uint, location int, bufSize int, params ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_bufSize := C.GLsizei(bufSize)
	panic()
	C.glGetnUniformuivARB(_program, _location, _bufSize, _params)
}*/
/*func GetnUniformdvARB(program uint, location int, bufSize int, params ???){
	_program := C.GLuint(program)
	_location := C.GLint(location)
	_bufSize := C.GLsizei(bufSize)
	panic()
	C.glGetnUniformdvARB(_program, _location, _bufSize, _params)
}*/
func MinSampleShadingARB(value float32){
	_value := C.GLfloat(value)
	C.glMinSampleShadingARB(_value)
}
/*func NamedStringARB(whichtype int, namelen int, name ???, stringlen int, string ???){
	_whichtype := C.GLenum(whichtype)
	_namelen := C.GLint(namelen)
	panic()
	_stringlen := C.GLint(stringlen)
	panic()
	C.glNamedStringARB(_whichtype, _namelen, _name, _stringlen, _string)
}*/
/*func DeleteNamedStringARB(namelen int, name ???){
	_namelen := C.GLint(namelen)
	panic()
	C.glDeleteNamedStringARB(_namelen, _name)
}*/
/*func CompileShaderIncludeARB(shader uint, count int, path ???, length ???){
	_shader := C.GLuint(shader)
	_count := C.GLsizei(count)
	panic()
	panic()
	C.glCompileShaderIncludeARB(_shader, _count, _path, _length)
}*/
/*func IsNamedStringARB(namelen int, name ???)bool{
	_namelen := C.GLint(namelen)
	panic()
	returnvalue := C.glIsNamedStringARB(_namelen, _name)
	panic()
return convreturnvalue
}*/
/*func GetNamedStringARB(namelen int, name ???, bufSize int, stringlen ???, string int){
	_namelen := C.GLint(namelen)
	panic()
	_bufSize := C.GLsizei(bufSize)
	panic()
	panic()
	C.glGetNamedStringARB(_namelen, _name, _bufSize, _stringlen, _string)
}*/
/*func GetNamedStringivARB(namelen int, name ???, pname int, params ???){
	_namelen := C.GLint(namelen)
	panic()
	_pname := C.GLenum(pname)
	panic()
	C.glGetNamedStringivARB(_namelen, _name, _pname, _params)
}*/
/*func TexPageCommitmentARB(target int, level int, xoffset int, yoffset int, zoffset int, width int, height int, depth int, resident bool){
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_xoffset := C.GLint(xoffset)
	_yoffset := C.GLint(yoffset)
	_zoffset := C.GLint(zoffset)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_depth := C.GLsizei(depth)
	_resident := C.GLboolean(1);if !resident{_resident=C.GLboolean(0)}
	C.glTexPageCommitmentARB(_target, _level, _xoffset, _yoffset, _zoffset, _width, _height, _depth, _resident)
}*/
//failed:  376
//succeeded:  211
