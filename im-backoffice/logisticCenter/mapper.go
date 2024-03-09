package logisticCenter

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	db "im-backoffice/db/sqlc"
)

func ToLogisticCenterDto(model db.ImLogisticCenter) LogisticCenterDTO {
	return LogisticCenterDTO{
		ID:                 model.ID.String(),
		LogisticCenterName: model.LogisticCenterName,
	}
}

func ToLogisticCenterDTOArray(models []db.ImLogisticCenter) (result []LogisticCenterDTO) {
	for _, val := range models {
		result = append(result, ToLogisticCenterDto(val))
	}
	return
}

func ToAllLogisticCentersResponse(dto []json.RawMessage, nextPage, totalPages int) AllLogisticCentersResponse {
	return AllLogisticCentersResponse{
		LogisticCenters: dto,
		NextPage:        nextPage,
		TotalPages:      totalPages,
	}
}

func ToLogisticCenterModel(dto LogisticCenterDTO) db.ImLogisticCenter {
	return db.ImLogisticCenter{
		ID:                 uuid.FromStringOrNil(dto.ID),
		LogisticCenterName: dto.LogisticCenterName,
	}
}

func ToLogisticCenterModelArray(dtoArray []LogisticCenterDTO) (result []db.ImLogisticCenter) {
	for _, val := range dtoArray {
		result = append(result, ToLogisticCenterModel(val))
	}
	return
}
