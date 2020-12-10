# Day 10: Alvelekene

Hvert fjerde år arrangeres Alvelekene, med alt det tilhørende av premiering, pomp og prakt. Tradisjonen tro så skal Julenissen dele ut premien til vinneren av Alvelekene, men han trenger hjelp til å finne ut hvem som faktisk vant.

Til å hjelpe seg har Julenissen [resultatlisten](./leker.txt) for lekene, men ikke poengene som alvene har fått i løpet av lekene.

Hver linje i resultatene representerer én øvelse, hver alv er representert med de samme 2 bokstavene i alle øvelsene, og for hver linje er vinneren skrevet ned først og sisteplassen skrevet ned sist.

Hver øvelse er verdt like mange poeng som antall deltakere minus plassering, som betyr at hvis man får andreplass i en øvelse hvor det er 12 deltakere, vil man få 12-2=10 poeng. Hvis man ikke deltar får man 0 poeng.

## Oppgave

Hvilken alv vant, og hvor mye poeng fikk vinneren? Svaret skal være på formen vinner-poengsum.

## Eksempel

Med resultatlisten under, så vil svaret være **ae-11**.

```
ae,af,aa,ab,ad,ac
aa,ac,ab,ad,af,ae
ad,ae,ab,aa,af
ad,ac,aa,ab
af,ae,ab,aa,ac
```