1. 	Go für Windows herunterladen (MSI Installer)
	https://golang.org/dl/

2. 	Öffne den MSI Installer und befolge die Installationsanleitung
	Per Default wird die Go Distribution in C:/Go gespeichert
	
3.	Der Installer sollte den C:\Go\bin Pfad auch als PATH Environment aufnehmen.
	Restarte deine Powershell und teste die Installation mit:
	> go version
	
	Output sollte etwa so aussehen
	go version go1.9.2 windows/amd64
	
*	Falls dies nicht der fall ist muss C:\Go\bin der PATH Environment hinzugefügt werden.