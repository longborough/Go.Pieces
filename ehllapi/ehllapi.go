// ehllapi: An interface to IBM and other 3270 terminal emulators.
package ehllapi

import (
	"errors"
	"fmt" 
	"syscall"
	"unsafe"
    "strings"
)

const version string = "0.1.0"

// GetVersion returns the current version of the package.
func GetVersion() string {
	return version
}

// ElFunc: EHLLAPI Functions
type ElFunc uint32

const (
	eConnect    ElFunc = 1  // Connect to a terminal
	eDisconnect ElFunc = 2  // Disconnect from a terminal
	eSend       ElFunc = 3  // Send keystrokes to a terminal
	eWait       ElFunc = 4  // Wait for a terminal to unlock
	eReadScrn   ElFunc = 8  // Read the screen from a terminal
	dllName     string = `ehlapi32.dll`
	dllProc     string = "HLLAPI"
)

// Session: a 3270 session handle
type Session struct {
	SessId     string
	SysId      string
	Headers    int
	Footers    int
	ScrollTop  bool
	ScrollMid  bool
	ScrollBot  bool
	Proc       *syscall.LazyProc
}

// String: is the interface for printing Session structures
func (s Session) String() string {
	return fmt.Sprintf("%1s: %8s",s.SessId,s.SysId) 
}

// NewSession: sets up a Session and connects with its ID 
func NewSession(sessid string) (Session, error) {
	if len(sessid) != 1 || sessid < "A" || sessid > "Z" {
		return Session{}, errors.New("1 Invalid session ID:" + sessid)
	}
	sid := strings.ToUpper(sessid)
	
	dll := syscall.NewLazyDLL(dllName)
	prc := dll.NewProc(dllProc)
	
	self := Session{
		SessId:    sid ,
		SysId:     "" ,
		Headers:   0 ,  
		Footers:   1 ,
		ScrollTop: true , 
		ScrollMid: true ,
		ScrollBot: true ,
		Proc:      prc ,
	}
	self.Connect()
	return self, nil
} 

// ecall: provides simplified EHLLAPI DLL calls

func (s Session) ecall(fun ElFunc, buffer string, count uint32, ps uint32) (uint32) {
	rc := ps
	_, _, _ = s.Proc.Call(
		uintptr(unsafe.Pointer(uintptr(uint32(fun)))),
		uintptr(unsafe.Pointer(&buffer)),
		uintptr(unsafe.Pointer(uintptr(count))),
		uintptr(unsafe.Pointer(uintptr(rc))),
	)
	return rc
}

// Connect: connects to a presentation space and determines system ID

func (s Session) Connect() (uint32) {
	rc := s.ecall(eConnect,s.SessId,1,0)
	if rc == 0 {
		defer s.Disconnect()
	}
	return rc
}

// Disconnect: disconnect from the current presentation space

func (s Session) Disconnect() (uint32) {
	rc := s.ecall(eDisconnect," ",1,0)
	return rc
}
