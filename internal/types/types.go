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

	// One Routine has many RoutinePoses
	RoutinePoses []Pose `json:"routine_poses" gorm:"many2many:routine_poses;"`
}
