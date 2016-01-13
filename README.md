# UXSS PoC

PoC for a [UXSS exploit](https://blog.innerht.ml/ie-uxss/).

## Building

Get Goop if you don't already have it:

    $ go get -u github.com/nitrous-io/goop

Now build:

    $ make

## Running it

    $ bin/uxss -addr 192.169.1.100:8080 -protected_url https://target-url.com/

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

## Authors

* [Will Jessop](mailto:will@willj.net)

## LICENSE

See LICENSE.md
