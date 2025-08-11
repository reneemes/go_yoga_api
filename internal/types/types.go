package types

type Pose struct {
	ID   						int    `json:"id" gorm:"primaryKey"`
	Name 						string `json:"name"`
	SanskritName 		string `json:"sanskrit_name"`
	TranslationName string `json:"translation_name"`
	Description 		string `json:"description"`
	PoseBenefits 		string `json:"pose_benefits"`
	ImageURL 				string `json:"image_url"`	
}

type Routine struct {
	ID          	int    	`json:"id" gorm:"primaryKey"`
	Name					string	`json:"name"`
	Description 	string 	`json:"description"`
	Difficulty 		string 	`json:"difficulty"`

	// A routine has many poses (through routine_poses)
	// A pose can be in many routines (through routine_poses)
	RoutinePoses []Pose `json:"routine_poses" gorm:"many2many:routine_poses;"`
	// Because we are using AutoMigrate() in the database,
	// a joins table (Routine_Poses) is created for us
}
