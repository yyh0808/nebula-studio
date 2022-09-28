// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	favorite "github.com/vesoft-inc/nebula-studio/server/api/studio/internal/handler/favorite"
	file "github.com/vesoft-inc/nebula-studio/server/api/studio/internal/handler/file"
	gateway "github.com/vesoft-inc/nebula-studio/server/api/studio/internal/handler/gateway"
	health "github.com/vesoft-inc/nebula-studio/server/api/studio/internal/handler/health"
	importtask "github.com/vesoft-inc/nebula-studio/server/api/studio/internal/handler/importtask"
	schema "github.com/vesoft-inc/nebula-studio/server/api/studio/internal/handler/schema"
	sketches "github.com/vesoft-inc/nebula-studio/server/api/studio/internal/handler/sketches"
	"github.com/vesoft-inc/nebula-studio/server/api/studio/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/health",
				Handler: health.GetHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/exec",
				Handler: gateway.ExecNGQLHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/batchExec",
				Handler: gateway.BatchExecNGQLHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api-nebula/db"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/connect",
				Handler: gateway.ConnectHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/disconnect",
				Handler: gateway.DisonnectHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api-nebula/db"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPut,
				Path:    "/api/files",
				Handler: file.FileUploadHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/files/:name",
				Handler: file.FileDestroyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/files",
				Handler: file.FilesIndexHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/import-tasks",
				Handler: importtask.CreateImportTaskHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/import-tasks/:id",
				Handler: importtask.GetImportTaskHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/import-tasks",
				Handler: importtask.GetManyImportTaskHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/import-tasks/:id/logs",
				Handler: importtask.GetManyImportTaskLogHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/import-tasks/:id/task-log-names",
				Handler: importtask.GetImportTaskLogNamesHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/import-tasks/:id",
				Handler: importtask.DeleteImportTaskHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/import-tasks/:id/stop",
				Handler: importtask.StopImportTaskHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/import-tasks/:id/download-logs",
				Handler: importtask.DownloadLogsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/import-tasks/:id/download-config",
				Handler: importtask.DownloadConfigHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/import-tasks/working-dir",
				Handler: importtask.GetWorkingDirHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sketches/sketch",
				Handler: sketches.InitHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/sketches/list",
				Handler: sketches.GetListHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/sketches/:id",
				Handler: sketches.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/sketches/:id",
				Handler: sketches.UpdateHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPut,
				Path:    "/api/schema/:space/snapshot",
				Handler: schema.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/schema/:space/snapshot",
				Handler: schema.GetSnapshotHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/favorites",
				Handler: favorite.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/favorites/list",
				Handler: favorite.GetListHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/favorites/:id",
				Handler: favorite.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/favorites",
				Handler: favorite.DeleteAllHandler(serverCtx),
			},
		},
	)
}
