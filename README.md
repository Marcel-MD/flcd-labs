# Laboratories on Formal Languages and Compiler Design

## Table of Contents

- [Introduction](#introduction)
- [Literature to read](#literature-to-read)
- [Laboratory tasks](#laboratory-tasks)
  - [Lab 1 - Regular Grammar to Finite Automata](#lab-1---regular-grammar-to-finite-automata)
  - [Lab 2 - NFA to DFA conversion](#lab-2---nfa-to-dfa-conversion)
  - [Lab 3 - Lexer](#lab-3---lexer)
  - [Lab 4 - Chomsky Natural Form](#lab-4---chomsky-natural-form)
  - [Lab 5 - LL1 (top-down) and Simple Precedence (bottom-up) parsing](#lab-5---ll1-top-down-and-simple-precedence-bottom-up-parsing)
  - [Lab 6 - Parser and AST construction](#lab-6---parser-and-ast-construction)

## Introduction

The main objective of the course is to introduce the fundamental notions about formal languages (Chomsky classification of Languages, Regular Languages, Automata, Context Free Grammars, Push Down Automata, Turing Machine) and understand the mechanisms governing the analysis and synthesis of programming languages. Students will learn the most important techniques for the representation and generation of Languages (in particular, regular and context-free languages).

Those techniques will be applied to the construction of a compiler for a programming language. In particular, during this course the student will learn **how to build the different parts of a Compiler** with a particular emphasis on **Lexical Analysers, Parsers**.

## Literature to read

- Peter Linz - "Formal Languages and Automata"
- John E. Hopcroft, Jeffrey D. Ullman - "Introduction to Automata Theory, Languages, and Computation"
- [Andrew M. Pitts - "Regular Languages and Finite Automata"](https://www.cl.cam.ac.uk/teaching/0910/RLFA/reglfa.pdf)
- [Alfred V. Aho, Monica S. Lam, Ravi Sethi, Jeffrey D. Ullman - "Compilers: Principles, Techniques, & Tools"](<http://ce.sharif.edu/courses/94-95/1/ce414-2/resources/root/Text%20Books/Compiler%20Design/Alfred%20V.%20Aho,%20Monica%20S.%20Lam,%20Ravi%20Sethi,%20Jeffrey%20D.%20Ullman-Compilers%20-%20Principles,%20Techniques,%20and%20Tools-Pearson_Addison%20Wesley%20(2006).pdf>)
- [Thorsten Ball - "Writing an Interpreter in Go"](https://edu.anarcho-copy.org/Programming%20Languages/Go/writing%20an%20INTERPRETER%20in%20go.pdf)
- [Robert Nystrom - "Crafting Interpreters"](http://craftinginterpreters.com/contents.html)

## Laboratory tasks

### Lab 1 - Regular Grammar to Finite Automata

1. Write a program which converts regular grammar to Finite Automaton (FA).
2. Using Finite Automaton (FA) check if some input string is accepted by FA (meaning you could generate that string by traversing FA).
3. Determine the grammar type by the Chromsky classification.

### Lab 2 - NFA to DFA conversion

1. Convert NFA from your variant to DFA on paper, writing all transitions and drawing converted automato.
2. Write program which converts Nondeterministic Finite Automato (NFA) to Deterministic Finite Automato (DFA).
3. Display converted automato in form of graph or transition table.

### Lab 3 - Lexer

1. Define syntax of your programming languages and compute sample code which consist of 2 functions, first **main function** and second **function which performs some computations**. Sample should include call of second function from main function.
2. Define grammar of your programming language. Grammar should include sets of terminal and non-terminal symbols, and rules of derivation.
3. Write a program which will perform tokenization of your sample code, in other words write basic lexer of your programming language.

### Lab 4 - Chomsky Natural Form

Context Free Grammar (CFG) -> Chomsky Normal Form (CNF)

1. Eliminate ε-productions (D -> ε)
2. Eliminate any renaming / unit production (S -> A)
3. Eliminate Unproductive Symbols
4. Eliminate Inaccessible Symbols
5. Obtain the Chomsky Normal Form

### Lab 5 - LL1 (top-down) and Simple Precedence (bottom-up) parsing

LL1 parsing:

1. For the given grammar build the parsing table.
2. Analyse an word.

### Lab 6 - Parser and AST construction
