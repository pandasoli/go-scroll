<div align='center'>

# [GoScroll](https://github.com/pandasoli/go-scroll) <img width='32' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-victorious.png' alt='Icone de Gopher'/>

Uma boa biblioteca para boas barras de rolagem!
<img width='20' src='https://raw.githubusercontent.com/pandasoli/twemojis/master/1f693.svg' alt='TV twemoji'/>
</div>

Depois de surtar durante alguns belos dias <img width='20' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-tired.png' alt='Icone de Gopher'/> eu finalmente consegui reinventar a roda, ou melhor a matemática por trás de uma scrollbar. Então decidi fazer uma biblioteca pra parar de ficar copiando código entre projetos <img width='20' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-idea.png' alt='Icone de Gopher'/>.

<br/>

## Matemática <img width='27' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-not-sure-if.png' alt='Icone de Gopher'/>

Para calcular o tamanho da thumb:
```
bigger(1, viewport_size * (viewport_size / content_size))
```

<br/>

Para a posição dela dentro da scrollbar:
```
scroll / (content_size - viewport_size) * (viewport_size - thumb_size)
```

<br/>

Você pode muito bem avançar um pixel/row por vez,  
mas caso vc queira avançar a quantia que a thumb representa, essa é a quantia:
```
1 / track_size * content_size
```

<br/>

Para saber se é possível continuar rolando:
```
scroll + viewport_size < content_size
```

<br/>
<div align='right'>

## Como usar <img width='27' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-blushing.png' alt='Icone de Gopher'/>
</div>

Apesar da biblioteca ser bem simples e ter um funcionamento simples, acho que seria dificil exemplificar com tão poco código (já que ela foi feita para códigos CLI mais completos), então deixo [um exemplo](./tests).

<br/>

## Algumas informações <img width='27' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-idea.png' alt='Icone de Gopher'/>

Icones de Gopher por [Egon Elbre](https://github.com/egonelbre/gophers) <img width='20' src='https://raw.githubusercontent.com/egonelbre/gophers/master/icon/emoji/gopher-mind-blown.png' alt='Gopher icon'/>
