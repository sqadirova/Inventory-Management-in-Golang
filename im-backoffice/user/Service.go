package user

import (
	"context"
	"github.com/Nerzal/gocloak/v12"
	"github.com/gofiber/fiber/v2"
	"im-backoffice/CONSTANTS"
	"im-backoffice/config"
	myErrors "im-backoffice/errors"
	"im-backoffice/keycloak"
	"im-backoffice/util"
	"log"
)

type IUserService interface {
	getAllUsers(ctx context.Context, token string) ([]*UserResp, error)
	GetUserInfo(ctx context.Context, token string) (*gocloak.UserInfo, *gocloak.Role, error)
	getRolesFromDB(ctx context.Context, token string) ([]*gocloak.Role, error)
	getRoleByIDFromDB(ctx context.Context, roleId string, token string) (*gocloak.Role, error)
	createUser(ctx context.Context, userReq UserDTO, token string) (*UserResp, error)
	updateUser(ctx context.Context, userId string, userReq UpdateUserReq, token string) (*UserResp, error)
	updateUserRole(ctx context.Context, token string, roleId string, userId string) (*gocloak.Role, error)
	AddToUserRealmManagementRoles(ctx context.Context, adminRoles []string, realmManagement string, token string, userId string) error
	AddClientRoleToUser(ctx context.Context, token string, idOfClient string, userId string, rolesForUser []gocloak.Role) error
	UpdateUserRealmManagementRoles(ctx context.Context, newAdminRoles []string, realmManagement string, token string, userId string) error
}

type UserServiceImpl struct {
	keycloak *keycloak.Keycloak
}

func GetNewUserService(keycloak *keycloak.Keycloak) *UserServiceImpl {
	return &UserServiceImpl{
		keycloak: keycloak,
	}
}

func (u *UserServiceImpl) getRolesByUserId(ctx context.Context, userId string, token string) ([]*gocloak.Role, error) {
	clients, err := u.keycloak.Gocloak.GetClients(ctx, token, u.keycloak.Realm, gocloak.GetClientsParams{ClientID: &u.keycloak.ClientId})

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("not_found_client", CONSTANTS.LANGUAGE))
	}

	roles, err := u.keycloak.Gocloak.GetClientRolesByUserID(ctx, token, u.keycloak.Realm, *clients[0].ID, userId)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	return roles, nil
}

func (u *UserServiceImpl) GetUserInfo(ctx context.Context, token string) (*gocloak.UserInfo, *gocloak.Role, error) {
	userInfo, err := u.keycloak.Gocloak.GetUserInfo(ctx, token, u.keycloak.Realm)

	if err != nil {
		log.Println(err)
		return nil, nil, myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	roles, err := u.getRolesByUserId(ctx, *userInfo.Sub, token)

	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	var userRole *gocloak.Role

	for _, role := range roles {
		if !*role.Composite {
			userRole = role
		}
	}

	return userInfo, userRole, nil
}

func (u *UserServiceImpl) getAllUsers(ctx context.Context, token string) ([]*UserResp, error) {
	users, err := u.keycloak.Gocloak.GetUsers(ctx, token, u.keycloak.Realm, gocloak.GetUsersParams{})

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
	}

	var usersResponse []*UserResp

	for _, user := range users {
		roles, err := u.getRolesByUserId(ctx, *user.ID, token)

		if err != nil {
			log.Println(err)
			return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("cannot_find_data", CONSTANTS.LANGUAGE))
		}

		for _, role := range roles {
			userInfo := &UserResp{
				Id:        gocloak.PString(user.ID),
				Username:  gocloak.PString(user.Username),
				Firstname: gocloak.PString(user.FirstName),
				Lastname:  gocloak.PString(user.LastName),
				Enabled:   gocloak.PBool(user.Enabled),
				Role: RolesResp{
					Id:       gocloak.PString(role.ID),
					RoleType: gocloak.PString(role.Name),
				},
			}

			usersResponse = append(usersResponse, userInfo)
		}
	}

	return usersResponse, nil
}

