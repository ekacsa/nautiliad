package rpchandlers

import (
	"github.com/Nautilus-Network/nautiliad/app/appmessage"
	"github.com/Nautilus-Network/nautiliad/app/rpc/rpccontext"
	"github.com/Nautilus-Network/nautiliad/infrastructure/network/netadapter/router"
)

// HandleNotifyVirtualSelectedParentBlueScoreChanged handles the respectively named RPC command
func HandleNotifyVirtualSelectedParentBlueScoreChanged(context *rpccontext.Context, router *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	listener, err := context.NotificationManager.Listener(router)
	if err != nil {
		return nil, err
	}
	listener.PropagateVirtualSelectedParentBlueScoreChangedNotifications()

	response := appmessage.NewNotifyVirtualSelectedParentBlueScoreChangedResponseMessage()
	return response, nil
}
