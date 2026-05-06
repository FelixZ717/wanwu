package orm

import (
	"context"
	"errors"

	errs "github.com/UnicomAI/wanwu/api/proto/err-code"
	"github.com/UnicomAI/wanwu/internal/mcp-service/client/model"
	"github.com/UnicomAI/wanwu/internal/mcp-service/client/orm/sqlopt"
	"gorm.io/gorm"
)

func (c *Client) CreateCustomSkillVar(ctx context.Context, userId, orgId string, variable *model.CustomSkillVariable) (uint32, *errs.Status) {
	if variable == nil || variable.SkillID == "" || variable.Name == "" {
		return 0, toErrStatus("mcp_skill_var_invalid_arg")
	}
	skill, st := c.GetCustomSkill(ctx, variable.SkillID)
	if st != nil {
		return 0, st
	}
	if skill.UserID != userId || skill.OrgID != orgId {
		return 0, toErrStatus("mcp_custom_skill_var_identity_mismatch")
	}
	variable.UserID = userId
	variable.OrgID = orgId
	if err := c.db.WithContext(ctx).Create(variable).Error; err != nil {
		return 0, toErrStatus("mcp_custom_skill_var_create", err.Error())
	}
	return variable.ID, nil
}

func (c *Client) UpdateCustomSkillVar(ctx context.Context, userId, orgId string, id uint32, variable *model.CustomSkillVariable) *errs.Status {
	if id == 0 || variable == nil || variable.Name == "" {
		return toErrStatus("mcp_skill_var_invalid_arg")
	}
	var row model.CustomSkillVariable
	if err := sqlopt.SQLOptions(
		sqlopt.WithID(id),
		sqlopt.WithUserID(userId),
		sqlopt.WithOrgID(orgId),
	).Apply(c.db).WithContext(ctx).First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return toErrStatus("mcp_custom_skill_var_not_found")
		}
		return toErrStatus("mcp_custom_skill_var_update", err.Error())
	}
	skill, st := c.GetCustomSkill(ctx, row.SkillID)
	if st != nil {
		return st
	}
	if skill.UserID != userId || skill.OrgID != orgId {
		return toErrStatus("mcp_custom_skill_var_identity_mismatch")
	}
	if err := sqlopt.SQLOptions(
		sqlopt.WithID(id),
		sqlopt.WithUserID(userId),
		sqlopt.WithOrgID(orgId),
	).Apply(c.db.WithContext(ctx).Model(&model.CustomSkillVariable{})).
		Updates(map[string]interface{}{
			"name":           variable.Name,
			"desc":           variable.Desc,
			"variable_key":   variable.VariableKey,
			"variable_value": variable.VariableValue,
		}).Error; err != nil {
		return toErrStatus("mcp_custom_skill_var_update", err.Error())
	}
	return nil
}

func (c *Client) DeleteCustomSkillVar(ctx context.Context, userId, orgId string, id uint32) *errs.Status {
	if id == 0 {
		return toErrStatus("mcp_skill_var_invalid_arg")
	}
	var row model.CustomSkillVariable
	if err := sqlopt.SQLOptions(
		sqlopt.WithID(id),
		sqlopt.WithUserID(userId),
		sqlopt.WithOrgID(orgId),
	).Apply(c.db).WithContext(ctx).First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return toErrStatus("mcp_custom_skill_var_not_found")
		}
		return toErrStatus("mcp_custom_skill_var_delete", err.Error())
	}
	skill, st := c.GetCustomSkill(ctx, row.SkillID)
	if st != nil {
		return st
	}
	if skill.UserID != userId || skill.OrgID != orgId {
		return toErrStatus("mcp_custom_skill_var_identity_mismatch")
	}
	if err := sqlopt.SQLOptions(
		sqlopt.WithID(id),
		sqlopt.WithUserID(userId),
		sqlopt.WithOrgID(orgId),
	).Apply(c.db).WithContext(ctx).Delete(&model.CustomSkillVariable{}).Error; err != nil {
		return toErrStatus("mcp_custom_skill_var_delete", err.Error())
	}
	return nil
}

