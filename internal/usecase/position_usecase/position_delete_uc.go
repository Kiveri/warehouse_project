package position_usecase

func (u *PositionUseCase) DeletePositionUC(id int) error {
	return u.r.DeletePosition(id)
}
