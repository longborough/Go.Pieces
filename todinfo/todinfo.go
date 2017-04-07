// todinfo: TOD clock conversion routines
package todinfo

import (
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)

const ParsDayZero int64 = 2082931200000000 // Pars Day Zero microseconds TOD

type Todinfo struct {
	Intod                       string
	Offset                      float64
	Hextod                      string
	Numtod                      int64
	Year, Month, Day, Yday      int64
	Hour, Minute, Second, Musec int64
	Wkday                       int
	Pmc                         int64
}

// -----------------------------------------------------------
// String: provides the interface for printing Todinfo structures
func (t Todinfo) String() string {
	var offset string
	var pmc string
	wday := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}[t.Wkday]
	if nearInteger(t.Offset) {
		offset = fmt.Sprintf("%+d", int(round(t.Offset)))
	} else {
		offset = fmt.Sprintf("%+0.1f", t.Offset)
	}
	if t.Pmc < 0 || t.Pmc > 4294967295 { 
		pmc = "---" 
	} else {
		pmc = fmt.Sprintf("%08x", t.Pmc)
	}
	return fmt.Sprintf("%3s %8s %4s--- : %4d-%02d-%02d %02d:%02d:%02d.%06d GMT %s %4d.%03d %s %s\n",
		t.Hextod[0:3], t.Hextod[3:11], t.Hextod[11:16], t.Year, t.Month, t.Day, 
		t.Hour, t.Minute, t.Second, t.Musec, offset, t.Year, t.Yday, wday, pmc)
}

// -----------------------------------------------------------
// nearInteger: tests whether a float is near to a whole number
func nearInteger(num float64) bool {
	return (math.Abs(num-round(num)) < 0.01)
}

// -----------------------------------------------------------
// round: rounds a float64 to the nearest whole number
func round(num float64) float64 {
	return float64(int(num + float64(math.Copysign(0.5, float64(num)))))
}

// -----------------------------------------------------------
// Todsearch:
//          finds the TOD clock value corresponding to
//          a timestamp and offset
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
func Todsearch(sdate string, stime string, offset float64) *Todinfo {
	var a *Todinfo
	var mdate string
	jdf := ( len(sdate) == 8 )
	bot := int64(0)
	top := int64(1) << 57
	lmid := bot
	counter := 0
	for mid := (bot + top) / 2; mid != lmid; mid, lmid = ((bot + top) / 2), mid {
		a = Todcalc(fmt.Sprintf("%x", mid), offset, true, false, false)
		if jdf {
			mdate = fmt.Sprintf("%04d.%03d", a.Year, a.Yday)
		} else {
			mdate = fmt.Sprintf("%04d-%02d-%02d", a.Year, a.Month, a.Day)
		}
		mtime := fmt.Sprintf("%02d:%02d:%02d.%06d", a.Hour, a.Minute, a.Second, a.Musec)
		switch {
		case sdate < mdate:
			{
				top = mid
			}
		case sdate > mdate:
			{
				bot = mid
			}
		case stime < mtime:
			{
				top = mid
			}
		case stime > mtime:
			{
				bot = mid
			}
		}
		counter++
		if counter > 65 {
			break
		}
	}
	return a
}

// -----------------------------------------------------------
// Todcalc:
//          converts a [16]uint8 tod clock or a PARS perpetual minute clock
//          and an hours float offset
//          to a date and time string
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
func Todcalc(s string, offset float64, lpad bool, rpad bool, pmcin bool) *Todinfo {
	const micro int64 = 24 * 3600 * 1000000
	var nx int64
	var pmc int64
	var input string
	ot := 1000000 * int64(round(offset*3600)) // offset in microseconds
	if pmcin {
		input = "00000000"[len(s):8] + s
		pmc =  tod2int((input))
		nx = ParsDayZero + pmc * 60000000 - ot 
		input = fmt.Sprintf("%016x", nx)
		// pmc = (nx - ParsDayZero) / 60000000
	} else {
		input = todpad(s, lpad, rpad)
		nx = tod2int(input)
		pmc = (nx - ParsDayZero + ot) / 60000000
	}
	dn := int64(ot + nx + 5*micro)            // extra five days to allow -ve offset
	mms := dn % micro
	sss := int64(mms / 1000000)
	hh := int64(sss / 3600)
	nn := int64(sss/60) % 60
	ss := sss % 60
	mms %= 1000000
	dn = int64(dn/micro) - 5
	wkday := int((dn + 1) % 7)
	// g = dn + 693901
	g := dn + 693901
	// y = 1 + int((10000*g + 14780)/3652425)
	y := int64((10000*g+14780)/3652425) + 1
	dd := int64(-1)
	for dd < 0 {
		y--
		dd = g - (365*y + y/4 - y/100 + y/400)
	}
	yd := dd + 60
	// mi = int((100*dd + 52)/3060)
	mi := int64((100*dd + 52) / 3060)
	// mm = (mi + 2)%12 + 1
	mm := (mi+2)%12 + 1
	// y = y + int((mi + 2)/12)
	if mi > 9 {
		y++
		yd -= 365
	}
	// Correct day-in-year for leap year
	if mm > 2 && (y%4 == 0 || y%400 == 0) && (y%100 != 0) {
		yd++
	}
	// dd = dd - int((mi*306 + 5)/10) + 1
	dd -= (mi*306+5)/10 - 1
	return &Todinfo{
		Intod:  s,
		Offset: offset,
		Hextod: input,
		Numtod: nx,
		Year:   y,
		Month:  mm,
		Day:    dd,
		Hour:   hh,
		Minute: nn,
		Second: ss,
		Musec:  mms,
		Yday:   yd,
		Wkday:  wkday,
		Pmc:    pmc,
	}
}

// -----------------------------------------------------------
// todpad:
//          pads a string of hex digits
//          to a 16-character string according to some
//          weird rules-with-a-purpose
//          digits not in {0..9A..F}: invalid - return 16x"f"
//          more than 16 digits: invalid - return 16x"f"
//          13 digits or more: pad to 16 on left with zeros
//          12 digits or less:
//             if first character > "B" prepend "0"
//             then pad to 13 on right with zeros
//             then pad to 16 on left with zeros
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
func todpad(s string, lp bool, rp bool) string {
	ff := "ffffffffffffffff"
	zz := "0000000000000000"
	sn := strings.ToLower(s)
	if len(sn) > 16 {
		return ff
	}
	switch {
	case rp:
		{
			if len(sn) < 16 {
				sn = (sn + zz)[0:16]
			}
		}
	case lp:
		{
			if len(sn) < 16 {
				sn = zz[len(sn):16] + sn
			}
		}
	default:
		{
			if len(sn) <= 12 {
				switch string(sn[0]) {
				case "c", "d", "e", "f":
					{
						sn = "0" + sn
					}
				}
				sn = (sn + zz)[0:14]
			}
			if len(sn) < 16 {
				sn = zz[len(sn):16] + sn
			}
		}
	}
	return sn
}

// -----------------------------------------------------------
// tod2int:
//          converts a string of hex digits
//          to a int64 number of milliseconds
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
func tod2int(s string) int64 {
	var r int64
	var x byte
	ss, err := hex.DecodeString(s)
	if err != nil {
		return 0
	}
	r = 0
	for _, x = range ss {
		r = r * 256 + int64(x)
	}
	return r
}
