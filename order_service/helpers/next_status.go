package helpers

import (
	"errors"
	"order_service/utils"
)

func GetNextStatus(currentStatus string) (string, error) {
	switch currentStatus {
	case utils.OrderStatusWasherPreparing:
		return utils.OrderStatusWasherOnGoing, nil
	case utils.OrderStatusWasherOnGoing:
		return utils.OrderStatusWasherArrived, nil
	case utils.OrderStatusWasherArrived:
		return utils.OrderStatusWasherWashing, nil
	case utils.OrderStatusWasherWashing:
		return utils.OrderStatusWasherFinished, nil
	default:
		return "", errors.New("order status cannot be updated further")
	}
}
