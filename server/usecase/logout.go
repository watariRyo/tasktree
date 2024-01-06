package usecase

func (u *UseCaseImpl) Logout(ssid string) error {
	err := u.repo.RedisClient.DeleteSession(ssid)
	if err != nil {
		return err
	}
	return nil
}
