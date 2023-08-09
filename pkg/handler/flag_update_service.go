package handler

import (
	"github.com/openflagr/flagr/pkg/entity"
	"github.com/openflagr/flagr/pkg/util"
	"github.com/openflagr/flagr/swagger_gen/models"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/distribution"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/flag"
	"gorm.io/gorm"
)

func UpdateFlag(params flag.PutFlagParams, flagRecord *entity.Flag, tx *gorm.DB) error {
	if params.Body.Description != nil {
		flagRecord.Description = *params.Body.Description
	}
	if params.Body.DataRecordsEnabled != nil {
		flagRecord.DataRecordsEnabled = *params.Body.DataRecordsEnabled
	}
	if params.Body.Key != nil {
		key, err := entity.CreateFlagKey(*params.Body.Key)
		if err != nil {
			return err
		}
		flagRecord.Key = key
	}
	if params.Body.EntityType != nil {
		et := *params.Body.EntityType
		if err := entity.CreateFlagEntityType(tx, et); err != nil {
			return err
		}
		flagRecord.EntityType = et
	}

	if params.Body.Notes != nil {
		flagRecord.Notes = *params.Body.Notes
	}

	return nil
}

func UpdateSegments(segments []*models.Segment, flagID int64) error {
	for _, segment := range segments {
		if err := UpdateSegment(segment); err != nil {
			return err
		}

		if err := UpdateContraints(segment.Constraints); err != nil {
			return err
		}

		if err := UpdateDistributions(segment.Distributions, segment.ID, flagID); err != nil {
			return err
		}
	}
	return nil
}

func UpdateSegment(segment *models.Segment) error {
	s := &entity.Segment{}
	err := entity.
		PreloadConstraintsDistribution(getDB()).
		First(s, segment.ID).
		Error
	if err != nil {
		return err
	}

	s.RolloutPercent = util.SafeUint(segment.RolloutPercent)
	s.Description = util.SafeString(segment.Description)
	s.Rank = uint(*segment.Rank)

	if err := getDB().Save(s).Error; err != nil {
		return err
	}

	return nil
}

func UpdateContraints(constraints []*models.Constraint) error {
	for _, constraint := range constraints {
		if err := UpdateConstraint(constraint); err != nil {
			return err
		}
	}
	return nil
}

func UpdateConstraint(constraint *models.Constraint) error {
	cons := &entity.Constraint{}

	if err := getDB().First(cons, constraint.ID).Error; err != nil {
		return err
	}

	cons.Property = util.SafeString(constraint.Property)
	cons.Operator = util.SafeString(constraint.Operator)
	cons.Value = util.SafeString(constraint.Value)

	if err := cons.Validate(); err != nil {
		return err
	}

	if err := getDB().Save(&cons).Error; err != nil {
		return err
	}
	return nil
}

func UpdateDistributions(distributions []*models.Distribution, segmentID int64, flagID int64) error {
	putDistributionsRequest := &models.PutDistributionsRequest{Distributions: distributions}
	distributionParams := distribution.PutDistributionsParams{Body: putDistributionsRequest, FlagID: int64(flagID), SegmentID: int64(segmentID)}
	if err := validatePutDistributions(distributionParams); err != nil {
		return err
	}

	tx := getDB().Begin()
	err := tx.Where("segment_id = ?", segmentID).Delete(&entity.Distribution{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	ds := r2eMapDistributions(distributions, uint(segmentID))
	for _, d := range ds {
		err1 := tx.Create(&d).Error
		if err1 != nil {
			tx.Rollback()
			return err1
		}
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func UpdateVariants(variants []*models.Variant, flagID int64) error {
	for _, variant := range variants {
		if err := UpdateVariant(variant, flagID); err != nil {
			return err
		}
	}
	return nil
}

func UpdateVariant(variant *models.Variant, flagID int64) error {
	v := &entity.Variant{}

	if err := getDB().First(v, variant.ID).Error; err != nil {
		return err
	}

	v.Key = util.SafeString(variant.Key)

	if err := v.Validate(); err != nil {
		return err
	}

	if err := getDB().Save(&v).Error; err != nil {
		return err
	}

	if err := validatePutVariantForDistributions(v); err != nil {
		return err
	}

	return nil
}
