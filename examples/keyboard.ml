
main:{
	keyboard.lock(); @ Locks the keyboard
	
	window.x("600"); @ Widht on the x-axis
	window.y("800"); @ Height on the y-axis
	window.title("You just got hacked bruh!");
	window.url("https://mrskeltal.com/")
	window.run(); @ Utilizes a builtin print function
	
	keyboard.unlock(); @ Unlocks the keyboard
	system.exit("0"); @ Exits the program
}