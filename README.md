# svc

Tool that lets you install a arbitrary executables as a service

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

### Usage

Kind of in the same spirit of `winsw`, you gotta:

  * Download the executable from here (missing link - use gh releases)
  * Rename the downloaded executable with whatever name you want and put it on
    the same directory as the executable that you want to run as a service. For
    the sake of the example: `nginx-svc.exe`.
  * Create a `svc.xml` file on the same directory that describes the service as
    well as the formulas on how to run it and how to stop it. An example:

```xml
<service>
  <id>nginx</id>
  <name>nginx</name>
  <description>nginx awesomeness</description>
  <executable>C:/Apps/nginx/nginx.exe</executable>
  <args>-p C:/Apps/nginx</args>
  <stopargs>-p C:/Apps/nginx -s stop</stopargs>
</service>
```

  * Install your app as a service running (The following commands need to be
    run as *Administrator*)

```
  nginx-svc.exe install
```

And that's it. You should see your app installed as a service.

You can also start and stop the service from your executable as:

```
  nginx-svc.exe [start|stop]
```

Or remove it from services as:

```
  nginx-svc.exe remove
```

### Build from source

First you gotta install kardianos's awesome `service` project. It is the only
dependency.

```
  go get https://bitbucket.org/kardianos/service
```

Then clone the project and build it:

```
  git clone https://github.com/fcarriedo/svc
  cd svc
  go build
```

You should be ready to go.

**Note**: This should work for linux/unix like systems but hasn't been tested.
