# svc

Tool that lets you install an arbitrary executables as a service

### Why

This was a day I was trying to install fukn' `nginx` as a windows service and
since it doesn't provide a way (as of July 2013) to do it, I thought I would
find an easy way to do it. I tried a couple of Windows Service wrappers like
[kohsuke's winsw](https://github.com/kohsuke/winsw) and what seems to be it's
[predecessor](https://kenai.com/projects/winsw) to all fail on me.

After spending a whole day on this mundane task, I found [kardianos's awesome
project](https://bitbucket.org/kardianos/service) which lets you, albeit
programatically, install arbitrary programs as windows services. It seems to
support other unix-like systems as well.

### Prerequisites

You have to have kardianos's awesome `service` project. It is the only
dependency.

```
  go get https://bitbucket.org/kardianos/service
```

### Usage

Clone the project:

```
  git clone https://github.com/fcarriedo/svc
```

Put the `svc.go` file on the directory where your *to-be-service* is located.

Open the `svc.go` file and update the `svcDescriptor` with the details of your
executable.  For the sake of the example, lets say we want to embed `nginx` in
our `nginx-svc.exe` service wrapper:

```go
var svcDescriptor = `
<service>
  <id>nginx</id>
  <name>nginx</name>
  <desc>nginx awesomeness</desc>
  <exec>C:/Apps/nginx/nginx.exe</exec>
  <args>-p C:/Apps/nginx</args>
  <stopexec>C:/Apps/nginx/nginx.exe</stopexec>
  <stopargs>-p C:/Apps/nginx -s stop</stopargs>
</service>
`
```

Compile it:

```
  go build svc.go -o nginx-svc.exe
```

Now your ready to install your service and/or manipulate it.

To install it:

```
  nginx-svc.exe install
```

You can also start and stop the service from your executable as:

```
  nginx-svc.exe [start|stop]
```

Or remove it from services as:

```
  nginx-svc.exe remove
```

**Note**: This should work for linux/unix like systems but hasn't been tested.

**Note 2**: I couln't put the service descriptor as a file as of now. It was
causing a failure when starting it from the serivces (I believe it is
permission related). For now, you'll just have to compile it. I know it sucks
but I'll leave it like that until I can investigate further.
