
package main

import (
	"fmt"
	"os"
	"net"
	"strconv" 
	"strings" 
)






func gererClientCalcul(conn net.Conn) {

	defer conn.Close()

	conn.Write([]byte("Bienvenue sur le serveur !\n"))

	datas := make([]byte, 1024)

	n, err := conn.Read(datas)

	if err != nil{
		fmt.Println("Erreur au moment de reception des donnees client :", err)
	}


	//recuperation est traitement du donnees

	textRecu := string(datas[:n])
	textRecu = strings.TrimSpace(textRecu)
	parties := strings.Split(textRecu, "+")

	fmt.Printf("LE TEXTE RECU EST : %v et est divisé en deux parties comme : %v et %v", textRecu, parties[0],parties[1])

	if len(parties) != 2 {

		conn.Write([]byte("❌ Erreur : Utilise le format nombre+nombre (ex: 12+8)\n"))

	}
	
	n1, err1 := strconv.Atoi(strings.TrimSpace(parties[0]))
	n2, err2 := strconv.Atoi(strings.TrimSpace(parties[1]))

	if err1 == nil && err2 == nil {

		response := n1 +n2

		messageReponse := fmt.Sprintf("LA REPONSE EST  : %v \n", response)

		conn.Write([]byte(messageReponse))
		return

	}








}




func main(){



	listener, err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil {

		fmt.Printf("une erreur survenue au moment de l'ouverture du serveur")
		os.Exit(1)

	}

	defer listener.Close()

	fmt.Printf("🟢 OUVERT EST PRET A RECEVOIR LES APPEL")


	for {

		conn,err := listener.Accept()
		
		if err != nil {
			fmt.Println("Erreur connexion client:", err)
			continue
		}

		go gererClientCalcul(conn)



	}









}