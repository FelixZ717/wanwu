package mcp

import (
	"context"

	errs "github.com/UnicomAI/wanwu/api/proto/err-code"
	mcp_service "github.com/UnicomAI/wanwu/api/proto/mcp-service"
	"github.com/UnicomAI/wanwu/internal/mcp-service/client/model"
	"github.com/UnicomAI/wanwu/pkg/util"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) CustomSkillCreate(ctx context.Context, req *mcp_service.CustomSkillCreateReq) (*mcp_service.CustomSkillCreateResp, error) {
	skillId, err := s.cli.CreateCustomSkill(ctx, &model.CustomSkill{
		Name:            req.Name,
		Avatar:          req.Avatar,
		Author:          req.Author,
		Desc:            req.Desc,
		SourceType:      req.SourceType,
		WgaThreadID:     req.WgaThreadId,
		PreviewThreadID: req.PreviewThreadId,
		UserID:          req.Identity.UserId,
		OrgID:           req.Identity.OrgId,
	})
	if err != nil {
		return nil, errStatus(errs.Code_MCPCustomSkillErr, err)
	}

	return &mcp_service.CustomSkillCreateResp{SkillId: skillId}, nil
}

func (s *Service) CustomSkillDelete(ctx context.Context, req *mcp_service.CustomSkillDeleteReq) (*emptypb.Empty, error) {
	err := s.cli.DeleteCustomSkill(ctx, req.SkillId)
	if err != nil {
		return nil, errStatus(errs.Code_MCPCustomSkillErr, err)
	}

	return &emptypb.Empty{}, nil
}

func (s *Service) CustomSkillGet(ctx context.Context, req *mcp_service.CustomSkillGetReq) (*mcp_service.CustomSkill, error) {
	customSkill, err := s.cli.GetCustomSkill(ctx, req.SkillId)
	if err != nil {
		return nil, errStatus(errs.Code_MCPCustomSkillErr, err)
	}

	variables, err := s.cli.GetCustomSkillVars(ctx, customSkill.UserID, customSkill.OrgID, req.SkillId)
	if err != nil {
		return nil, errStatus(errs.Code_MCPCustomSkillErr, err)
	}

	return toCustomSkillInfo(customSkill, toCustomSkillVariables(variables)), nil
}

func (s *Service) GetCustomSkillByThreadID(ctx context.Context, req *mcp_service.GetCustomSkillByThreadIDReq) (*mcp_service.GetCustomSkillByThreadIDResp, error) {
	skillId, st := s.cli.GetCustomSkillIDByWgaThreadID(ctx, req.GetIdentity().GetUserId(), req.GetIdentity().GetOrgId(), req.GetWgaThreadId())
	if st != nil {
		return nil, errStatus(errs.Code_MCPCustomSkillErr, st)
	}
	return &mcp_service.GetCustomSkillByThreadIDResp{SkillId: skillId}, nil
}

func (s *Service) CustomSkillGetList(ctx context.Context, req *mcp_service.CustomSkillGetListReq) (*mcp_service.CustomSkillGetListResp, error) {
	customSkills, total, err := s.cli.GetCustomSkillList(ctx, req.GetIdentity().GetUserId(), req.GetIdentity().GetOrgId(), req.Name)
	if err != nil {
		return nil, errStatus(errs.Code_MCPCustomSkillErr, err)
	}

	customSkillList := make([]*mcp_service.CustomSkill, 0, len(customSkills))
	for _, customSkill := range customSkills {
		variables, err := s.cli.GetCustomSkillVars(ctx, customSkill.UserID, customSkill.OrgID, util.Int2Str(customSkill.ID))
		if err != nil {
			return nil, errStatus(errs.Code_MCPCustomSkillErr, err)
		}
		customSkillList = append(customSkillList, toCustomSkillInfo(customSkill, toCustomSkillVariables(variables)))
	}

	return &mcp_service.CustomSkillGetListResp{
		List:  customSkillList,
		Total: total,
	}, nil
}

func (s *Service) CustomSkillGetBySaveIds(ctx context.Context, req *mcp_service.CustomSkillGetBySaveIdsReq) (*mcp_service.CustomSkillSaveIdsResp, error) {
	customSkills, err := s.cli.GetCustomSkillBySaveIds(ctx, req.SaveIds)
	if err != nil {
		return nil, errStatus(errs.Code_MCPCustomSkillErr, err)
	}

	saveIds := make([]string, 0, len(customSkills))
	for _, customSkill := range customSkills {
		saveIds = append(saveIds, customSkill.SaveId)
	}

	return &mcp_service.CustomSkillSaveIdsResp{
		SaveIds: saveIds,
	}, nil
}

func (s *Service) GetCustomSkillDetailByIdList(ctx context.Context, req *mcp_service.CustomSkillDetailByIdListReq) (*mcp_service.CustomSkillDetailByIdListResp, error) {
	customSkills, err := s.cli.GetCustomSkillBySkillIds(ctx, req.SkillIds)
	if err != nil {
		return nil, errStatus(errs.Code_MCPCustomSkillErr, err)
	}

	skillDetails := make([]*mcp_service.CustomSkill, 0, len(customSkills))
	for _, customSkill := range customSkills {
		variables, err := s.cli.GetCustomSkillVars(ctx, customSkill.UserID, customSkill.OrgID, util.Int2Str(customSkill.ID))
		if err != nil {
			return nil, errStatus(errs.Code_MCPCustomSkillErr, err)
		}
		skillDetails = append(skillDetails, toCustomSkillInfo(customSkill, toCustomSkillVariables(variables)))
	}

	return &mcp_service.CustomSkillDetailByIdListResp{
		SkillDetails: skillDetails,
	}, nil
}

func toCustomSkillInfo(customSkill *model.CustomSkill, variables []*mcp_service.Variable) *mcp_service.CustomSkill {
	if customSkill == nil {
		return nil
	}
	return &mcp_service.CustomSkill{
		SkillId:         util.Int2Str(customSkill.ID),
		Name:            customSkill.Name,
		Avatar:          customSkill.Avatar,
		Author:          customSkill.Author,
		Desc:            customSkill.Desc,
		ObjectPath:      customSkill.ObjectPath,
		WgaThreadId:     customSkill.WgaThreadID,
		PreviewThreadId: customSkill.PreviewThreadID,
		Variables:       variables,
		CreatedAt:       customSkill.CreatedAt,
		UpdatedAt:       customSkill.UpdatedAt,
	}
}

func toCustomSkillVariables(variables []*model.CustomSkillVariable) []*mcp_service.Variable {
	ret := make([]*mcp_service.Variable, 0, len(variables))
	for _, variable := range variables {
		ret = append(ret, &mcp_service.Variable{
			Id:            util.Int2Str(variable.ID),
			Name:          variable.Name,
			Desc:          variable.Desc,
			VariableKey:   variable.VariableKey,
			VariableValue: variable.VariableValue,
		})
	}
	return ret
}
