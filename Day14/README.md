# Day 14: Minus-pluss-sekvens

Vi definerer en minus-pluss-sekvens som en sekvens der første element er 0 og andre element er 1. Resten av sekvensen blir generert på følgende vis:

* Først prøver man å substrahere n fra det (n-2)'te elementet. Dersom resultatet her er positivt og ikke finnes i sekvensen fra før av legges det til i sekvensen.
* Dersom resultatet er negativt eller i sekvensen fra før av adderer man n til det (n-2)'te elementet, og legger det til i sekvensen uansett om det finnes fra før av eller ikke.

## Oppgave

Hvor mange primtall finnes i en minus-pluss-sekvens med lengde 1800813?

## Eksempel

```
a(0) = 0
a(1) = 1
a(2) = 0 - 2 = -2. Det fungerer ikke fordi tallet er negativt.
a(2) = 0 + 2 = 2. Det fungerer med pluss.
a(3) = 1 - 3 = -2. Det fungerer ikke fordi tallet er negativt.
a(3) = 1 + 3 = 4. Det fungerer med pluss.
a(4) = 2 - 4 = -1. Det fungerer ikke fordi tallet er negativt.
a(4) = 2 + 4 = 6. Det fungerer med pluss.
a(5) = 4 - 5 = -1. Det fungerer ikke fordi tallet er negativt.
a(5) = 4 + 5 = 9. Det fungerer med pluss.
a(6) = 6 - 6 = 0. Det fungerer ikke fordi tallet allerede eksisterer i sekvensen.
a(6) = 6 + 6 = 12. Det fungerer med pluss.
a(7) = 9 - 7 = 2. Det fungerer ikke fordi tallet allerede eksisterer i sekvensen.
a(7) = 9 + 7 = 16. Det fungerer med pluss.
a(8) = 12 - 8 = 4. Det fungerer ikke fordi tallet allerede eksisterer i sekvensen.
a(8) = 12 + 8 = 20. Det fungerer med pluss.

...

a(13) = 18 - 13 = 5. Det fungerer med minus.

...

a(15) = 5 - 15 = -10. Det fungerer ikke fordi tallet er negativt.
a(15) = 5 + 15 = 20. Her blir tallet lagt til i sekvensen, selv om det allerede eksisterer.
```

Som gir en sekvens på 0, 1, 2, 4, 6, 9, 12, 16, 20, ..., 5, ..., 20