package position_usecase

func (pu *PositionUseCase) DeletePosition(id int) error {
	return pu.pr.DeletePosition(id)
}
