package api

import "CRUDUSERS/internal/database/store"

func validateCreateUser(user store.CreateUserParams) []map[string]string {
	errs := []map[string]string{}

	if user.FirstName == "" {
		errs = append(errs, map[string]string{"Nome": "É obrigatório informar o nome"})
	}

	if user.LastName == "" {
		errs = append(errs, map[string]string{"Sobrenome": "É obrigatório informar o sobrenome"})
	}

	return errs
}

func validateUpdateUser(user store.UpdateUserParams) []map[string]string {
	errs := []map[string]string{}

	if user.FirstName == "" {
		errs = append(errs, map[string]string{"Nome": "É obrigatório informar o nome"})
	}

	if user.LastName == "" {
		errs = append(errs, map[string]string{"Sobrenome": "É obrigatório informar o sobrenome"})
	}

	return errs
}
