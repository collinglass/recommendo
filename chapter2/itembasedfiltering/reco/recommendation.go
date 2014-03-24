package reco

import (
	"github.com/collinglass/recommendo/itembasedfiltering/algo"
	"github.com/collinglass/recommendo/itembasedfiltering/data"
)

func Recommend(userId int, users map[int]data.User, simFunc algo.SimFunc) {
	algo.GetSimilar(users, simFunc)

	reco := make(map[int]float64)
	simSum := make(map[int]float64)

	// for all users
	for id, value := range users {
		// if not me AND if similarity of userId is bigger than 0
		if id != userId && value.Similars[userId].Score > 0 {
			// for all book ratings
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

	// Create Normalized list
	for book, recomendation := range reco {
		users[userId].Recommended[book] = data.Recommendation{book, recomendation / simSum[book]}
	}
}
