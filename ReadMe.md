# Writing an interpreter in Go

Repository contains code for my implmentation of a basic tree-walk-interpreter written in Go, referenced from the text [W.A.I.I.G. by Thorsten Ball](https://interpreterbook.com/)

> It’s difficult to make generic statements about interpreters since the variety is so high and none are alike. What can be said is that the one fundamental attribute they all share is that they **take source code and evaluate it without producing some visible, intermediate result that can later be executed.** That’s in contrast to compilers, which take source code and produce output in another language that the underlying system can understand.

- VARIETY
    - some interpreter may skip the parsing stage entirely, like brainfuck
    - more elaborate interpreter like java make bytecode out of input and evaluate this
    - even more advanced are JIT interpreters, compile input just in time to machine code
    - tree walking interpreters, parse source code and build abstract tree out of it, then walk this tree
- IMPLEMENTATION (tree walk bottom-up)
    - **the lexer** (Lexing done by a lexer, tokenizer, scanner. Basically identiyfing what the raw input contains and change the representation of source code from `TEXT -> TOKENS`)
    - the parser
    - the Abstract Syntax Tree (AST)
    - the internal object system
    - the evaluator