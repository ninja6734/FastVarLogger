# UDP-Codec

Der UDP-Codec ist ein Protokoll zur Übertragung von Daten über das Internet. Er besteht aus vier Teilen:

*   Magic Number: Ein 4-Bit-Wert, der den Anfang des Pakets markiert.
*   Daten: Die eigentlichen Daten, die übertragen werden sollen.
*   End Codec: Ein 2-Bit-Wert, der das Ende des Pakets markiert.

## Funktionsweise

Der UDP-Codec arbeitet wie folgt:

1.  Der Sender erstellt ein neues Paket und fügt den Magic Number hinzu.
2.  Anschließend werden die Daten hinzugefügt.
3.  Schließlich wird der End Codec hinzugefügt, um das Ende des Pakets zu markieren.

## Beispiel

Ein Beispiel für eine UDP-Datenübertragung könnte wie folgt aussehen:

*   Magic Number: `FVLR`
*   Daten: `Hello, Welt!`
*   End Codec: `\x00\x00`

Der Empfänger kann dann das Paket entpacken und die Daten lesen.
