## Alien ransomware
![Cover photo](https://github.com/s9rA16Bf4/go-evil/blob/v2/examples/ransomware/Alien/background.jpg?raw=true)

### Purpose
The purpose behind is to provide a fully showcasing example of how you can utilize goevil
to create what could be a ransomware.

<b>Disclaimer: We take no liability of how you use it!</b>

### What it does
The ransomware will target the current running users Desktop directory, and each file will be encrypted with a AES crypto.
<br/>
Each encrypted file will be called <original_file_name>.hive
<br/>

The AES key being used is 'hiveXXXXXXXXXXXX'.

A decryption process will be created after entering 'hive' into the pop up windows text box and hitting submit.

The ransomware will only work on a nix based system but can with ease be changed to work on windows aswell.

### Detection
The best way forward to detect it is to look for files that have the extension '.hive' in the Desktop directory.

As of 2022-10-29, virustotal will not detect the ransomware as something malicious.
![Result](https://github.com/s9rA16Bf4/go-evil/blob/v2/images/Virus_total_alien_ransomware_result.png?raw=true)

[Link to virustotal result](https://www.virustotal.com/gui/file/13a66f4f61a4b9fe932ea1a69aa4fad0eb50fdac0f4a5ec4ed5987257820c118/detection)
