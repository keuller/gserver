Simple GO Web Server
====================

The idea of this utility is to start a simple web server that returns JSON data on REST-endpoints defined by the name of the JSON files.

Motivation
-------------
Sometimes we need a web server to prototype using static files and validate some communication with the server. The **gserver** can transform flat JSON files to REST endpoints to simulate reading data from a backend server.

Usage
-------
Create some JSON files in a folder, by default called **data** in the current directory. When started, **gserver** will transform those files into REST endpoints. In the  **data** folder, we have some JSON files, named:

```
Using the api_v1_contacts.json file will generate an endpoint like: /api/v1/contacts
Using the api_v1_todos.json file will generate an endpoint like: /api/v1/todos
Using the api_v1_todos_{id:[0-9]+}.json file will generate an endpoint like: /api/v1/todos/<number>
```

Note: if you're running *gserver* and change the content of any file or add a new one, you must stop and start the server to get the new endpoints and content.

To run the utility, link to the correct version for your operating system (located in the ```bin``` folder) and put inside ```/usr/local/bin``` (Linux and OSX) or ```c:\windows\system32``` for Windows.

Note: In case of Linux or OSX, you might need to set permission mode ```chmod +x gserver``` to execution.

Run
----
To run the server:
```
$ gserver
```
You will get the output like this:
```
> ./bin/gserver-osx -v
[gserver] 2016/04/06 16:12:09 Server is running at http://0.0.0.0:9000
[gserver] 2016/04/06 16:13:13 Simple Go Server version 1.2.0
[gserver] 2016/04/06 16:13:13 (build e31ec5c3f0e6e8041273473f1f91405118c49f23)
[gserver] 2016/04/06 16:13:13 Adding handler for /api/v1/todos
[gserver] 2016/04/06 16:13:13 Adding handler for /api/v1/todos/{id:[0-9]+}
[gserver] 2016/04/06 16:13:13 Adding handler for /api/v1/contacts
[gserver] 2016/04/06 16:13:13 Server is running at http://0.0.0.0:9000
```

By default ```gserver``` starts running on ```9000``` port listening to all IPv4 on ```0.0.0.0```. But if you prefer to change that, just pass an argument ```-addr 127.0.0.1``` or ```-port 9090``` to set address or port.

Web Sockets
-----------
You can test a websocket connection against the ```/echo``` endpoint if you start the server with the ```-websocket``` flag. There is a ***websocket.html*** file that you can use for tests.

Suggestions for improvements are appreciated!
