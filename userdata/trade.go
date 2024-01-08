package userdata

import (
	"elichika/client"
	"elichika/model"
	"elichika/utils"
)

func (session *Session) GetTradeProductUser(productId int) int {
	result := 0
	exist, err := session.Db.Table("u_trade_product").
		Where("user_id = ? AND product_id = ?", session.UserId, productId).
		Cols("traded_count").Get(&result)
	utils.CheckErr(err)
	if !exist {
		result = 0
	}
	return result
}

func (session *Session) SetTradeProductUser(productId, newTradedCount int) {
	record := model.TradeProductUser{
		ProductId:   productId,
		TradedCount: newTradedCount,
	}
	exist, err := session.Db.Table("u_trade_product").
		Where("user_id = ? AND product_id = ?", session.UserId, productId).
		Update(record)
	utils.CheckErr(err)
	if exist == 0 {
		genericDatabaseInsert(session, "u_trade_product", record)
	}
}

func (session *Session) GetTrades(tradeType int32) []model.Trade {
	trades := []model.Trade{}
	for _, trade_ptr := range session.Gamedata.TradesByType[tradeType] {
		trade := *trade_ptr
		for j, product := range trade.Products {
			product.TradedCount = session.GetTradeProductUser(product.ProductId)
			trade.Products[j] = product
		}
		trades = append(trades, trade)
	}
	return trades
}

// return whether the item is added to present box
func (session *Session) ExecuteTrade(productId int, tradeCount int) bool {
	// update count
	tradedCount := session.GetTradeProductUser(productId)
	tradedCount += tradeCount
	session.SetTradeProductUser(productId, tradedCount)

	// award items and take away source item
	product := session.Gamedata.TradeProduct[productId]
	trade := session.Gamedata.Trade[product.TradeId]
	content := product.ActualContent
	content.ContentAmount *= int32(tradeCount)
	session.AddResource(content)
	session.RemoveResource(client.Content{
		ContentType:   trade.SourceContentType,
		ContentId:     int32(trade.SourceContentId),
		ContentAmount: int32(product.SourceAmount) * int32(tradeCount),
	})
	return true
}
