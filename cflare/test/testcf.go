// test: test bed for CloudFlare tools
package main

import (
//    "github.com/longborough/cflare"
    "fmt"
    "log"
    cloudflare "github.com/cloudflare/cloudflare-go"
)

func main() {
api, err := cloudflare.New("30c7cae7a81b320413944634f25a0c6fd8694", "brent@longborough.org")
if err != nil {
    log.Fatal(err)
}

zoneID, err := api.ZoneIDByName("llwyd-consulting.cymru")
if err != nil {
    log.Fatal(err)
}

// Fetch only A type records
aaaa := cloudflare.DNSRecord{Type: "A"}
recs, err := api.DNSRecords(zoneID, aaaa)
if err != nil {
    log.Fatal(err)
}

for _, r := range recs {
    fmt.Printf("%-40s: %s -- %7d %s\n", r.Name, r.Content, r.TTL, r.ID)
    r.Content = "87.113.224.192"
    r.Content = "64.98.145.30"
    r.TTL = 900
    api.UpdateDNSRecord(zoneID, r.ID, r)
}
}
