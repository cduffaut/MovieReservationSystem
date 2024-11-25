# ToDoList:

- Tester les fonctions : avec la fonction interface pour comprendre la logique:
	1. Avec des demandes POST, GET, DELETE au bon format, étape par étape
	2. Avec des demandes au mauvais format, idem pour le process
	3. Recherche de nouvelles méthodes à ajouter

- A regler :

	Comprendre pourquoi cette ligne:
	
	query := `INSERT INTO reservation_list (FirstName, Name, Mail, Date, Time, MovieName) VALUES (reservation.FirstName, reservation.Name, reservation.Mail, reservation.Date, reservation.Time, reservation.MovieName)`

	Produit cette erreur:

	pq: missing FROM-clause entry for table "reservation"

	Tester tous les types de requetes
	Tester tous les types de requetes avec differents types d'erreurs
