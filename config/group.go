package config

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

var (
	isInitAddUserGroups       bool
	addUserGroups             bool
	isInitAddUserGroupsExcept bool
	addUserGroupsExcept       []string
	isInitAddGroups           bool
	addGroups                 []user.Group
	isInitAddGroupsFinal      bool
	addGroupsFinal            []user.Group
)

func AddUserGroups() bool {
	if !isInitAddUserGroups {
		addUserGroups = boolFromEnv(EnvPrefix()+"_ADD_USER_GROUPS", false)
		isInitAddUserGroups = true
	}

	return addUserGroups
}

func AddUserGroupsExcept() []string {
	if !isInitAddUserGroupsExcept {
		addUserGroupsExcept = listFromEnv(EnvPrefix()+"_ADD_USER_GROUPS_EXCEPT", ",")
		for groupName, add := range listFromEnvs(EnvPrefix() + "_ADD_USER_GROUPS_EXCEPT_") {
			groupName = strings.ReplaceAll(groupName, "__", "-")
			if !boolFromStr(add, false) {
				if add != "" {
					addUserGroupsExcept = removeStringFromSlice(groupName, addUserGroupsExcept)
				}
				continue
			}
			addUserGroupsExcept = append(addUserGroupsExcept, groupName)
		}
		isInitAddUserGroupsExcept = true
	}

	return addUserGroupsExcept
}

func AddGroups() []user.Group {
	if !isInitAddGroups {
		addGroupNames := listFromEnv(EnvPrefix()+"_ADD_GROUPS", ",")
		for groupName, add := range listFromEnvs(EnvPrefix() + "_ADD_GROUP_") {
			groupName = strings.ReplaceAll(groupName, "__", "-")
			if !boolFromStr(add, false) {
				if add != "" {
					addGroupNames = removeStringFromSlice(groupName, addGroupNames)
				}
				continue
			}
			addGroupNames = append(addGroupNames, groupName)
		}
		for _, groupName := range addGroupNames {
			group, err := user.LookupGroup(groupName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: error: %s\n", BinName(), err.Error())
				os.Exit(1)
			}
			if groupInSlice(group.Name, addGroups) {
				continue
			}
			addGroups = append(addGroups, *group)
		}
		isInitAddGroups = true
	}

	return addGroups
}

func AddGroupsFinal() []user.Group {
	if !isInitAddGroupsFinal {
		if AddUserGroups() {
			for _, group := range Groups() {
				if stringInSlice(group.Name, AddUserGroupsExcept()) {
					continue
				}
				addGroupsFinal = append(addGroupsFinal, group)
			}
		}

		for _, group := range AddGroups() {
			if groupInSlice(group.Name, addGroupsFinal) {
				continue
			}
			addGroupsFinal = append(addGroupsFinal, group)
		}

		isInitAddGroupsFinal = true
	}

	return addGroupsFinal
}

func groupInSlice(name string, slice []user.Group) bool {
	for _, group := range slice {
		if group.Name == name {
			return true
		}
	}
	return false
}
