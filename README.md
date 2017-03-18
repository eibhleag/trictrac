**trictrac** is a local database for tracking, well, anything (see: [motivation](#motivation)).

## usage

To build: `go build -o tt`

To run: `tt [command] [arguments]`

### commands
- **new** key [value] - creates a new counter
- **del** key - deletes a counter
- **inc** key [step] - increments a counter (by step, if provided)
- **dec** key [step] - decrements a counter (by step, if provided)
- **set** key value - sets the value of a counter
- **sum** prefix - sums the value of 'prefix.*'
- **list** [prefix]

## motivation

I want to easily track of disparate bits of information without creating (and structuring) a spreadsheet, script or notebook to do it.

**trictrac** should help me test some related hypotheses. Given the means to track this information:
- I gain some value from doing so 
- I'll actually keep up with it
- It's useful to be able to script it
- The command line is a good place to do it

It's also a good opportunity for me to get better acquainted with Go :-)

(Which is to say, if you're new to Go yourself, this may not be the best repo to learn from...)

## collaboration

With motivation in mind, I'm not trying to build a widely applicable tool at present. However, if you want to try **trictrac** out for yourself, freel free to raise an issue with any feedback you might have.

At this stage I probably won't be accepting PRs.

## design

Currently, **trictrac** is a very thin wrapper over [Bolt](https://github.com/boltdb/bolt) supporting one data type - 64 bit Counters. The interface is a set of [cobra](https://github.com/spf13/cobra) commands, and what error handling there is reflects that. In short, there's not much in the way of _technical_ design.

The main design principle behind **trictrac** _as a tool_ is to have a small set of composable primitives that balance ease of manual use (command line invocation) with ease of automation (running in commit hooks, etc). As such, there's currently a small set of operations over one data type.