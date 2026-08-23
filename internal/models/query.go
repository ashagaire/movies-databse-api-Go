package models

type QueryFilter struct {
	Page   int    
	Size   int    // Items per page
	Year   int    
	GenreID int64 
}