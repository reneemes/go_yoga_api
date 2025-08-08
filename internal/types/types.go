package types

type Pose struct {
	ID   						int    `json:"id" gorm:"primaryKey"`
	Name 						string `json:"name"`
	SanskritName 		string `json:"sanskrit_name"`
	TranslationName string `json:"translation_name"`
	Description 		string `json:"description"`
	PoseBenefits 		string `json:"pose_benefits"`
	ImageURL 				string `json:"image_url"`

	// One Pose can appear in many RoutinePoses
	// RoutinePoses []RoutinePose `gorm:"foreignKey:PoseID"`

	// Many-to-many link via RoutinePoses
	// Routines []Routine `gorm:"many2many:routine_poses;joinForeignKey:PoseID;joinReferences:RoutineID"`
}

type Routine struct {
	ID          	int    	`json:"id" gorm:"primaryKey"`
	Name					string	`json:"name"`
	Description 	string 	`json:"description"`
	Difficulty 		string 	`json:"difficulty"`

	// One Routine can have many RoutinePoses
	// RoutinePoses []RoutinePose `gorm:"foreignKey:RoutineID"`

	// Many-to-many link via RoutinePoses
	RoutinePoses []Pose `json:"routine_poses" gorm:"many2many:routine_poses;joinForeignKey:RoutineID;joinReferences:PoseID"`
}

type RoutinePose struct {
	RoutineID 	int `gorm:"primaryKey"`
	PoseID    	int `gorm:"primaryKey"`

	// Belongs to Routine
	Routine Routine `gorm:"foreignKey:RoutineID"`

	// Belongs to Pose
	Pose Pose `gorm:"foreignKey:PoseID"`
}
/*
	RoutinePoses []RoutinePose on both models lets you query the join table directly if needed.
	many2many:routine_poses; tells GORM the name of the join table for higher-level association queries.
	joinForeignKey and joinReferences make it explicit which fields map — good practice when your foreign keys aren’t named in GORM’s default style.
*/