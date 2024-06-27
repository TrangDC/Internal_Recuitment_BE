package dto

import (
	"sort"
	"trec/ent"

	"github.com/samber/lo"
)

type EntitySkillDto interface {
	GroupSkillType(input []*ent.EntitySkill) []*ent.EntitySkillType
}

type entitySkillDtoImpl struct {
}

func NewEntitySkillDto() EntitySkillDto {
	return &entitySkillDtoImpl{}
}

func (dto *entitySkillDtoImpl) GroupSkillType(input []*ent.EntitySkill) []*ent.EntitySkillType {
	var result []*ent.EntitySkillType
	for _, entity := range input {
		_, exist := lo.Find(result, func(entity2 *ent.EntitySkillType) bool {
			return entity2.ID == entity.Edges.SkillEdge.Edges.SkillTypeEdge.ID.String()
		})
		if !exist {
			result = append(result, &ent.EntitySkillType{
				ID:      entity.Edges.SkillEdge.Edges.SkillTypeEdge.ID.String(),
				Name:    entity.Edges.SkillEdge.Edges.SkillTypeEdge.Name,
				OrderID: entity.OrderID / 1000,
			})
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].OrderID < result[j].OrderID
	})
	for _, entity := range result {
		skills := lo.Filter(input, func(entity2 *ent.EntitySkill, index int) bool {
			return entity2.Edges.SkillEdge.Edges.SkillTypeEdge.ID.String() == entity.ID
		})
		sort.Slice(skills, func(i, j int) bool {
			return skills[i].OrderID < skills[j].OrderID
		})
		newSkills := lo.Map(skills, func(entity2 *ent.EntitySkill, index int) *ent.EntitySkillRecord {
			return &ent.EntitySkillRecord{
				ID:      entity2.ID.String(),
				Name:    entity2.Edges.SkillEdge.Name,
				OrderID: entity2.OrderID,
			}
		})
		entity.Skills = newSkills
	}
	return result
}
