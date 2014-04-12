package main

import (
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	//"code.google.com/p/plotinum/plotutil"
	"github.com/collinglass/recommendo/cluster"
	"github.com/collinglass/recommendo/data"
	"github.com/collinglass/recommendo/reco"
	"log"
)

type recoFunc func() (map[int]data.Recommendation, map[int]data.Recommendation)
type clusterFunc func() map[int]map[int]int

func main() {
	log.Println("Starting Algorithms...")

	plotrunner("user_based_reco", reco.UserRunner)

	plotrunner("k_means_cluster", cluster.ClusterRunner)

	plotrunner("item_based_reco", reco.ItemRunner)

	log.Println("Finishing Algorithms...")
}

func plotrunner(filename string, funcer recoFunc) {

	recoData, _ := funcer()
	scatterData := createPlot(recoData)

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
	if err := p.Save(4, 4, filename+".png"); err != nil {
		panic(err)
	}
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

func plotrunner(filename string, funcer clusterFunc) {

	recoData := funcer()
	scatterData := createPlot(recoData)

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
	if err := p.Save(4, 4, filename+".png"); err != nil {
		panic(err)
	}
}

func createPlot(input map[int]map[int]int) plotter.XYs {
	pts := make(plotter.XYs, len(input))
	j := 0
	for i, cluster := range input {
		for _, value := range cluster {
			pts[j].X = float64(value)
			pts[j].Y = float64(value)
			j++
		}
	}
	return pts
}
