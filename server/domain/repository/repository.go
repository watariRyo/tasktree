package repository

type AllRepository struct {
	DBConnection        DBConnection
	DBTransaction       DBTransaction
	RedisClient         RedisClient
	FirebaseApp         FirebaseApp
	BaseTasksRepository BaseTasksRepository
	UsersRepository     UsersRepository
}
