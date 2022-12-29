# Dummy Process

Provides a simple executable with the simple capabilities of

- Time to wait before self terminating
- Listen to any OS signal and exit when a signal is received
- Configure exit code for normal termination
- Configure exit code for OS signal reception

## Requirements

You need the following software installed in order to compile this project

- A `go` compiler supporting the go version stated in `go.mod`
- GNU `make` or compatible


## Testing

Sorry, No unit test present at the moment. But the make commands for running them are in place :)

```shell
make test          # run all tests
make coverage      # like "test" but additionally show code coverage
make viewcoverage  # like "coverage" but opens a report in your default browser
```

## Building


To build a native executable run in a terminal

```shell
make
```

you may as well generate executables for Linux, Mac and Windows by running

```shell
make all
```

The newly created executable fille `dummy` (or `dummy.exe`) should now exist

## Running

You can see the usage page by running

```shell
./dummy --help
```

It should print something like this:

```
./dummy <duration> <exit_code> <signal_code>
    <duration> : time to wait before terminating the process, supported suffixes are ms,s,m,h,d. Default: 0ms
    <exit_code> : the code to return when the process ends sleeping. Default: 0
    <signal_code> : the code to return when a signal is captured by the process. Default: 65

Specifying "-" as a parameter will cause to use its default value.
```
Basically calling the executable without arguments will result in it exiting immediately after waiting 0ms and return exit code 0.


## Examples

Wait for 0ms, return code 0 for normal termination or code 65 if interrupted.

```shell
./dummy
```

Same as before

```
./dummy 0ms 0 65
```

Wait for 10 seconds using the default exit code values

```shell
./dummy 10s
```

Wait for 3 minutes, return code 10 if not interrupted and code 99 if interrupted.

```shell
./dummy 3m 10 99
```

Duration can be expressed as decimal numbers or composed by different values and units, all the following examples wait for one and half hours.

```shell
./dummy 1.5h
./dummy 90m
./dummy 1h30m
```
