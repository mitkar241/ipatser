package movies

/*
##########
# Movie related Struct
##########
*/
type Movie struct {
	MovieID   string `json:"movieid"`
	MovieName string `json:"moviename"`
}

/*
##########
# REST API response Struct
##########
*/
type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Movie `json:"data"`
	Message string  `json:"message"`
}
