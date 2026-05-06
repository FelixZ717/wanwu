package model

type AcquiredSkillVariable struct {
	ID            uint32 `gorm:"primarykey"`
	SkillID       string `gorm:"column:skill_id;index:idx_acquired_skill_variable_skill_id;comment:skill id"`
	UserID        string `gorm:"column:user_id;index:idx_acquired_skill_variable_user_id;comment:用户id"`
	OrgID         string `gorm:"column:org_id;index:idx_acquired_skill_variable_org_id;comment:组织id"`
	Name          string `gorm:"column:name;index:idx_acquired_skill_variable_name;comment:工具名称"`
	Desc          string `gorm:"column:desc;type:text;comment:描述"`
	VariableKey   string `gorm:"column:variable_key;comment:变量Key"`
	VariableValue string `gorm:"column:variable_value;comment:变量Value"`
	CreatedAt     int64  `gorm:"column:created_at;default:0;comment:创建时间"`
	UpdatedAt     int64  `gorm:"column:updated_at;default:0;comment:更新时间"`
}
