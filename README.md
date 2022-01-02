[<!--lint ignore no-dead-urls-->![GitHub Actions status](https://github.com/s9rA16Bf4/go-evil/workflows/Building%20on%20latest%20ubuntu/badge.svg)](https://github.com/s9rA16Bf4/go-evil/actions?workflow=Building+on+latest+ubuntu)
[<!--lint ignore no-dead-urls-->![GitHub Actions status](https://github.com/s9rA16Bf4/go-evil/workflows/Building%20on%20latest%20macos/badge.svg)](https://github.com/s9rA16Bf4/go-evil/actions?workflow=Building+on+latest+macos)
[<!--lint ignore no-dead-urls-->![GitHub Actions status](https://github.com/s9rA16Bf4/go-evil/workflows/Building%20on%20latest%20windows/badge.svg)](https://github.com/s9rA16Bf4/go-evil/actions?workflow=Building+on+latest+windows)

# Go evil

Customizing evil has never been so easy

# Note: there is a lot of stuff not done with this project and therefore there really isn't much you can currently do.
# Note: A lot of changes will occur before a v.1 release

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


### Legal notice
I take no responsibility for what you create and do with this project. But do share your art as it's fun to see!

