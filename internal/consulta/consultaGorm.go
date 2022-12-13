package consulta

import (
	"OliveiraJardelBkend3Final/internal/configs"
	"OliveiraJardelBkend3Final/internal/domain"
	"context"
	"gorm.io/gorm"
)

type consultaGorm struct {
	db *gorm.DB
}

func NewConsultaGorm() IConsultaRepository {
	return &consultaGorm{
		db: configs.GetGormDB(),
	}
}

func (pg *consultaGorm) Save(consulta domain.Consulta, ctx context.Context) (p domain.Consulta, err error) {
	err = pg.db.WithContext(ctx).Save(&consulta).Error
	pg.db.WithContext(ctx).Unscoped().Delete(&domain.Endereco{})
	if err != nil {
		return p, err
	}
	return consulta, nil
}

func (pg *consultaGorm) FindAll(ctx context.Context) (list []domain.Consulta, err error) {
	err = pg.db.WithContext(ctx).Model(&domain.Consulta{}).Preload("Endereco").Preload("Prontuario").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (pg *consultaGorm) FindById(id uint, ctx context.Context) (p domain.Consulta, err error) {
	err = pg.db.WithContext(ctx).Model(&p).Preload("Endereco").Preload("Prontuario").First(&p, "id = ?", id).Error
	if err != nil {
		return p, err
	}
	return p, nil
}

func (pg *consultaGorm) Update(consulta domain.Consulta, ctx context.Context) (p domain.Consulta, err error) {
	err = pg.db.WithContext(ctx).Model(&consulta).Updates(consulta).Error
	pg.db.WithContext(ctx).Unscoped().Delete(&domain.Endereco{})
	if err != nil {
		return p, err
	}
	return consulta, nil
}

func (pg *consultaGorm) Delete(id uint, ctx context.Context) (err error) {
	err = pg.db.WithContext(ctx).Delete(&domain.Consulta{}, "id = ?", id).Error
	//pg.db.WithContext(ctx).Unscoped().Delete(&domain.Endereco{})
	if err != nil {
		return err
	}
	return nil
}
