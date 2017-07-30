// ehllapi: An interface to IBM and other 3270 terminal emulators.
package ehllapi

import (
	"encoding/binary"
	"errors"
	"fmt" 
	"syscall"
	"unsafe"
    "strings"
	"regexp"
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
	eQuerySess  ElFunc = 22 // Query session status
)

const (
	dllName     string = `ehlapi32.dll`
	dllProc     string = "HLLAPI"
)

var (
	reSysId,
	reKeys      *regexp.Regexp	
)

// Session: a 3270 session handle
type Session struct {
	SessId     string
	SysId      string
	Entry      string
	Result     []string
	Headers    int
	Footers    int
	Rows       int
	Cols       int
	Scroll     int
	Proc       *syscall.LazyProc
}

func init() {
	reSysId = regexp.MustCompile("ALCS VTAM application name \\.\\.\\.\\. +([A-Z0-9]{1,8}) ")
	reKeys  = regexp.MustCompile("\\[[]\\]")
}

// String: is the interface for printing Session structures
func (s Session) String() string {
	return fmt.Sprintf("%1s: %s",s.SessId,s.SysId) 
}

// NewSession: sets up a Session and connects with its ID 
func NewSession(sessid string) (Session, error) {
	if len(sessid) != 1 || sessid < "A" || sessid > "Z" {
		return Session{}, errors.New("1 Invalid session ID:" + sessid)
	}
	sid := strings.ToUpper(sessid)
	self := Session{
		SessId:    sid ,
		SysId:     "UNKNOWN" ,
		Headers:   0 ,  
		Footers:   1 ,
		Rows:      0 ,
	    Cols:      0 ,
		Scroll:    0 , 
	}
	return self, nil
} 

// Connect: connects to a presentation space and determines system ID

func (s Session) Connect() (uint32) {
// Load EHLLAPI DLL
	dll := syscall.NewLazyDLL(dllName)
	s.Proc = dll.NewProc(dllProc)
	
// Connect to session
	rc := s.ecall(eConnect,s.SessId,1,0)
	if rc == 0 {
		defer s.Disconnect()
	}

// Get rows and columns
	sessparm := "000000000000000000"[0:18]
	rc = s.ecall(eQuerySess,sessparm,18,0)
	s.Rows = int(binary.LittleEndian.Uint16([]byte(sessparm[11:13])))
	s.Cols = int(binary.LittleEndian.Uint16([]byte(sessparm[13:15])))

// Attempt to discover system name
	var eps uint32 = 0
	var buf = string(make([]byte,s.Rows*s.Cols))
	rc = s.ecall(eSend,"@C",2,eps)             // Clear screen
	rc = s.ecall(eWait,buf,1,eps)

	rc = s.ecall(eSend,"zdcom alcs@E",12,eps)  // Query ALCS and read output
	rc = s.ecall(eWait,buf,1,eps)
	rc = s.ecall(eReadScrn,buf,uint32(s.Rows*s.Cols),eps)

	m := reSysId.FindStringSubmatch(buf)
	if m != nil {
		s.SysId = m[1]
	}
	return rc
}

// Disconnect: disconnect from the current presentation space

func (s Session) Disconnect() (uint32) {
	rc := s.ecall(eDisconnect," ",1,0)
	return rc
}

// RunEntry: run an entry and capture the output

func (s Session) Write(entry string) (uint32) {
	var rc uint32 = 0
	return rc
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
