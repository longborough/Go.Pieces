// ST Pool analysis
package stplib

import (
	"math"
	"fmt"
	"errors"
)

const (
	delta float64 = 0.1E-9
	maxiter int = 5000
	NoConverge string = "Did not converge"
	Overloaded string = "Pool overloaded"
)

const DataHeader string = "Np,N1,N2,Rd,Pr,Tr,Pa,Ta,Nu,Nr,R0,R1,R2,R3,R4,Pu,Pp,Tc,Tt,Cc"

type STinfo struct {
// Input parameters
	Np float64      // Pool size
	N1 int        // ST Pool N1 (cycles to timeout)
	N2 int        // ST Pool N2 (cycles from release to dispense)
	Rd float64      // Dispense rate 
	Pr float64      // Probability of normal release
	Tr float64      // Mean time from get to release
	Pa float64      // Probability of peek-after-release
	Ta float64      // Mean time for peeking
// Output parameters
	Nu float64      // Mean number in use (also input as starting guess)
	Nr float64      // Mean cycles to normal release
	R0 float64      // Total release rate (Release + Timeouts)
	R1 float64      // Normal release rate
	R2 float64      // Never-release timeout rate
	R3 float64      // Unlucky timeout release rate 
	R4 float64      // Peeking rate
	Pu float64      // Probability a record will get an unlucky timeout
	Pp float64      // Probability a record will be peeked after redispense
	Tc float64      // Average single cycle time
	Tt float64      // Average timeout time (Recycle time)
	Cc int        // Convergence counter
}

// -----------------------------------------------------------
// String: provides the interface for printing STinfo structures
func (t STinfo) String() string {
	return fmt.Sprintf("%0.5f,%0d,%0d,%0.5f,%0.5f,%0.5f,%0.5f,%0.5f,%0.5f,%0.5f,%0.5f,%0.5f,%0.5f,%0.5f,%0.5f,%0.5f,%0.5f,%0.5f,%0.5f,%0d",
		t.Np,t.N1,t.N2,t.Rd,t.Pr,t.Tr,t.Pa,t.Ta,t.Nu,t.Nr,t.R0,t.R1,t.R2,t.R3,t.R4,t.Pu,t.Pp,t.Tc,t.Tt,t.Cc)
}

func STpoolan(in STinfo) (*STinfo, error) {
	out := in
	var last float64 = out.Tt
	if out.Nu <= 0.0 {
		out.Nu = out.Rd * out.Tr        // Initial guess at number in use if not provided
	}
	for out.Cc = 1 ; out.Cc <= maxiter ; out.Cc++ {
		if out.Nu >= out.Np {
			return &out, errors.New(Overloaded)
		}
		out.Tc = (out.Np - out.Nu) / out.Rd  // Average single cycle time
		out.Nr = out.Tr / out.Tc             // Average number of cycles
		out.Tt = float64(out.N1) * out.Tc      // Average recycle time
		out.Pu = Pdf(out.Nr,out.N1)          // Poisson probability of release before unlucky
		out.R0 = out.Rd                      // Total releases = total dispenses (steady-state)
		out.R1 = out.R0 * out.Pr             // Releases (RELF + unlucky timeout)
		out.R2 = out.Rd - out.R1             // Non-release timeouts
		out.R3 = out.Pu * out.R1             // Unlucky timeouts
		out.R1 = out.R1 - out.R3             // Actual RELFs
		out.Nu = out.R1 * out.Tr / float64(out.N1) + (out.R2 + out.R3) * out.Tc
//  Check for convergence
		if math.Abs(out.Tt - last) < delta * out.Tt {
			return &out, nil
		}
		last = out.Tt
	}
	return &out, errors.New(NoConverge)
}      

func Pdf(ncy float64, n1 int) float64 {
	var i int
	var x float64 = 1.0
	var p float64 = 1.0
	lm := ncy 
	for i = 1 ; i <= n1 ; i++ {
		x = x * lm / float64(i)
		p += x
	}
	return 1.0 - math.Exp(-lm) * p 
} 
