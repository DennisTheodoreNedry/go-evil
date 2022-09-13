# go-evil
Customizing evil has never been so easy

Go-evil is a red teams wet dream, a tool to beat all other tools of it's kind.<br>
What is go-evil, I hear you ask? Go-evil is a project all about the art of creating malware with a simpel language.<br>
The programming language we utilize is called evil, which only purpose is to translate ideas like "Hey I want a backdoor" into working code without the malware artist needing to know every every nook and cranny.<br> 

### Needed dependencies
The project currently needs two packages that need to be installed by hand, these are
1. webkit2gtk-4.0 
2. gtk3 

For the debian based systems, this should be a simple as the following.<br/>
`sudo apt install libwebkit2gtk-4.0-dev libgtk-3-dev`<br/>

### Compiling
Compiling the project is as simple as running `make` and run `make install_dependencies` if you get any dependency errors.<br>

To check if the compiler is working correctly either run the following code section or run `make test`
```
[version <major_version>]

use system;

main:{
    system.io.out("Hello world");
    system.exit("0");
}
```
<b>Tip:</b> The `major version` is obtained through `gevil -v`

### Example code
Basic example showcasing how we can display a website.
```
[version 1.0] @ Compiler version

use window;
use system;

main:{
    window.set.x("600"); @ Set the size on the x axis
    window.set.y("800"); @ Set the size on the y axis
    window.set.title("You just got hacked bruh!");
    window.goto("https://mrskeltal.com/"); @ Goes to the webiste specified

    system.exit("0"); @ Preventions the never ending loop
}
```


### Legal notice
I take no responsibility for what you create and do with this project. But do share your art as it's fun to see!



