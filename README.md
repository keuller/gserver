Simple GO Web Server
====================

The idea of this utility is start a simple Web server to provide a current directory as static dir.
This utility was written in Golang, that means you can run it in Windows, Linux and OSX.

Motivation
-------------
Sometimes we need a web server to prototype using static files and validate some communication with the server. The **gserver** utility brings it to us. **gserver** can also transform flat JSON files in "Rest endpoints" to simulate load some data from the server.

Usage
-------
You can achieve that, creating JSON files inside the **data** folder located at current directory and **gserver** will transform those files in REST endpoints. In the  **data** folder, we have 3 JSON files, named:

```
Using the api_v1_contacts.json file will be generated an endpoint like: /api/v1/contacts
Using the api_v1_todos_1.json file will be generated an endpoint like: /api/v1/todos/1
Using the api_v1_todos.json file will be generated an endpoint like: /api/v1/todos
```

Note: if you're running *gserver* and change the content of any file, you just must hit Ctrl+F5 on Windows to get the updated data. But if you create a new file while gserver is running, you must stop and start the server to load a new file.

To run the utility, we recommend you copy the specific version of operating system (located at ```dist``` folder) and put inside ```/usr/local/bin``` (Linux and OSX) or ```c:\windows\system32``` for Windows.

Note: In case of Linux or OSX, maybe you need to set permission mode ```chmod +x gserver``` to execution.

Run
----
To run the server, just type in your terminal:
```
$ gserver
```
You will get the output like this:
```
$ gserver
Go Server version 1.1.0
Static directory file /Users/keuller/Development/sample
creating handler for /doc
Server is running at http://0.0.0.0:9000

Current directory: /Users/keuller/Development/sample
Adding handler for /doc
Adding handler for /echo
```

By default ```gserver``` starts running on ```9000``` port listen all IP ```0.0.0.0```. But if you prefer to change that, just pass an argument ```--addr=127.0.0.1``` or ```--port=9090```, informing the specific IP address or port to bind the server.

Web Sockets
-----------

Help me with good sugestions to include in **gserver** or fork the project.

Any suggestion ?
