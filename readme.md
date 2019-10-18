# Introduction
This tool is specially provided for bolt merchants who need to check their account flow regularly.
Although we provide API interface for merchants like you to query, it still seems hard for you to use the data efficiently, especially when you don't have technical support to develop a system connecting to the API.

Therefore, we open sourced a client application for you to convert API data into csv files.
This tool seems more helpful to accountancy. But you also get the idea of developing you own client UI or website from our open-sourced code.

# How to build it
1. Install golang from https://golang.org/
2. Install git to your system
3. Open a terminal 
3. run ```git clone https://github.com/sdcex/historder```
4. run ```cd historder & make build```
Then you can find the excutable binaries in ./out folder
We build three binaries for mac(darwin), linux and windows respectively,  
please make sure to use the right binary according to your operating system.

# How to use it
1. Put the config folder to the path where you put the binary
2. Modify the yaml file in the config folder
 - Fill your clientId and clientSecret from 'auth' field
 - Change search condition from 'search' field
3. Run the binary
Every time you run the binary, You will find a new csv report named with datetime in current folder.

Wish you enjoy it! 
If you have any problem with it, feel free to contact dev@sdce.com.au