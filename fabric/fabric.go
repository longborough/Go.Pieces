// Fabric cutting from the roll
package main

import (
	"fmt"
	"log"
	"github.com/gedex/bp3d"
)

func NewPiece(n string, w, h float64) *bp3d.Item {
	const shrink float64 = 1
	return bp3d.NewItem(n, shrink*w+1, shrink*h+1, 1, 0.01)
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
	fmt.Printf("Name,X,Y,Z,DX,DY,R,%d\n",len(bins[0].Items))
	for q, b := range bins {
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
	}
}

func main() {
	var rollwidth float64 = 148
	var rollength float64 = 1210
	p := bp3d.NewPacker()

	// One Roll
	p.AddBin(bp3d.NewBin("Fabric Roll 1 ", rollwidth, rollength, 1, 100))

	// Add items.
	// Footstool
	p.AddItem(NewPiece("Stool Main   1 ",   58,    59))
	p.AddItem(NewPiece("Stool Back   1 ",   19,    59))
	p.AddItem(NewPiece("Stool Back   2 ",   19,    59))
	p.AddItem(NewPiece("Stool Edge   1 ",   19,    35))
	p.AddItem(NewPiece("Stool Edge   2 ",   19,    35))
	// Cushions
	p.AddItem(NewPiece("Cushion Main 1 ",   73,    67))
	p.AddItem(NewPiece("Cushion Main 2 ",   73,    67))
	p.AddItem(NewPiece("Cushion Main 3 ",   73,    67))
	p.AddItem(NewPiece("Cushion Main 4 ",   73,    67))
	p.AddItem(NewPiece("Cushion Main 5 ",   73,    67))
	p.AddItem(NewPiece("Cushion Main 6 ",   73,    67))
	p.AddItem(NewPiece("Cushion Main 7 ",   73,    67))
	p.AddItem(NewPiece("Cushion Main 8 ",   73,    67))
	p.AddItem(NewPiece("Cushion Side 1 ",   16,   209))
	p.AddItem(NewPiece("Cushion Side 2 ",   16,   209))
	p.AddItem(NewPiece("Cushion Side 3 ",   16,   209))
	p.AddItem(NewPiece("Cushion Side 4 ",   16,   209))
	p.AddItem(NewPiece("Cushion Zip  1 ",   73,    16))
	p.AddItem(NewPiece("Cushion Zip  2 ",   73,    16))
	p.AddItem(NewPiece("Cushion Zip  3 ",   73,    16))
	p.AddItem(NewPiece("Cushion Zip  4 ",   73,    16))
	// Armchairs
	p.AddItem(NewPiece("Chair Seat   1 ",   46,   196))
	p.AddItem(NewPiece("Chair Seat   2 ",   46,   196))
	p.AddItem(NewPiece("Chair Seat   3 ",   46,   196))
	p.AddItem(NewPiece("Chair Seat   4 ",   46,   196))
	p.AddItem(NewPiece("Chair Back   1 ",   48,    88))
	p.AddItem(NewPiece("Chair Back   2 ",   48,    88))
	p.AddItem(NewPiece("Chair Back   3 ",   48,    88))
	p.AddItem(NewPiece("Chair Back   4 ",   48,    88))
   	// Throne
	p.AddItem(NewPiece("Throne Seat  1 ",   46,   259))
	p.AddItem(NewPiece("Throne Seat  2 ",   46,   259))
	p.AddItem(NewPiece("Throne Back  1 ",   48,   146))
	p.AddItem(NewPiece("Throne Back  2 ",   48,   146))

	// Pack items to bins.
	if err := p.Pack(); err != nil {
		log.Fatal(err)
	}
	displayPackedCSV(p.Bins)
}
