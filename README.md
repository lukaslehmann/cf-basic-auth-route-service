# Basic Authentication Route Service

**NOTE** This is still a work in progress, so the password is not yet the reverse of the URL as promised below. It is either defaulted to `letmein` or can be overridden as an environment variable.
Refer to the example `config.yml` file in the `servicebroker` folder. All other functionality works as stated, although the tests for the router are bit of a mess and work in progress!

Using the new route services functionality available in Cloud Foundry, you can now bind applications to routing services.
Traffic sent to your application is routed through the bound routing service before continuing onto your service.

This allows you to perform actions on the HTTP traffic, such as enforcing basic authentication (this sample app), rate limiting or logging.

For more details see:
* [Route Services Documentation](http://docs.cloudfoundry.org/services/route-services.html)

## Getting Started

There are two components and thus steps to getting this up and running.

### Deploying the service broker

First navigate to the `servicebroker` directory and `$cf push` this will use the provided example manifest.
This will register the application with the URL `https://basic-auth-broker.local.pcfdev.io`

Next you need to register this as a private service broker in your space, you can change the default username and password in the `config.yml` file.
By default they are `admin` and `letmein`

```
cf create-service-broker basic-auth-broker admin letmein https://basic-auth-broker.local.pcfdev.io --space-scoped
```

You should now be able to see the service in the marketplace if you run `$cf m`

### Deploying the routing service

First navigate to the `routeserver` directory and `$cf push` this will use the provided example manifest.
This will register the application with the URL `https://basic-auth-router.local.pcfdev.io`

### Protecting an application with basic authentication

Now you have setup the supporting components, you can now protect your application with basic auth!

First create an instance of the service from the marketplace, here we are calling our instance `authy`
```
$cf create-service p-basic-auth reverse-name authy
```

Next, identify the application and its URL which you wish to protect. Here we have an application called `hello` with a URL of `https://hello.local.pcfdev.io`
```
⇒  cf a
Getting apps in org pcfdev-org / space pcfdev-space as admin...
OK

name                requested state   instances   memory   disk   urls
hello               started           1/1         256M     512M   hello.local.pcfdev.io
basic-auth-broker   started           1/1         512M     512M   basic-auth-broker.local.pcfdev.io
basic-auth-router   started           1/1         256M     512M   basic-auth-router.local.pcfdev.io
```
We can validate that we can access this URL without providing any credentials
```
⇒  curl https://hello.local.pcfdev.io -k
Hello world! I should be protected by basic auth%
```

Then you need to bind the service instance you created called `authy` to the `hello.local.pcfdev.io` route
```
⇒  cf bind-route-service local.pcfdev.io authy --hostname hello

Binding may cause requests for route hello.local.pcfdev.io to be altered by service instance authy. Do you want to proceed?> y
Binding route hello.local.pcfdev.io to service instance authy in org pcfdev-org / space pcfdev-space as admin...
OK
```

You can validate the route for `hello` is now bound to the `authy` service instance
```
⇒  cf routes
Getting routes for org pcfdev-org / space pcfdev-space as admin ...

space          host                domain            port   path   type   apps                service
pcfdev-space   hello               local.pcfdev.io                        hello               authy
pcfdev-space   basic-auth-broker   local.pcfdev.io                        basic-auth-broker
pcfdev-space   basic-auth-router   local.pcfdev.io                        basic-auth-router
```

All of that looks good, so the last step is to validate we can no longer view the `hello` application without providing credentials!

```
⇒  curl -k https://hello.local.pcfdev.io
Unauthorized
```

and with credentials
```
⇒  curl -k https://hello.local.pcfdev.io -u admin:letmein
Hello world! I should be protected by basic auth%
```

Success!

## Overview

### Service Broker

This is a service broker which conforms to the Services API. 

This registers a service in the market place called `p-basic-auth`. It currently only has one service plan called `reverse-name`. This service will protect you application with basic authentication.
The username is hard coded to `admin` and the password is based upon your applications URL.

It uses the URL before the first . (dot), without any special characters or spaces and then in reverse. 

For example:
`https://hello-world.cfapps.io` will have a password of `dlrowolleh`
`https://pivotal.cfapps.io` will have a password `latovip`

Why is the password is in reverse!? Because it's a pretty simple implementation, it also means the routing service can be stateless as we can determine what the password should be during runtime, which makes for a nice simple reference implementation.

# TODO

* Rewrite the actual route-service and refactor
* Add tests
* Finish writing readme and getting started
* Integration tests?
