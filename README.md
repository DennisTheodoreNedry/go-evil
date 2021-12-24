# Go evil

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
