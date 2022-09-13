# go-evil
Customizing evil has never been so easy

Go-evil is a red teams wet dream, a tool to beat all other tools of it's kind.<br>
What is go-evil, I hear you ask? Go-evil is a project all about the art of creating malware with a simpel language.<br>
The programming language we utilize is called evil, which only purpose is to translate ideas like "Hey I want a backdoor" into working code without the malware artist needing to know every every nook and cranny.<br> 

### Needed dependencies
The project needs a couple of packages and modules to work

Those that must be installed by hand are,
1. webkit2gtk-4.0, `sudo apt install libwebkit2gtk-4.0-dev`
2. gtk3, `sudo apt install libgtk-3-dev`

The rest of them can be installed by running `make dependencies`

### Compiling
Compiling the ide is as simple as running `make` in the root directory of the project <br>

So what can you now do? Well.. you can either compile a file containing evil code or launching the text editor.

Run the following if you want to compile a file, `./gevil -f <path/to/file>` and if you would like to open the text editor run,
<br/>`./gevil -f <path/to/file> -t`

### Changes in version 2
There are many changes in v2, some are directly related to the syntaxes but the majority is how the internals work.

Some notable features arriving in v2 are,
1. Evil arrays `${...}`, i.e. `${"value1", "value"}$`.
2. Functions! Which have different meanings depending on their function type.
2. A fully integrated `webview` into the source code.
3. You can now customize your malware by using a `compiler configuration section` aka `[]` in your code.
4. Obfuscating the source code is now a built in option.
5. The line terminator `;` is no longer needed.
6. Comments must now be terminated, i.e. `@ this is a comment @`.
7. Variables must now be terminated, i.e. `$1$`.

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

// Global variables
...

// Function def
...

func main(){
    // Calling `boot` functions
    ...

    for {
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
    system::out("Hello from boot!")
}

l main_func {
    system::out("Hello from main!")
    system::exit("0")
}
```

This example shows how you can bind evil functions to javascript functions
```
@ Showcasing how you can bind an evil function to a javascript call @

use system
use webview

[
    version 2.0
    output bind
    os linux
    arch amd64
    obfuscate false
]

c test1 {
    system::out("I've been summoned!")
}

b bind {
    @ Essentially saying that when the js function button is called, call test1 @
    webview::bind(${"button", "test1"}$) @ The ${...}$ is known as an evil array @ 

    webview::js("function testing() {window.button();}")

    webview::html("<p>Watch the console after pressing the button!</p>")
    webview::html("<input type='submit' onclick='testing()'>")

    webview::run()
}

l loop {
    system::exit("0")
}
```
Running the code above will result in the following
![image](https://user-images.githubusercontent.com/14398606/189844664-0d870f9a-4a27-401f-a6ec-619fb8556cd4.png)


### Legal notice
I take no responsibility for what you create and do with this project. But do share your art as it's fun to see!



