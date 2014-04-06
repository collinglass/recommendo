package main

import (
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	//"code.google.com/p/plotinum/plotutil"
	"github.com/collinglass/recommendo/data"
	"github.com/collinglass/recommendo/reco"
	"log"
)

func main() {
	log.Println("Starting Algorithms...")

	ureco, _ := reco.UserRunner()
	//ireco, ireco2 := reco.ItemRunner()
	scatterData := createPlot(ureco)

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "User: Book Ratings"
	p.X.Label.Text = "Book"
	p.Y.Label.Text = "Rating"

	s, err := plotter.NewScatter(scatterData)
	if err != nil {
		panic(err)
	}
	p.Add(s)

	// Save the plot to a PNG file.
	if err := p.Save(4, 4, "points.png"); err != nil {
		panic(err)
	}
	log.Println("Finishing Algorithms...")
}

func createPlot(input map[int]data.Recommendation) plotter.XYs {
	pts := make(plotter.XYs, len(input))
	j := 0
	for i := range input {
		pts[j].X = float64(input[i].Book)
		pts[j].Y = input[i].Score
		j++
	}
	return pts
}
