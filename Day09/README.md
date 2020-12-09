# Day 9: Dei arktiske alvane sine arbitrære avstandsregler

I eit forsøk på å stanse spreiinga av Tuborg-viruset har Julenissen bestemt at alle alvar skal halde seg innanfor kvart sit oppteikna kvadrat til smitta stoppar opp. Diverre er kvadrata litt for små, slik at viruset likevel kan smitte om ein er så uheldig at ein har minst to naboar (i horisontal og vertikal retning) som sjølv er sjuke. Om ein blir smitta, er ein sjølv ikkje smittsam før neste dag. I [denne fila](./elves.txt) ser vi oversikta over dei friske (f) og sjuke (s) alvane. Kor mange dagar går det før smitta stansar opp?

## Eksempel

Vi startar med alvane som vist under.

```
FFFSF  
FSFFF  
FSFSF  
SFFSF  
FSFFF
```

Alle alvane som har to eller fleire tilstøtande sjuke naboar blir sjølv sjuke og smittsame dagen derpå:

```
FFFSF
FSFSF
SSSSF
SSFSF
SSFFF
```

Og så vidare ...

```
FFFSF
SSSSF
SSSSF
SSSSF
SSFFF
```

```
FFSSF
SSSSF
SSSSF
SSSSF
SSSFF
```

```
FSSSF
SSSSF
SSSSF
SSSSF
SSSSF
```

Til slutt endar vi i ein tilstand der ingen fleire kan blir smitta:

```
SSSSF
SSSSF
SSSSF
SSSSF
SSSSF
```

Dette tek **6 dagar**, som då blir svaret for dette eksempelet.