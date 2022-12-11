package migrations

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) (err error) {
	err = db.Migrator().DropTable(
		&domain.Bairro{},
		&domain.Cidade{},
		&domain.Clinica{},
		&domain.Paciente{},
		&domain.Dentista{},
		&domain.Endereco{},
		&domain.Prontuario{},
		&domain.Consulta{})

	err = db.AutoMigrate(
		&domain.Bairro{},
		&domain.Cidade{},
		&domain.Clinica{},
		&domain.Paciente{},
		&domain.Dentista{},
		&domain.Endereco{},
		&domain.Prontuario{},
		&domain.Consulta{})

	//db.Migrator().CreateConstraint(&domain.Bairro{}, "fk_bairro_endereco")
	//db.Migrator().CreateConstraint(&domain.Cidade{}, "fk_ciadade_endereco")
	//db.Migrator().CreateConstraint(&domain.Clinica{}, "fk_clinica_endereco")
	//db.Migrator().CreateConstraint(&domain.Clinica{}, "fk_clinica_dentista")
	//db.Migrator().CreateConstraint(&domain.Dentista{}, "fk_dentista_consulta")
	//db.Migrator().CreateConstraint(&domain.Paciente{}, "fk_paciente_consulta")

	return err
}
