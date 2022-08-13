# go-evil
Customizing evil has never been so easy


Example
```

    f init:{
        @ This section should contain everything that needs to be done before the malware actually can start

    }

    f infect:{
        @ This section will handle how the malware spread
    }

    f loop:{
        @ The loop function is the core of each malware
        @ and just as the name says it will loop until an `sys:exit` call is reached

    }
```


