package impl

func checkUser(id string) {
	for _, m := range LiveMatches {
		for c, _ := range m.Clients {
			if c.ID == id {
				c.LiveMatch.Unregister <- c
				c.End <- true
			}
		}
	}
}
