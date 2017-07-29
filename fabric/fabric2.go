// Fabric cutting from the roll
package main

import (
	"fmt"
	"log"
	"github.com/gedex/bp3d"
)

func NewPiece(n string, w, h float64) *bp3d.Item {
//	const shrink float64 = 1
	return bp3d.NewItem(n, h+1, w+1, 1, 1)
}

func displayPacked(bins []*bp3d.Bin) {
	fmt.Println("X,Y,DX,DY")
	for _, b := range bins {
		for _, i := range b.Items {
			fmt.Printf("%s,%s,%0.1f,%0.1f,%d\n", i.GetName(), i.Position, i.GetWidth(), i.GetHeight(),i.RotationType)
		}
	}
}

func displayPackedCSV(bins []*bp3d.Bin) {
	for q, b := range bins {
		fmt.Printf("Name,X,Y,DX,DY,R,B,%d\n",len(bins[q].Items))
		for _, i := range b.Items {
			if i.RotationType == bp3d.RotationType_WHD {
				fmt.Printf("%s,%6.1f,%6.1f,%6.1f,%6.1f,%d,%d\n", i.GetName(),
					i.Position[1], i.Position[0],
					i.GetHeight(), i.GetWidth(),i.RotationType, q)
			} else {
				fmt.Printf("%s,%6.1f,%6.1f,%6.1f,%6.1f,%d,%d\n", i.GetName(),
					i.Position[1], i.Position[0],
					i.GetWidth(), i.GetHeight(), i.RotationType, q)
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	var rollwidth float64 = 148
	var rollength float64 = 1600
	p := bp3d.NewPacker()

	// One Roll
	p.AddBin(bp3d.NewBin("Fabric Roll 1 ", rollwidth, rollength, 1, 1000))
//	p.AddBin(bp3d.NewBin("Fabric Roll 2 ", rollwidth, rollength, 1, 1000))
//	p.AddBin(bp3d.NewBin("Fabric Roll 3 ", rollwidth, rollength, 1, 1000))

	// Add items.
	// Footstool (S.)
	p.AddItem(NewPiece("S.Zip    1 ",   35,    61))
	p.AddItem(NewPiece("S.NZip   1 ",   64,    61))
	p.AddItem(NewPiece("S.End    1 ",   18,    32))
	p.AddItem(NewPiece("S.End    2 ",   18,    32))

	// Throne (T.)
	p.AddItem(NewPiece("T.LSeat 1F ",   45,    71))
	p.AddItem(NewPiece("T.LBack 1F ",   72,    55))
	p.AddItem(NewPiece("T.LArm  1F ",   61,    37))
	p.AddItem(NewPiece("T.LSeat 1B ",   45,    71))
	p.AddItem(NewPiece("T.LBack 1B ",   72,    55))
	p.AddItem(NewPiece("T.LArm  1B ",   61,    37))
	p.AddItem(NewPiece("T.RSeat 1F ",   45,    71))
	p.AddItem(NewPiece("T.RBack 1F ",   72,    55))
	p.AddItem(NewPiece("T.RArm  1F ",   61,    37))
	p.AddItem(NewPiece("T.RSeat 1B ",   45,    71))
	p.AddItem(NewPiece("T.RBack 1B ",   72,    55))
	p.AddItem(NewPiece("T.RArm  1B ",   61,    37))

	// Throne Cushions (T.)
	p.AddItem(NewPiece("T.Cush Zip  1  ",   55,    19))
	p.AddItem(NewPiece("T.Cush NZip 1  ",  214,    19))
	p.AddItem(NewPiece("T.Cush Main 1F ",   75,    62))
	p.AddItem(NewPiece("T.Cush Main 1B ",   75,    62))

	p.AddItem(NewPiece("T.Cush Zip  2  ",   55,    19))
	p.AddItem(NewPiece("T.Cush NZip 2  ",  214,    19))
	p.AddItem(NewPiece("T.Cush Main 2F ",   75,    62))
	p.AddItem(NewPiece("T.Cush Main 2B ",   75,    62))
	

	// Armchairs (A.)
	p.AddItem(NewPiece("A.Back   1B ",   56,   87))
	p.AddItem(NewPiece("A.Back   1F ",   56,   87))
	p.AddItem(NewPiece("A.LArm   1B ",   36,   59))
	p.AddItem(NewPiece("A.LArm   1F ",   36,   59))
	p.AddItem(NewPiece("A.RArm   1B ",   36,   59))
	p.AddItem(NewPiece("A.RArm   1F ",   36,   59))
	p.AddItem(NewPiece("A.Seat   1B ",   45,   87))
	p.AddItem(NewPiece("A.Seat   1F ",   45,   87))

	p.AddItem(NewPiece("A.Back   2B ",   56,   87))
	p.AddItem(NewPiece("A.Back   2F ",   56,   87))
	p.AddItem(NewPiece("A.LArm   2B ",   36,   59))
	p.AddItem(NewPiece("A.LArm   2F ",   36,   59))
	p.AddItem(NewPiece("A.RArm   2B ",   36,   59))
	p.AddItem(NewPiece("A.RArm   2F ",   36,   59))
	p.AddItem(NewPiece("A.Seat   2B ",   45,   87))
	p.AddItem(NewPiece("A.Seat   2F ",   45,   87))

	// Armchair Cushions (A.)
	p.AddItem(NewPiece("A.Cush Zip  1  ",   80,    17))
	p.AddItem(NewPiece("A.Cush NZip 1  ",  170,    17))
	p.AddItem(NewPiece("A.Cush Main 1F ",   72,    63))
	p.AddItem(NewPiece("A.Cush Main 1B ",   72,    63))

	p.AddItem(NewPiece("A.Cush Zip  2  ",   80,    17))
	p.AddItem(NewPiece("A.Cush NZip 2  ",  170,    17))
	p.AddItem(NewPiece("A.Cush Main 2F ",   72,    63))
	p.AddItem(NewPiece("A.Cush Main 2B ",   72,    63))

	// Pack items to bins.
	if err := p.Pack(); err != nil {
		log.Fatal(err)
	}
	displayPackedCSV(p.Bins)
}