func (c *Client) GetCustomSkillVars(ctx context.Context, userId, orgId, skillId string) ([]*model.CustomSkillVariable, *errs.Status) {
	if skillId == "" {
		return nil, toErrStatus("mcp_skill_var_invalid_arg")
	}
	var list []*model.CustomSkillVariable
	if err := sqlopt.SQLOptions(
		sqlopt.WithSkillID(skillId),
		sqlopt.WithUserID(userId),
		sqlopt.WithOrgID(orgId),
	).Apply(c.db).WithContext(ctx).Find(&list).Error; err != nil {
		return nil, toErrStatus("mcp_custom_skill_var_list", err.Error())
	}
	return list, nil
}

func (c *Client) CreateAcquiredSkillVar(ctx context.Context, userId, orgId string, variable *model.AcquiredSkillVariable) (uint32, *errs.Status) {
	if variable == nil || variable.SkillID == "" || variable.Name == "" {
		return 0, toErrStatus("mcp_skill_var_invalid_arg")
	}
	acquired, st := c.GetAcquiredSkill(ctx, variable.SkillID)
	if st != nil {
		return 0, st
	}
	if acquired.UserID != userId || acquired.OrgID != orgId {
		return 0, toErrStatus("mcp_acquired_skill_var_identity_mismatch")
	}
	variable.UserID = userId
	variable.OrgID = orgId
	if err := c.db.WithContext(ctx).Create(variable).Error; err != nil {
		return 0, toErrStatus("mcp_acquired_skill_var_create", err.Error())
	}
	return variable.ID, nil
}

func (c *Client) UpdateAcquiredSkillVar(ctx context.Context, userId, orgId string, id uint32, variable *model.AcquiredSkillVariable) *errs.Status {
	if id == 0 || variable == nil || variable.Name == "" {
		return toErrStatus("mcp_skill_var_invalid_arg")
	}
	var row model.AcquiredSkillVariable
	if err := sqlopt.SQLOptions(
		sqlopt.WithID(id),
		sqlopt.WithUserID(userId),
		sqlopt.WithOrgID(orgId),
	).Apply(c.db).WithContext(ctx).First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return toErrStatus("mcp_acquired_skill_var_not_found")
		}
		return toErrStatus("mcp_acquired_skill_var_update", err.Error())
	}
	acquired, st := c.GetAcquiredSkill(ctx, row.SkillID)
	if st != nil {
		return st
	}
	if acquired.UserID != userId || acquired.OrgID != orgId {
		return toErrStatus("mcp_acquired_skill_var_identity_mismatch")
	}
	if err := sqlopt.SQLOptions(
		sqlopt.WithID(id),
		sqlopt.WithUserID(userId),
		sqlopt.WithOrgID(orgId),
	).Apply(c.db.WithContext(ctx).Model(&model.AcquiredSkillVariable{})).
		Updates(map[string]interface{}{
			"name":           variable.Name,
			"desc":           variable.Desc,
			"variable_key":   variable.VariableKey,
			"variable_value": variable.VariableValue,
		}).Error; err != nil {
		return toErrStatus("mcp_acquired_skill_var_update", err.Error())
	}
	return nil
}

func (c *Client) DeleteAcquiredSkillVar(ctx context.Context, userId, orgId string, id uint32) *errs.Status {
	if id == 0 {
		return toErrStatus("mcp_skill_var_invalid_arg")
	}
	var row model.AcquiredSkillVariable
	if err := sqlopt.SQLOptions(
		sqlopt.WithID(id),
		sqlopt.WithUserID(userId),
		sqlopt.WithOrgID(orgId),
	).Apply(c.db).WithContext(ctx).First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return toErrStatus("mcp_acquired_skill_var_not_found")
		}
		return toErrStatus("mcp_acquired_skill_var_delete", err.Error())
	}
	acquired, st := c.GetAcquiredSkill(ctx, row.SkillID)
	if st != nil {
		return st
	}
	if acquired.UserID != userId || acquired.OrgID != orgId {
		return toErrStatus("mcp_acquired_skill_var_identity_mismatch")
	}
	if err := sqlopt.SQLOptions(
		sqlopt.WithID(id),
		sqlopt.WithUserID(userId),
		sqlopt.WithOrgID(orgId),
	).Apply(c.db).WithContext(ctx).Delete(&model.AcquiredSkillVariable{}).Error; err != nil {
		return toErrStatus("mcp_acquired_skill_var_delete", err.Error())
	}
	return nil
}

