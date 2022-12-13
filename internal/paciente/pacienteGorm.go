package paciente

import (
	"OliveiraJardelBkend3Final/internal/configs"
	"OliveiraJardelBkend3Final/internal/domain"
	"context"
	"gorm.io/gorm"
)

type pacienteGorm struct {
	db *gorm.DB
}

func NewPacienteGorm() IPacienteRepository {
	return &pacienteGorm{
		db: configs.GetGormDB(),
	}
}

func (pg *pacienteGorm) Save(paciente domain.Paciente, ctx context.Context) (p domain.Paciente, err error) {
	err = pg.db.WithContext(ctx).Save(&paciente).Error
	pg.db.WithContext(ctx).Unscoped().Delete(&domain.Endereco{})
	if err != nil {
		return p, err
	}
	return paciente, nil
}

func (pg *pacienteGorm) FindAll(ctx context.Context) (list []domain.Paciente, err error) {
	err = pg.db.WithContext(ctx).Model(&domain.Paciente{}).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (pg *pacienteGorm) FindByRG(rg string, ctx context.Context) (p domain.Paciente, err error) {
	err = pg.db.WithContext(ctx).Model(&p).Where("RG", rg).First(&p).Error
	if err != nil {
		return p, err
	}
	return p, nil
}

func (pg *pacienteGorm) FindById(id uint, ctx context.Context) (p domain.Paciente, err error) {
	err = pg.db.WithContext(ctx).Model(&p).First(&p, "id = ?", id).Error
	if err != nil {
		return p, err
	}
	return p, nil
}

func (pg *pacienteGorm) Update(paciente domain.Paciente, ctx context.Context) (p domain.Paciente, err error) {
	err = pg.db.WithContext(ctx).Model(&paciente).Updates(paciente).Error
	pg.db.WithContext(ctx).Unscoped().Delete(&domain.Endereco{})
	if err != nil {
		return p, err
	}
	return paciente, nil
}

func (pg *pacienteGorm) Delete(id uint, ctx context.Context) (err error) {
	err = pg.db.WithContext(ctx).Delete(&domain.Paciente{}, "id = ?", id).Error
	//pg.db.WithContext(ctx).Unscoped().Delete(&domain.Endereco{})
	if err != nil {
		return err
	}
	return nil
}
