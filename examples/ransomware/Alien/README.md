## Alien ransomware

### Purpose
The purpose behind is to provide a fully showcasing example of how you can utilize goevil
to create what could be a ransomware.

Disclaimer: This is only an example and we will not take any liability if used!

### What it does
The ransomware will target the current running users Desktop directory, and each file will be encrypted with a AES crypto.
<br/>
Each encrypted file will be called <original_file_name>.hive
<br/>

The AES key being used is 'hiveXXXXXXXXXXXX'.

A decryption process will be created after entering 'hive' into the pop up windows text box and hitting submit.

### Detection
