#include "funcdata.h"
#include "textflag.h"

// func getg() unsafe.Pointer
TEXT ·getG(SB), NOSPLIT, $0-8
    MOVQ (TLS), AX
    MOVQ AX, ret+0(FP)
    RET


// func getg() interface{}
TEXT ·getGInterface(SB), NOSPLIT, $32-16
    NO_LOCAL_POINTERS

    MOVQ $0, ret_type+0(FP)
    MOVQ $0, ret_data+8(FP)
    GO_RESULTS_INITIALIZED

    // get runtime.g
    MOVQ (TLS), AX

    // get runtime.g type
    MOVQ $type·runtime·g(SB), BX

    // convert (*g) to interface{}
    MOVQ AX, 8(SP)
    MOVQ BX, 0(SP)
    CALL runtime·convT2E(SB)
    MOVQ 16(SP), AX
    MOVQ 24(SP), BX

    // return interface{}
    MOVQ AX, ret_type+0(FP)
    MOVQ BX, ret_data+8(FP)
    RET
