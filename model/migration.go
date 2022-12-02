package model

func migration() {
	err := DB.AutoMigrate(&User{}, &Task{}, &Team{})
	if err != nil {
		return
	}
}
