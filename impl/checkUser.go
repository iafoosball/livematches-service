package impl

func checkUser(id string) {
	for _, m := range LiveMatches {
		for c := range m.Clients {
			if c.ID == id {
				closeUser(c)
			}
		}
	}
}
