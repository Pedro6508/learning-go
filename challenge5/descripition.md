# Desafio: escrever e ler assincronamente o valor de um channel em Go
> Implemente as funções ```verOppenheimer```, ```verBarbie``` e ```comecar```
no arquivo ```movieWar.go``` seguindo esses objetivos:
>> - ```comecar```: Deve chamar as funções ```verOppenheimer``` e ```verBarbie```
de forma assíncrona e também preencher constantemente o channel ```nome``` com
um nome aleatório.
>> - ```verOppenheimer```: Caso exista uma string de um nome no channel ```nome```,
deve imprimir uma frase dizendo que alguém com esse nome
viu o filme Oppenheimer e fazer a rotina em questão "dormir" por 3 segundos.
Caso contrário, a função deve imprimir uma frase indicando que ninguém está
vendo o filme Oppenheimer.
>> - ```verBarbie```: Análoga à função ```verOppenheimer```, mudando o filme
em questão para Barbie e o tempo que a rotina irá "dormir" para 2 segundos.

# Material de Apoio
- [Como ler um channel](https://www.bogotobogo.com/GoLang/GoLang_Channel_with_Select.php#:~:text=Combining%20goroutines%20and%20channels%20with,wait%20on%20multiple%20communication%20operations.)
- [Como colocar uma rotina para "dormir"](https://www.geeksforgeeks.org/time-sleep-function-in-golang-with-examples/)
- Lembre-se, caso apareça qualquer dúvida você pode e deve pedir ajuda no grupo
ou no privado.