func (c *Client) GetAcquiredSkillVars(ctx context.Context, userId, orgId, skillId string) ([]*model.AcquiredSkillVariable, *errs.Status) {
	if skillId == "" {
		return nil, toErrStatus("mcp_skill_var_invalid_arg")
	}
	var list []*model.AcquiredSkillVariable
	if err := sqlopt.SQLOptions(
		sqlopt.WithSkillID(skillId),
		sqlopt.WithUserID(userId),
		sqlopt.WithOrgID(orgId),
	).Apply(c.db).WithContext(ctx).Find(&list).Error; err != nil {
		return nil, toErrStatus("mcp_acquired_skill_var_list", err.Error())
	}
	return list, nil
}

func (c *Client) CreateBuiltinSkillVar(ctx context.Context, userId, orgId string, variable *model.BuiltinSkillVariable) (uint32, *errs.Status) {
	if variable == nil || variable.SkillID == "" || variable.Name == "" {
		return 0, toErrStatus("mcp_skill_var_invalid_arg")
	}
	variable.UserID = userId
	variable.OrgID = orgId
	if err := c.db.WithContext(ctx).Create(variable).Error; err != nil {
		return 0, toErrStatus("mcp_builtin_skill_var_create", err.Error())
	}
	return variable.ID, nil
}

func (c *Client) UpdateBuiltinSkillVar(ctx context.Context, userId, orgId string, id uint32, variable *model.BuiltinSkillVariable) *errs.Status {
	if id == 0 || variable == nil || variable.Name == "" {
		return toErrStatus("mcp_skill_var_invalid_arg")
	}
	var row model.BuiltinSkillVariable
	if err := sqlopt.SQLOptions(
		sqlopt.WithID(id),
		sqlopt.WithUserID(userId),
		sqlopt.WithOrgID(orgId),
	).Apply(c.db).WithContext(ctx).First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return toErrStatus("mcp_builtin_skill_var_not_found")
		}
		return toErrStatus("mcp_builtin_skill_var_update", err.Error())
	}
	if err := sqlopt.SQLOptions(
		sqlopt.WithID(id),
		sqlopt.WithUserID(userId),
		sqlopt.WithOrgID(orgId),
	).Apply(c.db.WithContext(ctx).Model(&model.BuiltinSkillVariable{})).
		Updates(map[string]interface{}{
			"name":           variable.Name,
			"desc":           variable.Desc,
			"variable_key":   variable.VariableKey,
			"variable_value": variable.VariableValue,
		}).Error; err != nil {
		return toErrStatus("mcp_builtin_skill_var_update", err.Error())
	}
	return nil
}

func (c *Client) DeleteBuiltinSkillVar(ctx context.Context, userId, orgId string, id uint32) *errs.Status {
	if id == 0 {
		return toErrStatus("mcp_skill_var_invalid_arg")
	}
	var row model.BuiltinSkillVariable
	if err := sqlopt.SQLOptions(
		sqlopt.WithID(id),
		sqlopt.WithUserID(userId),
		sqlopt.WithOrgID(orgId),
	).Apply(c.db).WithContext(ctx).First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return toErrStatus("mcp_builtin_skill_var_not_found")
		}
		return toErrStatus("mcp_builtin_skill_var_delete", err.Error())
	}
	if err := sqlopt.SQLOptions(
		sqlopt.WithID(id),
		sqlopt.WithUserID(userId),
		sqlopt.WithOrgID(orgId),
	).Apply(c.db).WithContext(ctx).Delete(&model.BuiltinSkillVariable{}).Error; err != nil {
		return toErrStatus("mcp_builtin_skill_var_delete", err.Error())
	}
	return nil
}

// GetBuiltinSkillVars 返回的 total：无分页，表示当前 skill 下变量全量条数，与 len(variables) 一致。
func (c *Client) GetBuiltinSkillVars(ctx context.Context, userId, orgId, skillId string) ([]*model.BuiltinSkillVariable, int64, *errs.Status) {
	if skillId == "" {
		return nil, 0, toErrStatus("mcp_skill_var_invalid_arg")
	}
	var list []*model.BuiltinSkillVariable
	if err := sqlopt.SQLOptions(
		sqlopt.WithSkillID(skillId),
		sqlopt.WithUserID(userId),
		sqlopt.WithOrgID(orgId),
	).Apply(c.db).WithContext(ctx).Find(&list).Error; err != nil {
		return nil, 0, toErrStatus("mcp_builtin_skill_var_list", err.Error())
	}
	return list, int64(len(list)), nil
}
