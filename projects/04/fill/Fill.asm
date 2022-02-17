// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.
(MAIN)
    @i
    M=0

    @8191
    D=A
    @n
    M=D

    @KBD
    D=M

    @BLACKSCREEN
    D;JGT

    @WHITESCREEN
    0;JMP

(BLACKSCREEN)
    @i
    D=M
    @n
    D=D-M
    @MAIN
    D;JGT

    @SCREEN
    D=A

    @i
    A=D+M
    M=-1

    @i
    M=M+1

    @BLACKSCREEN
    0;JMP

(WHITESCREEN)
    @i
    D=M
    @n
    D=D-M
    @MAIN
    D;JGT

    @SCREEN
    D=A

    @i
    A=D+M
    M=0

    @i
    M=M+1

    @WHITESCREEN
    0;JMP