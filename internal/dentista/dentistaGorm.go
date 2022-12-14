package dentista

import (
	"OliveiraJardelBkend3Final/internal/configs"
	"OliveiraJardelBkend3Final/internal/domain"
	"context"
	"gorm.io/gorm"
	"log"
)

type dentistaGorm struct {
	db *gorm.DB
}

func NewDentistaGorm() IDentistaRepository {
	return &dentistaGorm{
		db: configs.GetGormDB(),
	}
}

func (pg *dentistaGorm) Save(dentista domain.Dentista, ctx context.Context) (d domain.Dentista, err error) {
	err = pg.db.WithContext(ctx).Save(&dentista).Error
	pg.db.WithContext(ctx).Unscoped().Delete(&domain.Endereco{})
	if err != nil {
		return d, err
	}
	return dentista, nil
}

func (pg *dentistaGorm) FindAll(ctx context.Context) (list []domain.Dentista, err error) {
	err = pg.db.WithContext(ctx).Model(&domain.Dentista{}).Preload("Consultas").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (pg *dentistaGorm) FindByMatricula(matricula string, ctx context.Context) (d domain.Dentista, err error) {
	err = pg.db.WithContext(ctx).Model(&d).Preload("Consultas").First(&d, "matricula = ?", matricula).Error
	if err != nil {
		return d, err
	}
	return d, nil
}

func (pg *dentistaGorm) FindById(id uint, ctx context.Context) (d domain.Dentista, err error) {
	err = pg.db.WithContext(ctx).Model(&d).Preload("Consultas").First(&d, "id = ?", id).Error
	if err != nil {
		return d, err
	}
	return d, nil
}

func (pg *dentistaGorm) Update(dentista domain.Dentista, ctx context.Context) (d domain.Dentista, err error) {
	log.Printf("\ndentista update: %v\n", dentista.Model.ID)
	err = pg.db.WithContext(ctx).Table("dentistas").Where("id = ?", dentista.Model.ID).Updates(dentista).Error
	pg.db.WithContext(ctx).Unscoped().Delete(&domain.Endereco{})
	if err != nil {
		return d, err
	}
	return dentista, nil
}

func (pg *dentistaGorm) Delete(id uint, ctx context.Context) (err error) {
	err = pg.db.WithContext(ctx).Delete(&domain.Dentista{}, "id = ?", id).Error
	//pg.db.WithContext(ctx).Unscoped().Delete(&domain.Endereco{})
	if err != nil {
		return err
	}
	return nil
}
