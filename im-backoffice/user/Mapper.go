package user

import "github.com/Nerzal/gocloak/v12"

func userToUserResp(user UserDTO, role *gocloak.Role, userId string) *UserResp {
	return &UserResp{
		Id:        userId,
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Enabled:   user.Enabled,
		Role: RolesResp{
			Id:       *role.ID,
			RoleType: *role.Name,
		},
	}
}

//	func usersToUserRespArr(users []*gocloak.User) []UserResp {
//		var userResponses []UserResp
//		for _, user := range users {
//			userResponses = append(userResponses, userToUserResp(&user))
//		}
//
//		return userResponses
//	}

func userToUserMeResp(userInfo *gocloak.UserInfo, userRole *gocloak.Role) UserMeResp {
	return UserMeResp{
		Id:        *userInfo.Sub,
		Firstname: *userInfo.GivenName,
		Lastname:  *userInfo.FamilyName,
		Role: RolesResp{
			Id:       *userRole.ID,
			RoleType: *userRole.Name,
		},
		Username: *userInfo.PreferredUsername,
	}
}

func roleToRolesResp(role *gocloak.Role) RolesResp {
	return RolesResp{
		Id:       *role.ID,
		RoleType: *role.Name,
	}
}

func rolesToRolesRespArr(roles []*gocloak.Role) []RolesResp {
	var rolesResponse []RolesResp
	for _, role := range roles {
		rolesResponse = append(rolesResponse, roleToRolesResp(role))
	}

	return rolesResponse
}
