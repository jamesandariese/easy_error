# easy_error
A very simple library to make using x, err easier from both ends.

[![Build Status](https://travis-ci.org/jamesandariese/easy_error.svg)](https://travis-ci.org/jamesandariese/easy_error)

## Example

Change this code

    if f, err := os.Open(*flagfile); err != nil {
            panic(nil)
    } else {
            defer f.Close()
            if c, err := ioutil.ReadAll(f); err != nil {
                    panic(nil)
            } else {
                    return c
            }
    }

To this code

    f := easy_error.WrapFile(os.Open(*flagfile))
    defer f.Close()
    c := easy_error.WrapBytes(ioutil.ReadAll(f))
    return c

This becomes useful when you have a few things that might return an error
but you don't care about the error beyond it meaning you should panic
or do similar such as might be the case in a command line utility.

`easy_error.Wrap` will panic with the error being reported so that the specific
error can be reported elsewhere in a custom fashion.
