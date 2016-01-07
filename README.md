Simple GO Web Server
====================

The idea behind this utility is to instantiate a simple Web server to provide a current directory as static dir.

Motivation
-------------
Sometime we need to create a simple prototype using static files and validate communication with server. ```gserver``` utility can help with it. ```gserver``` also transform JSON files in "Rest endpoints" to simulate load some data from the server.

Usage
-------
You can achieve that, creating files inside ```data``` folder located at current directory and ```gserver``` will transform those files in REST endpoints. Let's take a look at example:

```
In data folder inside the repo, we have 3 JSON files, named:
api_v1_contacts.json wil be generated an endpoint like: /api/v1/contacts
api_v1_todos_1.json wil be generated an endpoint like: /api/v1/todos/1
api_v1_todos.json wil be generated an endpoint like: /api/v1/todos
```

Note: if you're running gserver and change the content of any file, you just must hit Ctrl+F5 on Windows to get the updated data. But if you create a new file while gserver is running, you must stop and start the server to load a new file.

This utility was written in Golang, that means you can run it in Windows, Linux and OSX.

To run the utility, we recommend you copy the specific version of operating system (located at ```dist``` folder) and put inside ```/usr/local/bin``` (Linux and OSX) or ```c:\windows\system32``` for Windows.

Enjoy.
