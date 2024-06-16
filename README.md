# GO-Restaurant_Concurrency

## Opis ogólny

Zadaniem kodu jest uproszczona symulacja pracy restauracji.
Klienci wchodzą do restauracji i szukają stolika. Kiedy go znajdą siadają przy nim i składają zamówienie. Kiedy zamówienie zostanie złożone , wolny kuchaż zaczyna przygotowywać zamówienie. Kiedy zamówienie jest gotowe jest ono dostarczane do klienta przez wolnego kelnera. Następnie kliet zjada dostarczone mu danie poczym opuszcza restauracje.
Kod jest napisany w języku Go (Golang) i wykorzystuje współbieżność do równoczesnego wykonywania operacji.


## Opis i struktura Kodu

### Funkcja `main()`

W funkcji `main()` inicjalizowane są zmienne symulacji takie jak: liczba klientów, liczba dostępnych stolików, liczba kucharzy i kelnerów. 
 - `numCustomers` - zmienna przedstawia liczbą wszystkich **klientów** którzy pojawię się w restauracji podczas symulacji.
 - `numTables` - zmienna przedstawia liczbą wszystkich dostępnych **stolików** które dostępne są dla klientów restauracji podczas symulacji.
 - `numChefs` - zmienna przedstawia liczbą wszystkich **kucharzy** którzy pracują w restauracji podczas symulacji.
 - `numWaiters` - zmienna przedstawia liczbą wszystkich **kelnerów** którzy pracują w restauracji podczas symulacji.


Tworzone są również kanały do komunikacji między gorutynami:
- `tables` - kanał przechowujący dostępne stoliki,
- `orders` - kanał przechowujący zamówienia składane przez klientów,
- `preparedOrders` - kanał przechowujący zamówienia gotowe do dostarczenia,
- `deliveredOrders` - kanał przechowujący zamówienia dostarczone do klientów.

Następnie stoliki są inicjalizowane w pętli.

Dalej uruchamiane są gorutyny symulujące pracę kucharzy i kelnerów. 

Na koniec Klienci są symulowani poprzez uruchomienie gorutyn w pętli, z losowym opóźnieniem które ma za zadnie zasymulować wchodzenie klientów aby wszyscy nie wchodzili do restauracji w tym samym czasie.

### Funkcja `customer()`

**Funkcja symuluje zachowanie klienta w restauracji:**
1. Klient wchodzi do restauracji.
2. Klient czeka na wolny stolik.
3. Klient siada przy wolnym stoliku
4. Klient składa zamówienie.
5. Klient następnie czeka na zrealizowanie i dostarczenie zamówionego przez niego jedzenia.
6. Po otrzymaniu jedzenia, klient symuluje jedzenie.
7. Po zakończeniu jedzenia klient opuszcza stolik i restauracje.

### Funkcje `chef()` i `waiter()`

Funkcje te wykonują symulacje pracy kucharzy i kelnerów w restauracji:
- `chef(id int, orders chan int, preparedOrders chan int)`: Kucharz pobiera zamówienia z kanału `orders`, przygotowuje jedzenie przez określony czas, a następnie wysyła gotowe zamówienie do kanału `preparedOrders`.
- `waiter(id int, preparedOrders chan int, deliveredOrders chan int)`: Kelner pobiera gotowe zamówienia z kanału `preparedOrders`, dostarcza je do klientów poprzez kanał `deliveredOrders`.

## Współbieżność

W kodzie wykorzystywana jest współbieżność za pomocą gorutyn w celu równoczesnego wykonania różnych operacji:
- Gorutyny dla każdego klienta pozwalają na jednoczesne obsługiwania wielu klientów w restauracji.
- Kucharze i kelnerzy działają równolegle, co umożliwia jednoczesne przygotowywanie i dostarczanie zamówień.

