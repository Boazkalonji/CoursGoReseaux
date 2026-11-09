package main

import(


	"fmt"
	"os"
	"net"
	//"strings" 


)





func main() {
	call, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Erreur du serveur:", err)
		os.Exit(1)
	}
	defer call.Close()

	// --- 1. GESTION DU PSEUDO ---
	fmt.Printf("🟢 VEUILLEZ SAISIR VOTRE PSEUDO :\n")
	pseudo := make([]byte, 1024)
	n, err := os.Stdin.Read(pseudo)
	if err != nil {
		fmt.Println("Erreur de lecture du pseudo :", err)
		return
	}

	// On envoie le pseudo au serveur
	_, err = call.Write(pseudo[:n])
	if err != nil {
		fmt.Println("Erreur d'envoi du pseudo :", err)
		return
	}

	go func(){
		serveurBuff := make([]byte,1024)
		n,err = call.Read(serveurBuff)
		if err != nil {
			fmt.Println("\n🔴 Déconnecté du serveur.")
			os.Exit(0)
		}
		os.Stdout.Write(serveurBuff[:n])
	}()


	for{

		// --- 2. GESTION DU MESSAGE ---
		fmt.Printf("🟢 VEUILLEZ SAISIR VOTRE MESSAGE :\n")
		message := make([]byte, 1024)
		
		// Utilisation de '=' car 'err' et 'n' existent déjà au-dessus
		n, err = os.Stdin.Read(message)
		if err != nil {
			fmt.Println("Erreur de lecture du message :", err)
			return
		}


		


		// On envoie le message au serveur
		_, err = call.Write(message[:n])
		if err != nil {
			fmt.Println("Erreur d'envoi du message :", err)
			return
		}
	}
}
/*

func main(){


	call,err := net.Dial("tcp","127.0.0.1:8080")

	if err != nil {
		fmt.Println("Erreur du serveur:", err)
		os.Exit(1)
	}
	defer call.Close()


	message := make([]byte, 1024)
	n, err := call.Read(message)
	if err != nil {
		fmt.Println("Erreur de lecture du message d'accueil:", err)
		return
	}
	fmt.Printf("voici le message du serveur :  %v", string(message[:n]))


	operations := make([]byte, 1024)

	fmt.Printf("VEUILLEZ SAISIR L'operation sous cette forme (50+50) \n")

	nInput, err := os.Stdin.Read(operations)
	if err != nil {
		fmt.Println("Erreur lors de la saisie clavier:", err)
		return
	}

	
	texteNettoye := strings.TrimSpace(string(operations[:nInput]))
	operations = []byte(texteNettoye)
	_, err = call.Write(operations)

	if err != nil {
		fmt.Printf("une erreur au moment de envoie : %v", err)
		return
	}

	ResponseBuffer := make([]byte, 1024)
	Nresponse, err := call.Read(ResponseBuffer)

	if err != nil {
		fmt.Println("Erreur de lecture de la réponse:", err)
		return
	}

	fmt.Printf("✨ %v", string(ResponseBuffer[:Nresponse]))


	








}

*/