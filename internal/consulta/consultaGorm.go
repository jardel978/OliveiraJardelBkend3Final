package consulta

import (
	"OliveiraJardelBkend3Final/internal/configs"
	"OliveiraJardelBkend3Final/internal/domain"
	"context"
	"gorm.io/gorm"
	"log"
)

type consultaGorm struct {
	db *gorm.DB
}

func NewConsultaGorm() IConsultaRepository {
	return &consultaGorm{
		db: configs.GetGormDB(),
	}
}

func (cg *consultaGorm) Save(consulta domain.Consulta, ctx context.Context) (c domain.Consulta, err error) {
	err = cg.db.WithContext(ctx).Save(&consulta).Error
	cg.db.WithContext(ctx).Unscoped().Delete(&domain.Endereco{})
	if err != nil {
		return c, err
	}
	return consulta, nil
}

func (cg *consultaGorm) FindAll(ctx context.Context) (list []domain.Consulta, err error) {
	err = cg.db.WithContext(ctx).Model(&domain.Consulta{}).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (cg *consultaGorm) FindById(id uint, ctx context.Context) (c domain.Consulta, err error) {
	err = cg.db.WithContext(ctx).Model(&c).First(&c, "id = ?", id).Error
	if err != nil {
		return c, err
	}
	return c, nil
}

func (cg *consultaGorm) FindAllByPacienteID(pacienteID uint, ctx context.Context) (list []domain.Consulta, err error) {
	err = cg.db.WithContext(ctx).Model(&domain.Consulta{}).Where("PacienteID", "paciente_id = ?", pacienteID).Find(&list).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func (cg *consultaGorm) Update(consulta domain.Consulta, ctx context.Context) (c domain.Consulta, err error) {
	log.Printf("\ngorm - consulta: %v\n", consulta)

	err = cg.db.WithContext(ctx).Table("consultas").Where("id = ?", consulta.Model.ID).Updates(consulta).Error
	cg.db.WithContext(ctx).Unscoped().Delete(&domain.Endereco{})
	if err != nil {
		return c, err
	}
	return consulta, nil
}

func (cg *consultaGorm) Delete(id uint, ctx context.Context) (err error) {
	err = cg.db.WithContext(ctx).Delete(&domain.Consulta{}, "id = ?", id).Error
	//cg.db.WithContext(ctx).Unscoped().Delete(&domain.Endereco{})
	if err != nil {
		return err
	}
	return nil
}
