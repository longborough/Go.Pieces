// untod converts TOD clock values to
// dates and times
package main

import "fmt"
import "math"
import "flag"
import "strconv"
import "os"
import "time"
import "github.com/longborough/todinfo"

// -----------------------------------------------------------
// envFloat:
//          Gets an hour offset from a named environment
//          variable and converts it to a float64.
//          If not a valid value, returns math.NaN()
//          Limits are +/- 24.5
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
func envFloat(s string) float64 {
	s = os.Getenv(s)
	if s == "" {
		return math.NaN()
	}
	off, err := strconv.ParseFloat(s, 64)
	if err != nil {
		println(s, "- Environment variable", err)
		return math.NaN()
	}
	if off > 24.5 {
		return math.NaN()
	}
	if off < -24.5 {
		return math.NaN()
	}
	return off
}

func main() {
	t := time.Now()
	_, z := t.Zone()
	sysloff := float64(z / 3600) // Local machine offset
	loff := envFloat("TODL")     // Try for local offset
	aoff := envFloat("TODA")     // Try for additional offset
	loffPtr := flag.Float64("zl", math.NaN(), "Zone Local: time zone offset")
	aoffPtr := flag.Float64("za", math.NaN(), "Zone Additional: time zone offset")
	revPtr := flag.Bool("r", false, "Reverse:   convert timestamp to TOD")
	revPmc := flag.Bool("m", false, "Minute:    convert PARS perpetual minute clock")
	ngPtr := flag.Bool("ng", false, "No GMT:    suppress GMT line")
	lpPtr := flag.Bool("pl", false, "Pad Left:  pad TOD with 0 on the left")
	rpPtr := flag.Bool("pr", false, "Pad Right: pad TOD with 0 on the right")
	flag.Parse()
	if math.IsNaN(*loffPtr) { // Local offset parameter not given
		if math.IsNaN(loff) { // Nor environment variable
			loff = sysloff // Default is local system offset
		}
	} else { // Local offset parameter given
		loff = *loffPtr // Use parameter value
	}
	if math.IsNaN(*aoffPtr) { // Additional offset parameter not given
		if math.IsNaN(aoff) { // Nor environment variable
			aoff = 0 // Default is zero (UTC)
		}
	} else { // Local offset parameter given
		aoff = *aoffPtr // Use parameter value
	}
	if *revPtr {
		var sdate, stime string
		switch flag.NArg() {
		case 0:
			{
				sdate, stime = t.Format("2006-01-02"), t.Format("15:04:05.999999")
			}
		case 1:
			{
				sdate, stime = flag.Arg(0), "00:00:00.000000"
			}
		case 2:
			{
				sdate, stime = flag.Arg(0), flag.Arg(1)
				switch {
				case len(stime) == 8:
					{
						stime += ".000000"
					}
				case len(stime) < 6:
					{
						println("Reverse needs exactly two arguments, date and time")
						os.Exit(9)
					}
				case stime[8:8] != ".":
					{
						println("Reverse needs exactly two arguments, date and time")
						os.Exit(9)
					}
				default:
					{
						stime = (stime + "000000")[0:16]
					}
				}
			}
		default:
			{
				println("Reverse needs exactly two arguments, date and time")
				os.Exit(9)
			}
		}
		if sdate < "1900-01-01" {
			println("Reverse needs exactly two arguments, date and time")
			os.Exit(9)
		}
		if sdate > "5999-99-99" {
			println("Reverse needs exactly two arguments, date and time")
			os.Exit(9)
		}
		if !*ngPtr {
			fmt.Print(todinfo.Todsearch(sdate, stime, 0))
		}
		if loff != 0 || *ngPtr {
			fmt.Print(todinfo.Todsearch(sdate, stime, loff))
		}
		if aoff != 0 && aoff != loff {
			fmt.Print(todinfo.Todsearch(sdate, stime, aoff))
		}
	} else {
		for i := 0; i < flag.NArg(); i++ {
			if !*revPmc {
				if !*ngPtr {
					fmt.Print(todinfo.Todcalc(flag.Arg(i), 0, *lpPtr, *rpPtr, *revPmc))
				}
				if loff != 0 || *ngPtr {
					fmt.Print(todinfo.Todcalc(flag.Arg(i), loff, *lpPtr, *rpPtr, *revPmc))
				}
				if aoff != 0 && aoff != loff {
					fmt.Print(todinfo.Todcalc(flag.Arg(i), aoff, *lpPtr, *rpPtr, *revPmc))
				}
			} else {
				fmt.Print(todinfo.Todcalc(flag.Arg(i), loff, *lpPtr, *rpPtr, *revPmc))				
			}
		}
	}
}
