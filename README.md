# go-tools
A utility tool library of Golang.

## Announcement
These packages only depend on the standard libraries, not any third packages.

## Installation
```shell
$ go get -u github.com/xgfone/go-tools
```

## Subpackages

subpackage   |   notice
-------------|-----------
atomics      | Some atomic types, such as Bool, Count, etc.
caller       | Get the filename and the line number where to call these functions.
checksum     | Calculate the checksum, such as `ICMP`.
compare      | Compare whether the first is greater than, less than, or equal to the second.
daemon       | Make the current process to the daemon process.
datetime     | Some convenient functions about datetime.
exception    | Exception handler like `"parent.child.sub-child...sub-child"`.
extremum     | Get the maximal or the minimal of both the values.
file         | Some convenient functions about the file operation.
function     | Call a function dynamically.
log/handler  | The logger handler, such as `TimedRotatingFile` like `logging.handlers.TimedRotatingFileHandler` in Python.
method       | Call the method of a type dynamically.
net/mac      | Standardize the mac address.
net/server   | The simple `TCP` / `UDP` server.
pagination   | It is usually used to compute the web pagination.
parse        | Convert something from a string to `bool`, `int`, `uint`, `float`, or from a certain type to string, etc.
pool         | Some simple convenient pools, such as `BufPool`, `GoPool`, etc.
tags         | Manage the tags in a struct.
slice        | Get a value from a slice and check whether a value exists in a slice.
tbucket      | The Simple Token Bucket like `HTB` in Linux `TC`.
values       | Get a value from a `slice` or `map`.
worker       | A worker pool with the dispatcher based on channel.

## Example
See the `test` file of each subpackage.
