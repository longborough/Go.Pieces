// cflare: Tools for loudFlare
package cflare

import (
    md "github.com/miekg/dns"
)

func MyIP() string {
    var Arecord string
    target := "myip.opendns.com"
    server := "resolver1.opendns.com"

    c := md.Client{}
    m := md.Msg{}
    m.SetQuestion(target+".", md.TypeA)
    r, _, err := c.Exchange(&m, server+":53")
    if err != nil {
        return "Shriek!"
    }
    
    for _, ans := range r.Answer {
        Arecord = ans.(*md.A).A.String()
    }
    return Arecord
}
