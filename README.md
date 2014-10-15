reprapi
=======

reprapi attempts to provide a REST/HTTP API for obtaining information about the 
contents of a [Debian package](http://www.debian.org/doc/manuals/debian-faq/ch-pkg_basics.en.html) 
hosting repository managed by the [reprepro](http://mirrorer.alioth.debian.org/) 
tool. reprapi is written in the [Go Programming Language](http://golang.org).

This tool was developed during a "Hack Day" while I was working for [Skybox](http://skybox.com). 
While I was there, we managed the lifecycle of internal services by building Debian 
packages and distributing them among various repositories depending on the environment. 
This tools allowed us to provide a more dynamic view into the stages of our packaging 
workflow.

## Endpoints

<table>
	<tr>
		<th>URI</th>
		<th>HTTP verb</th>
		<th>Pupose</th>
	</tr>
	<tr>
		<td>/package/{name}</td>
		<td>GET</td>
		<td>List the versions of the package and the distros they are associated with</td>
	</tr>
	<tr>
		<td>/distro/{name}</td>
		<td>GET</td>
		<td>List the packages in the given distro and their versions</td>
	</tr>
</table>

## Authentication

reprapi currently uses a very simple method for authenticating API requests. 
The daemon may be configured to require authentication at runtime using the `-t` 
switch. This token may be any arbitrary string at the time of this writing.

**Example Query with Parameter**:

`GET https://nodename.com:8080/reprapi/v2/package/foo?token=abc4c7c627376858`

## Requests

Requests to the API are simple HTTP requests against the API endpoints.

All request bodies should be in JSON, with Content-Type of `application/json`.

### Base URL

A few parameters may be set at runtime which will affect the base URL that you 
will use as the prefix for the desired endpoint.

* `-ssl`: Enforce encryption of communication with the API
* `-p`: Specify the port for communication with the API (defaults to 8080)

All endpoints should be prefixed with something similar to the following:

`{scheme}://{nodename}:{port}/reprapi/v2`

## Responses

All responses are in JSON, with Content-Type of `application/json`. A response 
is structured as follows:

`{ "resource_name": "resource value" }`

---

## Package

List all distros a package may be associated with and the versions of the 
package in that distro

**Endpoint**

`GET /package/{name}`

**Mandatory Parameters**

* `{name}`: Specify the name of the package you would like information about

**Response**

	{
	  "package": "foo",
	  "info": [
	    {
	      "distro": "unstable",
	      "version": "1.5.7-13101005539"
	    },
	    {
	      "distro": "staging",
	      "version": "1.4.0-13086201208"
	    },
	    {
	      "distro": "stable",
	      "version": "1.2.6-13099205144"
	    }
	  ]
	}

## Distro

List all packages associated with the given distro and their versions

**Endpoint**

`GET /distro/{name}`

**Mandatory Parameters**

* `{name}`: Specify the name of the distro you would like information about

**Response**


	{
	  "distro": "unstable",
	  "info": [
	    {
	      "package": "foo",
	      "version": "0.0.1-103306"
	    },
	    {
	      "package": "bar",
	      "version": "7.6.2.v20120308"
	    },
	    {
	      "package": "baz",
	      "version": "2.2.1-1"
	    },
	  ]
	}
