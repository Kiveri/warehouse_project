package position_usecase

func (pu *PositionUseCase) DeletePositionUC(id int) error {
	return pu.pr.DeletePosition(id)
}
