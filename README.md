Simple GO Web Server
====================

The idea behind this utility is to instantiate a simple Web server to provide a current directory as static dir.

+ Motivation
-------------

Sometime we need to create a simple prototype using static files and validate communication with server. ```gserver``` utility can help with it. ```gserver``` also transform JSON files in "Rest endpoints" to simulate load some data from the server.

You can achieve that, creating files inside ```data``` folder located at current directory and ```gserver``` will transform those files in REST endpoints.

This utility was written in Golang, that means you can run this utility in Windows, Linux and OSX.

To run the utility, we recommend you copy the specific version of operating system (located at ```dist``` folder) and put inside ```/usr/local/bin``` (Linux and OSX) or ```c:\windows\system32``` for Windows.

Enjoy.
