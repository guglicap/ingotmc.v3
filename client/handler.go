package client


// handleLogin handles authentication.
// TODO: should be a bigger deal, implement encryption, compression
// func handleLogin(cl *Client, nc action.NewConnection) {
// 	userUUID, err := cl.auth.Authenticate(nc.Username)
// 	if err != nil {
// 		pkt, _ := cl.packetFor(event.Disconnect{Reason: err})
// 		cl.cbound <- pkt
// 		return
// 	}
// 	pkt, _ := cl.packetFor(event.AuthSuccess{
// 		UUID:     userUUID,
// 		Username: nc.Username,
// 	})
// 	cl.cbound <- pkt
// 	cl.actions <- action.NewPlayer{
// 		UUID:     userUUID,
// 		Username: nc.Username,
// 	}
// 	fmt.Println("sent new player")
// }
