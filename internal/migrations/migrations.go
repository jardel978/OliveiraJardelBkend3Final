package migrations

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.Migrator().DropTable(
		&domain.Clinica{},
		&domain.Paciente{},
		&domain.Dentista{},
		&domain.Endereco{},
		&domain.Prontuario{},
		&domain.Consulta{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&domain.Clinica{},
		&domain.Paciente{},
		&domain.Dentista{},
		&domain.Endereco{},
		&domain.Prontuario{},
		&domain.Consulta{})

	//if err != nil {
	//	return err
	//}
	//
	//err = db.Migrator().CreateConstraint(&domain.Bairro{}, "fk_bairro_endereco")
	//if err != nil {
	//	return err
	//}
	//
	//err = db.Migrator().CreateConstraint(&domain.Cidade{}, "fk_ciadade_endereco")
	//if err != nil {
	//	return err
	//}
	//
	//err = db.Migrator().CreateConstraint(&domain.Clinica{}, "fk_clinica_endereco")
	//if err != nil {
	//	return err
	//}
	//
	//err = db.Migrator().CreateConstraint(&domain.Clinica{}, "fk_clinica_dentista")
	//if err != nil {
	//	return err
	//}
	//
	//err = db.Migrator().CreateConstraint(&domain.Dentista{}, "fk_dentista_consulta")
	//if err != nil {
	//	return err
	//}
	//
	//err = db.Migrator().CreateConstraint(&domain.Paciente{}, "fk_paciente_consulta")
	//if err != nil {
	//	return err
	//}

	return err
}
