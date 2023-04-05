# Go-evil - Customizing evil has never been so easy

<img src="./images/logo.png" width="1000"/>

Go-evil is a red teams wet dream, a tool to beat all other tools of it's kind.<br>
What is go-evil, I hear you ask? Go-evil is a project all about the art of creating malware with a simpel language.<br>
The programming language we utilize is called evil, which only purpose is to translate ideas like "Hey I want a backdoor" into working code without the malware artist needing to know every every nook and cranny.<br> 



### How to run
##### Docker
You can run the compiler through docker and by executing the statements shown below!

Run `make docker` to build the docker image.

Create a folder labeled `builds` in your local directory, this will be the directory where you want to place all your files to compile.

Run the following command `docker run -v $(pwd)/builds:/app/builds --rm goevil/gevil-2.1 <commands>`

An example of this can be `docker run -v $(pwd)/builds:/app/builds --rm goevil/gevil-2.1 -f builds/out.evil -bd builds` which will compile the `out.evil` file and create the binary `builds/out`.

But if you're lazy you can also utilize the script labeled `gevil` under the `tools` directory `bash tools/gevil <commands>` which will do all work for you.

##### Locally
If you still want to compile the project locally, just run `make dependencies` and `make` to compile the project. After this you can utilize the compiler as intended.

### Changes in version 2
There are many changes in v2, some are directly related to the syntaxes but the majority is how the internals work.

Some notable features arriving in v2 are,
1. Evil arrays `${...}`, i.e. `${"value1", "value"}$`.
2. Functions! Which have different meanings depending on their function type.
3. A fully integrated `window` into the source code.
4. You can now customize your malware by using a `compiler configuration section` aka `[]` in your code.
5. Obfuscating the source code is now a built in option.
6. The line terminator `;` is no longer needed.
7. Comments must now be terminated, i.e. `@ this is a comment @`.
8. Variables must now be terminated, i.e. `$1$`.
9. The malware now contains a behavior option when it detects an active debugger, this allows you to kill the malware, remove it from the disk or loop indefinitely.
10. You can add an `n` amount of randomly definied variables and functions to your source code through `self::add_random_var` and `self::add_random_func`

... and much more    
    
### Structure of a go-evil malware
We have currently three different functions, these are marked by a one character long string preceding the functions name.<br/>
And of course they have different meanings.

1. `b`- `boot`, these are functions that will be automatically called on boot of the program.
2. `l` - `loop`, these are functions that will be automatically called within the for-loop in the main function.
3. `c` - `call`, these are functions that will need to be called by the developer, i.e. `self::call("<function_name>")` or through a binding.

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
}

```

### Example code
In this example I will show you how you print messages to the CLI.
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

This example shows how you can bind evil functions to javascript functions
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

### Examples
There is no wiki page as each function/object have a corresponding example file under `examples/<domain>/`. <br/>
We strive to have an example for each new implementation, but we are also working on a so-called "cookbook" which will go into a further detail on how everything works.

### Legal notice
I take no responsibility for what you create and do with this project. And any mischief that you cause and happen to fall into is on yourself!
But do share your art, as it's fun to see!



