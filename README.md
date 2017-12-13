# Sky Island

**Experimental** *ALPHA stage* 

Sky Island is a FreeBSD based, Jail driven, ZFS backed, FaaS/serverless platform for running Go functions.  Interaction is via the HTTP JSON API.

<p align="left">
  <a href="https://github.com/briandowns/sky-island/releases"><img src="https://img.shields.io/badge/version-0.0.0-green.svg?" alt="Version"></a>
</p>

## How It Works

A request comes in to run a function. The request contains a Github URL to a Go repository containing the function. The request also contains the "call".  The call is what will be run including the arguments necessary to run the function.

Upon successfully accepting the inbound request, the server will check if the repo has already been cloned and if not, it will clone it. From there, it will generate a "main.go" file and compile a binary in the "build" jail. The "build" jail holds all of the cloned repositories and will be reused on each request unless otherwise told not to.  Once a binary is created, an execution jail is created, the binary is copied into it, and is subsequently executed. The binary's output is then returned to the caller via an HTTP response to the original request.

## Requirements

* ZFS
* Go version >= 1.9 
* Make sure that `jails_enabled="YES"` is present in the "/etc/rc.conf" file

## System Initialization

Initialzing the system does a number of things to make running Sky Island easier.  Sky Island will check to see if the base system packages and Go tarball have already been downloaded and if they have, they'll use those. 

* Create a ZFS dataset to work from
* Download the base package for the version of FreeBSD you have installed
* Extract those packages to the dataset where the base jail will be kept
* Update the base jail with `freebsd-update`
* Set some basic jail configuration
* Install Go and create a workspace
* Create a ZFS snapshot of the base jail

This is accomplished by running: 

`si-server -c config.json -i`

## Running Sky Island

To run Sky Island, run the command below.

`si-server -c config.json` 

## IP Address Management

A Sky Island config file has an IP4 section to configure how it handles jails IP addressing.  If a request is received that indicates it needs an IP address, Sky Island checks to see if there are available addresses and returns one to be assigned to the execution jail.  By default, all jails have no IP address.  For the time being, all assigned IP addresses need to exist on the same network as the host interface. [Issue 32]()https://github.com/briandowns/sky-island/issues/32 has been added to allow for multiple subnets but moreover segregate traffic.

## Caching

Rather than cloning the repo each time it's called, the previously cloned repo is cached in the build jail indefinitely. This cache can be busted however on a repo by repo basis by including `cache_bust=true` in payload of a "function run" POST request. This will force Sky Island to clone the repo and build a new binary.

## API

| Method | Resource                    | Description                                                            |
| :----- | :-------                    | :----------                                                            |
| GET    | /api/v1/healthcheck         | Verifies the service is up and running                                 | 
| POST   | /api/v1/function            | Endpoint that receives function run requests                           |
| GET    | /api/v1/admin/api-stats     | API statistics                                                         | 
| GET    | /api/v1/admin/jails         | Get a list of the running jails                                        |
| GET    | /api/v1/admin/jail/{id}     | Get the details for the given jail                                     |
| DELETE | /api/v1/admin/jail/{id}     | Kill the jail with the given ID                                        |
| DELETE | /api/v1/admin/jails         | Kill all jails                                                         |
| GET    | /api/v1/admin/ips           | Get a list of IP's filtered by param. `?state={available|unavailable}` |
| PUT    | /api/v1/admin/ips           | Update the state of a given IP                                         |


## Contact

Brian Downs [@bdowns328](http://twitter.com/bdowns328)

## License

Sky Island source code is available under the BSD 3 Clause [License](/LICENSE).

## 

![alt text](https://www.freebsd.org/gifs/powerani.gif)