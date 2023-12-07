package omc

import (
	"net/http"

	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/store"

	"github.com/heypkg/store/echohandler"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// HandleListGroups godoc
// @Summary List child groups
// @Tags OMC Groups
// @ID list-groups
// @Security Bearer
// @Success 200 {array} omc.Group
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/groups [get]
func HandleListGroups(c echo.Context) error {
	schema := cast.ToString(c.Get("schema"))
	db := store.GetDB()
	group := new(omc.Group)
	group.Schema = schema
	group.ID = 0
	if err := db.Preload("Parent").Preload("Children").First(group).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "group not found")
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}
	group.Children = omc.GetChildrenGroups(db, group.Schema, group.ID)
	return c.JSON(http.StatusOK, []omc.Group{*group})
}

// HandleGetGroup godoc
// @Summary Get group details
// @Tags OMC Groups
// @ID get-group
// @Security Bearer
// @Param id path int true "Group ID"
// @Success 200 {object} omc.Group
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/groups/{id} [get]
func HandleGetGroup(c echo.Context) error {
	group := echohandler.GetObjectFromEchoContext[omc.Group](c)

	return c.JSON(http.StatusOK, group)
}

// HandleGetGroupChild godoc
// @Summary Get child groups
// @Tags OMC Groups
// @ID get-group-children
// @Security Bearer
// @Param id path int true "Group ID"
// @Success 200 {array} omc.Group
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/groups/{id}/children [get]
func HandleGetGroupChild(c echo.Context) error {
	group := echohandler.GetObjectFromEchoContext[omc.Group](c)
	db := store.GetDB()
	group.Children = omc.GetChildrenGroups(db, group.Schema, group.ID)
	return c.JSON(http.StatusOK, group)
}

type createGroupBody struct {
	Name     string
	ParentID uint
}

// HandleCreateGroup godoc
// @Summary Create a group
// @Tags OMC Groups
// @ID create-group
// @Security Bearer
// @Param body body createGroupBody true "Create Group Body"
// @Success 200 {object} omc.Group
// @Failure 400 {object} echo.HTTPError "Bad Request"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/groups [post]
func HandleCreateGroup(c echo.Context) error {
	var data createGroupBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	db := store.GetDB()
	group := &omc.Group{
		Name:     data.Name,
		ParentID: &data.ParentID,
	}

	if data.ParentID != 0 {
		parent := new(omc.Group)
		if err := db.First(parent, data.ParentID).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid parent group")
		}
		group.Parent = parent
	}

	if err := db.Create(group).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, group)
}

type updateGroupInfoBody struct {
	Name string
}

// HandleUpdateGroupInfo godoc
// @Summary Update group information
// @Tags OMC Groups
// @ID update-group-info
// @Security Bearer
// @Param id path int true "Group ID"
// @Param body body updateGroupInfoBody true "Update Group Info Body"
// @Success 200 {object} omc.Group
// @Failure 400 {object} echo.HTTPError "Bad Request"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/groups/{id} [put]
func HandleUpdateGroupInfo(c echo.Context) error {
	var data updateGroupInfoBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}

	updateColumns := []string{"name"}
	updateData := &omc.Group{
		Name: data.Name,
	}

	db := store.GetDB()
	group := echohandler.GetObjectFromEchoContext[omc.Group](c)
	if err := db.Model(group).Select(updateColumns).Updates(updateData).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, group)
}

type setGroupParentBody struct {
	ParentID uint
}

// HandleSetGroupParent godoc
// @Summary Set group parent
// @Tags OMC Groups
// @ID set-group-parent
// @Security Bearer
// @Param id path int true "Group ID"
// @Param body body setGroupParentBody true "Set Group Parent Body"
// @Success 200 {object} omc.Group
// @Failure 400 {object} echo.HTTPError "Bad Request"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/groups/{id}/parent [put]
func HandleSetGroupParent(c echo.Context) error {
	var data setGroupParentBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	group := echohandler.GetObjectFromEchoContext[omc.Group](c)
	db := store.GetDB()
	if data.ParentID != 0 {
		parent := new(omc.Group)
		if err := db.First(parent, data.ParentID).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid parent group")
		}
		if parent.ID == group.ID {
			return echo.NewHTTPError(http.StatusForbidden, "invalid parent group")
		}
		group.Children = omc.GetChildrenGroups(db, group.Schema, group.ID)
		if v := group.FindChildByID(parent.ID); v != nil {
			return echo.NewHTTPError(http.StatusForbidden, "invalid parent group")
		}
		group.Parent = parent
	}
	if err := db.Save(group).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, group)
}

// @Summary Delete a group
// @Tags OMC Groups
// @ID delete-group
// @Security Bearer
// @Param id path int true "Group ID"
// @Success 204 "No Content"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/groups/{id} [delete]
func HandleDeleteGroup(c echo.Context) error {
	db := store.GetDB().Unscoped()
	group := echohandler.GetObjectFromEchoContext[omc.Group](c)
	if err := db.Delete(group).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