func (u *UserServiceImpl) getRolesFromDB(ctx context.Context, token string) ([]*gocloak.Role, error) {
	clients, err := u.keycloak.Gocloak.GetClients(ctx, token, u.keycloak.Realm, gocloak.GetClientsParams{ClientID: &u.keycloak.ClientId})

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("not_found_client", CONSTANTS.LANGUAGE))
	}

	roles, err := u.keycloak.Gocloak.GetClientRoles(ctx, token, u.keycloak.Realm, *clients[0].ID, gocloak.GetRoleParams{})

	if len(roles) == 0 {
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("no_role", CONSTANTS.LANGUAGE))
	}

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	return roles, nil
}

func (u *UserServiceImpl) getRoleByIDFromDB(ctx context.Context, roleId string, token string) (*gocloak.Role, error) {
	role, err := u.keycloak.Gocloak.GetClientRoleByID(ctx, token, u.keycloak.Realm, roleId)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("no_role", CONSTANTS.LANGUAGE))
	}

	return role, nil
}

func (u *UserServiceImpl) createUser(ctx context.Context, userReq UserDTO, token string) (*UserResp, error) {
	clients, err := u.keycloak.Gocloak.GetClients(ctx, token, u.keycloak.Realm, gocloak.GetClientsParams{ClientID: &u.keycloak.ClientId})

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("not_found_client", CONSTANTS.LANGUAGE))
	}

	if !util.IsValidUsername(userReq.Username) {
		return nil, myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("invalid_username", CONSTANTS.LANGUAGE))
	}

	if !util.IsValidPassword(userReq.Password) {
		return nil, myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("invalid_password", CONSTANTS.LANGUAGE))
	}

	role, err := u.keycloak.Gocloak.GetClientRoleByID(ctx, token, u.keycloak.Realm, userReq.RoleId)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("no_role", CONSTANTS.LANGUAGE))
	}

	cloakUser := gocloak.User{
		FirstName: gocloak.StringP(userReq.Firstname),
		LastName:  gocloak.StringP(userReq.Lastname),
		Enabled:   gocloak.BoolP(userReq.Enabled),
		Username:  gocloak.StringP(userReq.Username),
	}

	createdUserId, err := u.keycloak.Gocloak.CreateUser(ctx, token, u.keycloak.Realm, cloakUser)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	var rolesForUser []gocloak.Role
	rolesForUser = append(rolesForUser, *role)
	err = u.AddClientRoleToUser(ctx, token, *clients[0].ID, createdUserId, rolesForUser)

	if err != nil {
		return nil, myErrors.NewHttpError(err.(myErrors.HttpError).Code, myErrors.NewResponseByKey(err.Error(), CONSTANTS.LANGUAGE))
	}

	realmManagement := "realm-management"
	if *role.Name == CONSTANTS.ADMIN {
		realmAdmin := []string{"realm-admin"}

		err = u.AddToUserRealmManagementRoles(ctx, realmAdmin, realmManagement, token, createdUserId)

		if err != nil {
			return nil, myErrors.NewHttpError(err.(myErrors.HttpError).Code, myErrors.NewResponseByKey(err.Error(), CONSTANTS.LANGUAGE))
		}
	} else {
		realmManagmentRoles := []string{"view-clients", "view-users"}

		err = u.AddToUserRealmManagementRoles(ctx, realmManagmentRoles, realmManagement, token, createdUserId)

		if err != nil {
			return nil, myErrors.NewHttpError(err.(myErrors.HttpError).Code, myErrors.NewResponseByKey(err.Error(), CONSTANTS.LANGUAGE))
		}
	}

	err = u.keycloak.Gocloak.SetPassword(ctx, token, createdUserId, u.keycloak.Realm, userReq.Password, false)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	return userToUserResp(userReq, role, createdUserId), nil
}

