# IAAF Intervals [![Build Status](https://travis-ci.org/Voles/go-iaaf-intervals.svg?branch=master)](https://travis-ci.org/Voles/go-iaaf-intervals)

Golang (de)serialization of interval notations using the [IAAF Standard Representation of Running Training](http://www.newintervaltraining.com/iaaf-standardised-sessions-www-newintervaltraining-com.pdf).

## Usage

```go
package main
    
    import (
    	"fmt"
    	interval "github.com/Voles/go-iaaf-intervals"
    )
    
    func main() {
    	set, err := interval.Parse("'2 x 6 x 400 (72”) [2’]'")
    	if err != nil {
    		panic(err)
    	}
    
    	fmt.Printf("total distance: %v", set.TotalDistance())
    }
```

## Spec

> sets x repetitions x distance (intensity/pace) [recovery between reps, then recovery between sets]

## Credits

This implementation is based on the [Python implementation](https://github.com/bwind/iaaf-intervals) created by [Bas Wind](https://github.com/bwind).
