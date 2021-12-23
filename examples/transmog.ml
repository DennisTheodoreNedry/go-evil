
main:{
    @object.set_header_target("png"); @ Will result in the end product being identified as a png
    @object.set_extension(".png"); @ Sets the program extension
    
    malware.extension("png");
    window.x("600");
    window.y("800");
    window.title("You just got hacked bruh!");
    window.url("https://mrskeltal.com/");
    window.run();

    system.exit("0"); @ Exits the program
}
