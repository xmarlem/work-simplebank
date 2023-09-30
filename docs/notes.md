

L'ultimo video che stavo vedendo e':

#9 relativo gli isolation levels...

In questo video spiega solo gli isolation levels di postgres (facendo anche il comparison con mysql).
Fa anche vedere praticamente cosa succede settando i vari isolation level direttamente nel terminale di psql...

Fa partire due terminal e gioca con due transazioni...
`psql -h localhost -U postgres -d work_simplebank`

Facendo vedere come funziona con i vari isolation level... e.g. una transazione vede o non vede quello che fa l'altra...



## go mock

voglio capire in quali situazioni mi conviene usare go mock.
Da quello che ho capito serve per rimuovere dipendenze da altri sottosistemi.
Eg.. se ho un api call che chiama uno storage... sottostante oppure un altro service....si puo' simulare questa chiamata.


Con go mock possiamo generare delle response che matchano con le nostre dipendenze estena senza usare tali dipendenze.

