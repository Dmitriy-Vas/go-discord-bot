package utils

import (
	mod "../models"
	"errors"
	dgo "github.com/bwmarrin/discordgo"
)

func GuildFromMessage(m *mod.Message) (guild *dgo.Guild, err error) {
	gID := m.MessageCreate.GuildID
	guild, err = m.State.Guild(gID)
	if err != nil {
		guild, err = m.State.Guild(gID)
	}
	return
}

func MemberGuild(m *mod.Message, user *dgo.User) (member *dgo.Member, err error) {
	member, err = m.State.Member(m.MessageCreate.GuildID, user.ID)
	if err != nil {
		member, err = m.GuildMember(m.GuildID, user.ID)
	}
	return
}

func RolesFromGuild(m *mod.Message) (roles []*dgo.Role, err error) {
	guild, err := GuildFromMessage(m)
	if err != nil {
		return
	}
	roles = guild.Roles
	return
}

func RoleFromID(m *mod.Message, id string) (*dgo.Role, error) {
	return m.State.Role(m.MessageCreate.GuildID, id)
}

func RoleFromName(m *mod.Message, name string) (*dgo.Role, error) {
	roles, err := RolesFromGuild(m)
	if err != nil {
		return nil, err
	}
	for _, role := range roles {
		if role.Name == name {
			return role, nil
		}
	}
	return nil, errors.New("role with name " + name + " not found")
}

func RolesFromUser(m *mod.Message, user *dgo.User) ([]*dgo.Role, error) {
	member, err := MemberGuild(m, user)
	if err != nil {
		return nil, err
	}
	var roles []*dgo.Role
	for _, id := range member.Roles {
		role, err := RoleFromID(m, id)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func HasPermissions(m *mod.Message, user *dgo.User, permissions int) (bool, error) {
	roles, err := RolesFromUser(m, user)
	if err != nil {
		return false, err
	}
	for _, role := range roles {
		if role.Permissions&permissions == 0 {
			return true, nil
		}
	}
	return false, nil
}

func HasRole(m *mod.Message, user *dgo.User, role *dgo.Role) (bool, error) {
	member, err := MemberGuild(m, user)
	if err != nil {
		return false, err
	}
	for _, r := range member.Roles {
		if r == role.ID {
			return true, nil
		}
	}
	return false, nil
}
