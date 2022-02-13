# Go evil

[<!--lint ignore no-dead-urls-->![GitHub Actions status](https://github.com/s9rA16Bf4/go-evil/workflows/Building%20on%20latest%20ubuntu/badge.svg)](https://github.com/s9rA16Bf4/go-evil/actions?workflow=Building+on+latest+ubuntu)
[<!--lint ignore no-dead-urls-->![GitHub Actions status](https://github.com/s9rA16Bf4/go-evil/workflows/Building%20on%20latest%20macos/badge.svg)](https://github.com/s9rA16Bf4/go-evil/actions?workflow=Building+on+latest+macos)
[<!--lint ignore no-dead-urls-->![GitHub Actions status](https://github.com/s9rA16Bf4/go-evil/workflows/Building%20on%20latest%20windows/badge.svg)](https://github.com/s9rA16Bf4/go-evil/actions?workflow=Building+on+latest+windows)

<img src="https://github.com/s9rA16Bf4/go-evil/blob/unstable/pictures/logo.png" width="650"/>

Go-evil is a red teams wet dream, a tool to beat all other tools of it's kind.<br>
What is go-evil, I hear you ask? Go-evil is a project all about the art of creating malware with a simpel language.<br>
The programming language we utilize is called evil, which only purpose is to translate ideas like "Hey I want a backdoor" into working code without the malware artist needing to know every every nook and cranny.<br> 

### Compiling
Compiling the project is as simple as running `make` and run `make install_dependencies` if you get any dependency errors.<br>

To check if the compiler is working correctly either run the following code section or run `make test/check`
```
[version <major_version>]

main:{
    system.io.out("Hello world");
    system.exit("0");
}
```
<b>Tip:</b> The compiler version is obtained through `gevil -v`

### Example code
Basic example showcasing how we can display a website.
```
[version 1.0] @ Compiler version

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

