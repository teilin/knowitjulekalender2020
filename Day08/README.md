# Day 8: Julenissen kommer

Det er allment kjent at Nissen er tung. Blytung. Absolutt hinsides tung faktisk. Han er et vesen skapt før universet, hans sjel bunnløs som tiden selv... Nissens tyngdekraft gjør at tiden går saktere jo nærmere han du kommer.

Vi vil finne ut hva slags kaos han skaper idet han reiser rundt på Julaften ved å gjøre et eksperiment. I fila [input.txt](./input.txt) ligger det først 50 lokasjoner `navn: (x, y)` han skal besøke, deretter ruten hans som en liste av lokasjonsnavn. Han kan besøke hver lokasjon opp til flere ganger.

Nissen beveger seg mellom ruter kun langs rutenes kanter ([Taxicab Geometry](https://en.wikipedia.org/wiki/Taxicab_geometry)). Han reiser først horisontalt fram til han er i riktig `x`, deretter reiser han vertikalt til riktig `y`. For hver rute han reiser langs rutenettet går ett sekund for alle punktene i nettet - men for de som er nærmere Nissen enn 50 ruter går det bare 0.75 sekunder, for de som er nærmere enn 20 ruter går det 0.5 sekunder og for de som er nærmere Nissen enn 5 ruter går det kun 0.25 sekunder. Der han står går ikke tiden i det hele tatt.

Når Nissen har reist hele ruta si, hva er den største tidsforskjellen mellom to av de navngitte punktene? Han begynner i posisjon (0, 0). Svaret gis med 2 desimaler og punktum som desimalskilletegn.

## Eksempel

```
...x <<- Nordpolen (3, 4)
....
....
.x.. <<- Rovaniemi (1, 1)
J...
```

Julenissen starter her på `J` i `(0, 0)` og han skal først til Rovaniemi, deretter videre til Nordpolen. Turen hans går som følger:

* Han drar østover til `(1, 0)`. Nordpolen på `(3, 4)` er nå 6 ruter unna i Taxicab Geometry så der går det 0.5 sekunder. Han er nærmere Rovaniemi enn 5 ruter, dermed har det i Rovaniemi kun gått 0.25 sekunder.
* Han reiser nordover til Rovaniemi på `(1, 1)`. Der går det nå 0 sekunder. Nordpolen er nå nøyaktig 5 ruter unna så der går det 0.5 sekunder.
* `(2, 1)`, Rovaniemi +0.25s. Nordpolen er nå nærmere enn 5 ruter, så der går det 0.25s.
* `(3, 1)`, Rovaniemi +0.25s, Nordpolen +0.25s.
* `(3, 2)`, Rovaniemi +0.25s, Nordpolen +0.25s.
* `(3, 3)`, Rovaniemi +0.25s, Nordpolen +0.25s.
* `(3, 4)`, Rovaniemi +0.5s, Nordpolen +0.00s.

Det har gått 1.75 sekunder i Rovaniemi og 2 sekunder på Nordpolen. Svaret er `0.25`.