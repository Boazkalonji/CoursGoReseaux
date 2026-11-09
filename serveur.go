package main

import (
	//"context"
	"fmt"
	"net"
	"os"
	//"strconv"
	"strings"
	//"time"
	"sync"
)




type client struct{

	conn net.Conn
	pseudo string

}

type hub struct{

	clients map[net.Conn]*client
	mu sync.Mutex
	broadcast chan string

}



func (h *hub) DiffuserMessages(){



	for{

		message := <- h.broadcast 

		h.mu.Lock()

		for conn, _ := range h.clients{

			_,err := conn.Write([]byte(message + "\n"))

			if err != nil {

				conn.Close()
				delete(h.clients, conn)
				
			}

		}

		h.mu.Unlock()

	}

}




func main(){

	listener, err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil {

		fmt.Printf("une erreur survenue au moment de l'ouverture du serveur")
		os.Exit(1)

	}

	defer listener.Close()
	hub := hub{
		clients: make(map[net.Conn]*client),
		broadcast: make(chan string,100),
	}
	go hub.DiffuserMessages()

	fmt.Printf("🟢 OUVERT EST PRET A RECEVOIR LES APPEL")

	for {

		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Erreur connexion client:", err)
			continue
		}
		fmt.Printf("🟢 CONNEXION REUSSIT ")



		go func(c net.Conn){



			
			pseudo := make([]byte, 1024)
			n,_ := c.Read(pseudo)
			pseudoRecu := strings.TrimSpace(string(pseudo[:n]))

			
			hub.mu.Lock()
			hub.clients[c] = &client{conn: c, pseudo: pseudoRecu}
			hub.mu.Unlock()


			for {

				message := make([]byte, 1024)
				n,err := c.Read(message)

				if err != nil {
					hub.mu.Lock()
					delete(hub.clients, c)
					hub.mu.Unlock()
					c.Close()
					break
				}

				textRecu := strings.TrimSpace(string(message[:n]))
				hub.broadcast <- fmt.Sprintf("[%s] : %s", pseudoRecu, textRecu)



			}

			




		}(conn)

		

		
	}




}


/*
func gererClientCalcul(conn net.Conn, canalLogs chan OperationLog, canalResultat chan bool) {

	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(20 * time.Second))

	conn.Write([]byte("Bienvenue sur le serveur !\n"))

	datas := make([]byte, 1024)

	conn.SetReadDeadline(time.Now().Add(2 * time.Second))

	n, err := conn.Read(datas)

	if err != nil {
		fmt.Println("Erreur au moment de reception des donnees client :", err)
	}

	//recuperation est traitement du donnees

	textRecu := string(datas[:n])
	textRecu = strings.TrimSpace(textRecu)
	parties := strings.Split(textRecu, "+")

	fmt.Printf("LE TEXTE RECU EST : %v et est divisé en deux parties comme : %v et %v", textRecu, parties[0], parties[1])

	resultat := make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {

		if len(parties) != 2 {

			conn.Write([]byte("❌ Erreur : Utilise le format nombre+nombre (ex: 12+8)\n"))
			canalResultat <- false
			return

		}

		n1, err1 := strconv.Atoi(strings.TrimSpace(parties[0]))
		n2, err2 := strconv.Atoi(strings.TrimSpace(parties[1]))

		if err1 == nil && err2 == nil {

			response := n1 + n2

			//time.Sleep(6 * time.Second)

			canalResultat <- true
			resultat <- response

			return

		}
	}()

	select {

	case <-ctx.Done():

		// Le contexte a expiré (Timeout applicatif de 3s)
		fmt.Printf("⏰ [CONTEXTE] Traitement trop long pour le client %s, abandon.\n", conn.RemoteAddr().String())
		conn.Write([]byte("Erreur : Temps de traitement dépassé.\n"))
		return

	case r := <-resultat:

		if r == -99999 {
			conn.Write([]byte("Erreur : Format de calcul invalide (utilisez A+B).\n"))
			return
		}

		canalLogs <- OperationLog{

			ClientIP: conn.RemoteAddr().String(),
			Calcul:   textRecu,
			Resultat: r,
		}

		conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
		messageReponse := fmt.Sprintf("LA REPONSE EST  : %v \n", r)
		conn.Write([]byte(messageReponse))

	}

}

type OperationLog struct {
	ClientIP string
	Calcul   string
	Resultat int
}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil {

		fmt.Printf("une erreur survenue au moment de l'ouverture du serveur")
		os.Exit(1)

	}

	defer listener.Close()

	fmt.Printf("🟢 OUVERT EST PRET A RECEVOIR LES APPEL")

	canalLogs := make(chan OperationLog, 2)
	canalResultat := make(chan bool)

	go func() {

		NbrClient := 0

		for logRecu := range canalLogs {

			NbrClient = NbrClient + 1
			fmt.Printf("📊 [ADMIN STATS #%d] IP: %s | Opération: %s | Résultat: %d\n",
				NbrClient,
				logRecu.ClientIP,
				logRecu.Calcul,
				logRecu.Resultat)

		}

	}()

	for {

		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Erreur connexion client:", err)
			continue
		}

		go gererClientCalcul(conn, canalLogs, canalResultat)

		status := <-canalResultat

		if status {

			defer listener.Close()
		}

	}


	
}
*/