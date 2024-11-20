# ToDoList:

- Tester les fonctions : avec la fonction interface pour comprendre la logique:
	1. Avec des demandes POST, GET, DELETE au bon format, étape par étape
	2. Avec des demandes au mauvais format, idem pour le process
	3. Recherche de nouvelles méthodes à ajouter

- A regler :

	Reprendre cette partie: 	
	
	if res := s.DoesTableExist("reservation_list"); !res {
		fmt.Println("La DB Reservation n'existe pas..............")
		return nil
	}

	Et faire une fonction qui crée la TABLE à la places

	Tester tous les types de requetes
	Tester tous les types de requetes avec differents types d'erreurs
