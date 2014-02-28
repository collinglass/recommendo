package main

import (
	"fmt"
	"github.com/collinglass/recommendo/algo"
	"github.com/collinglass/recommendo/data"
	"github.com/collinglass/recommendo/sort"
)

func recommend(userId int, users map[int]data.User, simFunc algo.SimFunc) []float64 {
	algo.GetSimilar(users, simFunc)

	reco := make(map[int]float64)
	simSum := make(map[int]float64)

	//	for all users
	// 		if not me,
	// 			Get similar users,
	//			if similarity > 0
	//				if I haven't read the book
	//					Similarity * book rating
	//					totalsimilarity[item] += user similarity
	//	Create Normalized list
	//	Return the sorted list.

	// for all users
	for id, value := range users {
		// if not me
		if id != userId {
			// if similarity of userId is bigger than 0
			if value.Similars[userId].Score > 0 {
				// for ratings
				for book, rating := range value.Ratings {
					// if I haven't read the book
					if users[userId].Ratings[book] == 0 {
						// set book recommendation value to user similarity * book rating
						reco[book] += value.Similars[userId].Score * rating
						// sum the similarity ratings for normalize
						simSum[book] += value.Similars[userId].Score
					}
				}
			}
		}
	}

	// Create Normalized list
	rankings := make([]float64, 6, 6)
	for book, recomendation := range reco {
		rankings[book] = recomendation / simSum[book]
	}

	// Return the sorted list

	sort.SortFloat(rankings)
	//rankings.Reverse()
	return rankings
}

func main() {
	// Populate data list
	users := data.Populate()

	// recommendation list
	fmt.Println(recommend(6, users, algo.Pearson))

	/* func(size of list, data.UserId, map of data.User, algo.SimFunc) */
	//fmt.Println("Pearson:")
	//fmt.Println(sort.TopMatches(3, 0, users, algo.Pearson))
	//fmt.Println("Euclidean:")
	//fmt.Println(sort.TopMatches(3, 0, users, algo.Euclidean))
}
