### REF: https://golang.org/doc/asm
### REF: https://blog.hackercat.ninja/post/quick_intro_to_go_assembly/

~~~
---------------------------------------
|  NASM x86  | NASM x86_64 |  Go ASM  |
---------------------------------------
|  eax       | rax         |  AX      |
---------------------------------------
|  ebx       | rbx         |  BX      |
---------------------------------------
|  ...       | ...         | ...      |
---------------------------------------


Go all architectures:

 - FP: Frame Pointer.
 - PC: Program Counter.
 - SB: Static Base pointer.
 - SP: Stack Pointer.

   Estos registros virtuales estan jugando
un papel muy importante en Go ASM, y se
usan constantemente, los mas importantes
son SB y FP.

   El pseudo registro SB puede considerarse
como el origen de la memoria, por lo que el
simbolo foo(SB) es el nombre foo como una
direccion de memoria. Esta sintaxis tiene
dos modificadores basicos, <> y +N donde N
es un numero entero. El primero foo<>(SB)
significa un elemento privado, al que solo
puede acceder desde el mismo archivo fuente,
como un nombre en minuscula en Go, el segundo
se usa para agregar un desplazamiento a la
direccion relativa a la que se refiere el
nombre, por lo que foo+8(SB) hay 8 bytes
despues del inicio de foo.

   El FP pseudo registro es un puntero de
trama virtual utilizado para referirse a los
argumentos del procedimiento, el compilador
mantiene estas referencias haciendo referencia
a los argumentos en la pila como compensaciones
del pseudo registro. En una maquina de 64bits,
0(FP) es el primer argumento, 8(FP) el segundo
y asi sucesivamente.

Para hacer referencia a estos argumentos, el
compilador impone el uso de un nombre para ello,
en aras de la claridad y la legibilidad, por lo
que MOVL foo+0(FP), CX se;ala el primer argumento
del FP registro virtual al registro fisico CX.


-------------------------------------------------------------
|  Intel            |  AT&T              |   Go             |
-------------------------------------------------------------
| mov eax, 1        | movl $1,      %eax | MOVQ $1,      AX |
-------------------------------------------------------------
| mov rbx, 0ffh     | movl $0xff,   %rbx | MOVQ $(0xff), BX |
-------------------------------------------------------------
| mov ecx, [ebx+3]  | movl 3(%ebx), %ecx | MOVQ 3(BX),   CX |
-------------------------------------------------------------

~~~

## Compilar
~~~
$ go tool asm asm_amd64.s
$ go tool compile -dolinkobj asm.go
$ go tool pack c asm.a asm.o asm_amd64.o
$ go tool link -o asm asm.a
~~~
