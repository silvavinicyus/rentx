package routes

import (
	"net/http"
	"rentx/src/controller"
)

var categoriesRoutes = []Route{
	{
		Uri:                    "/categories",
		Method:                 http.MethodPost,
		Function:               controller.CreateCategory,
		RequiresAuthentication: false,
	},
	{
		Uri:                    "/categories/{uuid}",
		Method:                 http.MethodGet,
		Function:               controller.FindCategoryBy,
		RequiresAuthentication: false,
	},
	{
		Uri:                    "/categories",
		Method:                 http.MethodGet,
		Function:               controller.FindAllCategories,
		RequiresAuthentication: false,
	},
	{
		Uri:                    "/categories/{uuid}",
		Method:                 http.MethodPut,
		Function:               controller.UpdateCategory,
		RequiresAuthentication: false,
	},
	{
		Uri:                    "/categories/{uuid}",
		Method:                 http.MethodDelete,
		Function:               controller.DeleteCategory,
		RequiresAuthentication: false,
	},
}
