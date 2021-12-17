use system; @ Imports exit
use keyboard; @ Imports lock, unlock
use window; @ Imports notify


main:{
	keyboard.lock(); @ Locks the keyboard
	
	window.notify("Oh hi mark!"); @ Utilizes a builtin print function
	
	keyboard.unlock(); @ Unlocks the keyboard
	system.exit(); @ Exits the program
}
