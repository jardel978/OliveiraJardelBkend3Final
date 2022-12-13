package clinica

import (
	"OliveiraJardelBkend3Final/internal/configs"
	"OliveiraJardelBkend3Final/internal/domain"
	"context"
	"gorm.io/gorm"
)

type clinicaGorm struct {
	db *gorm.DB
}

func NewClinicaGorm() IClinicaRepository {
	return &clinicaGorm{
		db: configs.GetGormDB(),
	}
}

func (cg *clinicaGorm) Save(clinica domain.Clinica, ctx context.Context) (c domain.Clinica, err error) {
	err = cg.db.WithContext(ctx).Save(&clinica).Error
	cg.db.WithContext(ctx).Unscoped().Delete(&domain.Endereco{})
	if err != nil {
		return c, err
	}
	return clinica, nil
}

func (cg *clinicaGorm) FindAll(ctx context.Context) (list []domain.Clinica, err error) {
	err = cg.db.WithContext(ctx).Model(&domain.Clinica{}).Preload("Endereco").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (cg *clinicaGorm) FindById(id uint, ctx context.Context) (c domain.Clinica, err error) {
	err = cg.db.WithContext(ctx).Model(&c).Preload("Endereco").First(&c, "id = ?", id).Error
	if err != nil {
		return c, err
	}
	return c, nil
}

func (cg *clinicaGorm) Update(clinica domain.Clinica, ctx context.Context) (c domain.Clinica, err error) {
	err = cg.db.WithContext(ctx).Model(&clinica).Updates(clinica).Error
	cg.db.WithContext(ctx).Unscoped().Delete(&domain.Endereco{})
	if err != nil {
		return c, err
	}
	return clinica, nil
}

func (cg *clinicaGorm) Delete(id uint, ctx context.Context) (err error) {
	err = cg.db.WithContext(ctx).Delete(&domain.Clinica{}, "id = ?", id).Error
	//cg.db.WithContext(ctx).Unscoped().Delete(&domain.Endereco{})
	if err != nil {
		return err
	}
	return nil
}