func (u *UserServiceImpl) updateUser(ctx context.Context, userId string, userReq UpdateUserReq, token string) (*UserResp, error) {
	user := &gocloak.User{
		ID:        &userId,
		FirstName: gocloak.StringP(userReq.Firstname),
		LastName:  gocloak.StringP(userReq.Lastname),
		Enabled:   gocloak.BoolP(userReq.Enabled),
	}

	err := u.keycloak.Gocloak.UpdateUser(ctx, token, u.keycloak.Realm, *user)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusNotModified, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	role, err := u.updateUserRole(ctx, token, userReq.RoleId, userId)

	if err != nil {
		return nil, myErrors.NewHttpError(err.(myErrors.HttpError).Code, myErrors.NewResponseByKey(err.Error(), CONSTANTS.LANGUAGE))
	}

	updatedUser, err := u.keycloak.Gocloak.GetUserByID(ctx, token, u.keycloak.Realm, userId)

	if err != nil {
		return nil, myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	userResp := &UserResp{
		Id:        userId,
		Username:  gocloak.PString(updatedUser.Username),
		Firstname: gocloak.PString(updatedUser.FirstName),
		Lastname:  gocloak.PString(updatedUser.LastName),
		Enabled:   gocloak.PBool(updatedUser.Enabled),
		Role: RolesResp{
			Id:       gocloak.PString(role.ID),
			RoleType: gocloak.PString(role.Name),
		},
	}

	return userResp, nil
}

func (u *UserServiceImpl) updateUserRole(ctx context.Context, token string, roleId string, userId string) (*gocloak.Role, error) {
	clients, err := u.keycloak.Gocloak.GetClients(ctx, token, u.keycloak.Realm, gocloak.GetClientsParams{ClientID: &u.keycloak.ClientId})

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("not_found_client", CONSTANTS.LANGUAGE))
	}

	oldRoles, err := u.keycloak.Gocloak.GetClientRolesByUserID(ctx, token, u.keycloak.Realm, *clients[0].ID, userId)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("no_role", CONSTANTS.LANGUAGE))
	}

	if len(oldRoles) != 1 {
		if config.ProfileConfiguration.Profile.Active == "dev" {
			log.Println("Please check user roles. Can not assigned 2 or more roles to user.")
		}
		return nil, myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("no_role", CONSTANTS.LANGUAGE))
	}

	newRole, err := u.keycloak.Gocloak.GetClientRoleByID(ctx, token, u.keycloak.Realm, roleId)

	if err != nil {
		log.Println(err)
		return nil, myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("no_role", CONSTANTS.LANGUAGE))
	}

	if oldRoles[0].ID != newRole.ID {
		if *newRole.Name == CONSTANTS.ADMIN {
			newRealmAdmin := []string{"realm-admin"}

			err := u.UpdateUserRealmManagementRoles(ctx, newRealmAdmin, "realm-management", token, userId)

			if err != nil {
				return nil, myErrors.NewHttpError(err.(myErrors.HttpError).Code, myErrors.NewResponseByKey(err.Error(), CONSTANTS.LANGUAGE))
			}
		} else {
			NewRealmManagRoles := []string{"view-clients", "view-users"}

			err := u.UpdateUserRealmManagementRoles(ctx, NewRealmManagRoles, "realm-management", token, userId)

			if err != nil {
				return nil, myErrors.NewHttpError(err.(myErrors.HttpError).Code, myErrors.NewResponseByKey(err.Error(), CONSTANTS.LANGUAGE))
			}
		}

		err := u.keycloak.Gocloak.DeleteClientRolesFromUser(ctx, token, u.keycloak.Realm,
			*clients[0].ID, userId, []gocloak.Role{*oldRoles[0]})

		if err != nil {
			log.Println(err)
			return nil, myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
		}

		err = u.keycloak.Gocloak.AddClientRolesToUser(ctx, token, u.keycloak.Realm,
			*clients[0].ID, userId, []gocloak.Role{*newRole})

		if err != nil {
			log.Println(err)
			return nil, myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
		}

		return newRole, nil
	}

	return oldRoles[0], nil
}

