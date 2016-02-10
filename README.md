# UXSS PoC

PoC for a [UXSS exploit](https://blog.innerht.ml/ie-uxss/).

There are [releases](https://github.com/wjessop/UXSS_PoC/releases) if you don't want to have to build this tool yourself.

## Building

Get some dependant utils:

    $ make setup

Now build:

    $ make

## Running it

    $ bin/uxss -addr 192.169.1.100:8080 -target_url https://target-url.com/

Now navigate to http://192.169.1.100:8080/ in IE.

### Options:

-addr:
> The address of the web server we will listen on

-target_url:
> The target URL you want to perform the UXSS on

-non_target_url
> A URL to use in the redirect, necessary for the exploit, but the default is probably fine

-use_sleep
> Set to no to use the alert() method of exploitation

-template
> Path to a custom template. See https://github.com/wjessop/UXSS_PoC/blob/master/data/templates/home.html for the variables that will be interpolated. This is useful for demonstrating a more detailed attack. If not specified a simple HTML template will be used.

## Authors

* [Will Jessop](mailto:will@willj.net)

## LICENSE

See LICENSE.md
