

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


# Authorization

We manage the authentication via middleware. More specifically, gin middleware, but the concept is always the same.

Via gin routes shortcuts (e.g. .GET, .POST ... ) we can provide multiple handlers where the last one is rhe real handler and previous ones are middleware handlers.


With authMiddleware we take a token maker (we have two different implementations, one with JWT and one with Paseto tokens).

To implement a middleware with gin is very simple. Just create a function (with all needed arguments) and return a gin.HanlderFunc.

Steps are:
- get authorization header from the request context (gin ctx)
- split header into fields and extract authorization type
- extract access token
- verify token (call to tokenMaker)
- set the payload into header
- call next handler (ctx.Next())


When we verify the token... we get back an authorization payload.
This contains username and info like, id, issuedAt and expiredAt fields.

BUSINESS RULE: The owner of an account must be set to the username of the authorized user.

