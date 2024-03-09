package location

func ToLocationResponse(dto []LocationDTO, nextPage, totalPages int) LocationResponse {
	return LocationResponse{
		Locations:  dto,
		NextPage:   nextPage,
		TotalPages: totalPages,
	}
}
