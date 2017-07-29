package main
import (
	"fmt"
	"github.com/llgcode/draw2d/draw2dpdf"
	"github.com/llgcode/draw2d/samples/android"
)

func main(){  
	// Initialize the graphic context on a pdf document
	dest := draw2dpdf.NewPdf("L", "mm", "A4")
	gc := draw2dpdf.NewGraphicContext(dest)
	// Draw Android logo
	fn, err := android.Main(gc, "pdf")
	if err != nil {
		fmt.Printf("Drawing %q failed: %v\n", fn, err)
		return
	}
	// Save to pdf
	err = draw2dpdf.SaveToPdfFile(fn, dest)
	if err != nil {
		fmt.Printf("Saving %q failed: %v\n", fn, err)
	}
}
