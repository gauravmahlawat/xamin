---
title: "Demo"
date: 2019-11-19T17:21:20+01:00
page: true
---


# Hugo Quiz

{{< quizdown >}}

---
primary_color: '#fc4c02'
secondary_color: lightgray
text_color: black
shuffle_questions: true
---

## who is prime minister of india?

> modi

- [x] Narendra modi
- [ ] gandhi
- [ ] nehru
- [ ] kejriwal

## who is chief minister of Delhi?

> kejri

- [ ] Narendra modi
- [ ] gandhi
- [ ] nehru
- [x] kejriwal

{{< /quizdown >}}

{{< rawhtml >}}
     <script 
     src="https://cdn.jsdelivr.net/npm/quizdown@latest/public/build/quizdown.js">
  </script>
  <script 
      src="https://cdn.jsdelivr.net/npm/quizdown@latest/public/build/extensions/quizdownKatex.js">
  </script>
  <script 
      src="https://cdn.jsdelivr.net/npm/quizdown@latest/public/build/extensions/quizdownHighlight.js">
  </script>
  <script>quizdown.register(quizdownHighlight).register(quizdownKatex).init()</script> 
{{< /rawhtml >}}