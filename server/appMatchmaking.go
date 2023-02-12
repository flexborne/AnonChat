package main

func (app *App) handleNewFindMatch() {
	for {
		select {
		case c := <-app.ClientsToMatch:
			c.Mutex.Lock()

			matchMakingParams := convertToMatchmakingParams(c)
			c.Mutex.Unlock()
			app.Clients.Store(matchMakingParams.Id, c)

			app.RandomMatchmakingCh <- matchMakingParams
		}
	}
}

func (app *App) connect(c1, c2 *Client) {
	// Connect the clients to each other
	c1.Conn.WriteJSON(map[string]string{"type": "event", "event": "connect", "id": c2.Id})
	c2.Conn.WriteJSON(map[string]string{"type": "event", "event": "connect", "id": c1.Id})

	pair := newPair(c1, c2)

	go pair.run()
}

func (app *App) handleMatchResult(matchResult <-chan ClientMatchmakingResult) {
	for {
		select {
		case bestMatch := <-matchResult:
			finderV, ok := app.Clients.Load(bestMatch.FinderId)
			if !ok {
				continue
			}
			finder := finderV.(*Client)

			bestMatchClientV, ok := app.Clients.Load(bestMatch.BestMatchId)
			if !ok {
				continue
			}

			bestMatchClient := bestMatchClientV.(*Client)

			finder.Mutex.Lock()
			bestMatchClient.Mutex.Lock()

			app.connect(finder, bestMatchClient)

			app.MatchmakingResponseCh <- NewOptional(convertToMatchmakingParams(bestMatchClient))

			bestMatchClient.Mutex.Unlock()
			finder.Mutex.Unlock()
		}
	}
}

func (app *App) runMatchmaking() {
	go app.RandomMatchmakingManager.run(app.RandomMatchmakingCh, app.MatchmakingResponseCh,
		app.MatchmakingResultCh,
		app.RemoveFromQueueCh)
	go app.handleMatchResult(app.MatchmakingResultCh)
}
