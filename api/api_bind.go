
package api

import (
    "github.com/labstack/echo/v4"
)

func handleMessage(c echo.Context) error {
    var msg Message
    if err := c.Bind(&msg); err != nil {
        return err
    }

    switch msg.Type {
    case "private":
        // 现有的私聊处理逻辑
    case "group":
        // 现有的群聊处理逻辑
    case "guild":
        // 新增的频道消息处理逻辑
        handleGuildMessage(msg)
    }

    return c.JSON(http.StatusOK, Response{"status": "ok"})
}
