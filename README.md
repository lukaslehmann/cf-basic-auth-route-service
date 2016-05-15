# Basic Authentication Route Service

Using the new route services functionality available in Cloud Foundry, you can now bind applications to routing services.
Traffic sent to your application is routed through the bound routing service before continuning onto your service.

This allows you to perform actions on the HTTP traffic, such as enforcing basic authentication (this sample app), rate limiting or logging.

For more details see:
* (Route Services Documentation)[http://docs.cloudfoundry.org/services/route-services.html]

## Getting Started

There are two parts to this repository


## Overview

### Service Broker - This is a service broker which conforms to the Services API. 

This registers a service in the market place called `p-basic-auth`. It currently only has one service plan called `reverse-name`. This service will protect you application with basic authentication.
The username is hard coded to `admin` and the password is based upon your applications URL.

It uses the URL before the first . (dot), without any special characters or spaces and then in reverse. 

For example:
`https://hello-world.cfapps.io` will have a password of `dlrowolleh`
`https://pivotal.cfapps.io` will have a password `latovip`

Why is the password is in reverse!? Because it's a pretty simple implementation, it also means the routing service can be stateless as we can determine what the password should be during runtime, which makes for a nice simple reference implementation.