func (u *UserServiceImpl) UpdateUserRealmManagementRoles(ctx context.Context, newAdminRoles []string, realmManagement string, token string, userId string) error {
	realmManagClient, err := u.keycloak.Gocloak.GetClients(ctx, token, u.keycloak.Realm, gocloak.GetClientsParams{ClientID: &realmManagement})

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("not_found_client", CONSTANTS.LANGUAGE))
	}

	oldAdminRoles, err := u.keycloak.Gocloak.GetClientRolesByUserID(ctx, token, u.keycloak.Realm, *realmManagClient[0].ID, userId)

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusBadRequest, myErrors.NewResponseByKey("no_role", CONSTANTS.LANGUAGE))
	}

	var forDeleteRealmAdminRoles []gocloak.Role
	for _, oldAdminRole := range oldAdminRoles {
		forDeleteRealmAdminRoles = append(forDeleteRealmAdminRoles, *oldAdminRole)
	}

	err = u.keycloak.Gocloak.DeleteClientRolesFromUser(ctx, token, u.keycloak.Realm,
		*realmManagClient[0].ID, userId, forDeleteRealmAdminRoles)

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	var newRealmAdminRoles []gocloak.Role
	for _, newAdminRole := range newAdminRoles {
		realmAdmin, err := u.keycloak.Gocloak.GetClientRole(ctx, token, u.keycloak.Realm, *realmManagClient[0].ID, newAdminRole)

		if err != nil {
			log.Println(err)
			return myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("not_found_realm_roles", CONSTANTS.LANGUAGE))
		}

		newRealmAdminRoles = append(newRealmAdminRoles, *realmAdmin)
	}

	err = u.keycloak.Gocloak.AddClientRolesToUser(ctx, token, u.keycloak.Realm,
		*realmManagClient[0].ID, userId, newRealmAdminRoles)

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	return nil
}

func (u *UserServiceImpl) AddToUserRealmManagementRoles(ctx context.Context, adminRoles []string, realmManagement string, token string, userId string) error {
	realmManagClient, err := u.keycloak.Gocloak.GetClients(ctx, token, u.keycloak.Realm, gocloak.GetClientsParams{ClientID: &realmManagement})

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("not_found_client", CONSTANTS.LANGUAGE))
	}

	var realmAdminRoles []gocloak.Role
	for _, adminRole := range adminRoles {
		realmAdmin, err := u.keycloak.Gocloak.GetClientRole(ctx, token, u.keycloak.Realm, *realmManagClient[0].ID, adminRole)

		if err != nil {
			log.Println(err)
			return myErrors.NewHttpError(fiber.StatusNotFound, myErrors.NewResponseByKey("not_found_realm_roles", CONSTANTS.LANGUAGE))
		}

		realmAdminRoles = append(realmAdminRoles, *realmAdmin)
	}

	err = u.keycloak.Gocloak.AddClientRolesToUser(ctx, token, u.keycloak.Realm,
		*realmManagClient[0].ID, userId, realmAdminRoles)

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	return nil
}

func (u *UserServiceImpl) AddClientRoleToUser(ctx context.Context, token string, idOfClient string, userId string, rolesForUser []gocloak.Role) error {
	err := u.keycloak.Gocloak.AddClientRolesToUser(ctx, token, u.keycloak.Realm, idOfClient, userId, rolesForUser)

	if err != nil {
		log.Println(err)
		return myErrors.NewHttpError(fiber.StatusInternalServerError, myErrors.NewResponseByKey("unexpected_error", CONSTANTS.LANGUAGE))
	}

	return nil
}
